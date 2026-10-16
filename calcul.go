package main

import (
	"strconv"
	"strings"
	"unicode"
)

func baseVersDecimal(mot string, base int) string {
	nombre, err := strconv.ParseInt(mot, base, 64)
	if err != nil {
		return mot
	}
	return strconv.FormatInt(nombre, 10)
}

func mettreMajuscule(mot string) string {
	lettres := []rune(strings.ToLower(mot))
	if len(lettres) == 0 {
		return mot
	}
	lettres[0] = unicode.ToUpper(lettres[0])
	return string(lettres)
}

func rapprocherTags(ligne string) string {
	for _, nom := range []string{"up", "low", "cap"} {
		ligne = strings.ReplaceAll(ligne, "("+nom+", ", "("+nom+",")
	}
	return ligne
}

func lireTag(mot string) (string, int, bool) {
	if !strings.HasPrefix(mot, "(") || !strings.HasSuffix(mot, ")") {
		return "", 0, false
	}
	morceaux := strings.Split(mot[1:len(mot)-1], ",")
	nom := morceaux[0]
	if len(morceaux) == 1 && (nom == "hex" || nom == "bin" || nom == "up" || nom == "low" || nom == "cap") {
		return nom, 1, true
	}
	if len(morceaux) == 2 && (nom == "up" || nom == "low" || nom == "cap") {
		nombre, err := strconv.Atoi(morceaux[1])
		if err == nil {
			return nom, nombre, true
		}
	}
	return "", 0, false
}

func appliquerTag(mots []string, nom string, nombre int) {
	debut := len(mots) - nombre
	if debut < 0 {
		debut = 0
	}
	for i := debut; i < len(mots); i++ {
		switch nom {
		case "hex":
			mots[i] = baseVersDecimal(mots[i], 16)
		case "bin":
			mots[i] = baseVersDecimal(mots[i], 2)
		case "up":
			mots[i] = strings.ToUpper(mots[i])
		case "low":
			mots[i] = strings.ToLower(mots[i])
		case "cap":
			mots[i] = mettreMajuscule(mots[i])
		}
	}
}

func appliquerTags(mots []string) []string {
	var resultat []string
	for _, mot := range mots {
		nom, nombre, estTag := lireTag(mot)
		if estTag {
			appliquerTag(resultat, nom, nombre)
		} else {
			resultat = append(resultat, mot)
		}
	}
	return resultat
}

func commenceParVoyelle(mot string) bool {
	if mot == "" {
		return false
	}
	premiere := unicode.ToLower([]rune(mot)[0])
	return strings.ContainsRune("aeiouh", premiere)
}

func corrigerArticles(mots []string) []string {
	for i := 0; i < len(mots)-1; i++ {
		if !commenceParVoyelle(mots[i+1]) {
			continue
		}
		if mots[i] == "a" {
			mots[i] = "an"
		}
		if mots[i] == "A" {
			mots[i] = "An"
		}
	}
	return mots
}

func estPonctuation(c rune) bool {
	return strings.ContainsRune(".,!?:;", c)
}

func corrigerPonctuation(texte string) string {
	resultat := ""
	apresPonctuation := false
	for _, c := range texte {
		if estPonctuation(c) {
			resultat = strings.TrimRight(resultat, " ")
			resultat += string(c)
			apresPonctuation = true
		} else if c == ' ' {
			resultat += " "
			apresPonctuation = false
		} else {
			if apresPonctuation {
				resultat += " "
			}
			resultat += string(c)
			apresPonctuation = false
		}
	}
	return resultat
}

func estApostrophe(lettres []rune, i int) bool {
	if i == 0 || i == len(lettres)-1 {
		return false
	}
	return unicode.IsLetter(lettres[i-1]) && unicode.IsLetter(lettres[i+1])
}

func corrigerGuillemets(texte string) string {
	lettres := []rune(texte)
	resultat := ""
	ouvert := false
	for i := 0; i < len(lettres); i++ {
		c := lettres[i]
		if c != '\'' || estApostrophe(lettres, i) {
			resultat += string(c)
		} else if !ouvert {
			resultat += "'"
			ouvert = true
			for i+1 < len(lettres) && lettres[i+1] == ' ' {
				i++
			}
		} else {
			resultat = strings.TrimRight(resultat, " ")
			resultat += "'"
			ouvert = false
		}
	}
	return resultat
}

func corrigerLigne(ligne string) string {
	mots := strings.Fields(rapprocherTags(ligne))
	mots = appliquerTags(mots)
	mots = corrigerArticles(mots)
	texte := strings.Join(mots, " ")
	texte = corrigerPonctuation(texte)
	texte = corrigerGuillemets(texte)
	return texte
}

func corrigerTexte(texte string) string {
	lignes := strings.Split(texte, "\n")
	for i, ligne := range lignes {
		lignes[i] = corrigerLigne(ligne)
	}
	return strings.Join(lignes, "\n")
}