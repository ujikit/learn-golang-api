package networking

import (
	"fmt"
	http "net/http"

	constant "expense-tracker-api/src/constant"

	httprouter "github.com/julienschmidt/httprouter"
)

const ServerPort string = "localhost:" + constant.ServerPort
const message string = "\nRunning on http://" + ServerPort + "\n\n"

func RunServer() {
	router := httprouter.New()
	router.GET("/", func(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
		fmt.Fprint(writer, "hellow yaes")
	})

	server := http.Server{
		Handler: router,
		Addr:    ServerPort,
	}
	fmt.Printf(message)
	server.ListenAndServe()
}
