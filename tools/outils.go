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
