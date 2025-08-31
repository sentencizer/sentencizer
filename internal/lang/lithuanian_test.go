package lang_test

import (
	"testing"

	"github.com/sentencizer/sentencizer"
)

func Test_Lithuanian(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "1) Direct speech ending with exclamation mark",
			args: args{
				text: "– Labas! – švelniai tarė.",
			},
			want: []string{"– Labas! – švelniai tarė."},
		},
		{
			name: "2) Direct speech ending with question mark",
			args: args{
				text: "– Kaip laikaisi? – žvaliai paklausė.",
			},
			want: []string{"– Kaip laikaisi? – žvaliai paklausė."},
		},
		{
			name: "3) Sentence ending with exclamation mark with two dots",
			args: args{
				text: "– A tu, žabali, ar nenustosi!.. Savo žmogaus nemato, – girdi jisai pažįstamą balsą.",
			},
			want: []string{
				"– A tu, žabali, ar nenustosi!..",
				"Savo žmogaus nemato, – girdi jisai pažįstamą balsą.",
			},
		},
		{
			name: "4) Sentence ending with question mark and two dots",
			args: args{
				text: "Manot – neįmanoma?! Pasakyti reikės, tačiau ką aš pasakysiu?.. Ar ji pati nesuvokia, kokia yra juokinga?..",
			},
			want: []string{
				"Manot – neįmanoma?!",
				"Pasakyti reikės, tačiau ką aš pasakysiu?..",
				"Ar ji pati nesuvokia, kokia yra juokinga?..",
			},
		},
		{
			name: "5) Sentence ending with ellipsis",
			args: args{
				text: "Juk jisai nekaltas… Brisius nori pasigerinti, suvizginti uodegą, bet iš baimės tupiasi ant paskutinių kojų, ir per jo snukį rieda gailios karčios ašaros.",
			},
			want: []string{
				"Juk jisai nekaltas…",
				"Brisius nori pasigerinti, suvizginti uodegą, bet iš baimės tupiasi ant paskutinių kojų, ir per jo snukį rieda gailios karčios ašaros.",
			},
		},
		{
			name: "6) Upper case after abbreviations",
			args: args{
				text: "Pasak prof. Petrausko, šilauoges valgyti labai sveika.",
			},
			want: []string{"Pasak prof. Petrausko, šilauoges valgyti labai sveika."},
		},
		{
			name: "7) Direct speech in double quotes",
			args: args{text: "„Argi mudu čia ne vienu du pasaulyje?“ – mąstė jis, drąsindamas pats save."},
			want: []string{"„Argi mudu čia ne vienu du pasaulyje?“ – mąstė jis, drąsindamas pats save."},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sg := sentencizer.NewSegmenter("lt")
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
