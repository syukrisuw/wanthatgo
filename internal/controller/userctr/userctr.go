package userctr

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func UserCtr(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["key"]

	if !ok || len(keys[0]) < 1 {
		log.Println("Url Param 'key' is missing")
		return
	}

	// Query()["key"] will return an array of items,
	// we only want the single item.
	key := keys[0]

	log.Println("Url Param 'key' is: " + string(key))
	fmt.Fprintf(w, "key = %s\n", key)
}

func UserCtrById(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/user/")

	log.Println("Url Param 'id' is: " + string(id))
	fmt.Fprintf(w, "ID = %s\n", id)
}
