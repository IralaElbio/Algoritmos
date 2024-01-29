import heapq
from grafo.grafo import Grafo


def tendido_minimo_prim(grafo, origen):
    ''' Calcula y regresa un árbol de tendido minimo para el grafo actual '''
    visitados = set()
    arbol_minimo = Grafo(True)
    heap = []
    visitados.add(origen)
    for adyacente in grafo.adyacentes(origen):
        heapq.heappush(
            heap, (grafo.peso(origen, adyacente), origen, adyacente))

    while len(heap) > 0:
        _, vertice, adyacente = heapq.heappop(heap)
        arbol_minimo.agregar_vertice(vertice)

        if adyacente not in visitados:
            arbol_minimo.agregar_vertice(adyacente)

            arbol_minimo.agregar_arista(
                vertice, adyacente, grafo.peso(vertice, adyacente))

            visitados.add(adyacente)
            for adyacente2 in grafo.adyacentes(adyacente):
                if adyacente2 not in visitados:
                    heapq.heappush(
                        heap, (grafo.peso(adyacente, adyacente2), adyacente, adyacente2))
    return arbol_minimo
