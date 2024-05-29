package main

import (
	"go-web-native/config"
	"go-web-native/controllers/homecontroller"
	"go-web-native/controllers/karyawancontroller"
	"go-web-native/controllers/pelanggancontroller"
	"go-web-native/controllers/stockcontroller"
	"log"
	"net/http"
)

func main() {
	config.ConnectDB()

	//HomePage
	http.HandleFunc("/", homecontroller.Welcome)
	
	//Karyawan
	http.HandleFunc("/karyawan", karyawancontroller.Index)
	http.HandleFunc("/karyawan/add", karyawancontroller.Add)
	http.HandleFunc("/karyawan/edit", karyawancontroller.Edit)
	http.HandleFunc("/karyawan/delete", karyawancontroller.Delete)

	//Pelanggan
	http.HandleFunc("/pelanggan", pelanggancontroller.Index)
	http.HandleFunc("/pelanggan/add", pelanggancontroller.Add)
	http.HandleFunc("/pelanggan/edit", pelanggancontroller.Edit)
	http.HandleFunc("/pelanggan/delete", pelanggancontroller.Delete)

	//Stok Obat
	http.HandleFunc("/stok", stockcontroller.Index)
	http.HandleFunc("/stok/add", stockcontroller.Add)
	http.HandleFunc("/stok/edit", stockcontroller.Edit)
	http.HandleFunc("/stok/delete", stockcontroller.Delete)

	log.Println("Server running on port 8080")
	http.ListenAndServe(":8080", nil)
}
