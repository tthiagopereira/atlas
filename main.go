package main

import "fmt"

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

func findWay(graphs map[string][]string, start, end string, cache map[string][]string) []string {
	row := [][]string{{start}}
	visited := make(map[string]bool)

	for len(row) > 0 {
		currentPath := row[0]
		row = row[1:]
		currentState := currentPath[len(currentPath)-1]

		if currentState == end {
			return currentPath
		}

		if !visited[currentState] {
			visited[currentState] = true

			// Verificar se o caminho já está no cache
			if way, ok := cache[currentState]; ok {
				// Se estiver no cache, usamos o caminho já encontrado
				for _, neighbor := range way {
					newWay := append([]string{}, currentPath...)
					newWay = append(newWay, neighbor)
					row = append(row, newWay)
				}
			} else {
				// Se não estiver no cache, calculamos o caminho e armazenamos no cache
				neighbors := graphs[currentState]
				cache[currentState] = neighbors

				for _, neighbor := range neighbors {
					newWay := append([]string{}, currentPath...)
					newWay = append(newWay, neighbor)
					row = append(row, newWay)
				}
			}
		}
	}

	return nil
}

func main() {
	graphStates := createGraphStates()
	cache := make(map[string][]string)

	way := findWay(graphStates, "SP", "RS", cache)

	if way != nil {
		fmt.Println("Caminhos processados: ", way)
	} else {
		fmt.Println("Não foi encontrado um caminho.")
	}
}
