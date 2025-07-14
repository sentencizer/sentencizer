package lang

import (
	"github.com/sentencizer/sentencizer/internal/rule"
	"github.com/sentencizer/sentencizer/internal/processor"
)

func newHebrew() *processor.Config {
	cfg := processor.Standard()
	cfg.Abbreviation.PrePositiveAbbreviations = nil
	cfg.Abbreviation.NumberAbbreviations = nil
	cfg.SentenceStarters = nil
	cfg.QuotationAtEndOfSentenceRegex = processor.QuotationAtEndOfSentenceHebrewRegex
	cfg.SplitSpaceQuotationAtEndOfSentenceRule = processor.SplitSpaceHebrewAtEndOfSentenceRule
	cfg.Ellipsis = processor.EllipsisHebrewRule
	cfg.SentenceBoundaryRules = processor.SentenceBoundaryRules{
		All: rule.Rules{
			processor.SentenceBoundaryRuleHebrew1,
		},
	}
	return cfg
}
