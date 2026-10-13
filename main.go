package main

import (
	"fmt"
	"go-reloaded/tools"
	"os"
	"strconv"
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
		if len(mots) > 0 {
			if mots[0] == "(hex)" || mots[0] == "(bin)" || mots[0] == "(up)" || mots[0] == "(low)" || mots[0] == "(cap)" || mots[0] == "(up," || mots[0] == "(low," || mots[0] == "(cap," {
				if len(mots) >= 2 && (mots[0] == "(cap," || mots[0] == "(low," || mots[0] == "(up,") {
					mots = append(mots[:0], mots[2:]...)
					lignes[i] = strings.Join(mots, " ")
				} else {
					mots = append(mots[:0], mots[1:]...)
					lignes[i] = strings.Join(mots, " ")
				}
			}
		}
		for j := 1; j < len(mots); j++ {
			if mots[j] == "(hex)" || mots[j] == "(bin)" {
				base := 16
				if mots[j] == "(bin)" {
					base = 2
				}
				mots[j-1], err = tools.HexOrBinToDecimal(mots[j-1], base)
				if err != nil {
					fmt.Println("Erreur lors de la conversion hexadécimale ou binaire:", err)
				}
				mots = append(mots[:j], mots[j+1:]...)
				j--
				lignes[i] = strings.Join(mots, " ")
			}
			if mots[j] == "(up)" || mots[j] == "(low)" || mots[j] == "(cap)" {
				mots[j-1] = tools.Transform(mots[j], mots[j-1])
				mots = append(mots[:j], mots[j+1:]...)
				j--
				lignes[i] = strings.Join(mots, " ")
			}
			if mots[j] == "(up," || mots[j] == "(low," || mots[j] == "(cap," {
				if j+1 >= len(mots) {
					continue
				}
				nbMots, err := strconv.Atoi(strings.TrimSuffix(mots[j+1], ")"))
				if err != nil {
					fmt.Println("Erreur : nombre invalide dans le marqueur:", err)
					continue
				}
				marqueur := strings.Replace(mots[j], ",", ")", 1)

				for k := j - 1; k >= 0 && k >= j-nbMots; k-- {
					mots[k] = tools.Transform(marqueur, mots[k])
				}
				if len(mots) >= 2 {
					mots = append(mots[:j], mots[j+2:]...)
				}
				j--
				lignes[i] = strings.Join(mots, " ")
			}
		}
	}
	return lignes
}
