package test

import (
	"reflect"
	"testing"

	conv "github.com/EvanJana/go-reloaded/conversion"
)

func TestLower(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  []string
	}{
		{
			name:  "minuscules avec low",
			input: []string{"HELLO", "(low)"},
			want:  []string{"hello"},
		},
		{
			name:  "mot déjà en minuscules",
			input: []string{"hello", "(low)"},
			want:  []string{"hello"},
		},
		{
			name:  "minuscules avec low et deux mots",
			input: []string{"HELLO", "WORLD", "(low,2)"},
			want:  []string{"hello", "world"},
		},
		{
			name:  "minuscules sur les trois derniers mots",
			input: []string{"BONJOUR", "TOUT", "LE", "MONDE", "(low,3)"},
			want:  []string{"BONJOUR", "tout", "le", "monde"},
		},
		{
			name:  "sans marqueur",
			input: []string{"HELLO", "WORLD"},
			want:  []string{"HELLO", "WORLD"},
		},
		{
			name:  "marqueur sans mot précédent",
			input: []string{"(low)"},
			want:  []string{"(low)"},
		},
		{
			name:  "nombre invalide dans le marqueur",
			input: []string{"HELLO", "(low,abc)"},
			want:  []string{"HELLO", "(low,abc)"},
		},
		{
			name:  "nombre supérieur aux mots disponibles",
			input: []string{"HELLO", "WORLD", "(low,5)"},
			want:  []string{"hello", "world"},
		},
		{
			name:  "liste vide",
			input: []string{},
			want:  []string{},
		},
		{
			name:  "marqueur low avec zéro",
			input: []string{"HELLO", "(low,0)"},
			want:  []string{"HELLO"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := conv.Lower(tt.input)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf(
					"Lower(%v) = %v, attendu %v",
					tt.input, got, tt.want,
				)
			}
		})
	}
}