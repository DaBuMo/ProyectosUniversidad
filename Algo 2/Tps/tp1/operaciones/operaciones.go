package operaciones

import (
	"errors"
	"math"
	"strconv"
	"strings"
	TDAPila "tdas/pila"
)

const (
	_NUMERO               int    = -1
	_1_OPERANDO           int    = 1
	_2_OPERANDO           int    = 2
	_3_OPERANDO           int    = 3
	_LETRAS               string = "abcdefghijklmnopqrstuvwyxz"
	_OPERADORES_1_DIGITO  string = "sqrt"
	_OPERADORES_2_DIGITOS string = "+-*/^log"
	_TERNARIO             string = "?"
	_SUMA                 string = "+"
	_DIVISION             string = "/"
	_RAIZ                 string = "sqrt"
	_RESTA                string = "-"
	_EXPONENCIAL          string = "^"
	_MULTIPLICACION       string = "*"
	_LOGARITMO            string = "log"
)

type Calculadora struct {
	operador TDAPila.Pila[int64]
}

func CrearCalculadora() *Calculadora {
	p := TDAPila.CrearPilaDinamica[int64]()
	return &Calculadora{operador: p}
}

func (c *Calculadora) IniciarCalculo(calculo string) (int64, error) {
	c.operador = TDAPila.CrearPilaDinamica[int64]()
	valores := strings.Fields(calculo)

	for _, elem := range valores {
		tipo, err := tipoOperador(elem)

		if err != nil {
			return 0, err
		}

		if tipo == _NUMERO {
			c.operador.Apilar(StringAInt64(elem))
		} else {
			err := c.operar(elem, tipo)

			if err != nil {
				return 0, err
			}
		}
	}

	return c.mostrarResultado()
}

func (c *Calculadora) mostrarResultado() (int64, error) {
	resultado := c.operador.Desapilar()

	if c.operador.EstaVacia() {
		return resultado, nil
	}

	return 0, errors.New("quedaron operandos sin operar")
}

func (c *Calculadora) operar(operacion string, tipo int) error {
	operandos := make([]int64, 3)

	for i := 0; i < tipo; i++ {
		if !c.operador.EstaVacia() {
			operandos[i] = c.operador.Desapilar()
		} else {
			return errors.New("faltan operandos")
		}
	}

	return c.calcular(operacion, operandos[0], operandos[1], operandos[2])
}

func (c *Calculadora) calcular(operacion string, operando1 int64, operando2 int64, operando3 int64) error {
	var err error = nil

	switch operacion {
	case _SUMA:
		c.operador.Apilar(operando2 + operando1)

	case _RESTA:
		c.operador.Apilar(operando2 - operando1)

	case _MULTIPLICACION:
		c.operador.Apilar(operando2 * operando1)

	case _DIVISION:
		if operando1 == 0 {
			err = errors.New("division de 0")
		} else {
			c.operador.Apilar(operando2 / operando1)
		}

	case _EXPONENCIAL:
		if operando1 < 0 {
			err = errors.New("exponencial negativo")
		} else {
			potencia := int64(math.Pow(float64(operando2), float64(operando1)))
			c.operador.Apilar(potencia)
		}

	case _LOGARITMO:
		if operando1 < 2 {
			err = errors.New("logaritmo menor a 2")
		} else {
			base := math.Log2(float64(operando1))
			resultado := int64(math.Log2(float64(operando2)) / base)
			c.operador.Apilar(resultado)
		}

	case _TERNARIO:
		if operando3 > 0 {
			c.operador.Apilar(operando2)
		} else {
			c.operador.Apilar(operando1)
		}

	case _RAIZ:
		if operando1 < 0 {
			err = errors.New("raiz negativa")
		} else {
			raiz := int64(math.Sqrt(float64(operando1)))
			c.operador.Apilar(raiz)
		}
	}

	return err
}

func StringAInt64(s string) int64 {
	var valor int64
	valor, _ = strconv.ParseInt(s, 10, 64)
	return valor
}

func tipoOperador(s string) (int, error) {
	valor := strings.ToLower(s)
	_, err := strconv.ParseInt(valor, 10, 64)

	if strings.Contains(_OPERADORES_1_DIGITO, valor) {
		return _1_OPERANDO, nil
	} else if strings.Contains(_OPERADORES_2_DIGITOS, valor) {
		return _2_OPERANDO, nil
	} else if strings.Contains(_TERNARIO, valor) {
		return _3_OPERANDO, nil
	} else if strings.Contains(_LETRAS, valor) || err != nil {
		return 0, errors.New("valor invalido para operar")
	}

	return _NUMERO, nil
}
