package pesquisa_binaria

func PesquisaBinaria(list []int, item int) int {
	baixo := 0
	alto := len(list) - 1

	for baixo <= alto {
		meio := (baixo + alto) / 2

		if list[meio] == item {
			return meio
		}

		if list[meio] > item {
			alto = meio - 1
		} else {
			baixo = meio + 1
		}
	}
	return -1
}
