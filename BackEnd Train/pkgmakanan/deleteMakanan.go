package pkgmakanan

import (
	"fmt"
	"gormstudy/db"
)

var idDltMakanan string

func KurangMakanan() {
	fmt.Println("ID Makanan: ")
	fmt.Scanln(&idDltMakanan)

	result := db.DB.Delete(&db.Makanan{}, idDltMakanan)
	err := result.Error
	if err != nil {
		panic(err.Error())
	}
}
