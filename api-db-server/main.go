package main

import (
	"go-was-example/api-db-server/database"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// DB Instance
	rdbHandler := database.RDBHandler{
		UserName:      "dev",
		Password:      "jeongchul!@#",
		ServerAddress: "[GOOGLE_CLOUD_SQL_INTERNAL_IP]",
		DbName:        "dev",
	}

	rdbHandler.Connect()

	ras := RestApiServer{
		rdbHandler: &rdbHandler,
	}

	ras.router = mux.NewRouter()
	ras.router.HandleFunc("/", ras.hello)

	port := "80"
	log.Println("Server Start " + port)
	log.Fatal(http.ListenAndServe(":"+port, ras.router))
}

type RestApiServer struct {
	rdbHandler *database.RDBHandler
	router     *mux.Router
}

// hello / hello 핸들러
func (ras RestApiServer) hello(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		exist, users, errDB := ras.rdbHandler.GetUser()
		if errDB != nil {
			w.Write([]byte("Internal Server error 500"))
		}
		log.Println("User Query Result=", users)

		resultMsg := ""
		if exist {
			resultMsg += "Hello"
			for _, user := range users {
				resultMsg += ", " + user.UserName
			}
		} else {
			resultMsg = "Hello, World!"
		}
		log.Println("/ api call hello handler response msg=" + resultMsg)
		w.Write([]byte(resultMsg))
	}
}
