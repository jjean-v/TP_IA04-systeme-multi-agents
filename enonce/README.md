# Quelques utilitaires et démos en `go` pour IA04

[![Go Report Card](https://goreportcard.com/badge/gitlab.utc.fr/lagruesy/ia04)](https://goreportcard.com/report/gitlab.utc.fr/lagruesy/ia04)
[![Go Reference](https://pkg.go.dev/badge/gitlab.utc.fr/lagruesy/ia04/utils.svg)](https://pkg.go.dev/gitlab.utc.fr/lagruesy/ia04/utils)

## Récupération du package `utils`

```bash
# création du module
go mod init monsupermodule

# récupération
go get gitlab.utc.fr/lagruesy/ia04/utils
```

## Utilisation des permutations 

### Exemple

```golang
package main

import (
	"fmt"

	"gitlab.utc.fr/lagruesy/ia04/utils"
)

func main() {
	const n = 4

	// création et affichage de la première permutation
	perm := utils.FirstPermutation(n)
	fmt.Println(perm)

	// compteur des permutations
	count := 1

	// itération et affichage de la permutation suivante
	perm, ok := utils.NextPermutation(perm)
	for ok {
		count++
		fmt.Println(perm)
		perm, ok = utils.NextPermutation(perm)
	}

	// on affiche la valeur de la factorielle...
	fmt.Printf("\n%d!=%d\n", n, count)
}
```
