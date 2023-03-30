package module

import (
	"context"
	"fmt"
	"os"

	"github.com/FarelND29/monitoring_orang_tua/model"
	"github.com/aiteung/atdb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var MongoString string = os.Getenv("MONGOSTRING")

var MongoInfo = atdb.DBInfo{
	DBString: MongoString,
	DBName:   "farel_db",
}

var MongoConn = atdb.MongoConnect(MongoInfo)

func InsertOneDoc(db *mongo.Database, collection string, doc interface{}) (insertedID interface{}) {
	insertResult, err := db.Collection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		fmt.Printf("InsertOneDoc: %v\n", err)
	}
	return insertResult.InsertedID
}



func InsertOrangTua(db *mongo.Database, col string ,nama_ot string,phone_number string, anak model.Mahasiswa) (InsertedID interface{}) {
	var orangtua model.OrangTua
	orangtua.Nama_OT = nama_ot
	orangtua.Phone_number = phone_number
	orangtua.Anak = anak
	return InsertOneDoc(db, col, orangtua)
}

func GetOrangTuaFromNamaMahasiswa(db *mongo.Database, col string, nama string) (nomor model.OrangTua) {
	orang_tua := db.Collection(col)
	filter := bson.M{"anak.nama": nama}
	err := orang_tua.FindOne(context.TODO(), filter).Decode(&nomor)
	if err != nil {
		fmt.Printf("getTemaFromTanggal: %v\n", err)
	}
	return nomor
}