package main

import (
	"errors"
	"fmt"
	"os"
)

func lireFichier(nom string) (string, error) {
	contenu, err := os.ReadFile(nom)
	if err != nil {
		return "", err
	}
	return string(contenu), nil
}

func ecrireFichier(nom string, texte string) error {
	return os.WriteFile(nom, []byte(texte), 0644)
}

func afficherErreur(err error) {
	fmt.Fprintln(os.Stderr, "Erreur :", err)
}

func main() {
	if len(os.Args) != 3 {
		afficherErreur(errors.New("usage : go run . entree.txt sortie.txt"))
		os.Exit(1)
	}
	texte, err := lireFichier(os.Args[1])
	if err != nil {
		afficherErreur(err)
		os.Exit(1)
	}
	err = ecrireFichier(os.Args[2], corrigerTexte(texte))
	if err != nil {
		afficherErreur(err)
		os.Exit(1)
	}
}
