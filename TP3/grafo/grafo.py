import random

class Grafo:
    def __init__(self, esDirigido):
        ''' Crea un grafo vacío, ya sea dirigio o no '''
        self.grafo = {}
        self.dirigido = esDirigido

    def agregar_vertice(self, vertice):
        ''' Agrega el vertice al grafo, si el vertice
            ya existia, no hace nada'''
        if not self.pertence(vertice):
            self.grafo[vertice] = {}

    def sacar_vertice(self, vertice):
        ''' Quita el vertice del grafo,
            si el vertice no existia no hace nada'''
        if self.pertence(vertice):
            self.grafo.pop(vertice)
            for _, adyacente in self.grafo.items():
                if vertice in adyacente:
                    adyacente.pop(vertice)

    def agregar_arista(self, inicio, fin, peso = 1):
        ''' Agrega una arista entre los dos vertices con
            peso 1 por default '''
        if self.pertence(inicio) and self.pertence(fin):
            self.grafo[inicio][fin] = peso

            if self.dirigido is False:
                self.grafo[fin][inicio] = peso

    def sacar_arista(self, inicio, fin):
        ''' Quita la arista que une a los vertices,
            si los vertices no estan unidos no hace nada'''
        if self.estan_unidos(inicio, fin):
            self.grafo[inicio].pop(fin)

            if self.dirigido is False:
                self.grafo[fin].pop(inicio)

    def estan_unidos(self, vertice1, vertice2):
        """ Devuelve True o False según si los
            vertices estan unidos"""
        if self.pertence(vertice1):
            return vertice2 in self.grafo[vertice1]
        return False

    def peso(self, vertice1, vertice2):
        ''' Devuelve el peso de la arista que une a
            los dos vertices'''
        if self.estan_unidos(vertice1, vertice2):
            return self.grafo[vertice1][vertice2]

    def adyacentes(self, vertice):
        ''' Devuelve una lista con los adyacentes
            del vertices'''
        if vertice in self.grafo:
            adyacentes = self.grafo[vertice].keys()
            if adyacentes is None:
                return []
            return adyacentes

    def vertices(self):
        ''' Devuelve una lista con los vertices
            del grafo'''
        return self.grafo.keys()

    def obtener_vertice_aleatorio(self):
        ''' Devuelve un vertice random del grafo'''
        return random.choice(list(self.grafo.keys()))

    def pertence(self, vertice):
        ''' Devuelve True o False dependiendo si el vertice
            pertenece al grafo'''
        return vertice in self.grafo
