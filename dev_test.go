package NPM

import (
	"fmt"
	"testing"

	"github.com/FarelND29/monitoring_orang_tua/model"
	"github.com/FarelND29/monitoring_orang_tua/module"
)


func TestInsertOrangTua(t *testing.T) {
	nama_ot := "Rini"
	phone_number := "0888321157026"
	anak := model.Mahasiswa{
		Nama : "Farel",
		NPM : 1214070,
		Phone_number : "083821157026",
		
	}
	hasil:=module.InsertOrangTua(module.MongoConn, "monitoring", nama_ot , phone_number , anak )
	fmt.Println(hasil)
}

func TestGetOrangTuaFromNamaMahasiswa(t *testing.T) {
	nama := "Farel"
	anak :=module.GetOrangTuaFromNamaMahasiswa(module.MongoConn, "monitoring", nama)
	fmt.Println(anak)
}
