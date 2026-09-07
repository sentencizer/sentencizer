package cleaner

import (
	"regexp"
	"strings"

	"github.com/sentencizer/sentencizer/internal/rule"
)

// Cleaner cleans noisy text before sentence segmentation.
// Ported from pySBD pysbd/cleaner.py and pysbd/clean/rules.py.
// Go's RE2 has no lookaround; lookaround rules are rewritten with captures.
type Cleaner struct {
	docType string // "" or "pdf"
}

type Option func(*Cleaner)

// WithDocType enables PDF-specific newline handling (pySBD doc_type='pdf').
func WithDocType(docType string) Option {
	return func(c *Cleaner) {
		c.docType = docType
	}
}

func NewCleaner(opts ...Option) *Cleaner {
	c := &Cleaner{}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

func (c *Cleaner) Clean(text string) string {
	if text == "" {
		return text
	}
	text = c.removeAllNewlines(text)
	text = c.replaceDoubleNewlines(text)
	text = c.replaceNewlines(text)
	text = c.replaceEscapedNewlines(text)
	text = htmlRules.Apply(text)
	text = c.replacePunctuationInBrackets(text)
	text = inlineFormattingRule.Apply(text)
	text = c.cleanQuotations(text)
	text = c.cleanTableOfContents(text)
	text = c.spaceSentences(text)
	text = c.cleanConsecutiveCharacters(text)
	return text
}

func (c *Cleaner) removeAllNewlines(text string) string {
	text = c.removeNewlineInMiddleOfSentence(text)
	text = c.removeNewlineInMiddleOfWord(text)
	return text
}

// removeNewlineInMiddleOfWord ports \n(?=[a-zA-Z]{1,2}\n) → "".
// Lookahead cannot be expressed in RE2; matching the full \nXX\n and deleting
// only the first \n must be applied iteratively so adjacent letter-lines join
// (e.g. "\nW\nA\nRN\nI\nNG\n" → "WARNING\n").
func (c *Cleaner) removeNewlineInMiddleOfWord(text string) string {
	for {
		loc := newLineInMiddleOfWordRegex.FindStringIndex(text)
		if loc == nil {
			return text
		}
		// drop the leading '\n' of the match
		text = text[:loc[0]] + text[loc[0]+1:]
	}
}

// removeNewlineInMiddleOfSentence removes newlines that sit between a space
// and a lowercase letter / '('. pySBD applies this only inside non-'.' runs;
// applying the capture rewrite globally is equivalent for the supported rules.
func (c *Cleaner) removeNewlineInMiddleOfSentence(text string) string {
	return newLineInMiddleOfSentenceRule.Apply(text)
}

func (c *Cleaner) replaceDoubleNewlines(text string) string {
	return rule.Rules{doubleNewLineWithSpaceRule, doubleNewLineRule}.Apply(text)
}

func (c *Cleaner) replaceNewlines(text string) string {
	if c.docType == "pdf" {
		return c.removePDFLineBreaks(text)
	}
	return rule.Rules{newLineFollowedByPeriodRule, replaceNewlineWithCarriageReturnRule}.Apply(text)
}

func (c *Cleaner) removePDFLineBreaks(text string) string {
	return rule.Rules{
		newLineFollowedByBulletRule,
		pdfNewLineInMiddleOfSentenceRule,
		pdfNewLineInMiddleOfSentenceNoSpacesRule,
	}.Apply(text)
}

func (c *Cleaner) replaceEscapedNewlines(text string) string {
	return rule.Rules{
		escapedNewLineRule,
		escapedCarriageReturnRule,
		typoEscapedNewLineRule,
		typoEscapedCarriageReturnRule,
	}.Apply(text)
}

func (c *Cleaner) replacePunctuationInBrackets(text string) string {
	return bracketsContentRegex.ReplaceAllStringFunc(text, func(match string) string {
		if strings.Contains(match, "?") {
			return strings.ReplaceAll(match, "?", "&ᓷ&")
		}
		return match
	})
}

func (c *Cleaner) cleanQuotations(text string) string {
	text = strings.ReplaceAll(text, "`", "'")
	return rule.Rules{quotationsFirstRule, quotationsSecondRule}.Apply(text)
}

func (c *Cleaner) cleanTableOfContents(text string) string {
	return rule.Rules{
		tableOfContentsRule,
		consecutivePeriodsRule,
		consecutiveForwardSlashRule,
	}.Apply(text)
}

func (c *Cleaner) spaceSentences(text string) string {
	words := strings.Split(text, " ")

	// Transform each token independently so spacing cannot alter a URL or email elsewhere.
	for i, word := range words {
		words[i] = c.spaceWord(word)
	}
	return strings.Join(words, " ")
}

func (c *Cleaner) spaceWord(word string) string {
	for _, k := range urlEmailKeywords {
		if strings.Contains(word, k) {
			return word
		}
	}

	return rule.Rules{noSpaceBetweenSentencesRule, noSpaceBetweenSentencesDigitRule}.Apply(word)
}

func (c *Cleaner) cleanConsecutiveCharacters(text string) string {
	return rule.Rules{consecutivePeriodsRule, consecutiveForwardSlashRule}.Apply(text)
}

var (
	// \n(?=[a-zA-Z]{1,2}\n) — full match used by removeNewlineInMiddleOfWord
	newLineInMiddleOfWordRegex = regexp.MustCompile(`\n[a-zA-Z]{1,2}\n`)

	doubleNewLineWithSpaceRule = rule.NewRule(regexp.MustCompile(`\n \n`), "\r")
	doubleNewLineRule          = rule.NewRule(regexp.MustCompile(`\n\n`), "\r")

	// \n(?=\.(\s|\n))
	newLineFollowedByPeriodRule          = rule.NewRule(regexp.MustCompile(`\n(\.(?:\s|\n))`), `$1`)
	replaceNewlineWithCarriageReturnRule = rule.NewRule(regexp.MustCompile(`\n`), "\r")

	escapedNewLineRule             = rule.NewRule(regexp.MustCompile(`\\n`), "\n")
	escapedCarriageReturnRule      = rule.NewRule(regexp.MustCompile(`\\r`), "\r")
	typoEscapedNewLineRule         = rule.NewRule(regexp.MustCompile(`\\\ n`), "\n")
	typoEscapedCarriageReturnRule  = rule.NewRule(regexp.MustCompile(`\\\ r`), "\r")

	inlineFormattingRule = rule.NewRule(regexp.MustCompile(`\{b\^&gt;\d*&lt;b\^\}|\{b\^>\d*<b\^\}`), "")

	tableOfContentsRule         = rule.NewRule(regexp.MustCompile(`\.{4,}\s*\d+-*\d*`), "\r")
	consecutivePeriodsRule      = rule.NewRule(regexp.MustCompile(`\.{5,}`), " ")
	consecutiveForwardSlashRule = rule.NewRule(regexp.MustCompile(`/{3}`), "")

	// (?<=[a-z])\.(?=[A-Z])
	noSpaceBetweenSentencesRegex = regexp.MustCompile(`([a-z])\.([A-Z])`)
	noSpaceBetweenSentencesRule  = rule.NewRule(noSpaceBetweenSentencesRegex, `$1. $2`)

	// (?<=\d)\.(?=[A-Z])
	noSpaceBetweenSentencesDigitRegex = regexp.MustCompile(`(\d)\.([A-Z])`)
	noSpaceBetweenSentencesDigitRule  = rule.NewRule(noSpaceBetweenSentencesDigitRegex, `$1. $2`)

	urlEmailKeywords = []string{"@", "http", ".com", "net", "www", "//"}

	// (?<=\s)\n(?=([a-z]|\())
	newLineInMiddleOfSentenceRule = rule.NewRule(regexp.MustCompile(`(\s)\n([a-z(])`), `$1$2`)

	// \n(?=•')
	newLineFollowedByBulletRule = rule.NewRule(regexp.MustCompile(`\n(•')`), "\r$1")

	quotationsFirstRule  = rule.NewRule(regexp.MustCompile(`''`), `"`)
	quotationsSecondRule = rule.NewRule(regexp.MustCompile("``"), `"`)

	bracketsContentRegex = regexp.MustCompile(`\[(?:[^\]])*\]`)

	htmlTagRule = rule.NewRule(
		regexp.MustCompile(`</?\w+((\s+\w+(\s*=\s*(".*?"|'.*?'|[^'">\s]+))?)+\s*|\s*)/?>`),
		"",
	)
	escapedHTMLTagRule = rule.NewRule(regexp.MustCompile(`&lt;/?[^gt;]*gt;`), "")
	htmlRules          = rule.Rules{htmlTagRule, escapedHTMLTagRule}

	// PDF (?<=[^\n]\s)\n(?=\S)
	pdfNewLineInMiddleOfSentenceRule = rule.NewRule(regexp.MustCompile(`([^\n]\s)\n(\S)`), `$1$2`)
	// PDF \n(?=[a-z])
	pdfNewLineInMiddleOfSentenceNoSpacesRule = rule.NewRule(regexp.MustCompile(`\n([a-z])`), ` $1`)
)
