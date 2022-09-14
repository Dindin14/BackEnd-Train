package pkgmakanan

import (
	"gormstudy/db"
)

// var namaMakanan string
// var hargaMakanan int

func TambahMakanan(namaMakanan string, hargaMakanan int) {

	// fmt.Println("Nama makanan: ")
	// fmt.Scanln(&namaMakanan)
	// fmt.Println("Harga: ")
	// fmt.Scanln(&hargaMakanan)
	result := db.DB.Create(&db.Makanan{Nama: namaMakanan, Harga: hargaMakanan})
	err := result.Error
	if err != nil {
		panic(err.Error())
	}
}
