/*
O bubble sort, ou ordenação por flutuação (literalmente "por bolha"), é um algoritmo de ordenação dos mais simples.
A ideia é percorrer o vector diversas vezes, e a cada passagem fazer flutuar para o topo o maior elemento da sequência.
Essa movimentação lembra a forma como as bolhas em um tanque de água procuram seu próprio nível, e disso vem o nome
do algoritmo.

No melhor caso, o algoritmo executa n operações relevantes, onde n representa o número de elementos do vetor.
No pior caso, são feitas n² operações.
A complexidade desse algoritmo é de ordem quadrática.
Por isso, ele não é recomendado para programas que precisem de velocidade e operem com quantidade elevada de dados.
*/

package main;

import (
  "time"
)

var count int

func Bubble_Sort(array []int) ([]int, int, float64) {
	arrayLength := len(array)
	var inicio int64	
	var final int64
	var tempofinal float64


	inicio = time.Now().UnixNano()

	for i := 0; i < arrayLength; i++ {
		count++
		for j := 0; j < arrayLength-i-1; j++ {
			count++
			if array[j] > array[j+1] {
				count++
				array[j], array[j+1] = array[j+1], array[j]
				count++
			}
		}
	}

  	tempofinal = float64(final-inicio)/ 1000000

	return array, count, tempofinal
}
