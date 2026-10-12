package main

import (
	"fmt"
	"go-reloaded/tools"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("le programme doit être lancé avec 2 fichier: le fichier d'entrée et le fichier de sortie")
		return
	}

	lignes, err := tools.Readsample(os.Args[1])
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



func Marqueurs(lignes []string) []string {
	err := error(nil)
	for i := 0; i < len(lignes); i++ {
		mots := strings.Fields(lignes[i])
		for j := 1; j < len(mots); j++ {
			if mots[j] == "(hex)" || mots[j] == "(bin)" {
				base := 16
				if mots[j] == "(bin)" {
					base = 2
				}
				mots[j-1], err = tools.HexOrBinToDecimal(mots[j-1], base)
				if err != nil {
					fmt.Println("Erreur lors de la conversion hexadécimale ou binaire:", err)
					continue
				}
				mots = append(mots[:j], mots[j+1:]...)
				j--
				lignes[i] = strings.Join(mots, " ")
			}

			// étapes 5 et 6 : ajouter (up), (low), (cap) et (up, 2)...
		}
	}
	return lignes
}

