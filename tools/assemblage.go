package tools

// Assemblage s'occupe de traité les lignes du texte
func Assemblage(lignes []string) []string {
	lignes = Marqueurs(lignes)
	for i := 0; i < len(lignes); i++ {
		lignes[i] = OccurrenceDeA(lignes[i])
		lignes[i] = Guillemets(lignes[i])
		lignes[i] = Ponctuation(lignes[i])
	}
	return lignes
}
