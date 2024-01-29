package diccionario

func sdbmHash(data []byte) uint64 {
	var hash uint64

	for _, b := range data {
		hash = uint64(b) + (hash << 6) + (hash << 16) - hash
	}

	return hash
}

func SdbmHash(data []byte) int {
	return int(sdbmHash(data))
}
