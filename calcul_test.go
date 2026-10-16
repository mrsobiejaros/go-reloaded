package main

import "testing"

func TestCorrigerTexte(t *testing.T) {
	tests := []struct {
		nom     string
		entree  string
		attendu string
	}{
		{
			nom:     "normal - hex",
			entree:  "1E (hex) files were added",
			attendu: "30 files were added",
		},
		{
			nom:     "normal - bin",
			entree:  "It has been 10 (bin) years",
			attendu: "It has been 2 years",
		},
		{
			nom:     "normal - up",
			entree:  "Ready, set, go (up) !",
			attendu: "Ready, set, GO!",
		},
		{
			nom:     "normal - low",
			entree:  "I should stop SHOUTING (low)",
			attendu: "I should stop shouting",
		},
		{
			nom:     "normal - cap",
			entree:  "Welcome to the Brooklyn bridge (cap)",
			attendu: "Welcome to the Brooklyn Bridge",
		},
		{
			nom:     "normal - up avec nombre",
			entree:  "This is so exciting (up, 2)",
			attendu: "This is SO EXCITING",
		},
		{
			nom:     "normal - low avec nombre",
			entree:  "I SHOULD STOP SHOUTING (low, 2)",
			attendu: "I SHOULD stop shouting",
		},
		{
			nom:     "normal - cap avec nombre",
			entree:  "this is a brooklyn bridge (cap, 2)",
			attendu: "this is a Brooklyn Bridge",
		},
		{
			nom:     "normal - ponctuation",
			entree:  "I was sitting over there ,and then BAMM !!",
			attendu: "I was sitting over there, and then BAMM!!",
		},
		{
			nom:     "normal - tous les signes",
			entree:  "one . two , three ! four ? five : six ; seven",
			attendu: "one. two, three! four? five: six; seven",
		},
		{
			nom:     "normal - groupe ...",
			entree:  "I was thinking ... You were right",
			attendu: "I was thinking... You were right",
		},
		{
			nom:     "normal - groupe !?",
			entree:  "Are you sure !? yes",
			attendu: "Are you sure!? yes",
		},
		{
			nom:     "normal - guillemets un mot",
			entree:  "I am exactly how they describe me: ' awesome '",
			attendu: "I am exactly how they describe me: 'awesome'",
		},
		{
			nom:     "normal - guillemets plusieurs mots",
			entree:  "As Elton John said: ' I am the most well-known homosexual in the world '",
			attendu: "As Elton John said: 'I am the most well-known homosexual in the world'",
		},
		{
			nom:     "normal - a devient an",
			entree:  "There it was. A amazing rock!",
			attendu: "There it was. An amazing rock!",
		},
		{
			nom:     "normal - a devant un h",
			entree:  "I waited a hour",
			attendu: "I waited an hour",
		},
		{
			nom:     "normal - a devant une consonne",
			entree:  "I saw a cat",
			attendu: "I saw a cat",
		},
		{
			nom:     "sujet - exemple 1",
			entree:  "it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom, it was the age of foolishness (cap, 6)",
			attendu: "It was the best of times, it was the worst of TIMES, it was the age of wisdom, It Was The Age Of Foolishness",
		},
		{
			nom:     "sujet - exemple 2",
			entree:  "Simply add 42 (hex) and 10 (bin) and you will see the result is 68.",
			attendu: "Simply add 66 and 2 and you will see the result is 68.",
		},
		{
			nom:     "sujet - exemple 3",
			entree:  "There is no greater agony than bearing a untold story inside you.",
			attendu: "There is no greater agony than bearing an untold story inside you.",
		},
		{
			nom:     "sujet - exemple 4",
			entree:  "Punctuation tests are ... kinda boring ,what do you think ?",
			attendu: "Punctuation tests are... kinda boring, what do you think?",
		},
		{
			nom:     "limite - texte vide",
			entree:  "",
			attendu: "",
		},
		{
			nom:     "limite - tag au début",
			entree:  "(up) hello",
			attendu: "hello",
		},
		{
			nom:     "limite - nombre trop grand",
			entree:  "hello world (up, 5)",
			attendu: "HELLO WORLD",
		},
		{
			nom:     "limite - a à la fin",
			entree:  "this is a",
			attendu: "this is a",
		},
		{
			nom:     "limite - ponctuation seule",
			entree:  "...",
			attendu: "...",
		},
		{
			nom:     "piege - hex invalide",
			entree:  "zz (hex) test",
			attendu: "zz test",
		},
		{
			nom:     "piege - bin invalide",
			entree:  "12 (bin) test",
			attendu: "12 test",
		},
		{
			nom:     "piege - tag inconnu",
			entree:  "hello (xyz) world",
			attendu: "hello (xyz) world",
		},
		{
			nom:     "piege - tag sans espace dans le nombre",
			entree:  "so exciting (up,2)",
			attendu: "SO EXCITING",
		},
		{
			nom:     "piege - cap sur un mot en majuscules",
			entree:  "HELLO (cap)",
			attendu: "Hello",
		},
		{
			nom:     "piege - apostrophe dans un mot",
			entree:  "I don't know",
			attendu: "I don't know",
		},
		{
			nom:     "piege - guillemets puis point",
			entree:  "He said : ' hello ' .",
			attendu: "He said: 'hello'.",
		},
		{
			nom:     "piege - plusieurs espaces",
			entree:  "hello    world",
			attendu: "hello world",
		},
		{
			nom:     "piege - plusieurs lignes",
			entree:  "it (cap) works\nthe end (up)",
			attendu: "It works\nthe END",
		},
		{
			nom:     "piege - retour à la ligne final",
			entree:  "hello ,world\n",
			attendu: "hello, world\n",
		},
	}

	for _, test := range tests {
		t.Run(test.nom, func(t *testing.T) {
			resultat := corrigerTexte(test.entree)
			if resultat != test.attendu {
				t.Errorf("entrée %q : obtenu %q, attendu %q", test.entree, resultat, test.attendu)
			}
		})
	}
}
