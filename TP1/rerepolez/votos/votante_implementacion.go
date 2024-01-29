package votos

import (
	"rerepolez/errores"
	TDApila "rerepolez/pila"
)

type votanteImplementacion struct {
	dni            int
	fraudulento    bool
	voto           Voto
	registoDeVotos TDApila.Pila[Voto]
}

func CrearVotante(dni int, esFraudulento bool) Votante {
	votanteStruct := new(votanteImplementacion)
	votanteStruct.dni = dni
	votanteStruct.registoDeVotos = TDApila.CrearPilaDinamica[Voto]()
	votanteStruct.fraudulento = esFraudulento
	var votante Votante = votanteStruct
	return votante
}

func (votante votanteImplementacion) LeerDNI() int {
	return votante.dni
}

func (votante *votanteImplementacion) Votar(tipo TipoVoto, alternativa int) error {

	votante.registoDeVotos.Apilar(votante.voto)

	switch {
	case votante.fraudulento:
		return errores.ErrorVotanteFraudulento{Dni: votante.dni}

	case alternativa == LISTA_IMPUGNA:
		votante.voto.Impugnado = true

	case votante.voto.Impugnado:
		return nil

	}

	votante.voto.VotoPorTipo[tipo] = alternativa

	return nil
}

func (votante *votanteImplementacion) Deshacer() error {
	switch {
	case votante.fraudulento:
		return errores.ErrorVotanteFraudulento{Dni: votante.dni}

	case votante.registoDeVotos.EstaVacia():
		return errores.ErrorNoHayVotosAnteriores{}
	}

	votante.voto = votante.registoDeVotos.Desapilar()
	return nil
}

func (votante *votanteImplementacion) FinVoto() (Voto, error) {
	if votante.fraudulento {
		return Voto{}, errores.ErrorVotanteFraudulento{Dni: votante.dni}
	}
	return votante.voto, nil
}
