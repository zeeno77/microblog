package routers

import (
	"encoding/json"
	"net/http"

	"github.com/zeeno77/microblog/db"
)

/*VerPerfil - permite extraer los valores del perfil*/
func VerPerfil(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "debe enviar el parámetro ID", http.StatusBadRequest)
		return
	}
	// si todo anduvo bien:
	perfil, err := db.BuscoPerfil(ID)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar buscar el registro"+err.Error(), 400)
		return
	}
	w.Header().Set("context-type", "encoding/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(perfil)
}
