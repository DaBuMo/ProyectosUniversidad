import random

class Grafo():
    def __init__(self) -> None:
        self._conexiones = dict()
    
    def agregar_vertice(self,vertice:str) -> None:
        if vertice not in self._conexiones:
            self._conexiones[vertice] = {}

    def sacar_vertice(self,vertice:str) -> None:
        if not self.pertenece(vertice):
            raise KeyError(f"El vertice: {vertice}, no se encuentra almacenado en las conexiones")

        self._conexiones.pop(vertice)
        self._cantidad -= 1
        for w in self._conexiones.keys():
            if self.existe_arista(w,vertice):
                self.eliminar_arista(w,vertice)

    def obtener_vertices(self) -> list:

        return list(self._conexiones.keys())

    def agregar_arista(self,vertice1:str,vertice2:str) -> None:
        for vertice in (vertice1,vertice2):
            if not self.pertenece(vertice):
                raise KeyError(f"El vertice: {vertice}, no se encuentra almacenado en las conexiones")
        
        self._conexiones[vertice1][vertice2] = True
    
    def eliminar_arista(self,vertice1:str,vertice2:str) -> str:
        if not self.pertenece(vertice1):
            raise KeyError(f"El vertice: {vertice1}, no se encuentra almacenado en las conexiones")
        
        if vertice2 not in self._conexiones[vertice1]:
            raise KeyError(f"La arista entre el vertice {vertice1} y el vertice {vertice2} no existe")
        
        return self._conexiones[vertice1].pop(vertice2)
    
    def existe_arista(self,vertice1:str,vertice2:str) -> bool:
        if not self.pertenece(vertice1):
            raise KeyError(f"El vertice: {vertice1}, no se encuentra almacenado en las conexiones")
        
        if not vertice2 in self._conexiones[vertice1]:
            return False
        
        return True

    def obtener_adyacentes(self,vertice:str) ->list:  
        adyacentes = []
        
        for adyacente in self._conexiones[vertice]:
            adyacentes.append(adyacente)
        return adyacentes
        
    def pertenece(self,vertice:str) -> bool:
        return vertice in self._conexiones
    
    def vertice_random(self) -> any:
        return random.choice(list(self._conexiones.keys()))