package usuario

import (
	TDAHash "algogram/tdas/hash"
	heap "algogram/tdas/heap"
)

type usuario struct {
	nombreUsuario string
	numAfinidad   int
	feed          heap.ColaPrioridad[*Post]
}

func Crearusuario(nombreUsuario string, numAfinidad int) Usuario {
	usuario := new(usuario)
	usuario.nombreUsuario = nombreUsuario
	usuario.numAfinidad = numAfinidad
	usuario.feed = heap.CrearHeap(usuario.compararPosts)
	return usuario
}

func (user usuario) Nombre() string {
	return user.nombreUsuario
}

func (user *usuario) compararPosts(post1, post2 *Post) int {
	afinidad1 := modulo(user.numAfinidad - post1.Autor.numAfinidad)
	afinidad2 := modulo(user.numAfinidad - post2.Autor.numAfinidad)
	// si la afinidad es la misma, evalua cual post se creo primero
	if afinidad1 == afinidad2 {
		return post2.ID - post1.ID
	}
	return afinidad2 - afinidad1
}

func (user *usuario) AgregarAlFeed(post *Post) {
	user.feed.Encolar(post)
}

func (user usuario) FeedVacio() bool {
	return user.feed.EstaVacia()
}

func (user *usuario) VerProximoFeed() Post {
	return *user.feed.Desencolar()
}

func (user *usuario) Postear(post *Post, usuarios TDAHash.Diccionario[string, Usuario]) {
	post.Autor = user
	usuarios.Iterar(func(nombre string, usuario Usuario) bool {
		if nombre != user.nombreUsuario {
			usuario.AgregarAlFeed(post)
		}
		return true
	})
}

func modulo(num int) int {
	if num < 0 {
		return num * -1
	}
	return num
}
