package main

import (
	pasienController "example/helloGop/controllers"
	"net/http"
)

func main() {
	http.HandleFunc("/", pasienController.Index)
	http.HandleFunc("/pasien", pasienController.Index)
	http.HandleFunc("/pasien/index", pasienController.Index)
	http.HandleFunc("/pasien/add", pasienController.Add)
	http.HandleFunc("/pasien/edit", pasienController.Edit)
	http.HandleFunc("/pasien/delete", pasienController.Delete)

	http.ListenAndServe(":3000", nil)
}
