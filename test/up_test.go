package test

import (
	"reflect"
	"testing"

	conv "github.com/EvanJana/go-reloaded/conversion"
)

func TestUpper(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  []string
	}{
		{
			name:  "majuscules avec up",
			input: []string{"hello", "(up)"},
			want:  []string{"HELLO"},
		},
		{
			name:  "mot déjà en majuscules",
			input: []string{"HELLO", "(up)"},
			want:  []string{"HELLO"},
		},
		{
			name:  "majuscules avec up et deux mots",
			input: []string{"hello", "world", "(up,2)"},
			want:  []string{"HELLO", "WORLD"},
		},
		{
			name:  "majuscules sur les trois derniers mots",
			input: []string{"bonjour", "tout", "le", "monde", "(up,3)"},
			want:  []string{"bonjour", "TOUT", "LE", "MONDE"},
		},
		{
			name:  "sans marqueur",
			input: []string{"hello", "world"},
			want:  []string{"hello", "world"},
		},
		{
			name:  "marqueur sans mot précédent",
			input: []string{"(up)"},
			want:  []string{"(up)"},
		},
		{
			name:  "nombre invalide dans le marqueur",
			input: []string{"hello", "(up,abc)"},
			want:  []string{"hello", "(up,abc)"},
		},
		{
			name:  "nombre supérieur aux mots disponibles",
			input: []string{"hello", "world", "(up,5)"},
			want:  []string{"HELLO", "WORLD"},
		},
		{
			name:  "liste vide",
			input: []string{},
			want:  []string{},
		},
		{
			name:  "marqueur up avec zéro",
			input: []string{"hello", "(up,0)"},
			want:  []string{"hello"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := conv.Upper(tt.input)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf(
					"Upper(%v) = %v, attendu %v",
					tt.input, got, tt.want,
				)
			}
		})
	}
}
