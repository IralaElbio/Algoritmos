package funciones

import (
	errores "algogram/errores"
	ABB "algogram/tdas/abb"
	TDAHash "algogram/tdas/hash"
	TDAUsuario "algogram/usuario"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const TAMANIOHISTORIAL = 20
const AFINIDADINICIAL = 1

// Abre la ruta y crea los elementos necesarios para el funcionamiento de la APP, caso contrario devuelve el error correspondiente
func IniciarApp(ruta string) (TDAHash.Diccionario[string, TDAUsuario.Usuario], []*TDAUsuario.Post, error) {
	archivo, err := os.Open(ruta)
	if err != nil {
		return nil, nil, errores.ErrorLeerArchivo{}
	}

	defer archivo.Close()

	dict := TDAHash.CrearHash[string, TDAUsuario.Usuario]()

	historialPost := make([]*TDAUsuario.Post, TAMANIOHISTORIAL)

	scan := bufio.NewScanner(archivo)

	numAfinidad := AFINIDADINICIAL

	for scan.Scan() {
		nombreUsuario := scan.Text()
		dict.Guardar(nombreUsuario, TDAUsuario.Crearusuario(nombreUsuario, numAfinidad))
		numAfinidad++

	}
	return dict, historialPost, nil

}

// Divide el texto que recibe en el primer espacio que encuentra
func LeerEntrada(texto string) (string, string) {

	parametros := strings.Split(texto, " ")

	parametro2 := strings.Join(parametros[1:], " ")

	return parametros[0], parametro2
}

// Verifica si el usuario existe o si su nombre se encuentra en el diccionario de usuarios,
// Si el usuario ya existe devuelve el error correspondiente, caso contrario
// devuelve el usuario al cual le pertenece dicho nombre
func AsignarUsuario(usuario TDAUsuario.Usuario, nombreUsuario string, usuarios TDAHash.Diccionario[string, TDAUsuario.Usuario]) (error, TDAUsuario.Usuario) {
	if usuario != nil {
		return errores.ErrorUsuarioLogeado{}, usuario
	}

	if usuarios.Pertenece(nombreUsuario) {
		usuario = usuarios.Obtener(nombreUsuario)
		return nil, usuario
	}
	return errores.ErrorUsuarioInexistente{}, usuario
}

// Recibe un usuario y verifica si existe en caso de cumplirse devuelve nil,
// Caso contrario el error correspondiente
func QuitarUsuario(usuario TDAUsuario.Usuario) (error, TDAUsuario.Usuario) {
	if usuario == nil {
		return errores.ErrorUsuarioNoLogeado{}, usuario
	}
	return nil, nil
}

// Crea un post para el usuario, lo publica y asigna dicho post al historial,
// Si el usuario no existe devuelve el error correspondiente
func CrearPost(usuario TDAUsuario.Usuario, texto string, idPost int, usuarios TDAHash.Diccionario[string, TDAUsuario.Usuario], historial *[]*TDAUsuario.Post) error {
	if usuario == nil {
		return errores.ErrorUsuarioNoLogeado{}
	}
	if idPost >= len(*historial) {
		redimensionarhistorial(historial)
	}
	post := new(TDAUsuario.Post)
	post.ID = idPost
	post.Texto = texto
	post.Likes = ABB.CrearABB[string, string](strings.Compare)

	(*historial)[idPost] = post
	usuario.Postear(post, usuarios)
	return nil
}

// Regresa el texto formateado para mostrar en pantalla del post que esta proximo al feed del usuario, si el usuario no existe o su feed esta vacio
// devuelve el mensaje correspondiente
func MostrarPostEnFeed(usuario TDAUsuario.Usuario, historial []*TDAUsuario.Post) string {
	if usuario == nil || usuario.FeedVacio() {
		return "Usuario no loggeado o no hay mas posts para ver"
	}
	post := usuario.VerProximoFeed()
	return fmt.Sprintf("Post ID %v\n%v dijo: %v\nLikes: %v", post.ID, post.Autor.Nombre(), post.Texto, historial[post.ID].CantidadLikes())
}

// Asigna el like del usuario al post con la id correspondiente,
// Si el usuario o el post no existen devuelve el error correspondiente,
func LikearPost(postID string, usuario TDAUsuario.Usuario, historial []*TDAUsuario.Post) error {
	index, err := strconv.Atoi(postID)
	if usuario == nil || err != nil || historial[index] == nil {
		return errores.ErrorPostUsuario{}

	}
	historial[index].LikearPost(usuario.Nombre())
	return nil
}

// Busca el post en el historial y regresa una lista con todas las personas que le dieron like a dicho post,
// si el post no existe o tiene 0 likes devolvera el error correspodiente
func BuscarLikes(postID string, usuario TDAUsuario.Usuario, historial []*TDAUsuario.Post) (error, []string) {
	index, err := strconv.Atoi(postID)
	if err != nil || historial[index] == nil || index >= len(historial) || historial[index].CantidadLikes() == 0 {
		return errores.ErrorPostInexistente{}, nil
	}
	likesEnPost := historial[index].ObtenerLikes()

	return nil, likesEnPost
}

// Regresa el texto formateado para mostrar en pantalla de los likes que tiene un post
// y los nombres de los usuarios que dieron like
func FormatearLikes(totalLikes []string) string {
	salidaTexto := fmt.Sprintf("El post tiene %v likes:", len(totalLikes))
	for _, usuario := range totalLikes {
		salidaTexto += fmt.Sprintf("\n\t%v", usuario)
	}
	return salidaTexto
}

// Verifica que no se haya cometido un error durante la ejecucion del programa,
// si un error se cometio devolvera el mensaje de dicho error, caso contrario
// el texto
func ObtenerResultado(err error, texto string) string {
	if err != nil {
		return err.Error()
	}
	return texto
}

// Redimensiona el historial de likes
func redimensionarhistorial(historial *[]*TDAUsuario.Post) {
	historialNuevo := make([]*TDAUsuario.Post, len(*historial)*2)
	copy(historialNuevo, *historial)
	*historial = historialNuevo
}
