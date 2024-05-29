package stockcontroller

import (
	"go-web-native/entities"
	"go-web-native/models/stockmodel"
	"html/template"
	"net/http"
	"strconv"
)

func Index(w http.ResponseWriter, r *http.Request) {
	stocks := stockmodel.GetAll()
	data := map[string]any{
		"stocks": stocks,
	}

	temp, err := template.ParseFiles("views/stock/index.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, data)
}

func Add(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/stock/create.html")
		if err != nil {
			panic(err)
		}
		temp.Execute(w, nil)
	}

	if r.Method == "POST" {
		var stok entities.Stock
		harga, err := strconv.Atoi(r.FormValue("harga"))
		if err != nil {
			panic(err)
		}

		stock, err := strconv.Atoi(r.FormValue("stock"))
		if err != nil {
			panic(err)
		}

		stok.Nama = r.FormValue("nama")
		stok.Jenis_obat = r.FormValue("jenisobat")
		stok.Harga = uint64(harga)
		stok.Stok = uint64(stock)

		if berhasil := stockmodel.Create(stok); !berhasil {
			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusTemporaryRedirect)
		}
		http.Redirect(w, r, "/stok", http.StatusSeeOther)
	}

}

func Edit(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/stock/edit.html")
		if err != nil {
			panic(err)
		}

		idString := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			panic(err)
		}

		stock := stockmodel.Detail(id)
		data := map[string]any{
			"stock": stock,
		}
		temp.Execute(w, data)
	}

	if r.Method == "POST" {
		var stok entities.Stock

		idString := r.FormValue("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			panic(err)
		}

		harga, err := strconv.Atoi(r.FormValue("harga"))
		if err != nil {
			panic(err)
		}

		stock, err := strconv.Atoi(r.FormValue("stock"))
		if err != nil {
			panic(err)
		}

		stok.Nama = r.FormValue("nama")
		stok.Jenis_obat = r.FormValue("jenisobat")
		stok.Harga = uint64(harga)
		stok.Stok = uint64(stock)

		if berhasil := stockmodel.Update(id, stok); !berhasil {
			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusSeeOther)
			return
		}
		http.Redirect(w, r, "/stok", http.StatusSeeOther)
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		panic(err)
	}

	if err := stockmodel.Delete(id); err != nil {
		panic(err)
	}
	http.Redirect(w, r, "/stok", http.StatusSeeOther)

}
