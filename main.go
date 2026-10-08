package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("le programme doit être lancé avec 2 fichier: le fichier d'entrée et le fichier de sortie")
		return
	}

	lignes, err := readsample(os.Args[1])
	if err != nil {
		fmt.Println("Erreur lors de la lecture du fichier d'entrée:", err)
		return
	}
	Marqueurs(lignes)

	err = os.WriteFile(os.Args[2], []byte(strings.Join(lignes, "\n")), 0644)
	if err != nil {
		fmt.Println("Erreur lors de l'écriture du fichier de sortie:", err)
	}
}

// readsample lit le fichier d'entrée et renvoie la liste de ses lignes.
func readsample(nomFichier string) ([]string, error) {
	contenuInit, err := os.ReadFile(nomFichier)
	if err != nil {
		return nil, err
	}
	return strings.Split(string(contenuInit), "\n"), nil
}



func Marqueurs(lignes []string) []string {

	for i := 0; i < len(lignes); i++ {
		mots := strings.Fields(lignes[i])
		for j := 1; j < len(mots); j++ {
			if mots[j] == "(hex)"{
				err := error(nil)
				mots[j-1], err = hexToDecimal(mots[j-1])
				if err != nil {
					fmt.Println("Erreur lors de la conversion hexadécimale:", err)
					return nil
				}
				mots = append(mots[:j], mots[j+1:]...)
				j--
				lignes[i] = strings.Join(mots, " ")
			}
			if mots[j] == "(bin)"{	
				// pareil que (hex), avec la base 2
			}
		// étapes 5 et 6 : ajouter (up), (low), (cap) et (up, 2)...
		}
	}
	return lignes
}


// transforme un hex en int puis en string
func hexToDecimal(mot string) (string, error) {
	nombre, err := strconv.ParseInt(mot, 16, 64)
	if err != nil {
		return "", err
	}
	mot = strconv.FormatInt(nombre, 10)
	return mot, nil
}