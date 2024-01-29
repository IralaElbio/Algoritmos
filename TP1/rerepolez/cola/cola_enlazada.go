package cola

type colaEnlazada[T any] struct {
	primero *nodo[T]
	ultimo  *nodo[T]
}

type nodo[T any] struct {
	dato T
	prox *nodo[T]
}

func crearNodo[T any]() *nodo[T] {
	nodo := new(nodo[T])
	return nodo
}

func CrearColaEnlazada[T any]() Cola[T] {
	colaStruct := new(colaEnlazada[T])
	var cola Cola[T] = colaStruct
	return cola
}

func (cola colaEnlazada[T]) EstaVacia() bool {
	return cola.primero == nil
}

func (cola *colaEnlazada[T]) Desencolar() T {
	if cola.EstaVacia() {
		panic("La cola esta vacia")
	}

	dato := cola.primero.dato
	cola.primero = cola.primero.prox
	if cola.primero == nil {
		cola.ultimo = nil
	}

	return dato
}

func (cola *colaEnlazada[T]) Encolar(elemento T) {
	nodo := crearNodo[T]()
	nodo.dato = elemento

	if cola.EstaVacia() {
		cola.primero = nodo
		cola.ultimo = nodo
	} else {
		cola.ultimo.prox = nodo
		cola.ultimo = cola.ultimo.prox
	}
}

func (cola colaEnlazada[T]) VerPrimero() T {
	if cola.EstaVacia() {
		panic("La cola esta vacia")
	}
	return cola.primero.dato
}
