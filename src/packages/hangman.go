package packages

import (
	"bufio"
	"math/rand"
	"os"
	"strings"
	"time"
)

type Hangman struct {
	WordStatus   []string
	LettersTried []string
	MaxTry       int
	NumTries     int
	Word         string
	Successful   bool
	Print        []string
}

var letter string

// Initialisation des données du jeu.
func (hang *Hangman) InitGame(difficulte string) {
	hang.WordStatus = []string{}
	hang.LettersTried = []string{}
	hang.MaxTry = 10
	hang.NumTries = 0
	hang.Word = GetWord(difficulte)
	hang.Successful = false
	hang.UpdateWordState(" ") // On met à jour l'état du mot.
}

// Obtention du mot à partir des fichiers wordsx.txt
func GetWord(difficulte string) string {
	wordfile, err := os.Open("words" + difficulte + ".txt")
	if err != nil {
		panic(err)
	}
	defer wordfile.Close()
	scanner := bufio.NewScanner(wordfile)
	scanner.Split(bufio.ScanLines)
	var line []string
	for scanner.Scan() {
		line = append(line, scanner.Text())
	}
	rand.Seed(time.Now().UnixNano())
	x := rand.Intn(len(line))
	return line[x]
}

// Cette fonction permet de vérifier qu'elle n'a pas déjà été essayée
func (hang *Hangman) VerifGuess(letter string) string {
	if strings.Contains(strings.Join(hang.LettersTried, ""), letter) { //si la lettre a déjà été essayée, on affiche un message d'erreur
		return "Vous avez déjà essayé cette lettre !"
	}
	if len(letter) > 1 { //si l'utilisateur a entré plus d'une lettre, on affiche un message d'erreur
		return "Vous devez entrer une seule lettre !"
	}
	if !strings.Contains("abcdefghijklmnopqrstuvwxyz", letter) { //si l'utilisateur n'a pas entré une lettre, on affiche un message d'erreur
		return "Vous devez entrer une lettre !"
	}
	hang.LettersTried = append(hang.LettersTried, letter) //on ajoute la lettre à la liste des lettres essayées
	return ""
}

// Cette fonction permet de mettre à jour l'état du mot en fonction de la lettre entrée par l'utilisateur.
func (hang *Hangman) UpdateWordState(letter string) {
	if letter == " " {
		for i := 0; i < len(hang.Word); i++ {
			hang.WordStatus = append(hang.WordStatus, "_") //on ajoute un "_" pour chaque lettre du mot
		}
	} else {
		for i, l := range hang.Word { //pour chaque lettre du mot on vérifie si elle correspond à la lettre entrée par l'utilisateur
			if string(l) == letter { //si oui, on remplace le "_" par la lettre
				hang.WordStatus[i] = letter
			}
		}
		if !strings.Contains(hang.Word, letter) { // Si la lettre n'est pas dans le mot
			hang.NumTries++
		}
	}
}

// Cette fonction permet de vérifier si le joueur a trouvé le mot, ne l'a pas trouvé, ou a perdu.
func (hang *Hangman) ContinueGame() string {
	if hang.NumTries >= hang.MaxTry { //si le nombre d'essais est supérieur au nombre d'essais maximum, on affiche un message d'erreur et on retourne lose
		return "lose"
	}
	if strings.Join(hang.WordStatus, "") == hang.Word { //si le mot est trouvé, on affiche un message de victoire et on retourne win
		return "win"
	}
	return ""
}
