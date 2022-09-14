package pkgminuman

import (
	"gormstudy/db"
)

// var namaMinuman string
// var hargaMinuman int
// var toping string

func TambahMinuman(namaMinuman string, hargaMinuman int, toping string) {
	// fmt.Println("Nama minuman: ")
	// fmt.Scanln(&namaMinuman)
	// fmt.Println("Toping: ")
	// fmt.Scanln(&toping)
	// fmt.Println("Harga: ")
	// fmt.Scanln(&hargaMinuman)

	result := db.DB.Create(&db.Minuman{Nama: namaMinuman, Toping: toping, Harga: hargaMinuman})
	err := result.Error
	if err != nil {
		panic(err.Error())
	}
}
