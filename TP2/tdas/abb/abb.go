package diccionario

import (
	TDApila "algogram/tdas/abb/pila"
)

// NodoAbb
type nodoabb[K comparable, V any] struct {
	izquierdo *nodoabb[K, V]
	derecho   *nodoabb[K, V]
	clave     K
	valor     V
}

func crearNodoAbb[K comparable, V any](clave K, valor V) *nodoabb[K, V] {
	nodoabb := new(nodoabb[K, V])
	nodoabb.clave = clave
	nodoabb.valor = valor
	return nodoabb
}

// ABB
type abb[K comparable, V any] struct {
	raiz               *nodoabb[K, V]
	cantidad           int
	funcionComparacion func(K, K) int
}

func CrearABB[K comparable, V any](funcionComparacion func(K, K) int) DiccionarioOrdenado[K, V] {
	nodoabb := new(abb[K, V])
	nodoabb.funcionComparacion = funcionComparacion
	return nodoabb
}

func (abb *abb[K, V]) Guardar(clave K, valor V) {
	if abb.raiz == nil {
		abb.raiz = crearNodoAbb(clave, valor)
	} else {
		nodo := abb.raiz.buscarNodo(clave, abb.funcionComparacion)

		// Si es una clave ya existente
		if *nodo != nil {
			(*nodo).valor = valor
			return
		}
		*nodo = crearNodoAbb(clave, valor)
	}
	abb.cantidad++
}
func (abb abb[K, V]) Pertenece(clave K) bool {
	return *abb.raiz.buscarNodo(clave, abb.funcionComparacion) != nil
}

func (abb abb[K, V]) Obtener(clave K) V {
	nodo := abb.raiz.buscarNodo(clave, abb.funcionComparacion)
	if *nodo == nil {
		panic("La clave no pertenece al diccionario")
	}
	return (*nodo).valor
}

func (abb abb[K, V]) Cantidad() int {
	return abb.cantidad
}

func (abb *abb[K, V]) Borrar(clave K) V {
	nodo := abb.raiz.buscarNodo(clave, abb.funcionComparacion)
	if *nodo == nil {
		panic("La clave no pertenece al diccionario")
	}
	valorRes := (*nodo).valor
	var reemplazo *nodoabb[K, V]

	switch {
	// Caso 2 hijos
	case (*nodo).izquierdo != nil && (*nodo).derecho != nil:
		nodomenorderecho := (*nodo).buscarMenorDerecho()
		clave := nodomenorderecho.clave
		valor := abb.Borrar(nodomenorderecho.clave)
		(*nodo).clave = clave
		(*nodo).valor = valor
		return valorRes

		// Casos con 1 hijo
	case (*nodo).izquierdo != nil:
		reemplazo = (*nodo).izquierdo

	case (*nodo).derecho != nil:
		reemplazo = (*nodo).derecho
	}

	if clave == abb.raiz.clave {
		abb.raiz = reemplazo
	}

	*nodo = reemplazo
	abb.cantidad--
	return valorRes
}

func (abb *abb[K, V]) Iterar(funcion func(clave K, valor V) bool) {
	abb.IterarRango(nil, nil, funcion)
}

func (abb *abb[K, V]) IterarRango(desde *K, hasta *K, funcion func(clave K, valor V) bool) {
	abb.raiz.iterarRango(desde, hasta, funcion, abb.funcionComparacion)

}

func (abb abb[K, V]) Iterador() IterDiccionario[K, V] {
	return abb.IteradorRango(nil, nil)
}

