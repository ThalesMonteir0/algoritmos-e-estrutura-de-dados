package pesquisa_binaria

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPesquisaBinaria(t *testing.T) {

	t.Run("encontrando_item_desejado", func(t *testing.T) {
		list := []int{1, 3, 5, 7, 9}
		indice := PesquisaBinaria(list, 3)

		assert.EqualValues(t, 1, indice)
	})

	t.Run("item_desejado_sendo_o_ultimo", func(t *testing.T) {
		list := []int{1, 3, 5, 7, 9}
		indice := PesquisaBinaria(list, 9)

		assert.EqualValues(t, 4, indice)
	})

	t.Run("nao_encontrando_item_desejado", func(t *testing.T) {
		list := []int{1, 3, 5, 7, 9}
		indice := PesquisaBinaria(list, 10)

		assert.EqualValues(t, -1, indice)
	})
}
