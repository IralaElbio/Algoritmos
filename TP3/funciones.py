from grafo.grafo import Grafo
from grafo.biblioteca_grafo.camino_minimo import camino_minimo_dijkstra
from grafo.biblioteca_grafo.orden_topologico import topologico_grados
from grafo.biblioteca_grafo.tendido_minimo import tendido_minimo_prim
from grafo.biblioteca_grafo.ciclo_euleriano import ciclo_euler
from grafo.biblioteca_grafo.extras import *


def iniciar_programa(ruta):
    ''' Lee la ruta pasada por parametro y inicial el programa,
        devolviendo los elementos necesarios para el programa'''

    with open(ruta, "r") as registro_mapa:
        ciudades = Grafo(False)
        registro_ciudades = {}
        cant_ciudades = int(registro_mapa.readline())

        for _ in range(cant_ciudades):
            ciudad, latitud, longitud = registro_mapa.readline().strip("\n").split(",")
            ciudades.agregar_vertice(ciudad)
            registro_ciudades[ciudad] = (latitud, longitud)

        conexiones = int(registro_mapa.readline())

        for _ in range(conexiones):
            ciudad1, ciudad2, peso = registro_mapa.readline().strip("\n").split(",")
            ciudades.agregar_arista(ciudad1, ciudad2, int(peso))

    return ciudades, registro_ciudades


def leer_entrada(entrada):
    ''' Separa el texto recibido por parametro en el primer espacio
        que encuentra'''

    for i in range(len(entrada)):
        if entrada[i] == " ":
            return entrada[:i], entrada[i + 1:]


def obtener_camino(ciudades, ciudad_inicial, ciudad_final, ruta, mapa):
    ''' Busca el camino minimo entre las ciudades, desde la ciudad inicial
        hasta la ciudad final, si lo encuentra devuelve el camino, el tiempo total
        que se tarda en hacer el recorrido y escribe un archivo.kms en la ruta indicada.
        Si no encuentra camino devuelve None '''

    camino, distancia = camino_minimo_dijkstra(ciudades, ciudad_inicial, ciudad_final)
    if camino is None:
        return  None, None

    recorrido = reconstruir_camino(camino, ciudad_final)

    escribir_kml(recorrido, mapa, ruta)

    return recorrido, distancia[ciudad_final]


def reconstruir_camino(camino, ciudad):
    ''' Reconstruye un camino a partir de un diccionario de ciudades para la ciudad en cuestion'''

    ciudad_actual = camino[ciudad]

    recorrido = [ciudad]

    while ciudad_actual != None:
        recorrido.append(ciudad_actual)

        ciudad_actual = camino[ciudad_actual]
    return recorrido[::-1]


def armar_recorrido(ciudades, tiempo = None):
    ''' Recibe una lista de ciudades y devuelve el recorrido en formato
        "ciudad_1" --> "ciudad_2" y el tiempo que se tarda si es que lo recibe'''

    if ciudades is None:
        return "No se encontro recorrido"
    
    if tiempo is not None:
        return f"{' -> '.join(ciudades)}\nTiempo total: {tiempo}"
    return " -> ".join(ciudades)


def obtener_camino_ordenado(ruta, mapa):
    ''' Devuelve un orden para visitar las ciudades dado el itinerario
        de la ruta que recibe por parametro y las ciudades que estan en el mapa'''

    with open(ruta, "r") as ciudades:
        recorrido_recomendado = Grafo(True)
        for ciudad in mapa.keys():
            recorrido_recomendado.agregar_vertice(ciudad)

        for linea in ciudades:
            ciudad1, ciudad2 = linea.strip("\n").split(",")
            recorrido_recomendado.agregar_arista(ciudad1, ciudad2)

    recorrido = topologico_grados(recorrido_recomendado)

    return recorrido


def camino_total(ciudades, ruta, mapa):
    ''' Escribe en la ruta un archivo .pajek con las ciudades extrictamentes necesarias
        y devuelve la cantidad de horas totales que toma hacer el recorrido'''

    ciudad_inicial = ciudades.obtener_vertice_aleatorio()

    ciudades_min = tendido_minimo_prim(ciudades, ciudad_inicial)

    horas_totales = peso_total(ciudades_min)

    escribir_pajek(ciudades_min, mapa, ruta)
    return horas_totales


def obtener_recorrido_completo(ciudades, ciudad_inicial, ruta, mapa):
    ''' Obtiene un recorrido para pasar por todas las ciudades sin pasar
        mas de una ves por los caminos que las unen, si lo encuentra,
        devuelve en formato el camino a recorrer y las horas totales,
        caso contrario None'''
    recorrido, horas_totales = ciclo_euler(ciudades, ciudad_inicial)
    if recorrido is not None:
        escribir_kml(recorrido, mapa, ruta)
    return recorrido, horas_totales


def escribir_pajek(ciudades, mapa, ruta):
    """ Escribe en la ruta pasada por parametro en formato texto
        lo siguiente:
        "cantidad de cuidaes (n)
        ciudad1,latitud1,longitud1
        ciudad2,latitud2,longitud2
        ...
        ciudad_n,latitud_n,longitud_n
        catidad de conexiones
        ciudad_i,ciudad_j,tiempo
        ...
        "

        """
    with open(ruta, "w") as salida:
        salida.write(f"{len(ciudades.vertices())}\n")
        for ciudad in ciudades.vertices():
            latitud, longitud = mapa[ciudad]

            salida.write(f"{ciudad},{latitud},{longitud}\n")

        salida.write(f"{cantidad_aristas_grafo_dirigido(ciudades)}\n")
        for ciudad in ciudades.vertices():
            for ciudad2 in ciudades.adyacentes(ciudad):

                salida.write(
                    f"{ciudad},{ciudad2},{ciudades.peso(ciudad, ciudad2)}\n")


def escribir_kml(camino, mapa, ruta):
    ''' Escribe en la ruta un archivo con formato .kml con el codigo necesario para poder
        visualizar la informacion geográfica del camino a hacer entre las ciudades'''
    
    ciudad_ubicada = set()

    with open(ruta, "w") as salida:
        salida.write('<?xml version="1.0" encoding="UTF-8"?>\n')
        salida.write('<kml xmlns="http://earth.google.com/kml/2.1">\n')
        salida.write('\t<Document>\n')
        for ciudad in camino:
            # para evitar escribir mas de una ves las ciudades en caso de que se repitan en el recorrido
            if ciudad in ciudad_ubicada:
                continue

            ciudad_ubicada.add(ciudad)
            
            latitud, longitud = mapa[ciudad]
            salida.write('\t\t<Placemark>\n')

            salida.write(f'\t\t\t<name>{ciudad}</name>\n')

            salida.write(f'\t\t\t<Point>\n')

            salida.write(f'\t\t\t\t<coordinates>{latitud}, {longitud}</coordinates>\n')

            salida.write(f'\t\t\t</Point>\n')

            salida.write('\t\t</Placemark>\n\n')

        for indice in range(len(camino)):
            if indice + 1 >= len(camino):
                break
            latitud_inicial, longitud_inicial = mapa[camino[indice]]
            latitud_final, longitud_final = mapa[camino[indice + 1]]
            salida.write('\t\t<Placemark>\n')

            salida.write(f'\t\t\t<LineString>\n')

            salida.write(f'\t\t\t\t<coordinates>{latitud_inicial}, {longitud_inicial} {latitud_final}, {longitud_final}</coordinates>\n')

            salida.write(f'\t\t\t</LineString>\n')

            salida.write('\t\t</Placemark>\n\n')

        salida.write('\t</Document>\n')
        salida.write('</kml>\n')