func (abb abb[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] {
	iteradorExterno := crearIteradorExterno[K, V]()
	iteradorExterno.desde = desde
	iteradorExterno.hasta = hasta
	iteradorExterno.funcionComparacion = abb.funcionComparacion
	nodoActual := abb.raiz

	if abb.raiz != nil {
		switch desde {
		// Si desde es nil el apilado es como un iterador sin rango
		case nil:
			for nodoActual != nil {
				iteradorExterno.pila.Apilar(nodoActual)
				nodoActual = nodoActual.izquierdo
			}
		default:
			for nodoActual != nil {
				comparacion := abb.funcionComparacion(nodoActual.clave, *desde)
				if comparacion >= 0 {
					iteradorExterno.pila.Apilar(nodoActual)
					nodoActual = nodoActual.izquierdo
					continue
				}
				nodoActual = nodoActual.derecho
			}
		}
	}
	var iterador IterDiccionario[K, V] = iteradorExterno
	return iterador
}

func (nodo *nodoabb[K, V]) iterarRango(desde *K, hasta *K, visitar func(clave K, valor V) bool, funcionComparacion func(K, K) int) bool {
	if nodo == nil {
		return false
	}
	// En caso de no recibir un desde o un hasta por paremetro los inicializo con un valor que cumpla con las condiciones
	compHasta := -1
	compDesde := 1
	var corte bool
	if desde != nil {
		compDesde = funcionComparacion(nodo.clave, *desde)
	}
	if hasta != nil {
		compHasta = funcionComparacion(nodo.clave, *hasta)
	}

	if compDesde > 0 && nodo.izquierdo != nil {
		corte = nodo.izquierdo.iterarRango(desde, hasta, visitar, funcionComparacion)
	}

	if corte {
		return true
	}

	if compDesde >= 0 && compHasta <= 0 {
		if !visitar(nodo.clave, nodo.valor) {
			return true
		}

	}

	if compHasta < 0 && nodo.derecho != nil {
		return nodo.derecho.iterarRango(desde, hasta, visitar, funcionComparacion)
	}
	return false
}

func (nodo nodoabb[K, V]) buscarMenorDerecho() *nodoabb[K, V] {
	nodoActual := nodo.derecho
	for nodoActual.izquierdo != nil {
		nodoActual = nodoActual.izquierdo

	}
	return nodoActual

}

func (nodo *nodoabb[K, V]) buscarNodo(clave K, funcionComparacion func(K, K) int) **nodoabb[K, V] {
	if nodo == nil {
		return &nodo
	}
	switch {
	case funcionComparacion(clave, nodo.clave) == 0:
		// Va a ser igual si el nodo recibido es el nodo raiz
		return &nodo
	case funcionComparacion(clave, nodo.clave) < 0:
		if nodo.izquierdo == nil || funcionComparacion(clave, nodo.izquierdo.clave) == 0 {
			return &nodo.izquierdo
		}
		return &(*(nodo.izquierdo).buscarNodo(clave, funcionComparacion))
	case funcionComparacion(clave, nodo.clave) > 0:
		if nodo.derecho == nil || funcionComparacion(clave, nodo.derecho.clave) == 0 {
			return &nodo.derecho
		}
		return &(*(nodo.derecho).buscarNodo(clave, funcionComparacion))
	}
	return nil
}

// IteradorExterno
type iteradorExterno[K comparable, V any] struct {
	pila               TDApila.Pila[*nodoabb[K, V]]
	nodoabb            *nodoabb[K, V]
	funcionComparacion func(K, K) int
	desde              *K
	hasta              *K
}

func crearIteradorExterno[K comparable, V any]() *iteradorExterno[K, V] {
	iteradorExterno := new(iteradorExterno[K, V])
	iteradorExterno.pila = TDApila.CrearPilaDinamica[*nodoabb[K, V]]()
	return iteradorExterno

}

func (iterador iteradorExterno[K, V]) VerActual() (K, V) {
	if !iterador.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	nodo := iterador.pila.VerTope()
	return nodo.clave, nodo.valor
}

func (iterador iteradorExterno[K, V]) HaySiguiente() bool {
	if iterador.hasta == nil || iterador.pila.EstaVacia() {
		return !iterador.pila.EstaVacia()
	}
	comparacion := iterador.funcionComparacion(iterador.pila.VerTope().clave, *iterador.hasta)
	if comparacion > 0 {
		return false
	}
	return true
}

func (iterador *iteradorExterno[K, V]) Siguiente() K {
	if !iterador.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	nodo := iterador.pila.Desapilar()
	nodoDerecho := nodo.derecho

	if nodoDerecho != nil {
		iterador.pila.Apilar(nodoDerecho)
		nodoActual := nodoDerecho.izquierdo
		switch iterador.desde {

		case nil:
			for nodoActual != nil {
				iterador.pila.Apilar(nodoActual)
				nodoActual = nodoActual.izquierdo
			}

		default:
			for nodoActual != nil {
				comparacion := iterador.funcionComparacion(nodoActual.clave, *iterador.desde)
				if comparacion >= 0 {
					iterador.pila.Apilar(nodoActual)
				}
				nodoActual = nodoActual.izquierdo
			}
		}
	}
	return nodo.clave
}
