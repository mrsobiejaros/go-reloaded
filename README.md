# Go-reloaded

Ce programme lit un fichier texte, corrige le texte, puis écrit le résultat dans un autre fichier.

## Utilisation

Installer Go 1.20 ou une version plus récente. Ouvrir un terminal dans ce dossier, puis lancer :

```sh
go run . sample.txt result.txt
```

Par exemple, si `sample.txt` contient `Simply add 42 (hex) and 10 (bin) and you will see the result is 68.`, alors `result.txt` contiendra `Simply add 66 and 2 and you will see the result is 68.`

## Ce que le programme corrige

- `(hex)` et `(bin)` : le mot avant devient un nombre décimal.
- `(up)`, `(low)`, `(cap)` : le mot avant passe en majuscules, en minuscules ou avec une majuscule initiale. Avec un nombre, comme `(up, 2)`, cela s'applique aux 2 mots avant.
- Ponctuation `. , ! ? : ;` : collée au mot avant, avec un espace après. Les groupes comme `...` ou `!?` restent ensemble.
- Guillemets `'` : collés aux mots qu'ils entourent, sans espace.
- `a` devient `an` devant une voyelle ou un `h`.

## Organisation

- `calcul.go` : toutes les corrections du texte.
- `main.go` : lecture du fichier d'entrée, écriture du fichier de sortie, messages d'erreur.
- `calcul_test.go` : 36 cas de test (normal, limite, piège et exemples du sujet).
- `go.mod` : déclaration du module Go.

Les fonctions de `calcul.go` :

- `corrigerTexte` : découpe le texte en lignes.
- `corrigerLigne` : enchaîne les corrections sur une ligne.
- `rapprocherTags`, `lireTag`, `appliquerTag`, `appliquerTags` : gèrent les tags.
- `baseVersDecimal`, `mettreMajuscule` : conversions d'un mot.
- `corrigerArticles`, `commenceParVoyelle` : gèrent `a` / `an`.
- `corrigerPonctuation`, `estPonctuation` : gèrent la ponctuation.
- `corrigerGuillemets`, `estApostrophe` : gèrent les guillemets simples.

Les corrections se font dans cet ordre : tags, `a` / `an`, ponctuation, guillemets.

## Décisions

- Un tag au tout début du texte, sans mot avant lui, est simplement supprimé.
- Si le nombre dans `(up, 5)` dépasse le nombre de mots disponibles, tous les mots avant sont modifiés.
- Un mot qui n'est pas un vrai nombre hexadécimal ou binaire reste inchangé.
- Un tag inconnu, comme `(xyz)`, reste dans le texte.
- Un `'` entre deux lettres (comme dans `don't`) est une apostrophe et n'est pas touché.
- Le texte est traité ligne par ligne et les retours à la ligne sont conservés.
- Les espaces en trop entre les mots sont réduits à un seul.
- Les tags doivent être séparés des mots par des espaces, comme dans les exemples du sujet.

## Vérifications

```sh
go test -v .
gofmt -w .
go vet .
```
