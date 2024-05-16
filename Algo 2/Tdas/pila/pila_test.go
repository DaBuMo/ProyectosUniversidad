package pila_test

import (
	TDAPila "tdas/pila"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPilaVacia(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	require.True(t, pila.EstaVacia())
	require.Panics(t, func() { pila.Desapilar() })
	require.Panics(t, func() { pila.VerTope() })
}

func TestApilarDesapilar(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	for i := 0; i < 10; i++ {
		pila.Apilar(i)
	}
	for i := 9; i > -1; i-- {
		require.EqualValues(t, pila.VerTope(), i)
		require.EqualValues(t, pila.Desapilar(), i)
	}
	require.True(t, pila.EstaVacia())
	require.Panics(t, func() { pila.Desapilar() })
	require.Panics(t, func() { pila.VerTope() })
}

func TestVolumen(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	for i := 0; i < 10000; i++ {
		pila.Apilar(i)
	}
	for i := 9999; i > -1; i-- {
		require.EqualValues(t, pila.VerTope(), i)
		require.EqualValues(t, pila.Desapilar(), i)
	}
}

func TestApilarDesapilarVolumen(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	for i := 0; i < 10000; i++ {
		pila.Apilar(i)
		require.EqualValues(t, pila.VerTope(), i)
		require.EqualValues(t, pila.Desapilar(), i)
	}
	require.True(t, pila.EstaVacia())
	require.Panics(t, func() { pila.Desapilar() })
	require.Panics(t, func() { pila.VerTope() })
}

func TestLAFOStr(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[rune]()
	letras := []rune("abcdefghijklmnñopqrstuvwxyz")
	for i := 0; i < len(letras); i++ {
		pila.Apilar(letras[i])
	}
	for i := len(letras) - 1; i > -1; i-- {
		require.EqualValues(t, pila.VerTope(), letras[i])
		require.EqualValues(t, pila.Desapilar(), letras[i])
	}
	require.True(t, pila.EstaVacia())
	require.Panics(t, func() { pila.Desapilar() })
	require.Panics(t, func() { pila.VerTope() })
}

func TestStrVolumen(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[rune]()
	letras := []rune("abcdefghijklmnñopqrstuvwxyzabcdefghijklmnñopqrstuvwxyzabcdefghijklmnñopqrstuvwxyzabcdefghijklmnñopqrstuvwxyzabcdefghijklmnñopqrstuvwxyzabcdefghijklmnñopqrstuvwxyzabcdefghijklmnñopqrstuvwxyzabcdefghijklmnñopqrstuvwxyzabcdefghijklmnñopqrstuvwxyzabcdefghijklmnñopqrstuvwxyzabcdefghijklmnñopqrstuvwxyzabcdefghijklmnñopqrstuvwxyz")
	for i := 0; i < len(letras); i++ {
		pila.Apilar(letras[i])
	}
	for i := len(letras) - 1; i > -1; i-- {
		require.EqualValues(t, pila.VerTope(), letras[i])
		require.EqualValues(t, pila.Desapilar(), letras[i])
	}
	require.True(t, pila.EstaVacia())
	require.Panics(t, func() { pila.Desapilar() })
	require.Panics(t, func() { pila.VerTope() })
}

func TestLAFOFloats(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[float32]()
	numeros := []float32{3.11, 2.23, 132.23, 1.454, 1.2341}
	for i := 0; i < len(numeros); i++ {
		pila.Apilar(numeros[i])
	}
	for i := len(numeros) - 1; i > -1; i-- {
		require.EqualValues(t, pila.VerTope(), numeros[i])
		require.EqualValues(t, pila.Desapilar(), numeros[i])
	}
	require.True(t, pila.EstaVacia())
	require.Panics(t, func() { pila.Desapilar() })
	require.Panics(t, func() { pila.VerTope() })
}
