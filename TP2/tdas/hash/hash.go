package diccionario

import (
	"fmt"
)

const (
	ALTERNATIVA_INICIAL = 0
	ALTERNATIVA_MAXIMA  = 3
)
const (
	LARGO_INICIAL = 97
	INCREMENTO    = 3
	DISMINUCION   = 2
	CARGA_MAXIMA  = 80
	CARGA_MINIMA  = 10
)

// STRUCT DATO
type dato[K comparable, T any] struct {
	clave             K
	valor             T
	alternativas      [ALTERNATIVA_MAXIMA]uint64
	alternativaActual int
}

func crearDato[K comparable, T any](clave K, valor T, largo uint64) dato[K, T] {
	dato := new(dato[K, T])
	dato.alternativas = hashearClave(clave, largo)
	dato.clave = clave
	dato.valor = valor
	return *dato

}

func (dato dato[K, T]) ubicacionActual() uint64 {
	return dato.alternativas[dato.alternativaActual]
}

func (dato dato[K, T]) hayAlternativas() bool {
	return dato.alternativaActual >= ALTERNATIVA_MAXIMA
}

// TDA HASH
type tDAhash[K comparable, T any] struct {
	largo              uint64
	cantidad           int
	vector             []*dato[K, T]
	posibleDisminucion bool
}

func CrearHash[K comparable, T any]() Diccionario[K, T] {
	hash := new(tDAhash[K, T])
	hash.largo = LARGO_INICIAL
	hash.vector = make([]*dato[K, T], LARGO_INICIAL)
	return hash
}

func (hash *tDAhash[K, T]) Guardar(clave K, valor T) {
	dato := crearDato(clave, valor, hash.largo)
	hash._guardar(&dato, &dato)

}

func (hash *tDAhash[K, T]) Pertenece(clave K) bool {
	return hash.obtenerDato(clave) != nil
}

func (hash *tDAhash[K, T]) Obtener(clave K) T {
	dato := hash.obtenerDato(clave)
	if dato != nil {
		return dato.valor
	}
	panic("La clave no pertenece al diccionario")
}

func (hash *tDAhash[K, T]) Borrar(clave K) T {
	dato := hash.obtenerDato(clave)
	if dato == nil {
		panic("La clave no pertenece al diccionario")
	}
	valor := dato.valor
	hash.vector[dato.ubicacionActual()] = nil
	hash.cantidad--
	if hash.cargaActual() <= CARGA_MINIMA && hash.posibleDisminucion {
		hash.redimensionar(hash.largo / DISMINUCION)
	}
	return valor
}

func (hash tDAhash[K, T]) Cantidad() int {
	return hash.cantidad
}

func (hash *tDAhash[K, T]) Iterador() IterDiccionario[K, T] {
	iteradorStruct := crearIteradorExterno[K, T]()
	iteradorStruct.diccionario = hash
	var iteradorExterno IterDiccionario[K, T] = iteradorStruct
	if hash.cantidad != 0 {
		for posicion, dato := range hash.vector {
			if dato != nil {
				iteradorStruct.posicionMaxima = int(hash.largo)
				iteradorStruct.posicionActual = posicion
				iteradorStruct.datoActual = dato
				return iteradorExterno
			}
		}
	}
	return iteradorExterno
}

func (hash *tDAhash[K, T]) Iterar(funcion func(clave K, dato T) bool) {
	for _, dato := range hash.vector {
		if dato != nil {
			resultado := funcion(dato.clave, dato.valor)
			if !resultado {
				return
			}
		}
	}
}

// Funciones Adicionales

func convertirABytes[K comparable](clave K) []byte {
	return []byte(fmt.Sprintf("%v", clave))
}

// Regresa un arreglo con el valor de la clave luego de haberle aplicado una funcion de hash
func hashearClave[K comparable](clave K, largo uint64) [ALTERNATIVA_MAXIMA]uint64 {
	claveBytes := convertirABytes(clave)
	hash3 := uint64(fnvUint32(claveBytes)) % largo
	hash2 := uint64(sdbmHash(claveBytes)) % largo
	hash1 := fnvUint64(claveBytes) % largo
	return [ALTERNATIVA_MAXIMA]uint64{hash1, hash2, hash3}

}

