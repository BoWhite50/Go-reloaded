# Go-reloaded

## Présentation

Go-reloaded est un outil en ligne de commande qui corrige et met en forme un texte automatiquement.
Il lit un fichier d'entrée, applique des règles de modification (conversions de nombres, majuscules, ponctuation, grammaire) et écrit le texte corrigé dans un fichier de sortie.

## Fonctionnalités

| Règle | Exemple d'entrée | Sortie |
|---|---|---|
| `(hex)` : hexadécimal → décimal | `1E (hex) files were added` | `30 files were added` |
| `(bin)` : binaire → décimal | `It has been 10 (bin) years` | `It has been 2 years` |
| `(up)` : majuscules | `Ready, set, go (up) !` | `Ready, set, GO!` |
| `(low)` : minuscules | `I should stop SHOUTING (low)` | `I should stop shouting` |
| `(cap)` : majuscule initiale | `Welcome to the Brooklyn bridge (cap)` | `Welcome to the Brooklyn Bridge` |
| `(up, N)`, `(low, N)`, `(cap, N)` : sur les N mots précédents | `This is so exciting (up, 2)` | `This is SO EXCITING` |
| Ponctuation `. , ! ? : ;` collée au mot précédent | `over there ,and then BAMM !!` | `over there, and then BAMM!!` |
| Groupes de ponctuation `...` `!?` | `I was thinking ... You were right` | `I was thinking... You were right` |
| Apostrophes `' '` autour d'un ou plusieurs mots | `describe me: ' awesome '` | `describe me: 'awesome'` |
| `a` → `an` devant une voyelle ou un `h` | `A amazing rock` | `An amazing rock` |
| guillemets doubles `" "` | `He said: " hello world "` | `He said: "hello world"` |

Une fois appliquée, chaque balise disparaît du texte, sans laisser d'espace en trop.
Les retours à la ligne du fichier d'entrée sont conservés.

## Technologies

- Go, uniquement la bibliothèque standard (`os`, `strings`, `strconv`, `fmt`, `testing`)

## Prérequis

- Go 1.27 ou plus récent (voir `go.mod`)

## Installation

```bash
git clone https://github.com/BoWhite50/Go-reloaded.git
cd Go-reloaded
```

Aucune dépendance à installer.

## Lancement

```bash
go run . <fichier_entree> <fichier_sortie>
```

Exemple :

```bash
$ cat sample.txt
Simply add 42 (hex) and 10 (bin) and you will see the result is 68.
$ go run . sample.txt result.txt
$ cat result.txt
Simply add 66 and 2 and you will see the result is 68.
```

Si le nombre d'arguments est incorrect, ou si le fichier d'entrée n'existe pas, le programme affiche un message d'erreur.

## Tests

```bash
go test ./...
```

Pour voir le détail de chaque test :

```bash
go test -v ./...
```

- `main_test.go` teste le programme complet (`tools.Assemblage`) avec les exemples du sujet et des cas limites.
- `tools/outils_test.go` teste chaque fonction du package `tools` séparément.

## Architecture

```
Go-reloaded/
├── go.mod              module go-reloaded
├── main.go             point d'entrée : lit les arguments, appelle tools.Assemblage, écrit le résultat
├── main_test.go        tests du programme complet
└── tools/
    ├── assemblage.go   Assemblage : enchaîne toutes les corrections dans le bon ordre
    ├── marqueurs.go    Marqueurs : applique et retire les balises (hex), (up, 2)...
    ├── outils.go       les autres fonctions de correction
    └── outils_test.go  tests unitaires des fonctions
```

Ordre des corrections appliquées à chaque ligne (fonction `Assemblage`) :

1. `Marqueurs` : applique et retire les balises `(hex)`, `(bin)`, `(up)`, `(low)`, `(cap)` et leurs versions avec un nombre
2. `OccurrenceDeA` : transforme `a` en `an` quand il le faut
3. `Guillemets` : colle les `'` et les `"` aux mots qu'ils entourent
4. `Ponctuation` : colle la ponctuation au mot précédent

Fonctions du package `tools` :

| Fonction | Rôle |
|---|---|
| `Assemblage` | enchaîne toutes les corrections sur chaque ligne |
| `Marqueurs` | applique et retire les balises `(hex)`, `(bin)`, `(up)`, `(low)`, `(cap)`, `(xxx, N)` |
| `Readsample` | lit le fichier d'entrée et renvoie ses lignes |
| `HexOrBinToDecimal` | convertit un nombre hexadécimal ou binaire en décimal |
| `Transform` | met un mot en majuscules, minuscules ou avec une majuscule initiale |
| `OccurrenceDeA` | remplace `a` par `an` devant une voyelle ou un `h` |
| `Guillemets` | colle les apostrophes et guillemets aux mots |
| `Ponctuation` | colle la ponctuation au mot précédent |

## Auteur

Yannis ([BoWhite50](https://github.com/BoWhite50))
