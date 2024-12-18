package main

import "fmt"

func main() {
	j1 := "Batata"
	j2 := ""

	resultMap := map[[2]string]string{
		{"Pedra", "Tesoura"}:   "V",
		{"Tesoura", "Papel"}:   "V",
		{"Papel", "Pedra"}:     "V",
		{"Tesoura", "Pedra"}:   "D",
		{"Pedra", "Papel"}:     "D",
		{"Papel", "Tesoura"}:   "D",
		{"Tesoura", "Tesoura"}: "E",
		{"Pedra", "Pedra"}:     "E",
		{"Papel", "Papel"}:     "E",
	}

	in := [2]string{j1, j2}

	result, ok := resultMap[in]
	if !ok {
		fmt.Printf("Jogada invalida J1=%s, J2=%s", j1, j2)
	} else {
		fmt.Printf("Jogador 1 %s\n", result)
	}

}

func resolve(j1, j2 string) string {
	if j1 == "Pedra" && j2 == "Tesoura" {
		return "j1"
	}
	if j1 == "Tesoura" && j2 == "Papel" {
		return "j1"
	}
	if j1 == "Papel" && j2 == "Pedra" {
		return "j1"
	}
	return "j2"
}
