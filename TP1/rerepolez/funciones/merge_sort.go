package funciones

func MergeSort(arreglo []int) []int {
	largo := len(arreglo)
	if largo == 1 {
		return arreglo
	}
	medio := largo / 2
	izquierda := arreglo[:medio]
	derecha := arreglo[medio:]
	return merge(MergeSort(izquierda), MergeSort(derecha))
}

func merge(izquierda, derecha []int) []int {
	arreglo := make([]int, len(izquierda)+len(derecha))
	i := 0
	for len(izquierda) > 0 && len(derecha) > 0 {
		if izquierda[0] < derecha[0] {
			arreglo[i] = izquierda[0]
			izquierda = izquierda[1:]
		} else {
			arreglo[i] = derecha[0]
			derecha = derecha[1:]
		}
		i++
	}

	for x := 0; x < len(izquierda); x++ {
		arreglo[i] = izquierda[x]
		i++
	}

	for x := 0; x < len(derecha); x++ {
		arreglo[i] = derecha[x]
		i++
	}
	return arreglo
}
