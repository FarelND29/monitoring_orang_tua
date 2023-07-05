package NPM

import (
	"fmt"
	"testing"

	"github.com/FarelND29/monitoring_orang_tua/model"
	"github.com/FarelND29/monitoring_orang_tua/module"
)

func TestInsertMahasiswa(t *testing.T) {
	nama := "Mamas Racing"
	npm := 1214090
	jenis_kelamin := "perempuan"
	phone_number := "08382098789"
	hasil := module.InsertMahasiswa(module.MongoConn, "mahasiswa", nama, npm, jenis_kelamin, phone_number)
	fmt.Println(hasil)
}

func TestInsertOrangTua(t *testing.T) {
	nama_ot := "Bimbim Metal"
	phone_number := "081342535636"
	anak := model.Mahasiswa{
		Nama:         	"Mamas Racing",
		NPM:          	1214090,
		Jekel: 			"Perempuan",	
		Phone_number: 	"08382098789",
	}
	hasil := module.InsertOrangTua(module.MongoConn, "orangtua", nama_ot, phone_number, anak)
	fmt.Println(hasil)
}

func TestInsertDosenWali(t *testing.T) {
	nama_dosen := "teteh Hardcore"
	alamat := "jalan tiktok"
	phone_number := "08386346775"
	email := "tetehaja12@gmail.com"
	hasil := module.InsertDosenWali(module.MongoConn, "dosenwali", nama_dosen, alamat, phone_number, email)
	fmt.Println(hasil)
}

func TestInsertTema(t *testing.T) {
	nama_tema := "Bahasa Jerman"
	hasil := module.InsertTema(module.MongoConn, "tema", nama_tema)
	fmt.Println(hasil)
}
func TestInsertMonitoring(t *testing.T) {
	orang_tua := model.OrangTua{
		Nama_OT:      "Bimbim Metal",
		Phone_number: "081342535636",
		Anak: model.Mahasiswa{
			Nama:         	"Mamas Racing",
			NPM:          	1214015,
			Jekel: 			"laki-laki",
			Phone_number: 	"083821132343",
		},
	}
	tema := model.Tema{
		Nama_Tema: "Bahasa Jerman",
	}
	dosen := model.DosenWali{
		Nama_Dosen:   "teteh Hardcore",
		Alamat:       "jalan tiktok",
		Phone_number: "08386346775",
		Email:        "tetehaja12@gmail.com",
	}
	tanggal := "17-11-2025"
	hari := "Senin"
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
	npm := 1214090
	data := module.GetMahasiswaFromNpm(module.MongoConn, "mahasiswa", npm)
	fmt.Println(data)
}

func TestGetOrangTuaFromNamaMahasiswa(t *testing.T) {
	nama := "Mamas Racing"
	data := module.GetOrangTuaFromNamaMahasiswa(module.MongoConn, "orangtua", nama)
	fmt.Println(data)
}

func TestGetDosenWaliFromNamaDosen(t *testing.T) {
	nama_dosen := "teteh Hardcore"
	data := module.GetDosenWaliFromNamaDosen(module.MongoConn, "dosenwali", nama_dosen)
	fmt.Println(data)
}

func TestGetTemaFromNamaTema(t *testing.T) {
	nama_tema := "Bahasa Jerman"
	data := module.GetTemaFromNamaTema(module.MongoConn, "tema", nama_tema)
	fmt.Println(data)
}

func TestGetMonitoringFromNamaMahasiswa(t *testing.T) {
	nama := "Mamas Racing"
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
