package webserver

import "net/http"

func MiWebServer() {
	http.HandleFunc("/", home)        //definimos una ruta y estamos escuchando home
	http.ListenAndServe(":3000", nil) //escuchar y servir el puerto 3000
}

func home(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./webserver/index.html")
}
