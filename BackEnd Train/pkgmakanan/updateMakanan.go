package pkgmakanan

import (
	"fmt"
	"gormstudy/db"
)

func UpdateMenuMakanan() {

	var update, hargaUpdate int
	var namaUpdate, idUptMakanan string

	fmt.Println("ID Makanan: ")
	fmt.Scanln(&idUptMakanan)
	fmt.Println("1. Ganti Nama Makanan")
	fmt.Println("2. Ganti Harga Makanan")
	fmt.Scanln(&update)

	switch update {
	case 1:
		fmt.Println("Nama Makanan: ")
		fmt.Scanln(&namaUpdate)
		db.DB.Model(&db.Makanan{}).Where("id = ?", idUptMakanan).Update("Nama", namaUpdate)
	case 2:
		fmt.Println("Harga Makanan: ")
		fmt.Scanln(&hargaUpdate)
		db.DB.Model(&db.Makanan{}).Where("id = ?", idUptMakanan).Update("Harga", hargaUpdate)
	}
}
