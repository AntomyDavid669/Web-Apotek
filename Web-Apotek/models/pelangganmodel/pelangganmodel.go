package pelangganmodel

import (
	"go-web-native/config"
	"go-web-native/entities"
)

func GetAll() []entities.Pelanggan{
	rows, err := config.DB.Query(`SELECT * FROM pelanggan`)
	if err != nil{
		panic(err)
	}

	defer rows.Close()

	var pelanggans []entities.Pelanggan

	for rows.Next(){
		var pelanggan entities.Pelanggan
		err := rows.Scan(&pelanggan.Id, &pelanggan.Nama, &pelanggan.Alamat, &pelanggan.Telepon)
		if err != nil{
			panic(err)
		}	

		pelanggans = append(pelanggans, pelanggan)
	}

	return pelanggans
}

func Create(pelanggan entities.Pelanggan) bool{
	var lastInsertID int64

	err := config.DB.QueryRow(`INSERT INTO pelanggan (nama, alamat, telepon)
	 VALUES($1, $2, $3) RETURNING id`,
	pelanggan.Nama, pelanggan.Alamat, pelanggan.Telepon).Scan(&lastInsertID)

	if err != nil{
		panic(err)
	}
	return lastInsertID > 0
}

func Detail(id int) entities.Pelanggan{
	row := config.DB.QueryRow(`SELECT id, nama, alamat, telepon FROM pelanggan WHERE id = $1`, id)

	var pelanggan entities.Pelanggan
	if err := row.Scan(&pelanggan.Id, &pelanggan.Nama, &pelanggan.Alamat, &pelanggan.Telepon); err != nil{
		panic(err.Error())
	}
	return pelanggan
}

func Update(id int, pelanggan entities.Pelanggan) bool{
	query, err := config.DB.Exec(`UPDATE pelanggan SET nama = $1, alamat = $2, telepon = $3 WHERE id = $4`,
pelanggan.Nama, pelanggan.Alamat, pelanggan.Telepon, id)
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
	_, err := config.DB.Exec(`DELETE FROM pelanggan WHERE id= $1`, id)
	return err
}