package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Mahasiswa struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama         string             `bson:"nama,omitempty" json:"nama,omitempty"`
	NPM          int                `bson:"npm,omitempty" json:"npm,omitempty"`
	Jekel        string             `bson:"jekel,omitempty" json:"jekel,omitempty"`
	Phone_number string             `bson:"phone_number,omitempty" json:"phone_number,omitempty"`
}

type OrangTua struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama_OT      string             `bson:"nama_ot,omitempty" json:"nama_ot,omitempty"`
	Phone_number string             `bson:"phone_number,omitempty" json:"phone_number,omitempty"`
	Anak         Mahasiswa          `bson:"anak,omitempty" json:"anak,omitempty"`
}

type DosenWali struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama_Dosen   string             `bson:"nama_dosen,omitempty" json:"nama_dosen,omitempty"`
	Alamat       string             `bson:"alamat,omitempty" json:"alamat,omitempty"`
	Phone_number string             `bson:"phone_number,omitempty" json:"phone_number,omitempty"`
	Email        string             `bson:"email,omitempty" json:"email,omitempty"`
}

type Tema struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama_Tema string             `bson:"nama_tema,omitempty" json:"nama_tema,omitempty"`
}

type Monitoring struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	OrangTua OrangTua           `bson:"ortu,omitempty" json:"ortu,omitempty"`
	Tema     Tema               `bson:"tema,omitempty" json:"tema,omitempty"`
	Dosen    DosenWali          `bson:"dosen,omitempty" json:"dosen,omitempty"`
	Tanggal  string             `bson:"tanggal,omitempty" json:"tanggal,omitempty"`
	Hari     string             `bson:"hari,omitempty" json:"hari,omitempty"`
}

type User struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Fullname        string             `bson:"fullname,omitempty" json:"fullname,omitempty"`
	Email           string             `bson:"email,omitempty" json:"email,omitempty"`
	Password        string             `bson:"password,omitempty" json:"password,omitempty"`
	Confirmpassword string             `bson:"confirmpass,omitempty" json:"confirmpass,omitempty"`
}