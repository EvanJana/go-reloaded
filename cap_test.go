package main

import (
	"reflect"
	"testing"

	conv "github.com/EvanJana/go-reloaded/conversion"
)

func TestCapitalize(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  []string
	}{
		{
			name:  "majuscule avec cap",
			input: []string{"hello", "(cap)"},
			want:  []string{"Hello"},
		},
		{
			name:  "première lettre en majuscule et reste en minuscule",
			input: []string{"hELLO", "(cap)"},
			want:  []string{"Hello"},
		},
		{
			name:  "capitaliser plusieurs mots",
			input: []string{"hello", "WORLD", "(cap,2)"},
			want:  []string{"Hello", "World"},
		},
		{
			name:  "capitaliser les trois derniers mots",
			input: []string{"bonjour", "TOUT", "LE", "MONDE", "(cap,3)"},
			want:  []string{"bonjour", "Tout", "Le", "Monde"},
		},
		{
			name:  "sans marqueur",
			input: []string{"hello", "WORLD"},
			want:  []string{"hello", "WORLD"},
		},
		{
			name:  "marqueur sans mot précédent",
			input: []string{"(cap)"},
			want:  []string{"(cap)"},
		},
		{
			name:  "nombre invalide dans le marqueur",
			input: []string{"hello", "(cap,abc)"},
			want:  []string{"hello", "(cap,abc)"},
		},
		{
			name:  "nombre supérieur aux mots disponibles",
			input: []string{"hello", "WORLD", "(cap,5)"},
			want:  []string{"Hello", "World"},
		},
		{
			name:  "mot vide",
			input: []string{"", "(cap)"},
			want:  []string{""},
		},
		{
			name:  "liste vide",
			input: []string{},
			want:  []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := conv.Capitalize(tt.input)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf(
					"Capitalize(%v) = %v, attendu %v",
					tt.input, got, tt.want,
				)
			}
		})
	}
}
