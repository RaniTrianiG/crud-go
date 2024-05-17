package models

import (
	"database/sql"
	"example/helloGop/config"
	"example/helloGop/entities"
	"fmt"
)

type PasienModel struct {
	conn *sql.DB
}

func NewPasienModel() *PasienModel {
	conn, err := config.DBConnection()
	if err != nil {
		panic(err)
	}

	return &PasienModel{
		conn: conn,
	}
}

func (p *PasienModel) FindAll() ([]entities.Pasien, error) {
	rows, err := p.conn.Query("select * from pasien")

	if err != nil {
		return []entities.Pasien{}, err
	}
	defer rows.Close()

	var dataPasien []entities.Pasien
	for rows.Next() {
		var pasien entities.Pasien
		rows.Scan(&pasien.Id, &pasien.NamaLengkap, &pasien.NIK, &pasien.Alamat, &pasien.JenisKelamin, &pasien.TanggalLahir, &pasien.TempatLahir, &pasien.NoHP)

	}
}

func (p *PasienModel) Create(pasien entities.Pasien) bool {
	result, err := p.conn.Exec("insert into pasien (nama_lengkap, nik, jenis_kelamin, tempat_lahir, tanggal_lahir, alamat, no_hp) values(?,?,?,?,?,?,?)", pasien.NamaLengkap, pasien.NIK, pasien.JenisKelamin, pasien.TempatLahir, pasien.TanggalLahir, pasien.Alamat, pasien.NoHP)

	if err != nil {
		fmt.Println(err)
		return false
	}

	lastInsertId, _ := result.LastInsertId()

	return lastInsertId > 0
}
