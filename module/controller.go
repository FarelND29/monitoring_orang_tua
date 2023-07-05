package module

import (
	"context"
	"fmt"
	"os"

	"github.com/FarelND29/monitoring_orang_tua/model"
	"github.com/aiteung/atdb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var MongoString string = os.Getenv("MONGOSTRING")

var MongoInfo = atdb.DBInfo{
	DBString: MongoString,
	DBName:   "monitoring_db",
}

var MongoConn = atdb.MongoConnect(MongoInfo)

func InsertOneDoc(db *mongo.Database, collection string, doc interface{}) (insertedID interface{}) {
	insertResult, err := db.Collection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		fmt.Printf("InsertOneDoc: %v\n", err)
	}
	return insertResult.InsertedID
}

func InsertMahasiswa(db *mongo.Database, col string, nama string, npm int, jenis_kelamin string, phone_number string) (InsertedID interface{}) {
	var mahasiswa model.Mahasiswa
	mahasiswa.Nama = nama
	mahasiswa.NPM = npm
	mahasiswa.Jekel = jenis_kelamin
	mahasiswa.Phone_number = phone_number
	return InsertOneDoc(db, col, mahasiswa)
}

func InsertOrangTua(db *mongo.Database, col string, nama_ot string, phone_number string, anak model.Mahasiswa) (InsertedID interface{}) {
	var orangtua model.OrangTua
	orangtua.Nama_OT = nama_ot
	orangtua.Phone_number = phone_number
	orangtua.Anak = anak
	return InsertOneDoc(db, col, orangtua)
}

func InsertDosenWali(db *mongo.Database, col string, nama_dosen string, alamat string, phone_number string, email string) (InsertedID interface{}) {
	var dosen model.DosenWali
	dosen.Nama_Dosen = nama_dosen
	dosen.Alamat = alamat
	dosen.Phone_number = phone_number
	dosen.Email = email
	return InsertOneDoc(db, col, dosen)
}

func InsertTema(db *mongo.Database, col string, nama_tema string) (InsertedID interface{}) {
	var tema model.Tema
	tema.Nama_Tema = nama_tema
	return InsertOneDoc(db, col, tema)
}

/* func InsertMonitoring(db *mongo.Database, col string, orang_tua model.OrangTua, tema model.Tema, dosen model.DosenWali, tanggal string, hari string) (insertedID primitive.ObjectID, err error) {
	presensi := bson.M{
	monitoring.OrangTua = orang_tua
	monitoring.Tema = tema
	monitoring.Dosen = dosen
	monitoring.Tanggal = tanggal
	monitoring.Hari = hari

	result, err := db.Collection(col).InsertOne(context.Background(), monitoring)
	if err != nil {
		fmt.Printf("InsertMonitoring: %v\n", err)
		return
	}
	insertedID = result.InsertedID.(primitive.ObjectID)
	return insertedID, nil
}
 */
 func InsertMonitoring(db *mongo.Database, col string, orang_tua model.OrangTua, tema model.Tema, dosen model.DosenWali, tanggal string, hari string) (insertedID primitive.ObjectID, err error) {
	monitoring := bson.M{
		"orangtua":    	orang_tua,
		"tema":     	tema,
		"dosen":     	dosen,
		"tanggal": 		tanggal,
		"hari":     	hari,
	}
	result, err := db.Collection(col).InsertOne(context.Background(), monitoring)
	if err != nil {
		fmt.Printf("InsertMonitoring: %v\n", err)
		return
	}
	insertedID = result.InsertedID.(primitive.ObjectID)
	return insertedID, nil
}


func GetMahasiswaFromNpm(db *mongo.Database, col string, npm int) (mhs model.Mahasiswa) {
	mahasiswa := db.Collection(col)
	filter := bson.M{"npm": npm}
	err := mahasiswa.FindOne(context.TODO(), filter).Decode(&mhs)
	if err != nil {
		fmt.Printf("GetMahasiswaFromNpm: %v\n", err)
	}
	return mhs
}

func GetOrangTuaFromNamaMahasiswa(db *mongo.Database, col string, nama string) (ortu model.OrangTua) {
	orang_tua := db.Collection(col)
	filter := bson.M{"anak.nama": nama}
	err := orang_tua.FindOne(context.TODO(), filter).Decode(&ortu)
	if err != nil {
		fmt.Printf("GetOrangTuaFromNamaMahasiswa: %v\n", err)
	}
	return ortu
}

func GetDosenWaliFromNamaDosen(db *mongo.Database, col string, nama_dosen string) (doswal model.DosenWali) {
	dosen_wali := db.Collection(col)
	filter := bson.M{"nama_dosen": nama_dosen}
	err := dosen_wali.FindOne(context.TODO(), filter).Decode(&doswal)
	if err != nil {
		fmt.Printf("GetDosenWaliFromNamaDosen: %v\n", err)
	}
	return doswal
}

func GetTemaFromNamaTema(db *mongo.Database, col string, nama_tema string) (tm model.Tema) {
	tema := db.Collection(col)
	filter := bson.M{"nama_tema": nama_tema}
	err := tema.FindOne(context.TODO(), filter).Decode(&tm)
	if err != nil {
		fmt.Printf("GetTemaFromNamaTema: %v\n", err)
	}
	return tm
}

func GetMonitoringFromNamaMahasiswa(db *mongo.Database, col string, nama string) (mntr model.Monitoring) {
	monitoring := db.Collection(col)
	filter := bson.M{"ortu.anak.nama": nama}
	err := monitoring.FindOne(context.TODO(), filter).Decode(&mntr)
	if err != nil {
		fmt.Printf("GetMonitoringFromNamaMahasiswa: %v\n", err)
	}
	return mntr
}

func GetAllMonitoring(db *mongo.Database, col string) (data []model.Monitoring) {
	monitoring := db.Collection(col)
	filter := bson.M{}
	cursor, err := monitoring.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetAllMonitoring :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return data
}

func GetAllMahasiswa(db *mongo.Database, col string) (data []model.Mahasiswa) {
	mahasiswa := db.Collection(col)
	filter := bson.M{}
	cursor, err := mahasiswa.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetAllMahasiswa :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return data
}

func GetAllOrangTua(db *mongo.Database, col string) (data []model.OrangTua) {
	orangtua := db.Collection(col)
	filter := bson.M{}
	cursor, err := orangtua.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetAllOrangTua :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return data
}

func GetAllDosenWali(db *mongo.Database, col string) (data []model.DosenWali) {
	dosenwali := db.Collection(col)
	filter := bson.M{}
	cursor, err := dosenwali.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetAllDosenWali :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return data
}

func GetAllTema(db *mongo.Database, col string) (data []model.Tema) {
	tema := db.Collection(col)
	filter := bson.M{}
	cursor, err := tema.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetAllTema :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return data
}
