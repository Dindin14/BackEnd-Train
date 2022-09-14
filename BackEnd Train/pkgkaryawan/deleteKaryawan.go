package pkgkaryawan

import (
	"fmt"
	"gormstudy/db"
)

var idKaryawan string

func KurangKaryawan() {
	fmt.Println("ID Karyawan: ")
	fmt.Scanln(&idKaryawan)

	result := db.DB.Delete(&db.Karyawan{}, idKaryawan)
	err := result.Error
	if err != nil {
		panic(err.Error())
	}
}
