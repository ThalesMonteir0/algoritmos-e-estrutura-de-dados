package main

func main() {
	arrDesordenado := []int{10, 7, 5, 8, 3, 2}
	arrOrdenado := ordenaPorSelecao(arrDesordenado)
	for _, item := range arrOrdenado {
		println(item)
	}

}

func buscaMenor(arr []int) int {
	menor := arr[0]
	menorIndice := 0

	for i := 1; i < len(arr); i++ {
		if arr[i] < menor {
			menor = arr[i]
			menorIndice = i
		}
	}

	return menorIndice
}

func ordenaPorSelecao(arr []int) []int {
	var novoArr []int

	for len(arr) > 0 {
		menor := buscaMenor(arr)
		novoArr = append(novoArr, arr[menor])
		arr = append(arr[:menor], arr[menor+1:]...)
	}

	return novoArr
}
