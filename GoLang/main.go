//Allan Diogo Bastos 03/07/2024

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var notesAvailable = []int{200, 100, 50, 20, 10, 5}

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Print("Digite o valor de saque: ")
	input, exception := inputReader.ReadString('\n')

	checkIfThereIsAnError(exception)

	input = strings.TrimSpace(input)
	amountValue, exception := strconv.Atoi(input)

	checkIfThereIsAnError(exception)

	if amountValue <= 0 {
		fmt.Printf("Não é possivel que o valor de saque seja %d", amountValue)
		return
	}

	arrayNoteCount, success := calcNotes(amountValue)

	if success {
		fmt.Printf("Para chegar no valor de %d, a saida de notas será:\n\n", amountValue)

		//Printando na tela a quantidade de notas
		for note, count := range arrayNoteCount {
			if count > 0 {
				fmt.Printf("%d notas de %d\n", count, note)
			}
		}
	} else {
		fmt.Println("Valor indisponível para saque com as notas disponíveis.")
	}
}

// Função reponsavel para calcular o minimo de cada nota disponivel
func calcNotes(amountParam int) (map[int]int, bool) {

	noteCount := make(map[int]int)
	amount := amountParam

	for _, note := range notesAvailable {
		if amount >= note {
			noteCount[note] = amount / note
			amount = amount % note
		}
	}

	if amount != 0 {
		return nil, false
	}

	return noteCount, true
}

func checkIfThereIsAnError(err error) {
	if err != nil {
		fmt.Println("Não foi possível ler o valor informado:", err)
		os.Exit(1)
	}
}
