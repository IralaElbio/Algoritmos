package diccionario

type DiccionarioOrdenado[K comparable, V any] interface {
	Diccionario[K, V]

	// IterarRango itera sólo incluyendo a los elementos que se encuentren comprendidos en el rango indicado,
	// incluyéndolos en caso de encontrarse
	IterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool)

	// IteradorRango crea un IterDiccionario que sólo itere por las claves que se encuentren en el rango indicado
	IteradorRango(desde *K, hasta *K) IterDiccionario[K, V]
}

type Diccionario[K comparable, V any] interface {
	// Guardar guarda el par clave-dato en el Diccionario. Si la clave ya se encontraba, se actualiza el dato asociado
	Guardar(clave K, dato V)

	// Pertenece determina si una clave ya se encuantra en el diccionario o no
	Pertenece(clave K) bool

	// Obtener devuelve el dato asociado a una clave. Si la clave no pertenece, entra en panico
	// con el mensaje "La clave no pertenece al diccionario"
	Obtener(clave K) V

	// Borrar borra del Diccionario la clave indicada, devolviendo el dato que se encontraba asociado.
	// Si la clave no pertenece al diccionario, entra en panico con el mensaje
	// "La clave no pertenece al diccionario"
	Borrar(clave K) V

	// Cantidad devuelve la cantidad de elementos dentro del diccionario
	Cantidad() int

	// Iterar itrea internamente el diccionario, aplicando la funcion pasada por parametro a todos los elementos
	Iterar(func(clave K, dato V) bool)

	// Iterador devuelve un IterDiccionario para este diccioario
	Iterador() IterDiccionario[K, V]
}

type IterDiccionario[K comparable, V any] interface {
	// HaySiguiente devualve si hay mas datos por ver. Esto es, si en el lugar donde se encuentra parado
	// el itrador hay un elemento
	HaySiguiente() bool

	// VerActual devuelve la clave y el dato del elemento actual en el que se encuentra posicionado el iterador.
	// Si no HaySiguiente, entra en panico con el mensaje "El iterador termino de iterar"
	VerActual() (K, V)

	// Siguiente si HaySiguiente, devuelve la clave actual y ademas avanza al siguiente elemento en el diccionario.
	// Si no HaySiguiente, entra en panico con el mensaje "El iterador termino de iterar"
	Siguiente() K
}
