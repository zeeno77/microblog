package routers

import (
	"encoding/json"
	"net/http"

	"github.com/zeeno77/microblog/db"
	"github.com/zeeno77/microblog/models"
)

/*ModificarPerfil - modifica el perfil de usuario*/
func BorrarTweet(w http.ResponseWriter, r *http.Request) {

	var twt models.Tweet

	err := json.NewDecoder(r.Body).Decode(&twt)
	if err != nil {
		// es un Json mal construído
		http.Error(w, "Datos incorrectos "+err.Error(), 400)
		return
	}

	var status bool

	println(twt.Mensaje)

	status, err = db.BorroTweet(twt, IDUsuario)
	//IDUsuario es la variable global que seteamos antes con el ID
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar modificar el registro. Intente nuevamente "+err.Error(), 400)
		return
	}
	if status == false {
		http.Error(w, "No se ha logrado modificar elregistro del usuario "+err.Error(), 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
