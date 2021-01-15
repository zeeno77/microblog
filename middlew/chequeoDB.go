package middlew

import (
	"net/http"

	"github.com/zeeno77/microblog/db"
)

/*ChequeoDB es el middleware que me permite conocer el estado de la BD */
func ChequeoDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if db.ChequeoConnection() == 0 {
			http.Error(w, "Conexión perdida con la base de datos", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
