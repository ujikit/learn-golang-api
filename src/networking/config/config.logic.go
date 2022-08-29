package networking

import (
	"fmt"
	http "net/http"
	json "encoding/json"
	httprouter "github.com/julienschmidt/httprouter"

	// Constant
	constant "expense-tracker-api/src/constant"

	// Service
	service "expense-tracker-api/src/service"
)

const ServerHostAndPort string = "localhost:" + constant.ServerPort
const message string = "\nRunning on http://" + ServerHostAndPort + "\n\n"

func RunServer() {
	router := httprouter.New()

	router.GET("/", func(w http.ResponseWriter, request *http.Request, param httprouter.Params) {
		languageCode := request.URL.Query().Get("languageCode")
		userData := service.UserDetail(languageCode)
		convert, _ := json.Marshal(userData)

		w.Header().Set("Content-Type", "application/json")
		w.Write(convert)
	})

	server := http.Server{
		Handler: router,
		Addr:    ServerHostAndPort,
	}

	fmt.Println(message)
	server.ListenAndServe()
}
