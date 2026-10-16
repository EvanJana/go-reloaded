package test

import (
	"reflect"
	"testing"
	 conv "github.com/EvanJana/go-reloaded/conversion"
)

func TestBinToDec(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  []string
		wantErr bool
	}{
		{
			name:    "conversion binaire simple",
			input:   []string{"101(bin)"},
			want:    []string{"101(bin)"},
			wantErr: false,
		},
		{
			name:    "conversion avec marqueur séparé",
			input:   []string{"101", "(bin)"},
			want:    []string{"5"},
			wantErr: false,
		},
		{
			name:    "conversion de plusieurs nombres",
			input:   []string{"101", "(bin)", "11", "(bin)"},
			want:    []string{"5", "3"},
			wantErr: false,
		},
		{
			name:    "nombre binaire invalide",
			input:   []string{"102", "(bin)"},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "liste vide",
			input:   []string{},
			want:    []string{},
			wantErr: false,
		},
		{
			name:    "marqueur sans nombre avant",
			input:   []string{"(bin)"},
			want:    []string{"(bin)"},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := conv.BinToDec(tt.input)

			if (err != nil) != tt.wantErr {
				t.Fatalf("BinToDec(%v) erreur = %v, wantErr = %v",
					tt.input, err, tt.wantErr)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BinToDec(%v) = %v, attendu %v",
					tt.input, got, tt.want)
			}
		})
	}
}