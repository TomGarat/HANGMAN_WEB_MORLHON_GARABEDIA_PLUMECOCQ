package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"src/packages"
)

const port = "8082"
const url = "localhost"

var lettre string
var mot_affiche string

type MonJeu struct {
	packages.Hangman
}

var jeu MonJeu

func Difficulte(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		difficulte := r.FormValue("difficulty")
		if difficulte != "" {
			fmt.Println("difficulté:", difficulte)
			jeu.InitGame(difficulte)
			http.Redirect(w, r, "/hangman", http.StatusSeeOther)
		}
	}
	template.Must(template.ParseFiles("templates/index.html")).Execute(w, nil)
}

func (jeu *MonJeu) HangmanW(w http.ResponseWriter, r *http.Request) {
	if jeu.Word == "" { // Si la difficulté n'est pas définie, renvoyer à la page dédiée.
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	if r.Method == http.MethodPost {
		if r.FormValue("letter") != "" {
			lettre = r.FormValue("letter")
			fmt.Println("lettre:", lettre)
			verif := jeu.VerifGuess(lettre)
			if verif != "" {
				fmt.Println("Message:", verif)
			}
			jeu.UpdateWordState(lettre)
			fmt.Println(jeu)
			gameState := jeu.ContinueGame()
			if gameState != "" {
				fmt.Println(gameState)
				http.Redirect(w, r, "/"+gameState, http.StatusSeeOther)
			}
		}
	}
	template.Must(template.ParseFiles("templates/hangman.html")).Execute(w, nil)
}

func Win(w http.ResponseWriter, r *http.Request) {
	template.Must(template.ParseFiles("templates/win.html")).Execute(w, nil)
}

func Lose(w http.ResponseWriter, r *http.Request) {
	template.Must(template.ParseFiles("templates/lose.html")).Execute(w, nil)
}

func main() {
	static := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", static))

	http.HandleFunc("/", Difficulte)
	http.HandleFunc("/hangman", jeu.HangmanW)
	http.HandleFunc("/win", Win)
	http.HandleFunc("/lose", Lose)

	fmt.Println("(http://"+url+":"+port+") - Server started on port", port)
	http.ListenAndServe(":"+port, nil)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}

}
