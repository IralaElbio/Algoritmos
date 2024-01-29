import heapq

def camino_minimo_dijkstra(grafo, origen, fin):
    ''' Calcula el camino minimo (si es que existe) para un grafo utilizando el algoritmo de dijkstra '''
    if hay_arista_negativa(grafo) or origen not in grafo.vertices():
        return None, None
    distancia = {}
    padres = {}

    for v in grafo.vertices():
        distancia[v] = float("inf")

    distancia[origen] = 0
    padres[origen] = None
    heap = [(0, origen)]
    camino = []
    peso = 0

    while len(heap) > 0:
        _, vertice = heapq.heappop(heap)
        peso += distancia[vertice]
        camino.append(vertice)
        if vertice == fin:
            return padres, distancia

        for adyacente in grafo.adyacentes(vertice):

            distacia_por_este_camino = distancia[vertice] + grafo.peso(vertice, adyacente)

            if distacia_por_este_camino < distancia[adyacente]:

                distancia[adyacente] = distacia_por_este_camino

                padres[adyacente] = vertice

                heapq.heappush(heap, (distacia_por_este_camino, adyacente))
    return None, None


def hay_arista_negativa(grafo):
    ''' Duvuelve True si encuentra una arista negativa en el grafo '''
    for v in grafo.vertices():
        for w in grafo.adyacentes(v):
            if grafo.peso(v, w) < 0:
                return True
    return False
