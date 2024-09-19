package main

import (
	"fmt"
	"math/rand"
)

func get_theme(themes []string) string {
	return themes[rand.Intn(len(themes))]
}

func get_letter(alphabet []int32) string {
	return string(alphabet[rand.Intn(len(alphabet))])
}

func main() {
	alphabet := []int32("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	letter := get_letter(alphabet)
	themes := [7]string{"CIDADE", "PAÍS", "ESTADO", "NOMES", "ANIMAIS", "FRUTAS", "MARCAS"}
	theme := get_theme(themes[:])
	fmt.Printf("Theme: %s\nLetter: %s", theme, letter)
}
