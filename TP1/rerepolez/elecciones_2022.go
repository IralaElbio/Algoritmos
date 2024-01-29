package main

import (
	"bufio"
	"fmt"
	"os"
	TDAcola "rerepolez/cola"
	errores "rerepolez/errores"
	funciones "rerepolez/funciones"
	elecciones "rerepolez/votos"
	"strings"
)

const (
	COMANDO int = iota
	PARAMETRO1
	PARAMETRO2
)
const (
	INGRESAR  = "ingresar"
	VOTAR     = "votar"
	DESHACER  = "deshacer"
	FINALIZAR = "fin-votar"
)

func main() {
	rutas := os.Args[1:]

	padron, listaCandidatos, err := funciones.IniciarMesa(rutas)
	if err != nil {
		fmt.Printf("%v\n", err.Error())
		return
	}

	cola := TDAcola.CrearColaEnlazada[elecciones.Votante]()

	scaner := bufio.NewScanner(os.Stdin)

	var (
		errorVoto        error
		entrada          []string
		comando          string
		registroVotantes []int
		votosImpugnados  int
		resultado        string
	)

	for scaner.Scan() {
		entrada = strings.Split(scaner.Text(), " ")
		comando = entrada[COMANDO]
		switch {
		case comando == INGRESAR:
			errorVoto = funciones.CrearVotante(entrada[PARAMETRO1], padron, cola, &registroVotantes)

		case cola.EstaVacia():
			errorVoto = errores.FilaVacia{}

		case comando == VOTAR:
			errorVoto = funciones.IngresarVoto(entrada[PARAMETRO1], entrada[PARAMETRO2], cola.VerPrimero(), len(listaCandidatos)-1)

		case comando == DESHACER:
			errorVoto = cola.VerPrimero().Deshacer()

		case comando == FINALIZAR:
			errorVoto = funciones.FinalizarVoto(cola, listaCandidatos, &votosImpugnados)
		}
		resultado = funciones.VerificarVotoDelVotante(errorVoto, cola)
		fmt.Printf("%v\n", resultado)
	}
	if !cola.EstaVacia() {
		error := errores.ErrorCiudadanosSinVotar{}
		fmt.Printf("%v\n", error)
	}
	funciones.FinalizarVotacion(listaCandidatos, votosImpugnados)
}
