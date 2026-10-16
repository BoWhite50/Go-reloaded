package tools

import "testing"

func TestHexOrBinToDecimal(t *testing.T) {
	tests := []struct {
		mot     string
		base    int
		attendu string
	}{
		{"1E", 16, "30"},
		{"ff", 16, "255"},
		{"10", 2, "2"},
		{"hello", 16, "hello"},
	}

	for _, test := range tests {
		resultat, _ := HexOrBinToDecimal(test.mot, test.base)
		if resultat != test.attendu {
			t.Errorf("HexOrBinToDecimal(%q, %d) = %q, attendu %q", test.mot, test.base, resultat, test.attendu)
		}
	}
}

func TestTransform(t *testing.T) {
	tests := []struct {
		marqueur string
		mot      string
		attendu  string
	}{
		{"(up)", "go", "GO"},
		{"(low)", "SHOUTING", "shouting"},
		{"(cap)", "bridge", "Bridge"},
		{"(cap)", "PARIS", "Paris"},
	}

	for _, test := range tests {
		resultat := Transform(test.marqueur, test.mot)
		if resultat != test.attendu {
			t.Errorf("Transform(%q, %q) = %q, attendu %q", test.marqueur, test.mot, resultat, test.attendu)
		}
	}
}

func TestMarqueurs(t *testing.T) {
	tests := []struct {
		entree  string
		attendu string
	}{
		{"1E (hex) files were added", "30 files were added"},
		{"It has been 10 (bin) years", "It has been 2 years"},
		{"Welcome to the Brooklyn bridge (cap)", "Welcome to the Brooklyn Bridge"},
		{"This is so exciting (up, 2)", "This is SO EXCITING"},
		{"one two three (up, 10)", "ONE TWO THREE"},
		{"(up) hello", "hello"},
		{"(up) (low) hello", "hello"},
	}

	for _, test := range tests {
		resultat := Marqueurs([]string{test.entree})[0]
		if resultat != test.attendu {
			t.Errorf("Marqueurs(%q) = %q, attendu %q", test.entree, resultat, test.attendu)
		}
	}
}

func TestOccurrenceDeA(t *testing.T) {
	tests := []struct {
		entree  string
		attendu string
	}{
		{"A amazing rock", "An amazing rock"},
		{"a hour", "an hour"},
		{"a book", "a book"},
		{"a egg and a apple", "an egg and an apple"},
		{"I want a", "I want a"},
		{"a 'amazing'", "an 'amazing'"},
		{"a ' amazing", "an ' amazing"},
	}

	for _, test := range tests {
		resultat := OccurrenceDeA(test.entree)
		if resultat != test.attendu {
			t.Errorf("OccurrenceDeA(%q) = %q, attendu %q", test.entree, resultat, test.attendu)
		}
	}
}

func TestPonctuation(t *testing.T) {
	tests := []struct {
		entree  string
		attendu string
	}{
		{"there ,and then BAMM !!", "there, and then BAMM!!"},
		{"I was thinking ... You", "I was thinking... You"},
		{"Really !? yes", "Really!? yes"},
		{"'hello .'", "'hello.'"},
	}

	for _, test := range tests {
		resultat := Ponctuation(test.entree)
		if resultat != test.attendu {
			t.Errorf("Ponctuation(%q) = %q, attendu %q", test.entree, resultat, test.attendu)
		}
	}
}

func TestGuillemets(t *testing.T) {
	tests := []struct {
		entree  string
		attendu string
	}{
		{"me: ' awesome '", "me: 'awesome'"},
		{"' I am the best '", "'I am the best'"},
		{`" hello "`, `"hello"`},
		{"I don't know", "I don't know"},
	}

	for _, test := range tests {
		resultat := Guillemets(test.entree)
		if resultat != test.attendu {
			t.Errorf("Guillemets(%q) = %q, attendu %q", test.entree, resultat, test.attendu)
		}
	}
}
