package karyawancontroller

import (
	"go-web-native/entities"
	"go-web-native/models/karyawanmodel"
	"html/template"
	"net/http"
	"strconv"
)

func Index(w http.ResponseWriter, r *http.Request) {
	karyawans := karyawanmodel.GetAll()
	data := map[string]any{
		"karyawans": karyawans,
	}

	temp, err := template.ParseFiles("views/karyawan/index.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, data)
}

func Add(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/karyawan/create.html")
		if err != nil {
			panic(err)
		}
		temp.Execute(w, nil)
	}

	if r.Method == "POST" {
		var karyawan entities.Karyawan
		karyawan.Nama = r.FormValue("nama")
		karyawan.Alamat = r.FormValue("alamat")
		karyawan.Posisi = r.FormValue("posisi")

		if berhasil := karyawanmodel.Create(karyawan); !berhasil {
			temp, _ := template.ParseFiles("views/karyawan/create.html")
			temp.Execute(w, nil)
		}
		http.Redirect(w, r, "/karyawan", http.StatusSeeOther)
	}
}

func Edit(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/karyawan/edit.html")
		if err != nil {
			panic(err)
		}

		idString := r.URL.Query().Get("id_karyawan")
		id, err := strconv.Atoi(idString)
		if err != nil {
			panic(err)
		}

		karyawan := karyawanmodel.Detail(id)
		data := map[string]any{
			"karyawan": karyawan,
		}
		temp.Execute(w, data)
	}

	if r.Method == "POST" {
		var karyawan entities.Karyawan

		idString := r.FormValue("id_karyawan")
		id, err := strconv.Atoi(idString)
		if err != nil {
			panic(err)
		}

		karyawan.Nama = r.FormValue("nama")
		karyawan.Alamat = r.FormValue("alamat")
		karyawan.Posisi = r.FormValue("posisi")

		if berhasil := karyawanmodel.Update(id, karyawan); !berhasil {
			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusSeeOther)
			return
		}
		http.Redirect(w, r, "/karyawan", http.StatusSeeOther)
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("id_karyawan")
	id, err := strconv.Atoi(idString)
	if err != nil {
		panic(err)
	}

	if err := karyawanmodel.Delete(id); err != nil {
		panic(err)
	}
	http.Redirect(w, r, "/karyawan", http.StatusSeeOther)
}
