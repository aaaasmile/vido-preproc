package srv

import (
	"fmt"
	"log"
	"net/http"
)

func HandleIndex(w http.ResponseWriter, req *http.Request) {
	var err error
	switch req.Method {
	case "GET":
		err = handleIndexGet(w, req)
	case "POST":
		err = handlePostOperations(w, req)
	}
	if err != nil {
		log.Println("Handler error: ", err)
		http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusInternalServerError)
	}
}

func InitFromConfig(debug bool) error {
	log.Println("Handler initialized", debug)
	return nil
}
