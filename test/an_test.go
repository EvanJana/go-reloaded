package test


import (
	"reflect"
	"testing"
	 conv "github.com/EvanJana/go-reloaded/conversion"
)

func TestAn(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  []string
	}{
		{
			name:  "voyelle a",
			input: []string{"(a)", "apple"},
			want:  []string{"an", "apple"},
		},
		{
			name:  "voyelle e",
			input: []string{"(a)", "elephant"},
			want:  []string{"an", "elephant"},
		},
		{
			name:  "voyelle i",
			input: []string{"(a)", "ice"},
			want:  []string{"an", "ice"},
		},
		{
			name:  "voyelle o",
			input: []string{"(a)", "orange"},
			want:  []string{"an", "orange"},
		},
		{
			name:  "voyelle u",
			input: []string{"(a)", "umbrella"},
			want:  []string{"an", "umbrella"},
		},
		{
			name:  "consonne",
			input: []string{"(a)", "book"},
			want:  []string{"(a)", "book"},
		},
		{
			name:  "sans (a)",
			input: []string{"hello", "world"},
			want:  []string{"hello", "world"},
		},
		{
			name:  "aucun mot après (a)",
			input: []string{"(a)"},
			want:  []string{"(a)"},
		},
		{
			name:  "mot suivant vide",
			input: []string{"(a)", ""},
			want:  []string{"(a)", ""},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := conv.An(tt.input)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("An(%v) = %v, attendu %v",
					tt.input, got, tt.want)
			}
		})
	}
}