package cola_prioridad

const (
	PRIMERDATO      = 0
	CANTIDADINICIAL = 25
)

type heap[T comparable] struct {
	datos       []T
	cantidad    int
	comparacion func(T, T) int
}

func CrearHeap[T comparable](funcion_cmp func(T, T) int) ColaPrioridad[T] {
	heap := new(heap[T])
	heap.comparacion = funcion_cmp
	heap.datos = make([]T, CANTIDADINICIAL)
	var colaPrioridad ColaPrioridad[T] = heap
	return colaPrioridad
}

func CrearHeapArr[T comparable](arreglo []T, funcion_cmp func(T, T) int) ColaPrioridad[T] {
	datos := make([]T, len(arreglo)+CANTIDADINICIAL)
	copy(datos, arreglo)
	heapify(datos, funcion_cmp)
	heap := new(heap[T])
	heap.cantidad = len(arreglo)
	heap.datos = datos
	heap.comparacion = funcion_cmp
	var colaPrioridad ColaPrioridad[T] = heap
	return colaPrioridad
}

func (heap *heap[T]) Encolar(dato T) {
	heap.datos[heap.cantidad] = dato
	heap.cantidad++
	upHeap(heap.cantidad-1, heap.datos[:heap.cantidad], heap.comparacion)
	heap.verificarCapacidad()
}

func (heap *heap[T]) Desencolar() T {
	if heap.EstaVacia() {
		panic("La cola esta vacia")
	}
	dato := heap.datos[PRIMERDATO]
	heap.datos[PRIMERDATO], heap.datos[heap.cantidad-1] = heap.datos[heap.cantidad-1], heap.datos[PRIMERDATO]
	heap.cantidad--
	downHeap(PRIMERDATO, heap.datos[:heap.cantidad], heap.comparacion)
	heap.verificarCapacidad()
	return dato
}

func (heap heap[T]) EstaVacia() bool {
	return heap.cantidad == 0
}

func (heap heap[T]) VerMax() T {
	if heap.EstaVacia() {
		panic("La cola esta vacia")
	}
	return heap.datos[PRIMERDATO]
}

func (heap heap[T]) Cantidad() int {
	return heap.cantidad
}

func (heap *heap[T]) redimension(capacidad int) {
	datosNuevo := make([]T, capacidad)
	copy(datosNuevo, heap.datos)
	heap.datos = datosNuevo
}

func (heap *heap[T]) verificarCapacidad() {
	var capacidad int
	switch {
	case cap(heap.datos) == heap.cantidad:
		capacidad = cap(heap.datos) * 2

	case heap.cantidad > CANTIDADINICIAL && heap.cantidad*4 <= cap(heap.datos):
		capacidad = cap(heap.datos) / 2
	default:
		return
	}
	heap.redimension(capacidad)

}

func HeapSort[T comparable](elementos []T, funcion_cmp func(T, T) int) {
	heapify(elementos, funcion_cmp)
	for i := len(elementos) - 1; i > 0; i-- {
		elementos[PRIMERDATO], elementos[i] = elementos[i], elementos[PRIMERDATO]
		downHeap(PRIMERDATO, elementos[:i], funcion_cmp)

	}

}

func upHeap[T comparable](hijo int, arreglo []T, comparacion func(T, T) int) {
	if hijo == PRIMERDATO {
		return
	}
	padre := (hijo - 1) / 2
	if comparacion(arreglo[hijo], arreglo[padre]) > 0 {
		arreglo[hijo], arreglo[padre] = arreglo[padre], arreglo[hijo]
		upHeap(padre, arreglo, comparacion)
	}
	return
}

func heapify[T comparable](arreglo []T, funcion_cmp func(T, T) int) {
	for i := len(arreglo) / 2; i >= 0; i-- {
		downHeap(i, arreglo, funcion_cmp)
	}
}

func downHeap[T comparable](padre int, arreglo []T, comparacion func(T, T) int) {
	hijoIZQ := (padre * 2) + 1
	hijoDER := (padre * 2) + 2
	var hijoMayor int
	switch {
	case hijoDER >= len(arreglo) && hijoIZQ >= len(arreglo):
		return
	case hijoDER >= len(arreglo):
		hijoMayor = hijoIZQ
	case hijoIZQ >= len(arreglo):
		hijoMayor = hijoDER
	default:
		if comparacion(arreglo[hijoIZQ], arreglo[hijoDER]) > 0 {
			hijoMayor = hijoIZQ
		} else {
			hijoMayor = hijoDER
		}
	}

	if comparacion(arreglo[padre], arreglo[hijoMayor]) < 0 {
		arreglo[padre], arreglo[hijoMayor] = arreglo[hijoMayor], arreglo[padre]
		downHeap(hijoMayor, arreglo, comparacion)
	}
	return
}
