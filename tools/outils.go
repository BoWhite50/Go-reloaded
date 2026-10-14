package tools

import (
	"os"
	"strconv"
	"strings"
)

// readsample lit le fichier d'entrée et renvoie la liste de ses lignes.
func Readsample(nomFichier string) ([]string, error) {
	contenuInit, err := os.ReadFile(nomFichier)
	if err != nil {
		return nil, err
	}
	return strings.Split(string(contenuInit), "\n"), nil
}

// HexOrBinToDecimal transforme un hex ou bin en int puis en string
func HexOrBinToDecimal(mot string, base int) (string, error) {
	nombre, err := strconv.ParseInt(mot, base, 64)
	if err != nil {
		return mot, err
	}
	mot = strconv.FormatInt(nombre, 10)
	return mot, nil
}

// transforme un mot selon le marqueur
func Transform(marqueur string, mot string) string {
	switch marqueur {
	case "(up)":
		return strings.ToUpper(mot)
	case "(low)":
		return strings.ToLower(mot)
	case "(cap)":
		if len(mot) > 0 {
			capitalized := strings.ToUpper(mot[:1]) + strings.ToLower(mot[1:])
			return capitalized
		}
	}
	return mot
}
// OccurrenceDeA remplace "A"/"a" par "An"/"an" si le mot suivant commence par une voyelle ou un h
func OccurrenceDeA(lignes string) string {
	mots := strings.Fields(lignes)
	for l := 0; l < len(mots); l++ {
		if mots[l] == "A" {
			if l+1 < len(mots) && strings.ContainsRune("aeiouhAEIOUH", rune(mots[l+1][0])) {
				mots[l] = "An"
			}
		}
		if mots[l] == "a" {
			if l+1 < len(mots) && strings.ContainsRune("aeiouhAEIOUH", rune(mots[l+1][0])) {
				mots[l] = "an"
			}
		}
	}
	return strings.Join(mots, " ")
}

// ponctuation corrige les ponctuation du text
func Ponctuation(ligne string) string {
	mots := strings.Fields(ligne)

	for i := 1; i < len(mots); i++ {
		reste := strings.TrimLeft(mots[i], ".,!?:;")
		signe := ""
		if len(mots[i]) > len(reste) {
			signe = mots[i][:len(mots[i])-len(reste)]
		}
		if signe == "" {
			continue
		}
		mots[i-1] += signe
		if reste == "" {
			mots = append(mots[:i], mots[i+1:]...)
			i--
		} else {
			mots[i] = reste
		}
	}
	return strings.Join(mots, " ")
}

// Guillemets corrige les espace entre les apostrophes et guillemets
func Guillemets(ligne string) string {
	mots := strings.Fields(ligne)

	for _, signe := range []string{"'", `"`} {
		nb := 0
		for i := 0; i < len(mots); i++ {
			if mots[i] != signe {
				continue
			}

			if nb%2 == 0 && i+1 < len(mots) {
				mots[i+1] = signe + mots[i+1]
				mots = append(mots[:i], mots[i+1:]...)
				nb++
			} else if nb%2 == 1 && i > 0 {
				mots[i-1] += signe
				mots = append(mots[:i], mots[i+1:]...)
				i--
				nb++
			}

		}
	}

	return strings.Join(mots, " ")
}
