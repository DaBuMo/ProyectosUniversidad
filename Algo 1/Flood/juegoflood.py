from flood import Flood,COLUMNA_INICIAL,FILA_INICIAL
from pila import Pila
from cola import Cola
import random


class JuegoFlood:
    """
    Clase para administrar un Flood, junto con sus estados y acciones
    """

    def __init__(self, alto, ancho, n_colores):
        """
        Genera un nuevo JuegoFlood, el cual tiene un Flood y otros
        atributos para realizar las distintas acciones del juego.

        Argumentos:
            alto, ancho (int): Tamaño de la grilla del Flood.
            n_colores: Cantidad maxima de colores a incluir en la grilla.
        """
        self.n_movimientos = 0
        self.flood = Flood(alto, ancho)
        self.flood.mezclar_tablero(n_colores)
        self.anteriores = Pila()
        self.siguientes = Pila()
        self.pasos_solucion = Cola()
        self.mejor_n_movimientos, _ = self._calcular_movimientos()



    def cambiar_color(self, color):
        """
        Realiza la acción para seleccionar un color en el Flood, sumando a la
        cantidad de movimientos realizados y manejando las estructuras para
        deshacer y rehacer

        Argumentos:
            color (int): Nuevo color a seleccionar
        """

        if color != self.flood.obtener_color(FILA_INICIAL,COLUMNA_INICIAL):
            self.anteriores.apilar(self.flood.clonar())
            self.n_movimientos += 1
            self.flood.cambiar_color(color)
        
        if not self.siguientes.esta_vacia():
            self.siguientes = Pila()

        if not self.pasos_solucion.esta_vacia() and self.pasos_solucion.ver_frente() == color:
            self.pasos_solucion.desencolar()
        else:
            self.pasos_solucion = Cola()

    def deshacer(self):
        """
        Deshace el ultimo movimiento realizado si existen pasos previos,
        manejando las estructuras para deshacer y rehacer.
        """

        if not self.anteriores.esta_vacia():
            
            self.n_movimientos -= 1
            self.siguientes.apilar(self.flood.clonar())
            self.flood.tablero = self.anteriores.desapilar()
            self.pasos_solucion = Cola()

    def rehacer(self):
        """
        Rehace el movimiento que fue deshecho si existe, manejando las
        estructuras para deshacer y rehacer.
        """
            
        if not self.siguientes.esta_vacia():
            
            self.n_movimientos += 1
            self.anteriores.apilar(self.flood.clonar())
            self.flood.tablero = self.siguientes.desapilar()
            self.pasos_solucion = Cola()

    def _calcular_movimientos(self):
        """
        Wrapper de __calcular movimientos
        """
        cant_mov = 0
        col_soluc = Cola()
        copia_tab = self.flood.clonar()

        return self.__calcular_movimientos(cant_mov,copia_tab,col_soluc)

    def __calcular_movimientos(self,cant_mov,copia_tab,col_soluc):
        """
        Realiza una solución de pasos contra el Flood actual (en una Cola)
        y devuelve la cantidad de movimientos que llevó a esa solución.

        COMPLETAR CON EL CRITERIO DEL ALGORITMO DE SOLUCIÓN.

        Devuelve:
            int: Cantidad de movimientos que llevó a la solución encontrada.
            Cola: Pasos utilizados para llegar a dicha solución
        """
        if self.esta_completado():
            self.flood.tablero = copia_tab
            return cant_mov, col_soluc
        
        pos_val = {}
        max_color = None
        max_val = None

        for i in self.obtener_posibles_colores():    
            if self.obtener_color(FILA_INICIAL,COLUMNA_INICIAL) == i:
                continue
            tab_copia = self.flood.clonar()
            self.flood.cambiar_color(i)
            val = self.flood.contar_por_color()
            self.flood.tablero = tab_copia
            pos_val[i] = val

        for color, aparicion in pos_val.items():
            if max_val:
                if aparicion > max_val:
                    max_color = color
                    max_val = aparicion
            else:
                max_color = color
                max_val = aparicion

        col_soluc.encolar(max_color)
        self.flood.cambiar_color(max_color)

        cant_mov += 1
        return self.__calcular_movimientos(cant_mov,copia_tab,col_soluc)

    def hay_proximo_paso(self):
        """
        Devuelve un booleano indicando si hay una solución calculada
        """
        return not self.pasos_solucion.esta_vacia()


    def proximo_paso(self):
        """
        Si hay una solución calculada, devuelve el próximo paso.
        Caso contrario devuelve ValueError

        Devuelve:
            Color del próximo paso de la solución
        """
        return self.pasos_solucion.ver_frente()


    def calcular_nueva_solucion(self):
        """
        Calcula una secuencia de pasos que solucionan el estado actual
        del flood, de tal forma que se pueda llamar al método `proximo_paso()`
        """
        _, self.pasos_solucion = self._calcular_movimientos()


    def dimensiones(self):
        return self.flood.dimensiones()


    def obtener_color(self, fil, col):
        return self.flood.obtener_color(fil, col)


    def obtener_posibles_colores(self):
        return self.flood.obtener_posibles_colores()


    def esta_completado(self):
        return self.flood.esta_completado()
