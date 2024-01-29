package funciones

import (
	"bufio"
	"os"
	errores "rerepolez/errores"
	elecciones "rerepolez/votos"
	"strconv"
	"strings"
)

const (
	RUTA_CANDIDATOS int = iota
	RUTA_PADRON
	RUTAS_RECIBIDAS
)

// Crea un padron ordenado para con los DNI de los votantes
func crearPadron(ruta string) ([]int, error) {
	archivo, err := os.Open(ruta)
	if err != nil {
		return nil, errores.ErrorLeerArchivo{}
	}
	padron := []int{}

	scan := bufio.NewScanner(archivo)

	for scan.Scan() {
		valor, _ := strconv.Atoi(scan.Text())

		padron = append(padron, valor)

	}
	padron = MergeSort(padron)
	return padron, nil

}

// Crea una lista con los Partidos, para las elecciones,
// la primera posicion esta reservada para los votos en blanco
func crearPartidos(ruta string) ([]elecciones.Partido, error) {
	archivo, err := os.Open(ruta)
	if err != nil {
		return nil, errores.ErrorLeerArchivo{}
	}
	scan := bufio.NewScanner(archivo)

	listaEleccion := []elecciones.Partido{elecciones.CrearVotosEnBlanco()}

	for scan.Scan() {

		linea := strings.Split(scan.Text(), ",")

		partido := elecciones.CrearPartido(linea[0], linea[1:])

		listaEleccion = append(listaEleccion, partido)
	}
	return listaEleccion, nil
}

// Inicializa la mesa de votacion junto con el padron y la lista de candidatos,
// regresa un error si hubo un fallo al crear la mesa
func IniciarMesa(rutas []string) ([]int, []elecciones.Partido, error) {
	if len(rutas) < RUTAS_RECIBIDAS {
		return nil, nil, errores.ErrorParametros{}

	}
	listaCandidatos, err := crearPartidos(rutas[RUTA_CANDIDATOS])
	if err != nil {
		return nil, nil, err
	}
	padron, err := crearPadron(rutas[RUTA_PADRON])
	if err != nil {
		return nil, nil, err
	}

	return padron, listaCandidatos, nil

}
