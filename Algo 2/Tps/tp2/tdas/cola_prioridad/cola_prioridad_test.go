package cola_prioridad_test

import (
	"fmt"
	TDAHEAP "tdas/cola_prioridad"
	"testing"

	"github.com/stretchr/testify/require"
)

var TAMAÑOS_VOLUMEN = []int{13000, 25000, 60000, 100000, 200000, 500000}

const _VALOR_STRING = "a"
const _VALOR_INT = 2
const _INICIO_INTS = 100000

func heapString(a, b string) int {
	if a > b {
		return 1
	} else if b > a {
		return -1
	}
	return 0
}

func heapInt(a, b int) int {
	if a > b {
		return 1
	} else if b > a {
		return -1
	}
	return 0
}
func TestPanics(t *testing.T) {
	h := TDAHEAP.CrearHeap(heapString)
	require.True(t, h.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { h.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { h.Desencolar() })
	h.Encolar(_VALOR_STRING)
	require.False(t, h.EstaVacia())
}

func TestEncolarDesencolar(t *testing.T) {
	h := TDAHEAP.CrearHeap(heapInt)
	require.EqualValues(t, 0, h.Cantidad())
	h.Encolar(_VALOR_INT)
	require.EqualValues(t, 1, h.Cantidad())
	require.False(t, h.EstaVacia())
	require.EqualValues(t, _VALOR_INT, h.VerMax())
	require.EqualValues(t, _VALOR_INT, h.Desencolar())
	require.True(t, h.EstaVacia())
}

func TestEncolarVariosElementos(t *testing.T) {
	h := TDAHEAP.CrearHeap(heapInt)
	for i := 0; i <= _INICIO_INTS; i++ {
		h.Encolar(i)
	}
	for i := _INICIO_INTS; i >= 0; i-- {
		require.EqualValues(t, i, h.VerMax())
		require.EqualValues(t, i, h.Desencolar())
	}
	require.True(t, h.EstaVacia())
}

func TestHeapify(t *testing.T) {
	arr := []int{50, 204, 1000, 450, 205102, 2}
	heap := TDAHEAP.CrearHeapArr(arr, heapInt)
	require.EqualValues(t, 6, heap.Cantidad())
	require.EqualValues(t, 205102, heap.Desencolar())
	require.False(t, heap.EstaVacia())
	require.EqualValues(t, 5, heap.Cantidad())
	require.EqualValues(t, 1000, heap.Desencolar())
	require.False(t, heap.EstaVacia())
	require.EqualValues(t, 4, heap.Cantidad())
	require.EqualValues(t, 450, heap.Desencolar())
	require.False(t, heap.EstaVacia())
	require.EqualValues(t, 3, heap.Cantidad())
	require.EqualValues(t, 204, heap.Desencolar())
	require.False(t, heap.EstaVacia())
	require.EqualValues(t, 2, heap.Cantidad())
	require.EqualValues(t, 50, heap.Desencolar())
	require.False(t, heap.EstaVacia())
	require.EqualValues(t, 1, heap.Cantidad())
	require.EqualValues(t, 2, heap.Desencolar())
	require.True(t, heap.EstaVacia())
}

func TestHeapifyString(t *testing.T) {
	arr := []string{"b", "l", "a", "k", "al", "c"}
	heap := TDAHEAP.CrearHeapArr(arr, heapString)
	require.EqualValues(t, 6, heap.Cantidad())
	require.EqualValues(t, "l", heap.Desencolar())
	require.False(t, heap.EstaVacia())
	require.EqualValues(t, 5, heap.Cantidad())
	require.EqualValues(t, "k", heap.Desencolar())
	require.False(t, heap.EstaVacia())
	require.EqualValues(t, 4, heap.Cantidad())
	require.EqualValues(t, "c", heap.Desencolar())
	require.False(t, heap.EstaVacia())
	require.EqualValues(t, 3, heap.Cantidad())
	require.EqualValues(t, "b", heap.Desencolar())
	require.False(t, heap.EstaVacia())
	require.EqualValues(t, 2, heap.Cantidad())
	require.EqualValues(t, "al", heap.Desencolar())
	require.False(t, heap.EstaVacia())
	require.EqualValues(t, 1, heap.Cantidad())
	require.EqualValues(t, "a", heap.Desencolar())
	require.True(t, heap.EstaVacia())

}

func TestHeapVacio(t *testing.T) {
	arr := []string{}
	h := TDAHEAP.CrearHeapArr(arr, heapString)
	require.True(t, h.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { h.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { h.Desencolar() })
	h.Encolar(_VALOR_STRING)
	require.False(t, h.EstaVacia())
	h.Encolar("b")
	h.Encolar("c")
	require.EqualValues(t, "c", h.Desencolar())
	require.EqualValues(t, "b", h.Desencolar())
	require.EqualValues(t, "a", h.Desencolar())

}

func TestHeapsort(t *testing.T) {
	arr := []int{1, 20, 2, 60, 12, 102, 0}
	arr_ordenado := []int{0, 1, 2, 12, 20, 60, 102}
	TDAHEAP.HeapSort(arr, heapInt)
	for i := 0; i < len(arr); i++ {
		require.EqualValues(t, arr[i], arr_ordenado[i])
	}
}

func TestHeapSortVacio(t *testing.T) {
	arr := make([]int, 0)
	arr_ordenado := make([]int, 0)
	TDAHEAP.HeapSort(arr, heapInt)
	require.EqualValues(t, arr, arr_ordenado)
}

func TestHeapSortString(t *testing.T) {
	arr := []string{"mi", "mama", "me", "mima"}
	arr_ordenado := []string{"mama", "me", "mi", "mima"}
	TDAHEAP.HeapSort(arr, heapString)
	require.EqualValues(t, arr, arr_ordenado)
}

func ejecutarPruebaVolumen(b *testing.B, n int) {
	heap := TDAHEAP.CrearHeap[string](heapString)
	claves := make([]string, n)

	for i := 0; i < n; i++ {
		claves[i] = fmt.Sprintf("%08d", i)
		heap.Encolar(claves[i])
		require.EqualValues(b, claves[i], heap.VerMax())
	}

	require.EqualValues(b, n, heap.Cantidad())

	for i := 0; i < n; i++ {
		heap.Desencolar()
	}
	require.EqualValues(b, 0, heap.Cantidad())
	require.True(b, heap.EstaVacia())
}
func BenchmarkHeap(b *testing.B) {
	for _, n := range TAMAÑOS_VOLUMEN {
		b.Run(fmt.Sprintf("%d Elementos", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ejecutarPruebaVolumen(b, n)
			}
		})
	}
}
