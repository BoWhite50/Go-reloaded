package main

import (
	"fmt"
	"go-reloaded/tools"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr, "le programme doit être lancé avec 2 fichiers: le fichier d'entrée et le fichier de sortie")
		os.Exit(1)
	}

	lignes, err := tools.Readsample(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Erreur lors de la lecture du fichier d'entrée:", err)
		os.Exit(1)
	}
	tools.Assemblage(lignes)

	err = os.WriteFile(os.Args[2], []byte(strings.Join(lignes, "\n")), 0644)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Erreur lors de l'écriture du fichier de sortie:", err)
		os.Exit(1)
	}
}
