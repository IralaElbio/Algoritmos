package diccionario

// URL = https://github.com/theTardigrade/golang-hash

const (
	uint64Prime  uint64 = 0x00000100000001b3
	uint64Offset uint64 = 0xcbf29ce484222325
)
const (
	uint32Prime  uint32 = 0x01000193
	uint32Offset uint32 = 0x811c9dc5
)

func fnvUint64(data []byte) (hash uint64) {
	hash = uint64Offset

	for _, b := range data {
		hash ^= uint64(b)
		hash *= uint64Prime
	}

	return
}

func fnvUint32(data []byte) (hash uint32) {
	hash = uint32Offset

	for _, b := range data {
		hash ^= uint32(b)
		hash *= uint32Prime
	}

	return
}
