#!/usr/bin/python3
import csv
import sys
from biblioteca import minimos_seguimientos,mas_importantes,ciclos_hasta,persecucion_mas_rapida,divulgar_rumor,comunidades,cfc,SEPARADOR_LISTA
from grafo import Grafo

VERTICE_1 = 0
VERTICE_2 = 1
DIRECCION = 1
DELIMITADOR = "\t"
SALTO_LINEA = "\r\n"
LINEA_VACIA = "\n"
SEPARADOR_ENCUBIERTOS = ","
SIGUIENTE = " -> "

ERROR_APERTURA = "Error al abrir el archivo"
ERROR_PARAMETROS = "Error de parametros"
ERROR_COMANDO = "Error, comando invalido"
ARCHIVOS_MINIMOS = 2

PERSECUCION_RAPIDA = "persecucion"
MAS_IMPORTANTES = "mas_imp"
SEGUIMIENTOS_MINIMOS = "min_seguimientos"
COMUNIDADES = "comunidades"
DIVULGAR_RUMOR = "divulgar"
CICLO_MAS_CORTO = "divulgar_ciclo"
COMPONENTE_FUERTEMENTE_CONEXAS = "cfc"

COMANDO = 0
PARAMETRO_1 = 1
PARAMETRO_2 = 2

def stdout(*a):
    print(*a,file=sys.stdout)

def cargar_vertices_arista(grafo:Grafo,vertices:list) -> None:
    for vertice in vertices:
        if not grafo.pertenece(vertice):
            grafo.agregar_vertice(vertice)
    grafo.agregar_arista(vertices[VERTICE_1],vertices[VERTICE_2])

def cargar_grafo(direccion:str) -> Grafo:
    grafo = Grafo()
    try:
        with open(direccion) as file:
            tsv_file = csv.reader(file,delimiter=DELIMITADOR)
            for linea in tsv_file:
                cargar_vertices_arista(grafo,linea)
            return grafo
    except:
        stdout(ERROR_APERTURA)
        return
    
def main():
    if len(sys.argv) < ARCHIVOS_MINIMOS:
        stdout(ERROR_PARAMETROS)
        return
    
    grafo = cargar_grafo(sys.argv[DIRECCION])
    scores = []

    for linea in sys.stdin:
        if linea == SALTO_LINEA:
            continue

        directiva = linea.strip(SALTO_LINEA).split(" ")
        comando = directiva[COMANDO]

        if comando == PERSECUCION_RAPIDA:
            encubiertos = directiva[PARAMETRO_1].split(SEPARADOR_ENCUBIERTOS)
            cant_buscados = directiva[PARAMETRO_2]
            stdout(persecucion_mas_rapida(grafo,encubiertos,int(cant_buscados),scores))

        elif comando == MAS_IMPORTANTES:
            cant_imp = directiva[PARAMETRO_1]
            scores,mensaje = mas_importantes(grafo,int(cant_imp),scores)
            stdout(mensaje)

        elif comando == COMPONENTE_FUERTEMENTE_CONEXAS:
            componentes = cfc(grafo)
            for i in range(len(componentes)):
                resul = "CFC " + str(i+1) +": "+SEPARADOR_LISTA.join(componentes[i])
                stdout(resul)

        elif comando == SEGUIMIENTOS_MINIMOS:
            origen = directiva[PARAMETRO_1]
            destino = directiva[PARAMETRO_2]
            stdout(minimos_seguimientos(grafo,origen,destino))

        elif comando == COMUNIDADES:
            largo_min = directiva[PARAMETRO_1]
            conjuntos = comunidades(grafo,int(largo_min))
            for i in range(len(conjuntos)):
                resul = "Comunidad " + str(i+1)+ ": "+SEPARADOR_LISTA.join(conjuntos[i])
                stdout(resul)

        elif comando == DIVULGAR_RUMOR:
            delincuente = directiva[PARAMETRO_1]
            saltos = directiva[PARAMETRO_2]
            stdout(divulgar_rumor(grafo,delincuente,int(saltos)))

        elif comando == CICLO_MAS_CORTO:
            delincuente = directiva[PARAMETRO_1]
            stdout(ciclos_hasta(grafo,delincuente))
            
        else:
            stdout(ERROR_COMANDO)


main()
	
