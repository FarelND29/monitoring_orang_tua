package NPM

import (
	"fmt"
	"testing"

	"github.com/FarelND29/monitoring_orang_tua/model"
	"github.com/FarelND29/monitoring_orang_tua/module"
)

func TestInsertMahasiswa(t *testing.T) {
	nama := "Farel"
	npm := 1214070
	jenis_kelamin := "laki-laki"
	phone_number := "083821157026"
	hasil := module.InsertMahasiswa(module.MongoConn, "mahasiswa", nama, npm, jenis_kelamin, phone_number)
	fmt.Println(hasil)
}

func TestInsertOrangTua(t *testing.T) {
	nama_ot := "ronaldo"
	phone_number := "0888321157026"
	anak := model.Mahasiswa{
		Nama:         "banner",
		NPM:          1214038,
		Phone_number: "083821157026",
	}
	hasil := module.InsertOrangTua(module.MongoConn, "orangtua", nama_ot, phone_number, anak)
	fmt.Println(hasil)
}

func TestInsertDosenWali(t *testing.T) {
	nama_dosen := "Fatimah"
	alamat := "jalan kota kapur"
	phone_number := "083821157026"
	email := "fatimah.uhuy@gmail.com"
	hasil := module.InsertDosenWali(module.MongoConn, "dosenwali", nama_dosen, alamat, phone_number, email)
	fmt.Println(hasil)
}

func TestInsertTema(t *testing.T) {
	nama_tema := "Web Service"
	hasil := module.InsertTema(module.MongoConn, "tema", nama_tema)
	fmt.Println(hasil)
}
func TestInsertMonitoring(t *testing.T) {
	orang_tua := model.OrangTua{
		Nama_OT:      "Rini",
		Phone_number: "0888321157026",
		Anak: model.Mahasiswa{
			Nama:         "Farel",
			NPM:          1214038,
			Phone_number: "083821157026",
		},
	}
	tema := model.Tema{
		Nama_Tema: "Web Service",
	}
	tanggal := "06-10-2023"
	hari := "Kamis"
	hasil := module.InsertMonitoring(module.MongoConn, "monitoring", orang_tua, tema, tanggal, hari)
	fmt.Println(hasil)
}
func TestGetMahasiswaFromNpm(t *testing.T) {
	npm := 1214070
	data := module.GetMahasiswaFromNpm(module.MongoConn, "mahasiswa", npm)
	fmt.Println(data)
}

func TestGetOrangTuaFromNamaMahasiswa(t *testing.T) {
	nama := "toni"
	data := module.GetOrangTuaFromNamaMahasiswa(module.MongoConn, "orangtua", nama)
	fmt.Println(data)
}

func TestGetDosenWaliFromNamaDosen(t *testing.T) {
	nama_dosen := "Fatimah"
	data := module.GetDosenWaliFromNamaDosen(module.MongoConn, "dosenwali", nama_dosen)
	fmt.Println(data)
}

func TestGetTemaFromNamaTema(t *testing.T) {
	nama_tema := "Web Service"
	data := module.GetTemaFromNamaTema(module.MongoConn, "tema", nama_tema)
	fmt.Println(data)
}

func TestGetMonitoringFromNamaMahasiswa(t *testing.T) {
	nama := "Farel"
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
