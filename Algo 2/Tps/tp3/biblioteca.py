from grafo import Grafo
import collections
import random
import heapq

CICLO_INVALIDO = "No se encontro recorrido"
ERROR_SEGUIMIENTO = "Seguimiento imposible"
SIGUIENTE = " -> "
SEPARADOR_LISTA = ", "
VALOR_BASE = 1
VERTICE = 1
CONTADOR = 0
CANTIDAD_PAGE_RANK = 41
APARICIONES = 0
LABEL = 1
APARICION_LABEL = 1
RANGO_LABEL = 2
AMORTIGUACION = 0.85
# Funciones Principales
def minimos_seguimientos(grafo:Grafo,origen:str,hasta:str) -> str:
    caminos = caminos_minimos(grafo,origen,hasta)
    if caminos == ERROR_SEGUIMIENTO:
        return ERROR_SEGUIMIENTO
    else:
        return ""+SIGUIENTE.join(caminos)

def persecucion_mas_rapida(grafo:Grafo,encubiertos:list,cantidad:int,scores:list):
    if not scores or len(scores) < int(cantidad):
        scores = page_rank(grafo,cantidad)
    heap = []
    perseguidos = scores[:cantidad]
    for v in encubiertos:
        for w in perseguidos:
            camino = caminos_minimos(grafo,v,w)
            if camino != ERROR_SEGUIMIENTO:
                heapq.heappush(heap,(len(camino),camino))
    persecucion = heapq.heappop(heap)[VERTICE]
    return ""+SIGUIENTE.join(persecucion)

def mas_importantes(grafo:Grafo,cant:int,scores:list):
    if not scores or len(scores) < cant:
        scores = page_rank(grafo,cant)
    return scores,""+SEPARADOR_LISTA.join(scores)

def divulgar_rumor(grafo:Grafo,origen:str,saltos:int):
    cola = collections.deque()
    cola.append(origen)
    visitados = set()
    visitados.add(origen)
    resul= []
    contador = 0
    while contador != saltos:
        cola_aux = collections.deque()  

        while len(cola) > 0:

            v = cola.popleft()
            for w in grafo.obtener_adyacentes(v):
                if w not in visitados:
                    visitados.add(w)
                    resul.append(w)
                    cola_aux.append(w)

        cola = cola_aux.copy()
        contador += 1
  
    return ", ".join(resul)

def cfc(grafo:Grafo):
    visitados = set()
    apilados = set()
    orden = {}
    mas_bajo = {}
    cfcs = []
    contador_global = [0]
    pila = collections.deque()

    for v in grafo.obtener_vertices():
        if v not in visitados:
            dfs_cfc(grafo,v,visitados,orden,mas_bajo,pila,apilados,cfcs,contador_global)
    return cfcs

def comunidades(grafo:Grafo,largo_minimo:int):
    labels = label_propagation(grafo)
    posibles_comunidades = {}
    comunidades = []
    for vertice,label in labels.items():
        label = str(label)
        if label not in posibles_comunidades:
            posibles_comunidades[label] = [vertice]
        else:
            posibles_comunidades[label].append(vertice)
            if len(posibles_comunidades[label]) == largo_minimo and posibles_comunidades[label] not in comunidades:
                comunidades.append(posibles_comunidades[label])
    return comunidades

def ciclos_hasta(grafo:Grafo,buscado:str):
    if not grafo.pertenece(buscado):
        return CICLO_INVALIDO

    mas_corto = busqueda_ciclos(grafo,buscado)

    if mas_corto:
        return ""+SIGUIENTE.join(mas_corto)
    else:
        return CICLO_INVALIDO 
    
# Funciones auxiliares
def reconstruir_camino(padres:dict, buscado:str) -> list:
    w = buscado
    lista = collections.deque()
    while w != None:
        lista.appendleft(w)
        w = padres[w]
    return lista


