package pila

const TAMANIO_INICIAL int = 20

/* Definición del struct pila proporcionado por la cátedra. */

type pilaDinamica[T any] struct {
	datos    []T
	cantidad int
}

func CrearPilaDinamica[T any]() Pila[T] {
	pila_struck := new(pilaDinamica[T])
	pila_struck.datos = make([]T, TAMANIO_INICIAL)
	pila_struck.cantidad = 0
	var pila Pila[T] = pila_struck
	return pila
}

func (pila pilaDinamica[T]) EstaVacia() bool {
	return pila.cantidad == 0
}

func (pila pilaDinamica[T]) VerTope() T {
	if pila.EstaVacia() {
		panic("La pila esta vacia")
	}

	return pila.datos[pila.cantidad-1]
}

func (pila *pilaDinamica[T]) Desapilar() T {
	if pila.EstaVacia() {
		panic("La pila esta vacia")
	}

	defer pila.verificar_capacidad()
	defer pila.disminuir_cantidad()

	return pila.datos[pila.cantidad-1]
}

func (pila *pilaDinamica[T]) Apilar(elem T) {
	defer pila.verificar_capacidad()
	defer pila.aumentar_cantidad()

	pila.datos[pila.cantidad] = elem
}

func (pila *pilaDinamica[T]) aumentar_cantidad() {
	pila.cantidad++
}

func (pila *pilaDinamica[T]) disminuir_cantidad() {
	pila.cantidad--
}

func (pila *pilaDinamica[T]) verificar_capacidad() {
	redimension := 2
	cuadruple := 4
	var capacidad int
	switch {
	case cap(pila.datos) == pila.cantidad:
		capacidad = cap(pila.datos) * redimension

	case (pila.cantidad)*cuadruple <= cap(pila.datos):
		capacidad = cap(pila.datos) / redimension

	default:
		return
	}
	nuevo := make([]T, capacidad)
	copy(nuevo, pila.datos)
	pila.datos = nuevo
}
