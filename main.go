package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("le programme doit être lancé avec 2 fichier: le fichier d'entrée et le fichier de sortie")
		return
	}
	text, err := readsample()
	if err == nil {
		os.WriteFile(os.Args[2], []byte(strings.Join(text, "\n")), 0644)
	}

}
// lit le fichier d'entrée et resort une liste de mots avec les retour a la ligne
func readsample() ([]string, error) {
	ContenuInit, err := os.ReadFile(os.Args[1])
	if err != nil {
		return nil, err
	}
	ligne := strings.Split(string(ContenuInit), "\n")
	for i := 0; i < len(ligne); i++ {
		mots := strings.Fields(ligne[i])
		ligne[i] = strings.Join(mots, " ")
	}
	return ligne, nil
}
