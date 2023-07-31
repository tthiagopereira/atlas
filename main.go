package main

import (
	"fmt"
	"sync"
)

func createGraphStates() map[string][]string {
	graphs := map[string][]string{

		"RR": []string{"AM", "PA"},
		"AM": []string{"RR", "AP", "PA", "MT", "RO", "AC"},
		"AP": []string{"AM", "PA"},
		"PA": []string{"RR", "AM", "AP", "MT", "TO", "MA"},
		"TO": []string{"PA", "MA", "PI", "BA", "GO"},
		"MA": []string{"PA", "TO", "PI", "BA"},

		"PI": []string{"MA", "TO", "BA", "PE", "CE"},
		"CE": []string{"PI", "RN", "PB", "PE"},
		"RN": []string{"CE", "PB"},
		"PB": []string{"RN", "CE", "PE"},
		"PE": []string{"CE", "PB", "AL", "BA"},
		"AL": []string{"PE", "BA", "SE"},
		"SE": []string{"AL", "BA"},

		"BA": []string{"SE", "AL", "PE", "PI", "TO", "GO", "MG", "ES"},
		"MG": []string{"BA", "GO", "MS", "SP", "RJ", "ES"},
		"ES": []string{"BA", "MG", "RJ"},
		"RJ": []string{"MG", "ES", "SP"},
		"SP": []string{"MG", "RJ", "MS", "PR"},

		"PR": []string{"SP", "SC"},
		"SC": []string{"PR", "RS"},
		"RS": []string{"SC"},

		"MS": []string{"MG", "SP", "PR", "MT", "GO"},
		"MT": []string{"AM", "PA", "TO", "GO", "MS"},
		"GO": []string{"TO", "MT", "MS", "MG", "BA"},
	}
	return graphs
}

func findWay(graphs map[string][]string, start, end string, result chan []string, wg *sync.WaitGroup) {
	defer wg.Done()

	row := [][]string{{start}}
	visited := make(map[string]bool)

	for len(row) > 0 {
		currentPath := row[0]
		row = row[1:]
		currentState := currentPath[len(currentPath)-1]

		if currentState == end {
			result <- currentPath
			return
		}

		if !visited[currentState] {
			visited[currentState] = true

			for _, neighbor := range graphs[currentState] {
				newWay := append([]string{}, currentPath...)
				newWay = append(newWay, neighbor)
				row = append(row, newWay)
			}
		}
	}

	result <- nil
}

func main() {
	graphStates := createGraphStates()
	wayChan := make(chan []string)
	wg := sync.WaitGroup{}

	start := "RS"
	end := "GO"

	for _, neighbor := range graphStates[start] {
		wg.Add(1)
		go findWay(graphStates, neighbor, end, wayChan, &wg)
	}

	go func() {
		wg.Wait()
		close(wayChan)
	}()

	var way []string
	for foundWay := range wayChan {
		if foundWay != nil {
			way = foundWay
			break
		}
	}

	if way != nil {
		fmt.Println("Estados relacionados:", way)
	} else {
		fmt.Println("Não foi encontrado um caminho.")
	}
}
