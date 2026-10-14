
package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) >= 3 {
		return 
	}

	inputFile := os.Args[1]

	content, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Erreur lors de la lecture du fichier :", err)
		return
	}

	fmt.Println(string(content))
}
