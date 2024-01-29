package usuario

import (
	ABB "algogram/tdas/abb"
	TDAHash "algogram/tdas/hash"
)

type Usuario interface {
	// Postear, crea un post para el usuario y lo agrega al feed de los demas usuarios existentes
	Postear(post *Post, usuarios TDAHash.Diccionario[string, Usuario])

	// AgregarAlFeed, agrega el post recibido al feed de usuario
	AgregarAlFeed(post *Post)

	// FeedVacio, indica si el usuario ya no tiene mas post por ver
	FeedVacio() bool

	// VerProximoFeed, retorna el siguiente post para el usuario,
	// Basandose en la afinidad del usuario con los demas usuarios
	VerProximoFeed() Post

	// Nombre, retorna el nombre del usuario
	Nombre() string
}

type Post struct {
	ID    int
	Autor *usuario
	Texto string
	Likes ABB.Diccionario[string, string]
}

func (post *Post) CantidadLikes() int {
	return post.Likes.Cantidad()
}

func (post *Post) LikearPost(nombreUsuario string) {
	post.Likes.Guardar(nombreUsuario, nombreUsuario)
}

func (post *Post) ObtenerLikes() []string {
	likes := []string{}

	for i := post.Likes.Iterador(); i.HaySiguiente(); {
		nombreUsuario, _ := i.VerActual()
		i.Siguiente()
		likes = append(likes, nombreUsuario)
	}
	return likes
}
