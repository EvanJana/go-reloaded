package test

import (
	"reflect"
	"testing"

	conv "github.com/EvanJana/go-reloaded/conversion"
)

func TestPonctuation(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  []string
	}{
		{
			name:  "ponctuation avec point",
			input: []string{"Hello", "world", " ."},
			want:  []string{"Hello", "world", "."},
		},
		{
			name:  "espace avant la virgule",
			input: []string{"Hello", "world", ","},
			want:  []string{"Hello", "world", ","},
		},
		{
			name:  "point d'exclamation",
			input: []string{"Hello", "world", "!"},
			want:  []string{"Hello", "world", "!"},
		},
		{
			name:  "point d'interrogation",
			input: []string{"How", "are", "you", "?"},
			want:  []string{"How", "are", "you", "?"},
		},
		{
			name:  "ponctuation suivie d'un mot",
			input: []string{"Hello", ",", "world", "!"},
			want:  []string{"Hello", ",", "world", "!"},
		},
		{
			name:  "plusieurs signes de ponctuation",
			input: []string{"Hello", "world", "!", "?", ":"},
			want:  []string{"Hello", "world", "!", "?", ":"},
		},
		{
			name:  "apostrophes simples",
			input: []string{"'", "hello", "world", "'"},
			want:  []string{"'hello", "world'"},
		},
		{
			name:  "guillemets doubles",
			input: []string{`"`, "hello", "world", `"`},
			want:  []string{`"hello`, `world"`},
		},
		{
			name:  "sans ponctuation",
			input: []string{"Hello", "world"},
			want:  []string{"Hello", "world"},
		},
		{
			name:  "liste vide",
			input: []string{},
			want:  []string{},
		},
		{
			name:  "ponctuation seule",
			input: []string{"."},
			want:  []string{"."},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := conv.Ponctuation(tt.input)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf(
					"Ponctuation(%v) = %v, attendu %v",
					tt.input, got, tt.want,
				)
			}
		})
	}
}
