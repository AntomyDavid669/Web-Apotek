package karyawanmodel

import (
	"go-web-native/config"
	"go-web-native/entities"
)

func GetAll() []entities.Karyawan{ 
	rows, err := config.DB.Query(`SELECT * FROM karyawan`)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var karyawans []entities.Karyawan

	for rows.Next() {
		var karyawan entities.Karyawan
		err := rows.Scan(&karyawan.Id_karyawan, &karyawan.Nama, &karyawan.Alamat, &karyawan.Posisi)
		if err != nil {
			panic(err)
		}

		karyawans = append(karyawans, karyawan)
	}

	return karyawans
}

func Create(karyawan entities.Karyawan) bool{
	var lastInsertID int64

	err := config.DB.QueryRow(`INSERT INTO karyawan (nama, alamat, posisi)
	 VALUES($1, $2, $3) RETURNING id_karyawan`,
	karyawan.Nama, karyawan.Alamat, karyawan.Posisi).Scan(&lastInsertID)

	if err != nil{
		panic(err)
	}
	return lastInsertID > 0
}

func Detail(id_karyawan int) entities.Karyawan{
	row := config.DB.QueryRow(`SELECT id_karyawan, nama, alamat, posisi FROM karyawan WHERE id_karyawan = $1`, id_karyawan)

	var karyawan entities.Karyawan
	if err := row.Scan(&karyawan.Id_karyawan, &karyawan.Nama, &karyawan.Alamat, &karyawan.Posisi); err != nil{
		panic(err.Error())
	}
	return karyawan
}

func Update(id_karyawan int, karyawan entities.Karyawan) bool{
	query, err := config.DB.Exec(`UPDATE karyawan SET nama = $1, alamat = $2, posisi = $3 WHERE id_karyawan = $4`,
karyawan.Nama, karyawan.Alamat, karyawan.Posisi, id_karyawan)
	if err != nil{
		panic(err)
	}
	
	result, err := query.RowsAffected()
	if err != nil{
		panic(err)
	}
	return result > 0
}
func Delete(id_karyawan int) error{
	_, err := config.DB.Exec(`DELETE FROM karyawan WHERE id_karyawan= $1`, id_karyawan)
	return err
}
