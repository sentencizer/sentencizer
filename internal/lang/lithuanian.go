package lang

import (
	"regexp"

	"github.com/sentencizer/sentencizer/internal/processor"
	"github.com/sentencizer/sentencizer/internal/replacer"
	"github.com/sentencizer/sentencizer/internal/rule"
)

type DirectSpeechRules struct {
	All rule.Rules
}

var (
	questionExclamationMarkDirectSpeechRule = rule.NewRule(
		regexp.MustCompile(`\?!(\s–\s[a-ž])`), `&ᓷ&&ᓴ&$1`,
	)
	exclamationMarkDirectSpeechRule = rule.NewRule(
		regexp.MustCompile(`!(\s–\s[a-ž])`), `&ᓴ&$1`,
	)
	questionMarkDirectSpeechRule = rule.NewRule(
		regexp.MustCompile(`\?(\s–\s[a-ž])`), `&ᓷ&$1`,
	)
	exclamationMarkTwoRule = rule.NewRule(
		regexp.MustCompile(`!\.\.`), `&ᓴ&∯.`,
	)
	questionMarkTwoRule = rule.NewRule(
		regexp.MustCompile(`\?\.\.`), `&ᓷ&∯.`,
	)
	// https://rubular.com/r/zrmWh8sRQ4R8Ay
	horizontalEllipsisRule = rule.NewRule(
		regexp.MustCompile(`…(\s+[A-Ž])`), `☍☍.$1`,
	)
	reinsertHorizontalEllipsis = rule.NewRule(
		regexp.MustCompile(`☍☍.`), "…",
	)

	betweenLithuanianDoubleQuotesRegex = regexp.MustCompile(`„([^“\\]+|\\{2}|\\.)*“`)

	directSpeechRules = DirectSpeechRules{
		All: rule.Rules{
			questionExclamationMarkDirectSpeechRule,
			exclamationMarkDirectSpeechRule,
			questionMarkDirectSpeechRule,
		},
	}
)

type betweenPunctuationReplacerLithuanian struct {
	punctuationReplacer replacer.Punctuation
	betweenPunctuation  replacer.BetweenPunctuation
}

func (b *betweenPunctuationReplacerLithuanian) subPunctuationBetweenLithuanianDoubleQuotes(text string) string {
	return betweenLithuanianDoubleQuotesRegex.ReplaceAllStringFunc(
		text,
		b.punctuationReplacer.ReplaceFunc(processor.PunctuationMatchTypeNone),
	)
}

func (b *betweenPunctuationReplacerLithuanian) subPunctuationInDirectSpeech(text string) string {
	return directSpeechRules.All.Apply(text)
}

func (b *betweenPunctuationReplacerLithuanian) Replace(text string) string {
	text = b.betweenPunctuation.Replace(text)
	text = b.subPunctuationBetweenLithuanianDoubleQuotes(text)
	text = b.subPunctuationInDirectSpeech(text)
	return text
}

