package test

import (
	"reflect"
	"testing"

	conv "github.com/EvanJana/go-reloaded/conversion"
)

func TestHexToDec(t *testing.T) {
	tests := []struct {
		name    string
		input   []string
		want    []string
		wantErr bool
	}{
		{
			name:    "conversion hexadécimale simple",
			input:   []string{"A", "(hex)"},
			want:    []string{"10"},
			wantErr: false,
		},
		{
			name:    "conversion hexadécimale avec lettres",
			input:   []string{"FF", "(hex)"},
			want:    []string{"255"},
			wantErr: false,
		},
		{
			name:    "conversion hexadécimale minuscule",
			input:   []string{"a", "(hex)"},
			want:    []string{"10"},
			wantErr: false,
		},
		{
			name:    "conversion de plusieurs nombres",
			input:   []string{"A", "(hex)", "1F", "(hex)"},
			want:    []string{"10", "31"},
			wantErr: false,
		},
		{
			name:    "nombre hexadécimal invalide",
			input:   []string{"G", "(hex)"},
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
			input:   []string{"(hex)"},
			want:    []string{"(hex)"},
			wantErr: false,
		},
		{
			name:    "nombre sans marqueur",
			input:   []string{"FF"},
			want:    []string{"FF"},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := conv.HexToDec(tt.input)

			if (err != nil) != tt.wantErr {
				t.Fatalf(
					"HexToDec(%v) erreur = %v, wantErr = %v",
					tt.input, err, tt.wantErr,
				)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf(
					"HexToDec(%v) = %v, attendu %v",
					tt.input, got, tt.want,
				)
			}
		})
	}
}
