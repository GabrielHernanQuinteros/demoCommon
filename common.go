package common

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gorilla/mux"

	"strconv"
)

const AllowedCORSDomain = "http://localhost"

func Hola() {
	fmt.Println("test mod")
}

func ConectarDB(parConnectionString string) (*sql.DB, error) {
	return sql.Open("mysql", parConnectionString)
}

//===================================================================================================
// Funciones de CORS

func EnableCORS(parRouter *mux.Router) {

	parRouter.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", AllowedCORSDomain)
	}).Methods(http.MethodOptions)

	parRouter.Use(MiddlewareCors)

}

func MiddlewareCors(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, req *http.Request) {
			// Just put some headers to allow CORS...
			w.Header().Set("Access-Control-Allow-Origin", AllowedCORSDomain)
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
			// and call next handler!
			next.ServeHTTP(w, req)
		})

}

//===================================================================================================
// Funciones de respuesta

func RespondWithError(parError error, parWriter http.ResponseWriter) {

	parWriter.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(parWriter).Encode(parError.Error())

}

func RespondWithSuccess(parDato interface{}, parWriter http.ResponseWriter) {

	parWriter.WriteHeader(http.StatusOK)
	json.NewEncoder(parWriter).Encode(parDato)

}

//===================================================================================================

func StringToInt64(parCadena string) (int64, error) {

	auxNumero, err := strconv.ParseInt(parCadena, 0, 64)

	if err != nil {
		return 0, err
	}

	return auxNumero, err
}

func InterfaceToInt64(parInterface interface{}) (int64, error) {

	switch parInterface := parInterface.(type) { // This is a type switch.
	case int64:
		return parInterface, nil // All done if we got an int64.
	case int:
		return int64(parInterface), nil // This uses a conversion from int to int64
	case string:
		return strconv.ParseInt(parInterface, 10, 64)
	default:
		return 0, fmt.Errorf("Tipo %T no soportada", parInterface)
	}
}
