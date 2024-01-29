def peso_total(grafo):
    ''' Calcula el peso total de todas las aristas del grafo'''
    total = 0
    visitados = set()
    for v in grafo.vertices():
        visitados.add(v)
        for adyacente in grafo.adyacentes(v):
            if adyacente not in visitados:
                visitados.add(adyacente)
                total += grafo.peso(v, adyacente)
    return total


def cantidad_aristas_grafo_dirigido(grafo):
    ''' Regresa la cantidad de aristas que hay en un grafo dirigido'''
    total = 0
    for v in grafo.vertices():
        total += len(grafo.adyacentes(v))
    return total
