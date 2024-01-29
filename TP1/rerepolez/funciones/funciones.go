package funciones

import (
	"fmt"
	TDAcola "rerepolez/cola"
	errores "rerepolez/errores"
	elecciones "rerepolez/votos"
	"strconv"
)

const (
	PRESIDENTE_STR = "Presidente"
	GOBERNADOR_STR = "Gobernador"
	INTENDENTE_STR = "Intendente"
)

// Da por finalizada la votacion y escribe el resultado de las votaciones por Stdout.
func FinalizarVotacion(partidos []elecciones.Partido, votosImpugnados int) {
	for i := 0; i < int(elecciones.CANT_VOTACION); i++ {
		tipo_voto := elecciones.TipoVoto(i)
		fmt.Println(tipo_voto.String())
		for _, partido := range partidos {
			fmt.Println(partido.ObtenerResultado(tipo_voto))
		}
		fmt.Println("")
	}
	votoString := "votos"
	if votosImpugnados == 1 {
		votoString = "voto"
	}
	fmt.Printf("Votos Impugnados: %v %v\n", votosImpugnados, votoString)
}

// Termina el proceso de votación para el votante actual en la fila y emite su voto,
// si el votante cometio una infraccion regresa el error correspondiente
func FinalizarVoto(cola TDAcola.Cola[elecciones.Votante], partidos []elecciones.Partido, votosImpugnados *int) error {
	votoFinal, err := cola.VerPrimero().FinVoto()
	if err != nil {
		return err
	}
	cola.Desencolar()
	if votoFinal.Impugnado {
		*votosImpugnados++
		return nil
	}

	for categoria, alternativa := range votoFinal.VotoPorTipo {
		partidos[alternativa].VotadoPara(elecciones.TipoVoto(categoria))
	}
	return nil

}

// Crea a un votante y lo pone en la cola de votacion,
// si el DNI es invalio o no esta en el padron regresa el error correspondiente
func CrearVotante(dniString string, padron []int, cola TDAcola.Cola[elecciones.Votante], registro *[]int) error {
	dni, err := strconv.Atoi(dniString)
	if err != nil || dni < 0 {
		return errores.DNIError{}
	}
	var votanteFraudulento bool
	if BusquedaBinaria(padron, dni) {
		for _, dniRegistro := range *registro {
			if dni == dniRegistro {
				votanteFraudulento = true
			}
		}
		*registro = append(*registro, dni)
		cola.Encolar(elecciones.CrearVotante(dni, votanteFraudulento))
		return nil
	}
	return errores.DNIFueraPadron{}
}

// Vota por alguna alternativa, si el votante cometio una infraccion regresara el error correspondiente
func IngresarVoto(tipo_voto string, numeroLista string, votante elecciones.Votante, cantidadDeCandidatos int) error {
	var err error
	numero_lista, err := strconv.Atoi(numeroLista)

	if err != nil || numero_lista > cantidadDeCandidatos || numero_lista < 0 {
		return errores.ErrorAlternativaInvalida{}
	}
	var voto elecciones.TipoVoto

	switch tipo_voto {
	case PRESIDENTE_STR:
		voto = elecciones.PRESIDENTE
	case GOBERNADOR_STR:
		voto = elecciones.GOBERNADOR
	case INTENDENTE_STR:
		voto = elecciones.INTENDENTE
	default:
		return errores.ErrorTipoVoto{}
	}
	return votante.Votar(voto, numero_lista)
}

// Verifica si el votante cometio una infraccion durante su votacion,
// regresa "OK" y al votante si no se cometio infraccion,
// caso contrario el error correspondiente y la accion necesaria hacia el votante
func VerificarVotoDelVotante(err error, cola TDAcola.Cola[elecciones.Votante]) string {
	switch err.(type) {
	case errores.ErrorVotanteFraudulento:
		cola.Desencolar()
		return err.Error()
	}

	if err != nil {
		return err.Error()
	}
	return fmt.Sprint("OK")
}
