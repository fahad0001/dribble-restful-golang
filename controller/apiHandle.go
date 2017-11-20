package controller

import (
	"github.com/gorilla/mux"
	"net/http"
	"log"
	"gallery-app/dataDir"
	"encoding/json"
)

type photoCallHandler struct {
	dataSource string
}

func (ph photoCallHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	title := vars["title"]
	desc := vars["description"]
	resp, _ := dataDir.GetPhotoDetailByCriteria(ph.dataSource, title, desc)
	json.NewEncoder(writer).Encode(resp)
}

func APIBuilder(dataSource string) {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("This is goAPP which retrieves dribble image info"))
	})
	router.Handle("/api/photo/{title}/{description}", &photoCallHandler{dataSource})
	router.Handle("/api/photo/{title}", &photoCallHandler{dataSource})
	router.Handle("/api/photo/{description}", &photoCallHandler{dataSource})
	log.Fatal(http.ListenAndServe(":8000", router))
}