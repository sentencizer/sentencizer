package lang_test

import (
	"testing"

	"github.com/sentencizer/sentencizer"
)

func Test_Hebrew(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "BasicSplit",
			args: args{
				text: "שלום עולם. זה משפט אחר.",
			},
			want: []string{
				"שלום עולם.",
				"זה משפט אחר.",
			},
		},
		{
			name: "QuestionSplit",
			args: args{
				text: "מה שלומך? הכל בסדר.",
			},
			want: []string{
				"מה שלומך?",
				"הכל בסדר.",
			},
		},
		{
			name: "ExclamationSplit",
			args: args{
				text: "זה נהדר! איך לא חשבתי על זה קודם?",
			},
			want: []string{
				"זה נהדר!",
				"איך לא חשבתי על זה קודם?",
			},
		},
		{
			name: "MultiplePunctuation",
			args: args{
				text: "מה?! זה לא ייאמן?!",
			},
			want: []string{
				"מה?!",
				"זה לא ייאמן?!",
			},
		},
		{
			name: "QuotesWithQuestions",
			args: args{
				text: "\"מה שלומך?\" היא שאלה. \"יפה!\" הוא ענה.",
			},
			want: []string{
				"\"מה שלומך?\"",
				"היא שאלה.",
				"\"יפה!\"",
				"הוא ענה.",
			},
		},
		{
			name: "ParenthesesMidSentence",
			args: args{
				text: "משה (אבי של יעקב) אמר שלום. אני זוכר.",
			},
			want: []string{
				"משה (אבי של יעקב) אמר שלום.",
				"אני זוכר.",
			},
		},
		{
			name: "DoctorAbbrev",
			args: args{
				text: "ד\"ר לוי מונה למנהל. המרצים בירכו.",
			},
			want: []string{
				"ד\"ר לוי מונה למנהל.",
				"המרצים בירכו.",
			},
		},
		{
			name: "MrsAbbrev",
			args: args{
				text: "גב' לוי אמרה שלום. הוא חייך.",
			},
			want: []string{
				"גב' לוי אמרה שלום.",
				"הוא חייך.",
			},
		},
		{
			name: "DecimalNumber",
			args: args{
				text: "מטוס טס במהירות 3.14 ק\"מ לשעה. אנחנו שמחים.",
			},
			want: []string{
				"מטוס טס במהירות 3.14 ק\"מ לשעה.",
				"אנחנו שמחים.",
			},
		},
		{
			name: "DateNoSplit",
			args: args{
				text: "הסכם ב-16.12.2022 נחתם.",
			},
			want: []string{
				"הסכם ב-16.12.2022 נחתם.",
			},
		},
		{
			name: "EllipsisBreak",
			args: args{
				text: "היא חשבה... אבל המחשבה התפוגגה.",
			},
			want: []string{
				"היא חשבה...",
				"אבל המחשבה התפוגגה.",
			},
		},
		{
			name: "EllipsisExclaim",
			args: args{
				text: "וואו... היא אמרה!",
			},
			want: []string{
				"וואו...",
				"היא אמרה!",
			},
		},
		{
			name: "ColonNoSplit",
			args: args{
				text: "ישראל: מדינה קטנה, אבל חזקה.",
			},
			want: []string{
				"ישראל: מדינה קטנה, אבל חזקה.",
			},
		},
		{
			name: "ListNoSplit",
			args: args{
				text: "הוא קנה תפוח, בננה, אגס; בשל מחיר נמוך.",
			},
			want: []string{
				"הוא קנה תפוח, בננה, אגס; בשל מחיר נמוך.",
			},
		},
		{
			name: "SingleExclaim",
			args: args{
				text: "וואו!",
			},
			want: []string{
				"וואו!",
			},
		},
		{
			name: "SingleQuestion",
			args: args{
				text: "אה?",
			},
			want: []string{
				"אה?",
			},
		},
		{
			name: "QuoteStatement",
			args: args{
				text: "הוא אמר: 'שלום'.",
			},
			want: []string{
				"הוא אמר: 'שלום'.",
			},
		},
		{
			name: "LowerCaseStart",
			args: args{
				text: "שלום. זאת דרך ארוכה.",
			},
			want: []string{
				"שלום.",
				"זאת דרך ארוכה.",
			},
		},
		{
			name: "DashParentheses",
			args: args{
				text: "הוא לא היה - בסופו של דבר - בטוח.",
			},
			want: []string{
				"הוא לא היה - בסופו של דבר - בטוח.",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sg := sentencizer.NewSegmenter("he")
			got := sg.Segment(tt.args.text)
			for i, v := range got {
				if i >= len(tt.want) {
					break
				}
				if v != tt.want[i] {
					t.Errorf("Segmenter.Segment() = %#v, want %#v", v, tt.want[i])
				}
			}
			if len(got) != len(tt.want) {
				t.Errorf("Segmenter.Segment() = %#v, want %#v", got, tt.want)
			}
		})
	}
}
