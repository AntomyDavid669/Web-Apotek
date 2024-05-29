package stockmodel

import (
	"go-web-native/config"
	"go-web-native/entities"
)

func GetAll() []entities.Stock{
	rows, err := config.DB.Query(`SELECT * FROM obat`)
	if err != nil{
		panic(err)
	}

	defer rows.Close()

	var stocks []entities.Stock

	for rows.Next(){
		var stock entities.Stock
		err := rows.Scan(&stock.Id, &stock.Nama, &stock.Jenis_obat, &stock.Harga, &stock.Stok)
		if err != nil{
			panic(err)
		}	

		stocks = append(stocks, stock)
	}

	return stocks
}

func Create(stok entities.Stock) bool{
	var lastInsertID int64

	err := config.DB.QueryRow(`INSERT INTO obat (nama, jenis_obat, harga, stok)
	 VALUES($1, $2, $3, $4) RETURNING id`,
	stok.Nama, stok.Jenis_obat, stok.Harga, stok.Stok).Scan(&lastInsertID)

	if err != nil{
		panic(err)
	}
	return lastInsertID > 0
}

func Detail(id int) entities.Stock{
	row := config.DB.QueryRow(`SELECT id, nama, jenis_obat, harga, stok FROM obat WHERE id = $1`, id)

	var stock entities.Stock
	if err := row.Scan(&stock.Id, &stock.Nama, &stock.Jenis_obat, &stock.Harga, &stock.Stok); err != nil{
		panic(err.Error())
	}
	return stock
}

func Update(id int, stock entities.Stock) bool{
	query, err := config.DB.Exec(`UPDATE obat SET nama = $1, jenis_obat = $2, harga = $3, stok = $4 WHERE id = $5`,
stock.Nama, stock.Jenis_obat, stock.Harga, stock.Stok, id)
	if err != nil{
		panic(err)
	}
	
	result, err := query.RowsAffected()
	if err != nil{
		panic(err)
	}
	return result > 0
}

func Delete(id int) error{
	_, err := config.DB.Exec(`DELETE FROM obat WHERE id= $1`, id)
	return err
}

