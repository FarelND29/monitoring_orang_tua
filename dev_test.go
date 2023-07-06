package NPM

import (
	"fmt"
	"testing"

	"github.com/FarelND29/monitoring_orang_tua/model"
	"github.com/FarelND29/monitoring_orang_tua/module"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestInsertMahasiswa(t *testing.T) {
	nama := "Fahri Fahrian"
	npm := 1214011
	jenis_kelamin := "laki laki"
	phone_number := "083823545542"
	hasil := module.InsertMahasiswa(module.MongoConn, "mahasiswa", nama, npm, jenis_kelamin, phone_number)
	fmt.Println(hasil)
}

func TestInsertOrangTua(t *testing.T) {
	nama_ot := "Stephen Sabun"
	phone_number := "081342565098"
	anak := model.Mahasiswa{
		Nama:         	"Fahri Fahrian",
		NPM:          	1214011,
		Jekel: 			"laki laki",	
		Phone_number: 	"083823545542",
	}
	hasil := module.InsertOrangTua(module.MongoConn, "orangtua", nama_ot, phone_number, anak)
	fmt.Println(hasil)
}

func TestInsertDosenWali(t *testing.T) {
	nama_dosen := "Joji Yuda"
	alamat := "jalan setia"
	phone_number := "08386347643"
	email := "joji12@gmail.com"
	hasil := module.InsertDosenWali(module.MongoConn, "dosenwali", nama_dosen, alamat, phone_number, email)
	fmt.Println(hasil)
}

func TestInsertTema(t *testing.T) {
	nama_tema := "Bahasa Spanyol"
	hasil := module.InsertTema(module.MongoConn, "tema", nama_tema)
	fmt.Println(hasil)
}
func TestInsertMonitoring(t *testing.T) {
	orang_tua := model.OrangTua{
		Nama_OT:      "Stephen Sabun",
		Phone_number: "081342565098",
		Anak: model.Mahasiswa{
			Nama:         	"Fahri Fahrian",
			NPM:          	1214011,
			Jekel: 			"laki-laki",
			Phone_number: 	"083823545542",
		},
	}
	tema := model.Tema{
		Nama_Tema: "Bahasa Spanyol",
	}
	dosen := model.DosenWali{
		Nama_Dosen:   "Joji Yuda",
		Alamat:       "jalan setia",
		Phone_number: "08386347643",
		Email:        "joji12@gmail.com",
	}
	tanggal := "20-11-2025"
	hari := "Kamis"
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

func TestGetMahasiswaFromID(t *testing.T) {
	id := "64a4e2e694cb7dd7f0d0f9f5"
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		t.Fatalf("error converting id to ObjectID: %v", err)
	}
	mahasiswa, err := module.GetMahasiswaFromID(objectID, module.MongoConn, "mahasiswa")
	if err != nil {
		t.Fatalf("error calling GetMahasiswaFromID: %v", err)
	}
	fmt.Println(mahasiswa)
}

func TestGetOrangTuaFromID(t *testing.T) {
	id := "64134e6931b8ceb1ad63d3ad"
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		t.Fatalf("error converting id to ObjectID: %v", err)
	}
	orangtua, err := module.GetOrangTuaFromID(objectID, module.MongoConn, "orangtua")
	if err != nil {
		t.Fatalf("error calling GetOrangTuaFromID: %v", err)
	}
	fmt.Println(orangtua)
}

func TestGetDosenWaliFromID(t *testing.T) {
	id := "64a4e2e994cb7dd7f0d0f9f7"
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		t.Fatalf("error converting id to ObjectID: %v", err)
	}
	dosenwali, err := module.GetDosenWaliFromID(objectID, module.MongoConn, "dosenwali")
	if err != nil {
		t.Fatalf("error calling GetDosenWaliFromID: %v", err)
	}
	fmt.Println(dosenwali)
}

func TestGetTemaFromID(t *testing.T) {
	id := "64a4e2ea94cb7dd7f0d0f9f8"
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		t.Fatalf("error converting id to ObjectID: %v", err)
	}
	tema, err := module.GetTemaFromID(objectID, module.MongoConn, "tema")
	if err != nil {
		t.Fatalf("error calling GetTemaFromID: %v", err)
	}
	fmt.Println(tema)
}

func TestGetMonitoringFromID(t *testing.T) {
	id := "64a4e2ea94cb7dd7f0d0f9f9"
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		t.Fatalf("error converting id to ObjectID: %v", err)
	}
	monitoring, err := module.GetMonitoringFromID(objectID, module.MongoConn, "monitoring")
	if err != nil {
		t.Fatalf("error calling GetMonitoringFromID: %v", err)
	}
	fmt.Println(monitoring)
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

//Delete Monitoring

func TestDeleteMonitoringByID(t *testing.T) {
	id := "64a635e4b5a0cc779e9a01ae" // ID data yang ingin dihapus
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		t.Fatalf("error converting id to ObjectID: %v", err)
	}

	err = module.DeleteMonitoringByID(objectID, module.MongoConn, "monitoring")
	if err != nil {
		t.Fatalf("error calling DeleteMonitoringByID: %v", err)
	}

	// Verifikasi bahwa data telah dihapus dengan melakukan pengecekan menggunakan GetMonitoringFromID
	_, err = module.GetMonitoringFromID(objectID, module.MongoConn, "monitoring")
	if err == nil {
		t.Fatalf("expected data to be deleted, but it still exists")
	}
}