package NPM

import (
	"fmt"
	"testing"

	"github.com/FarelND29/monitoring_orang_tua/model"
	"github.com/FarelND29/monitoring_orang_tua/module"
)

func TestInsertMahasiswa(t *testing.T) {
	nama := "Rio Riyanto"
	npm := 1214015
	jenis_kelamin := "laki-laki"
	phone_number := "083821132343"
	hasil := module.InsertMahasiswa(module.MongoConn, "mahasiswa", nama, npm, jenis_kelamin, phone_number)
	fmt.Println(hasil)
}

func TestInsertOrangTua(t *testing.T) {
	nama_ot := "Haryanto Subathon"
	phone_number := "088821151212"
	anak := model.Mahasiswa{
		Nama:         "Rio Riyanto",
		NPM:          1214013,
		Phone_number: "083821132343",
	}
	hasil := module.InsertOrangTua(module.MongoConn, "orangtua", nama_ot, phone_number, anak)
	fmt.Println(hasil)
}

func TestInsertDosenWali(t *testing.T) {
	nama_dosen := "Munaroh Maija"
	alamat := "jalan Macan"
	phone_number := "08382114428"
	email := "munaroh24@gmail.com"
	hasil := module.InsertDosenWali(module.MongoConn, "dosenwali", nama_dosen, alamat, phone_number, email)
	fmt.Println(hasil)
}

func TestInsertTema(t *testing.T) {
	nama_tema := "Etika Manajemen"
	hasil := module.InsertTema(module.MongoConn, "tema", nama_tema)
	fmt.Println(hasil)
}
func TestInsertMonitoring(t *testing.T) {
	orang_tua := model.OrangTua{
		Nama_OT:      "Haryanto Subathon",
		Phone_number: "088821151212",
		Anak: model.Mahasiswa{
			Nama:         "Rio Riyanto",
			NPM:          1214015,
			Jekel: 		"laki-laki",
			Phone_number: "083821132343",
		},
	}
	tema := model.Tema{
		Nama_Tema: "Etika Manajemen",
	}
	dosen := model.DosenWali{
		Nama_Dosen:   "Munaroh Maija",
		Alamat:       "jalan Macan",
		Phone_number: "08382114428",
		Email:        "munaroh24@gmail.com",
	}
	tanggal := "10-12-2000"
	hari := "Minggu"
	insertedID, err := module.InsertMonitoring(module.MongoConn, "monitoring", orang_tua, tema, dosen, tanggal, hari)
	if err != nil {
		t.Errorf("Error inserting data: %v", err)
	}
	fmt.Printf("Data berhasil disimpan dengan id %s", insertedID.Hex())
}

/* func TestInsertMonitoring(t *testing.T) {
	var orang_tua = model.OrangTua{
			Nama_OT:      		"Haryanto Subathon",
			Phone_number: 		"088821151212",
			Anak: model.Mahasiswa{
				Nama:         	"Rio Riyanto",
				NPM:          	1214015,
				Jekel: 			"laki-laki",	
				Phone_number: 	"083821132343",
			},
	tema := model.Tema{
			Nama_Tema: "Etika Manajemen",
			}
	var dosen 	= model.DosenWali{
				Nama_Dosen:   "Munaroh Maija",
				Alamat:       "jalan Macan",
				Phone_number: "08382114428",
				Email:        "munaroh24@gmail.com",
			}

			tanggal := "10-12-2000"
			hari := "Minggu"
	}
	insertedID, err := module.InsertMonitoring(module.MongoConn, "monitoring", orang_tua, tema, dosen, tanggal, hari)
	if err != nil {
		t.Errorf("Error inserting data: %v", err)
	}
	fmt.Printf("Data berhasil disimpan dengan id %s", insertedID.Hex())
} */

func TestGetMahasiswaFromNpm(t *testing.T) {
	npm := 1214013
	data := module.GetMahasiswaFromNpm(module.MongoConn, "mahasiswa", npm)
	fmt.Println(data)
}

func TestGetOrangTuaFromNamaMahasiswa(t *testing.T) {
	nama := "Rio Riyanto"
	data := module.GetOrangTuaFromNamaMahasiswa(module.MongoConn, "orangtua", nama)
	fmt.Println(data)
}

func TestGetDosenWaliFromNamaDosen(t *testing.T) {
	nama_dosen := "Munaroh Maija"
	data := module.GetDosenWaliFromNamaDosen(module.MongoConn, "dosenwali", nama_dosen)
	fmt.Println(data)
}

func TestGetTemaFromNamaTema(t *testing.T) {
	nama_tema := "Front End"
	data := module.GetTemaFromNamaTema(module.MongoConn, "tema", nama_tema)
	fmt.Println(data)
}

func TestGetMonitoringFromNamaMahasiswa(t *testing.T) {
	nama := "Rio Riyanto"
	data := module.GetMonitoringFromNamaMahasiswa(module.MongoConn, "monitoring", nama)
	fmt.Println(data)
}

func TestGetAllMonitoring(t *testing.T) {
	data := module.GetAllMonitoring(module.MongoConn, "monitoring")
	fmt.Println(data)
}

func TestGetAllMahasiswa(t *testing.T) {
	data := module.GetAllMahasiswa(module.MongoConn, "mahasiswa")
	fmt.Println(data)
}

func TestGetAllOrangTua(t *testing.T) {
	data := module.GetAllOrangTua(module.MongoConn, "orangtua")
	fmt.Println(data)
}

func TestGetAllDosenWali(t *testing.T) {
	data := module.GetAllDosenWali(module.MongoConn, "dosenwali")
	fmt.Println(data)
}

func TestGetAllTema(t *testing.T) {
	data := module.GetAllTema(module.MongoConn, "tema")
	fmt.Println(data)
}