def dfs_cfc(grafo:Grafo,vertice:str,visitados:set,orden:dict,mas_bajo:dict,pila:collections.deque,apilados:set,cfcs:list,contador:list):
    orden[vertice] = mas_bajo[vertice] = contador[CONTADOR]
    contador[CONTADOR] +=1
    visitados.add(vertice)
    pila.appendleft(vertice)
    apilados.add(vertice)
    for w in grafo.obtener_adyacentes(vertice):
        if w not in visitados:
            dfs_cfc(grafo,w,visitados,orden,mas_bajo,pila,apilados,cfcs,contador)
        if w in apilados:
            mas_bajo[vertice] = min(mas_bajo[vertice],mas_bajo[w])

    if orden[vertice] == mas_bajo[vertice]:
        nueva_cfc = []
        while True:
            w = pila.popleft()
            apilados.remove(w)
            nueva_cfc.append(w)
            if w == vertice:
                break
        cfcs.append(nueva_cfc)

def obtener_entradas(grafo:Grafo) -> dict:
    entradas = {}
    for v in grafo.obtener_vertices():
        entradas[v] = list()
    for v in grafo.obtener_vertices():
        for w in grafo.obtener_adyacentes(v):
            entradas[w].append(v)
    return entradas

def max_frec(labels:dict,entradas:dict,valor)->str:
    resul = {}
    heap = []
    for v in entradas[valor]:
        label = labels[v]
        if label in resul:
            resul[label] +=1
        else:
            resul[label] = 1
    for key,value in resul.items():
        heapq.heappush(heap,(value,key))

    return heapq.nlargest(1,heap)[APARICIONES][LABEL]

def label_propagation(grafo:Grafo):
    labels = {}
    entradas = obtener_entradas(grafo)
    contador = 0    
    vertices = grafo.obtener_vertices()
    for v in vertices:
        labels[v] = contador
        contador += APARICION_LABEL
    while len(vertices) > len(vertices)/RANGO_LABEL:
        v = random.choice(vertices)
        vertices.remove(v)
        if len(entradas[v]) > 0: 
            labels[v] = max_frec(labels,entradas,v)
    return labels

def busqueda_ciclos(grafo:Grafo,origen:str):
    heap =[(0,origen,[origen])]
    visitados = set()
    while heap:
        (peso,actual,camino) = heapq.heappop(heap)

        if actual == origen and len(camino) > 1:
            return camino
        
        if actual not in visitados:
            visitados.add(actual)
            for w in grafo.obtener_adyacentes(actual):
                if w not in camino or w == origen:
                    nuevo_camino = camino + [w]
                    heapq.heappush(heap,(peso+1,w,nuevo_camino))

    return []


def caminos_minimos(grafo:Grafo, desde:str, hasta:str) -> list:
    cola = collections.deque()
    padres = {}
    visitados = set()
    padres[desde] = None
    visitados.add(desde)
    cola.append(desde)
    while len(cola) != 0:
        v = cola.popleft()
        for w in grafo.obtener_adyacentes(v):
            if w not in visitados:
                padres[w] = v
                visitados.add(w)
                if w == hasta:
                    return reconstruir_camino(padres,w)
                cola.append(w)
    return ERROR_SEGUIMIENTO

def page_rank(grafo:Grafo,buscados:int) -> collections:
    entradas = obtener_entradas(grafo)
    vertices = grafo.obtener_vertices()
    cantidad = len(vertices)
    centralidad = {}
    heap = []
    resul = []

    for v in grafo.obtener_vertices():
        centralidad[v] = (1/cantidad)

    for _ in range(CANTIDAD_PAGE_RANK):
        dic_aux = {}
        for vertice in vertices:
            total = 0
            for entrante in entradas[vertice]:
                valores  = grafo.obtener_adyacentes(entrante)
                if vertice in valores:
                    total += centralidad[entrante] /len(valores)

            dic_aux[vertice] = (1- AMORTIGUACION) / cantidad + AMORTIGUACION * total
        centralidad = dic_aux

    for key,value in centralidad.items():
        heapq.heappush(heap,(value,key))
    for _,v in heapq.nlargest(cantidad,heap):
        resul.append(v)

    return resul
