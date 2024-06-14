package lista_test

import (
	TDALISTA "tdas/lista"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	RANGO_VOLUMEN   = 100000
	RANGO_VOLUMEN_2 = 200000
)

func TestListaVAcia(t *testing.T) {
	lista := TDALISTA.CrearListaEnlazada[int]()
	require.True(t, lista.EstaVacia())
}
func TestPanicsLista(t *testing.T) {
	lista := TDALISTA.CrearListaEnlazada[int]()
	require.True(t, lista.EstaVacia())
	require.Panics(t, func() { lista.BorrarPrimero() })
	require.Panics(t, func() { lista.VerPrimero() })
	require.Panics(t, func() { lista.VerUltimo() })
}
func TestVerPrimero(t *testing.T) {
	lista := TDALISTA.CrearListaEnlazada[rune]()
	require.True(t, lista.EstaVacia())
}

func TestInsertarUltimo(t *testing.T) {
	lista := TDALISTA.CrearListaEnlazada[rune]()
	largo := 0
	letras := []rune("abcdefghijklmnñopqrstuvwxyz")
	for i := 0; i < len(letras); i++ {
		lista.InsertarUltimo(letras[i])
		largo++
		require.EqualValues(t, lista.VerUltimo(), letras[i])
		require.EqualValues(t, lista.Largo(), largo)
	}

}

func TestInsertarPrimero(t *testing.T) {
	lista := TDALISTA.CrearListaEnlazada[int]()
	largo := 0
	for i := 0; i < RANGO_VOLUMEN; i++ {
		lista.InsertarPrimero(i)
		largo++
		require.EqualValues(t, lista.VerPrimero(), i)
		require.EqualValues(t, lista.Largo(), largo)
	}
}

func TestActuarComoVacia(t *testing.T) {
	lista := TDALISTA.CrearListaEnlazada[float32]()
	numeros := []float32{3.11, 2.23, 132.23, 1.454, 1.2341}
	for i := 0; i < len(numeros); i++ {
		lista.InsertarUltimo(numeros[i])
	}
	for i := 0; i < len(numeros); i++ {
		require.EqualValues(t, lista.BorrarPrimero(), numeros[i])
	}
	require.True(t, lista.EstaVacia())
	require.Panics(t, func() { lista.BorrarPrimero() })
	require.Panics(t, func() { lista.VerPrimero() })
	require.Panics(t, func() { lista.VerUltimo() })
}

func TestInsertar(t *testing.T) {
	lista := TDALISTA.CrearListaEnlazada[int]()
	iterador := lista.Iterador()

	iterador.Insertar(7)
	iterador.Insertar(13)
	iterador.Insertar(9)
	iterador.Insertar(23)
	iterador.Insertar(34)

	// Insertar donde se crea el iterador
	iterador.Insertar(5)
	require.EqualValues(t, 5, lista.VerPrimero())
	iterador.Insertar(31)
	require.EqualValues(t, 31, lista.VerPrimero())

	// Insertar en el medio
	for i := 0; i < 5; i++ {
		iterador.Siguiente()
	}
	iterador.Insertar(11)
	require.EqualValues(t, 11, iterador.VerActual())

	// Insertar al final
	for iterador.HaySiguiente() {
		iterador.Siguiente()
	}
	iterador.Insertar(2)
	require.EqualValues(t, 2, iterador.VerActual())
	require.EqualValues(t, 2, lista.VerUltimo())
}

func TestBorrar(t *testing.T) {
	lista := TDALISTA.CrearListaEnlazada[int]()
	lista.InsertarUltimo(0)
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)
	lista.InsertarUltimo(4)
	require.False(t, lista.EstaVacia())

	iterador := lista.Iterador()
	iterador.VerActual()
	require.True(t, iterador.HaySiguiente())

	require.EqualValues(t, 0, iterador.Borrar())
	require.EqualValues(t, 1, iterador.VerActual())
	iterador.Siguiente()
	require.EqualValues(t, 1, lista.VerPrimero())

	require.EqualValues(t, 2, iterador.Borrar())
	require.EqualValues(t, 3, iterador.VerActual())
	iterador.Siguiente()

	require.EqualValues(t, 4, iterador.Borrar())
	require.False(t, iterador.HaySiguiente())
	require.False(t, lista.EstaVacia())
	require.EqualValues(t, 3, lista.VerUltimo())

}

func TestIteradorInterno(t *testing.T) {
	lista := TDALISTA.CrearListaEnlazada[int]()
	for i := 0; i < 12; i++ {
		lista.InsertarUltimo(i)
	}
	contador := 0
	lista.Iterar(func(n int) bool {
		if n < 10 {
			contador++
			return true
		}
		return false
	})
	require.EqualValues(t, contador, 10)

}

func TestIteradorInternoEnListaVacia(t *testing.T) {
	lista := TDALISTA.CrearListaEnlazada[int]()
	contador := 0

	contadorPrt := &contador
	lista.Iterar(func(elemento int) bool {
		*contadorPrt += 1
		return true
	})
	require.EqualValues(t, 0, contador)
}

func TestVolumenesDiferentes(t *testing.T) {
	lista := TDALISTA.CrearListaEnlazada[int]()
	for i := 0; i < RANGO_VOLUMEN; i++ {
		lista.InsertarUltimo(i)
	}

	require.False(t, lista.EstaVacia())
	iterador := lista.Iterador()
	require.True(t, iterador.HaySiguiente())
	for i := 0; iterador.HaySiguiente(); i++ {
		require.EqualValues(t, i, iterador.VerActual())
		require.EqualValues(t, i, iterador.Borrar())
	}

	require.True(t, lista.EstaVacia())
	for i := 0; i < RANGO_VOLUMEN_2; i++ {
		lista.InsertarPrimero(i)
	}
	contador := 0
	contadorPrt := &contador
	lista.Iterar(func(elemento int) bool {
		*contadorPrt += 1
		return true
	})
	require.EqualValues(t, RANGO_VOLUMEN_2, contador)

}
