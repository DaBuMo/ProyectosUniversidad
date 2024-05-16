package cola_test

import (
	TDACola "tdas/cola"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestColaVacia(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	require.True(t, cola.EstaVacia())
}
func TestPanicsCola(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	require.True(t, cola.EstaVacia())
	require.Panics(t, func() { cola.VerPrimero() })
	require.Panics(t, func() { cola.Desencolar() })
}
func TestVerPrimeroDesencolar(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[rune]()
	letras := []rune("abcdefghijklmnñopqrstuvwxyz")
	for i := 0; i < len(letras); i++ {
		cola.Encolar(letras[i])
	}
	for i := 0; i < len(letras); i++ {
		require.EqualValues(t, cola.VerPrimero(), letras[i])
		require.EqualValues(t, cola.Desencolar(), letras[i])
	}
	require.True(t, cola.EstaVacia())
}

func TestVolumen(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	for i := 0; i < 100000; i++ {
		cola.Encolar(i)
	}
	for i := 0; i < 100000; i++ {
		require.EqualValues(t, cola.Desencolar(), i)
	}
	require.True(t, cola.EstaVacia())
}

func TestDesencolarBorde(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	for i := 0; i < 1000; i++ {
		cola.Encolar(i)
	}
	for i := 0; i < 1000; i++ {
		require.EqualValues(t, cola.Desencolar(), i)
	}
	require.True(t, cola.EstaVacia())
}
func TestEncolarDesencolar(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	for i := 0; i < 1000; i++ {
		cola.Encolar(i)
		require.EqualValues(t, cola.Desencolar(), i)
	}
	require.True(t, cola.EstaVacia())
	require.Panics(t, func() { cola.VerPrimero() })
	require.Panics(t, func() { cola.Desencolar() })
}

func TestDesencolarComoNueva(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	for i := 0; i < 10000; i++ {
		cola.Encolar(i)
	}
	for !cola.EstaVacia() {
		cola.Desencolar()
	}
	require.Panics(t, func() { cola.VerPrimero() })
	require.Panics(t, func() { cola.Desencolar() })
}
