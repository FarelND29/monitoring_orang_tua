package module

import (
	"context"
	"fmt"
	"os"

	"errors"

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
		"ortu":    		orang_tua,
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

// UPDATE MONITORING ORANG TUA

func UpdateMonitoring(db *mongo.Database, col string,  id primitive.ObjectID, orang_tua model.OrangTua, tema model.Tema, dosen model.DosenWali, tanggal string, hari string) (err error) {
	filter := bson.M{"_id": id}
	update := bson.M{
		"$set": bson.M{
		"ortu":    		orang_tua,
		"tema":     	tema,
		"dosen":     	dosen,
		"tanggal": 		tanggal,
		"hari":     	hari,
		},
	}
	result, err := db.Collection(col).UpdateOne(context.Background(), filter, update)
	if err != nil {
		fmt.Printf("UpdateMonitoring: %v\n", err)
		return
	}
	if result.ModifiedCount == 0 {
		err = errors.New("No data has been changed with the specified ID")
		return
	}
	return nil
}

// DELETE MONITORING ORANG TUA

func DeleteMonitoringByID(_id primitive.ObjectID, db *mongo.Database, col string) error {
	orangtua := db.Collection(col)
	filter := bson.M{"_id": _id}

	result, err := orangtua.DeleteOne(context.TODO(), filter)
	if err != nil {
		return fmt.Errorf("error deleting data for ID %s: %s", _id, err.Error())
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("data with ID %s not found", _id)
	}

	return nil
}

func GetMahasiswaFromID(_id primitive.ObjectID, db *mongo.Database, col string) (mhs model.Monitoring, errs error) {
	siswa := db.Collection(col)
	filter := bson.M{"_id": _id}
	err := siswa.FindOne(context.TODO(), filter).Decode(&mhs)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return mhs, fmt.Errorf("no data found for ID %s", _id)
		}
		return mhs, fmt.Errorf("error retrieving data for ID %s: %s", _id, err.Error())
	}
	return mhs, nil
}

func GetOrangTuaFromID(_id primitive.ObjectID, db *mongo.Database, col string) (ot model.Monitoring, errs error) {
	ortu:= db.Collection(col)
	filter := bson.M{"_id": _id}
	err := ortu.FindOne(context.TODO(), filter).Decode(&ot)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return ot, fmt.Errorf("no data found for ID %s", _id)
		}
		return ot, fmt.Errorf("error retrieving data for ID %s: %s", _id, err.Error())
	}
	return ot, nil
}

func GetDosenWaliFromID(_id primitive.ObjectID, db *mongo.Database, col string) (dos model.Monitoring, errs error) {
	dosen := db.Collection(col)
	filter := bson.M{"_id": _id}
	err := dosen.FindOne(context.TODO(), filter).Decode(&dos)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return dos, fmt.Errorf("no data found for ID %s", _id)
		}
		return dos, fmt.Errorf("error retrieving data for ID %s: %s", _id, err.Error())
	}
	return dos, nil
}

func GetTemaFromID(_id primitive.ObjectID, db *mongo.Database, col string) (mkl model.Monitoring, errs error) {
	matkul := db.Collection(col)
	filter := bson.M{"_id": _id}
	err := matkul.FindOne(context.TODO(), filter).Decode(&mkl)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return mkl, fmt.Errorf("no data found for ID %s", _id)
		}
		return mkl, fmt.Errorf("error retrieving data for ID %s: %s", _id, err.Error())
	}
	return mkl, nil
}

//  FUNCTION GET MONITORING FROM ID 
func GetMonitoringFromID(_id primitive.ObjectID, db *mongo.Database, col string) (adm model.Monitoring, errs error) {
	admin := db.Collection(col)
	filter := bson.M{"_id": _id}
	err := admin.FindOne(context.TODO(), filter).Decode(&adm)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return adm, fmt.Errorf("no data found for ID %s", _id)
		}
		return adm, fmt.Errorf("error retrieving data for ID %s: %s", _id, err.Error())
	}
	return adm, nil
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