// Regresa el porcentaje de la carga actual del hash
func (hash tDAhash[K, T]) cargaActual() int {
	return ((hash.cantidad * 100) / int(hash.largo))
}

// Guarda el dato en la ubicacion correspondiente aplicando Cuckoo Hashing
func (hash *tDAhash[K, T]) _guardar(dato *dato[K, T], primerDato *dato[K, T]) {
	if hash.cargaActual() >= CARGA_MAXIMA || primerDato.hayAlternativas() {
		hash.redimensionar(hash.largo * INCREMENTO)
		dato.alternativaActual = ALTERNATIVA_INICIAL
		hash.Guardar(dato.clave, dato.valor)
		return
	}

	if dato.hayAlternativas() {
		dato.alternativaActual = ALTERNATIVA_INICIAL
	}

	if hash.Pertenece(dato.clave) {
		hash.vector[dato.ubicacionActual()].valor = dato.valor

	} else {
		datoAcutal := hash.vector[dato.ubicacionActual()]
		hash.vector[dato.ubicacionActual()] = dato
		if datoAcutal == nil {
			hash.cantidad++
			return
		}
		datoAcutal.alternativaActual++
		hash._guardar(datoAcutal, primerDato)
	}

}

func guardarEnTabla[K comparable, T any](dato *dato[K, T], vector []*dato[K, T]) {
	if dato.hayAlternativas() {
		dato.alternativaActual = ALTERNATIVA_INICIAL
	}
	datoAcutal := vector[dato.ubicacionActual()]
	vector[dato.ubicacionActual()] = dato
	if datoAcutal == nil {
		return
	}
	datoAcutal.alternativaActual++
	guardarEnTabla(datoAcutal, vector)
}

// Busca a un dato que posea la misma clave y lo regresa, caso contrario regresa nil
func (hash tDAhash[K, T]) obtenerDato(clave K) *dato[K, T] {
	for _, valor := range hashearClave(clave, hash.largo) {
		dato := hash.vector[valor]
		if dato != nil && dato.clave == clave {
			return dato
		}
	}
	return nil
}

func (hash *tDAhash[K, T]) redimensionar(capacidad uint64) {
	verctorNuevo := make([]*dato[K, T], capacidad)
	for _, dato := range hash.vector {
		if dato != nil {
			datonuevo := crearDato(dato.clave, dato.valor, capacidad)
			guardarEnTabla(&datonuevo, verctorNuevo)
		}
	}
	hash.vector = verctorNuevo
	hash.largo = capacidad
	hash.posibleDisminucion = true

}

// ITERADOR EXTERNO

func crearIteradorExterno[K comparable, T any]() *iteradorExterno[K, T] {
	iteradorExterno := new(iteradorExterno[K, T])
	return iteradorExterno

}

type iteradorExterno[K comparable, T any] struct {
	datoActual     *dato[K, T]
	diccionario    *tDAhash[K, T]
	posicionActual int
	posicionMaxima int
}

func (iterador iteradorExterno[K, T]) HaySiguiente() bool {
	return iterador.posicionActual < iterador.posicionMaxima
}
func (iterador iteradorExterno[K, T]) VerActual() (K, T) {
	if !iterador.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	return iterador.datoActual.clave, iterador.datoActual.valor
}

func (iterador *iteradorExterno[K, T]) Siguiente() K {
	if !iterador.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	clave := iterador.datoActual.clave
	iterador.posicionActual++
	for {
		if iterador.posicionActual >= iterador.posicionMaxima {
			iterador.datoActual = nil
			return clave
		}

		if iterador.diccionario.vector[iterador.posicionActual] != nil {
			iterador.datoActual = iterador.diccionario.vector[iterador.posicionActual]
			return clave

		}
		iterador.posicionActual++

	}
}
