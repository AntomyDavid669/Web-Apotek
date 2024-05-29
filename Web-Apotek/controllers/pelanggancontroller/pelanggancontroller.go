package pelanggancontroller

import (
	"go-web-native/entities"
	"go-web-native/models/pelangganmodel"
	"html/template"
	"net/http"
	"strconv"
)

func Index(w http.ResponseWriter, r *http.Request) {
	pelanggans := pelangganmodel.GetAll()
	data := map[string]any{
		"pelanggans": pelanggans,
	}

	temp, err := template.ParseFiles("views/pelanggan/index.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, data)
}

func Add(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/pelanggan/create.html")
		if err != nil {
			panic(err)
		}
		temp.Execute(w, nil)
	}

	if r.Method == "POST" {
		var pelanggan entities.Pelanggan
		pelanggan.Nama = r.FormValue("nama")
		pelanggan.Alamat = r.FormValue("alamat")
		pelanggan.Telepon = r.FormValue("telepon")

		if berhasil := pelangganmodel.Create(pelanggan); !berhasil {
			temp, _ := template.ParseFiles("views/pelanggan/create.html")
			temp.Execute(w, nil)
		}
		http.Redirect(w, r, "/pelanggan", http.StatusSeeOther)
	}
}

func Edit(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET"{
		temp, err := template.ParseFiles("views/pelanggan/edit.html")
		if err != nil {
			panic(err)
		}

		idString := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idString)
		if err != nil{
			panic(err)
		}

		pelanggan := pelangganmodel.Detail(id)
		data := map[string]any{
			"pelanggan": pelanggan,
		}
		temp.Execute(w, data)
	}

	if r.Method == "POST"{
		var pelanggan entities.Pelanggan

		idString := r.FormValue("id")
		id, err := strconv.Atoi(idString)
		if err != nil{
			panic(err)
		}

		pelanggan.Nama = r.FormValue("nama")
		pelanggan.Alamat = r.FormValue("alamat")
		pelanggan.Telepon = r.FormValue("telepon")
		
		if berhasil := pelangganmodel.Update(id, pelanggan); !berhasil {
			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusSeeOther)
			return
		}
		http.Redirect(w, r, "/pelanggan", http.StatusSeeOther)
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idString)
	if err != nil{
		panic(err)
	}

	if err := pelangganmodel.Delete(id); err != nil{
		panic(err)
	}
	http.Redirect(w, r, "/pelanggan", http.StatusSeeOther)
}