func newLithuanian() *processor.Config {
	cfg := processor.Standard()
	punctuationReplacer := replacer.NewPunctuationReplacer()
	cfg.BetweenPunctuationReplacer = &betweenPunctuationReplacerLithuanian{
		punctuationReplacer: punctuationReplacer,
		betweenPunctuation:  replacer.NewBetweenPunctuation(punctuationReplacer),
	}
	// Based on https://www.vlkk.lt/aktualiausios-temos/rasyba/sutrumpinimai
	cfg.Abbreviation.Abbreviations = []string{
		"a. k",
		"a. s",
		"b. k",
		"el. p",
		"e. p",
		"ir pan",
		"ir t. t",
		"J. E",  // Jo Ekscelencija
		"J. Em", // Jo Eminencija
		"k. a",  // kaip antai
		"l. e. p",
		"m. e", // mūsų eros
		"m. m", // mokslo metai
		"p. d",
		"p. m. e", // prieš mūsų erą
		"pr. Kr",
		"pr. m. e",
		"š. m",
		"t. p",
		"t. y",
		"ž. ū",
		"a",
		"adm",
		"adv",
		"agr",
		"akad",
		"aklg",
		"akt",
		"al",
		"angl",
		"aps",
		"apsk",
		"apyg",
		"apyl",
		"art",
		"asist",
		"asmv",
		"atsak",
		"aut",
		"avd", // asmenvardis
		"biol",
		"bkl",
		"bot",
		"bt",
		"buh",
		"buv",
		"chem",
		"d", // duktė; diena
		"dail",
		"dek",
		"dėst",
		"dir",
		"dirig",
		"doc",
		"dr",
		"drg",
		"drp", // durpynas
		"dš",
		"egz",
		"eil",
		"ekon",
		"e", // elektroninis
		"el",
		"etc",
		"ež",
		"fak",
		"faks",
		"filol",
		"filos",
		"g",
		"gen",
		"geol",
		"gerb",
		"gim", // gyvenvietė
		"gv",
		"gyd",
		"įl",
		"insp",
		"inž",
		"istor",
		"k", // kaimas
		"kand",
		"kat",
		"kl",
		"kln",
		"kn",
		"koresp",
		"kpt",
		"kr",
		"kt",
		"kun",
		"kyš",
		"lent",
		"ltn",
		"m",
		"mat",
		"med",
		"mėn",
		"mgr",
		"mjr",
		"mln",
		"mlrd",
		"mok",
		"mokyt",
		"mot",
		"mst",
		"mstl",
		"nkt", // nekaitomas
		"ntk",
		"nr",
		"p", // ponas, ponia, panelė; puslapis; punktas
		"pav",
		"pavad",
		"pėst",
		"pil",
		"pirm",
		"pl",
		"plg",
		"plk",
		"pr",
		"pranc",
		"prof",
		"prok",
		"prot", // protokolas
		"psl",
		"pss", // pusiasalis
		"pšt",
		"pvz",
		"r", // rajonas
		"raj",
		"red",
		"rež",
		"rš",
		"s", // sūnus, sąskaita
		"sąs",
		"sąsk",
		"sav",
		"saviv",
		"sekr",
		"sen",
		"serž",
		"sk",
		"skg",
		"skv",
		"skyr",
		"šnek",
		"šv",
		"sp",
		"spec",
		"sr", // sritis
		"st",
		"str",
		"stud",
		"t", // tomas
		"techn",
		"tel",
		"teol",
		"tir",
		"tūkst",
		"tšk",
		"up",
		"upl",
		"vad",
		"ved",
		"vet",
		"virš",
		"vlsč",
		"vnt",
		"vs",
		"vtv",
		"vv",
		"vyr",
		"vyriaus",
		"vyresn",
		"žml",
		"zool",
		"žr",
	}
	cfg.Abbreviation.PrePositiveAbbreviations = []string{
		"adm",
		"adv",
		"agr",
		"akad",
		"akt",
		"art",
		"asist",
		"aut",
		"buh",
		"dail",
		"dr",
		"dek",
		"dėst",
		"dir",
		"dirig",
		"doc",
		"drg",
		"d",
		"gen",
		"gyd",
		"insp",
		"inž",
		"kand",
		"kpt",
		"koresp",
		"ltn",
		"mjr",
		"mok",
		"mokyt",
		"pav",
		"pavad",
		"pėst",
		"pil",
		"pirm",
		"p",
		"prof",
		"prok",
		"plk",
		"red",
		"rež",
		"sekr",
		"serž",
		"stud",
		"s",
		"ved",
	}
	cfg.Abbreviation.NumberAbbreviations = []string{"Nr", "nr"}
	cfg.SentenceStarters = nil
	cfg.Ellipsis.All = append(
		cfg.Ellipsis.All,
		exclamationMarkTwoRule,
		questionMarkTwoRule,
		horizontalEllipsisRule,
	)
	cfg.ReinsertEllipsisRules.All = append(
		cfg.ReinsertEllipsisRules.All,
		reinsertHorizontalEllipsis,
	)
	return cfg
}
