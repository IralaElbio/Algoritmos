#!/usr/bin/python3
import sys
from funciones import *

IR = "ir"
ITINERARIO = "itinerario"
VIAJE = "viaje"
REDUCIR_CAMINOS = "reducir_caminos"


def main():
    ruta = sys.argv[1]
    ciudades, mapa = iniciar_programa(ruta)
    while True:
        try:
            entrada = input()
        except:
            break

        comando, parametro = leer_entrada(entrada)

        if comando == IR:
            desde, hasta, ruta = parametro.split(", ")
            camino, tiempo = obtener_camino(ciudades, desde, hasta, ruta, mapa)
            resultado = armar_recorrido(camino, tiempo)

        elif comando == ITINERARIO:
            camino = obtener_camino_ordenado(parametro, mapa)
            resultado = armar_recorrido(camino)

        elif comando == VIAJE:
            desde, ruta = parametro.split(", ")
            camino, tiempo = obtener_recorrido_completo(ciudades, desde, ruta, mapa)
            resultado = armar_recorrido(camino, tiempo)

        elif comando == REDUCIR_CAMINOS:
            horas = camino_total(ciudades, parametro, mapa)
            resultado = f"Peso total: {horas}"

        else:
            resultado = ""

        print(resultado, file = sys.stdout)


main()
