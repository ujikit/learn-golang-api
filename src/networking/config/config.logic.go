package networking

import (
	"fmt"
	context "context"
	http "net/http"
	json "encoding/json"
	httprouter "github.com/julienschmidt/httprouter"

	// Constant
	constant "learn-golang-api/src/constant"

	// Service
	service "learn-golang-api/src/service"
)

const ServerHostAndPort string = "localhost:" + constant.ServerPort
const message string = "\nRunning on http://" + ServerHostAndPort + "\n\n"

func RunServer() {
	router := httprouter.New()

	router.GET("/", func(w http.ResponseWriter, request *http.Request, param httprouter.Params) {
		languageCode := request.URL.Query().Get("languageCode")
		ctx := context.Background()
		ctx = context.WithValue(ctx, "languageCode", languageCode)
		userData := service.UserDetail(ctx)
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
