package main

import (
	funciones "algogram/funciones"
	TDAusuario "algogram/usuario"
	"bufio"
	"fmt"
	"os"
)

const (
	LOGIN    = "login"
	LOGOUT   = "logout"
	PUBLICAR = "publicar"
	VERFEED  = "ver_siguiente_feed"
	LIKE     = "likear_post"
	VERLIKES = "mostrar_likes"
)

func main() {
	ruta := os.Args[1:]

	usuarios, historialDePosts, err := funciones.IniciarApp(ruta[0])

	if err != nil {
		fmt.Println("Ruta inexistente")
		return
	}

	var usuario TDAusuario.Usuario
	var salidaTexto string
	var idpost int
	var likes []string

	scaner := bufio.NewScanner(os.Stdin)

	for scaner.Scan() {
		comando, parametro := funciones.LeerEntrada(scaner.Text())
		switch comando {
		case LOGIN:
			err, usuario = funciones.AsignarUsuario(usuario, parametro, usuarios)
			if usuario != nil {
				salidaTexto = fmt.Sprintf("Hola %v", usuario.Nombre())
			}

		case LOGOUT:
			err, usuario = funciones.QuitarUsuario(usuario)
			if err == nil {
				salidaTexto = "Adios"
			}

		case PUBLICAR:
			err = funciones.CrearPost(usuario, parametro, idpost, usuarios, &historialDePosts)
			if err == nil {
				idpost++
				salidaTexto = "Post publicado"

			}
		case VERFEED:
			salidaTexto = funciones.MostrarPostEnFeed(usuario, historialDePosts)
			err = nil

		case LIKE:
			err = funciones.LikearPost(parametro, usuario, historialDePosts)
			if err == nil {
				salidaTexto = "Post likeado"
			}

		case VERLIKES:
			err, likes = funciones.BuscarLikes(parametro, usuario, historialDePosts)
			if err == nil {
				salidaTexto = funciones.FormatearLikes(likes)
			}

		default:
			// Es meterle un reset a las variables, por si hubiera un caso donde no se cumple ninguna condicion de switch
			salidaTexto = ""
			err = nil
		}

		fmt.Println(funciones.ObtenerResultado(err, salidaTexto))
	}
}
