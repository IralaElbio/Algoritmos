package funciones

func BusquedaBinaria(arreglo []int, dato int) bool {
	return _busquedaBinaria(arreglo, 0, len(arreglo)-1, dato)
}

func _busquedaBinaria(arreglo []int, inicio, fin, dato int) bool {
	for inicio <= fin {
		mitad := (inicio + fin) / 2
		switch {
		case arreglo[mitad] == dato:
			return true

		case arreglo[mitad] < dato:
			return _busquedaBinaria(arreglo, mitad+1, fin, dato)

		case arreglo[mitad] > dato:
			return _busquedaBinaria(arreglo, inicio, mitad-1, dato)
		}

	}
	return false
}
