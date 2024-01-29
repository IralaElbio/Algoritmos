from tdas.cola import Cola
from tdas.pila import Pila
import sys

def ciclo_euler(grafo, inicio):
    ''' Calcula un ciclo de euler para el grafo actual si es que tiene.
        El algoritmo obtiene todas las aristas para cada vertice y mientras
        que haya vertices con aristas hace un llamado recursivo a un dfs donde
        saca una arista para el vertice en cuestion y lo agraga a un camino provisional
        '''
    sys.setrecursionlimit(2000)
    
    if not hay_ciclo_euleriano(grafo):
        return None, None

    aristas = obtener_aristas(grafo)

    camino_c = Pila()

    camino_c.apilar(inicio)

    resultado = []

    peso_total = 0
    while not camino_c.esta_vacia():

        vertice = camino_c.ver_tope()
        # verifico si el vertice que esta en el camino aun tiene aristas que lo unen con otros vertices
        if len(aristas[vertice]) > 0:
            peso_total += dfs_ciclo_euler(grafo, vertice, camino_c, aristas, vertice, 0)
        else:
            # si no tiene lo saco del camino y lo agrego al resultado
            camino_c.desapilar()
            resultado.append(vertice)

    return resultado, peso_total

def dfs_ciclo_euler(grafo, vertice, camino_c, aristas, vertice_inicial, suma_peso):
    # saco la primer arista del vertice y obtengo un adyacente
    adyacente = aristas[vertice].pop(0)

    suma_peso += grafo.peso(vertice, adyacente)
    # lo elimino de las aristas del adyacente para aseguarme no pasar por la misma arista
    aristas[adyacente].remove(vertice)

    camino_c.apilar(adyacente)

    if adyacente == vertice_inicial:
        return suma_peso

    return dfs_ciclo_euler(grafo, adyacente, camino_c, aristas, vertice_inicial, suma_peso)


def hay_ciclo_euleriano(grafo):
    ''' Regresa si para el grafo actual existe un ciclo euleriano '''
    return tiene_vertices_pares(grafo) and tiene_una_componente_conexa(grafo)


def obtener_aristas(grafo):
    ''' Regresa un diccionario en representacion para las aristas del grafo, donde la clave
        es el vertice de donde sale la arita y el valor una lista con todos los vertices que
        tienen una arista con el vertice clave'''
    aristas = {}
    visitados = set()
    for v in grafo.vertices():
        aristas[v] = []
        for w in grafo.adyacentes(v):
            if v not in visitados:
                aristas[v].append(w)
        visitados.add(v)
    return aristas


def tiene_vertices_pares(grafo):
    ''' Devuelve True si el grafo actual tiene solo vertices de grado par,
        caso contrario False'''
    for vertice in grafo.vertices():
        if len(grafo.adyacentes(vertice)) % 2 != 0:
            return False
    return True


def tiene_una_componente_conexa(grafo):
    ''' Devuelve true si el grafo tiene una sola componente conexa, caso contario False'''
    visitados = set()

    origen = grafo.obtener_vertice_aleatorio()

    cola = Cola()

    cola.encolar(origen)

    while not cola.esta_vacia():
        vertice = cola.desencolar()

        for adyacente in grafo.adyacentes(vertice):
            if adyacente not in visitados:
                visitados.add(adyacente)

                cola.encolar(adyacente)

    for vertice in grafo.vertices():
        if vertice not in visitados:
            return False
    return True
