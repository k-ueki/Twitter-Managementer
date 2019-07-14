package router

import (
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type Config struct {
	Port string
}

func main() {

}
func (rou *Config) NewRouter() (*mux.Router, func(http.Handler) http.Handler) {
	r := mux.NewRouter()

	str := "http://localhost"
	str += rou.Port

	allowedOrigins := handlers.AllowedOrigins([]string{str})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})
	allowedHeaders := handlers.AllowedHeaders([]string{"content-type"})

	return r, handlers.CORS(allowedOrigins, allowedMethods, allowedHeaders)
}
