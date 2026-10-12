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

// HexToDecimal transforme un hex en int puis en string
func HexOrBinToDecimal(mot string, base int) (string, error) {
	nombre, err := strconv.ParseInt(mot, base, 64)
	if err != nil {
		return "", err
	}
	mot = strconv.FormatInt(nombre, 10)
	return mot, nil
}

func BinToDecimal(mot string) (string, error) {
	nombre, err := strconv.ParseInt(mot, 2, 64)
	if err != nil {
		return "", err
	}
	mot = strconv.FormatInt(nombre, 10)
	return mot, nil
}
