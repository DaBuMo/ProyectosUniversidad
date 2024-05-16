import random

COLOR_BASE = 0
FILA_INICIAL = 0
COLUMNA_INICIAL = 0

FILAS = 0
COLUMNAS = 1

class Flood:
    """
    Clase para administrar un tablero de N colores.
    """

    def __init__(self, alto, ancho):
        """
        Genera un nuevo Flood de un mismo color con las dimensiones dadas.

        Argumentos:
            alto, ancho (int): Tamaño de la grilla.
        """
        self.alto = alto
        self.ancho = ancho
        self.tablero = [[COLOR_BASE for n in range(ancho)] for n in range(alto)]
        self.cantidad_colores = 0

    def mezclar_tablero(self, n_colores):
        """
        Asigna de forma completamente aleatoria hasta `n_colores` a lo largo de
        las casillas del tablero.

        Argumentos:
            n_colores (int): Cantidad maxima de colores a incluir en la grilla.
        """
        self.cantidad_colores = n_colores
        for fila in range(self.alto):
            self.tablero[fila] = random.choices([n for n in range(n_colores)],k=self.ancho)


    def obtener_color(self, fil, col):
        """
        Devuelve el color que se encuentra en las coordenadas solicitadas.

        Argumentos:
            fil, col (int): Posiciones de la fila y columna en la grilla.

        Devuelve:
            Color asignado.
        """
        return self.tablero[fil][col]


    def obtener_posibles_colores(self):
        """
        Devuelve una secuencia ordenada de todos los colores posibles del juego.
        La secuencia tendrá todos los colores posibles que fueron utilizados
        para generar el tablero, sin importar cuántos de estos colores queden
        actualmente en el tablero.

        Devuelve:
            iterable: secuencia ordenada de colores.
        """
        return [n for n in range(self.cantidad_colores)]


    def dimensiones(self):
        """
        Dimensiones de la grilla (filas y columnas)

        Devuelve:
            (int, int): alto y ancho de la grilla en ese orden.
        """
        return (self.alto,self.ancho)


    def cambiar_color(self, color_nuevo):
        """
        Asigna el nuevo color al Flood de la grilla. Es decir, a todas las
        coordenadas que formen un camino continuo del mismo color comenzando
        desde la coordenada origen en (0, 0) se les asignará `color_nuevo`

        Argumentos:
            color_nuevo: Valor del nuevo color a asignar al Flood.
        """
        color_viejo = self.obtener_color(FILA_INICIAL, COLUMNA_INICIAL)
        if color_viejo == color_nuevo:
            return
        max_fil,max_col = self.dimensiones()

        self._cambiar_color(color_viejo,color_nuevo,FILA_INICIAL,COLUMNA_INICIAL,max_fil,max_col)

    def _cambiar_color(self,color_viejo,color_nuevo,fil,col,max_fil,max_col):
        """
        Wrapper de la funcion cambiar_color
        """
        if fil < 0 or fil >= max_fil or col < 0 or col >= max_col or self.tablero[fil][col] != color_viejo:
            return
        self.tablero[fil][col] = color_nuevo
        self._cambiar_color(color_viejo,color_nuevo,fil + 1,col,max_fil,max_col)       
        self._cambiar_color(color_viejo,color_nuevo,fil - 1 ,col,max_fil,max_col)      
        self._cambiar_color(color_viejo,color_nuevo,fil,col + 1 ,max_fil,max_col)
        self._cambiar_color(color_viejo,color_nuevo,fil,col - 1,max_fil,max_col)
                
    def clonar(self):
        """
        Devuelve:
            Flood: Copia del Flood actual
        """
        tab_aux = []

        for i in range(len(self.tablero)):
            tab_aux.append([])
            for p in range(len(self.tablero[i])):
                tab_aux[i].append(self.tablero[i][p])
            
        return tab_aux

    def esta_completado(self):
        """
        Indica si todas las coordenadas de grilla tienen el mismo color

        Devuelve:
            bool: True si toda la grilla tiene el mismo color
        """
        for alto in range(self.alto):
            for ancho in range(self.ancho - 1):
                if self.obtener_color(alto,ancho) != self.obtener_color(alto,ancho + 1):
                    return False
        return True

    def contar_por_color(self):
        """
        Wrapper de la funcion _contar_por_color
        """
        color = self.obtener_color(FILA_INICIAL, COLUMNA_INICIAL)
        max_fil,max_col = self.dimensiones()
        contador = [0]
        visitados = set()
        self._contar_por_color(color,FILA_INICIAL,COLUMNA_INICIAL,max_fil,max_col,contador,visitados)
        return contador[0]

    def _contar_por_color(self,color,fil,col,max_fil,max_col,contador,visitados):
        """
        Indica cuantas cordenadas de la grilla tienen el mismo color siguiendo el sentido del flood, comparado al color inicial

        Devuelve:
            int: Cantidad de grillas conectadas desde el inicio con el mismo color al origen del tablero (Fil = 0, Col = 0)
        """
        if fil < 0 or fil >= max_fil or col < 0 or col >= max_col or self.obtener_color(fil,col) != color or (fil,col) in visitados:
            return
        contador[0] +=  1
        visitados.add((fil,col))
        self._contar_por_color(color,fil + 1,col,max_fil,max_col,contador,visitados)       
        self._contar_por_color(color,fil - 1 ,col,max_fil,max_col,contador,visitados)      
        self._contar_por_color(color,fil,col + 1 ,max_fil,max_col,contador,visitados)
        self._contar_por_color(color,fil,col - 1,max_fil,max_col,contador,visitados)



