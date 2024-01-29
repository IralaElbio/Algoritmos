package errores

type ErrorLeerArchivo struct{}

func (e ErrorLeerArchivo) Error() string {
	return "ERROR: Lectura de archivos"
}

type ErrorUsuarioInexistente struct{}

func (e ErrorUsuarioInexistente) Error() string {
	return "Error: usuario no existente"
}

type ErrorUsuarioLogeado struct{}

func (e ErrorUsuarioLogeado) Error() string {
	return "Error: Ya habia un usuario loggeado"
}

type ErrorUsuarioNoLogeado struct{}

func (e ErrorUsuarioNoLogeado) Error() string {
	return "Error: no habia usuario loggeado"
}

type ErrorPostUsuario struct{}

func (e ErrorPostUsuario) Error() string {
	return "Error: Usuario no loggeado o Post inexistente"
}

type ErrorPostInexistente struct{}

func (e ErrorPostInexistente) Error() string {
	return "Error: Post inexistente o sin likes"
}
