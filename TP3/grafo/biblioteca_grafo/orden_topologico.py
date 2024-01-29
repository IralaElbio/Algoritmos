from tdas.cola import Cola

def grado_entrada(grafo):
    ''' Calcula los grados de entrada para cada vertice del grafo'''
    grado_entrada = {}
    for vertice in grafo.vertices():
        grado_entrada[vertice] = grado_entrada.get(vertice, 0)
        for adyacente in grafo.adyacentes(vertice):
            grado_entrada[adyacente] = grado_entrada.get(adyacente, 0) + 1
    return grado_entrada


def topologico_grados(grafo):
    ''' Regresa un orden topologico para el grafo actual si es que lo encuentra'''
    resultado = []
    cola = Cola()
    grados_de_entrada = grado_entrada(grafo)
    for vertice, grado in grados_de_entrada.items():
        if grado == 0:
            cola.encolar(vertice)

    while not cola.esta_vacia():
        vertice = cola.desencolar()
        resultado.append(vertice)
        for adyacente in grafo.adyacentes(vertice):
            grados_de_entrada[adyacente] -= 1
            if grados_de_entrada[adyacente] == 0:
                cola.encolar(adyacente)

    if len(resultado) != len(grafo.vertices()):
        return None

    return resultado
