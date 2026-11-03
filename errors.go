package main

import (
	"fmt"
	"os"
)

func run() error {
	if len(os.Args) != 3 {
		return fmt.Errorf("utilisation : go run . fichier_entree fichier_sortie")
	}

	input := os.Args[1]
	output := os.Args[2]

	data, err := os.ReadFile(input)
	if err != nil {
		return fmt.Errorf("lecture de %q : %w", input, err)
	}

	result, err := ProcessText(string(data))
	if err != nil {
		return err
	}

	if err := os.WriteFile(output, []byte(result), 0644); err != nil {
		return fmt.Errorf("ecriture de %q : %w", output, err)
	}

	return nil
}
