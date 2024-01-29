package votos

import "fmt"

type partidoImplementacion struct {
	nombre      string
	integrantes []string
	votos       [CANT_VOTACION]int
}

type partidoEnBlanco struct {
	votos [CANT_VOTACION]int
}

func CrearPartido(nombre string, candidatos []string) Partido {
	patido_struct := new(partidoImplementacion)
	patido_struct.nombre = nombre
	patido_struct.integrantes = candidatos
	var partido Partido = patido_struct
	return partido
}

func CrearVotosEnBlanco() Partido {
	partidoblanco := new(partidoEnBlanco)
	var partido_blanco Partido = partidoblanco
	return partido_blanco
}

func (partido *partidoImplementacion) VotadoPara(tipo TipoVoto) {
	partido.votos[tipo] += 1
}

func (partido partidoImplementacion) ObtenerResultado(tipo TipoVoto) string {
	stringvoto := "votos"
	if partido.votos[tipo] == 1 {
		stringvoto = "voto"

	}
	return fmt.Sprintf("%v - %v: %v %v", partido.nombre, partido.integrantes[tipo], partido.votos[tipo], stringvoto)
}
func (blanco *partidoEnBlanco) VotadoPara(tipo TipoVoto) {
	blanco.votos[tipo] += 1
}

func (blanco partidoEnBlanco) ObtenerResultado(tipo TipoVoto) string {
	stringVoto := "votos"
	if blanco.votos[tipo] == 1 {
		stringVoto = "voto"

	}
	return fmt.Sprintf("Votos en Blanco: %v %v", blanco.votos[tipo], stringVoto)
}
