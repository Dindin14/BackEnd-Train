package pkgkaryawan

import (
	"gormstudy/db"
)

// var namaKaryawan string
// var umurKaryawan int
// var asalKaryawan string

func TambahKaryawan(namaKaryawan string, umurKaryawan int, asalKaryawan string) {
	// fmt.Println("Nama Karyawan: ")
	// fmt.Scanln(&namaKaryawan)
	// fmt.Println("Umur: ")
	// fmt.Scanln(&umurKaryawan)
	// fmt.Println("Asal: ")
	// fmt.Scanln(&asalKaryawan)

	result := db.DB.Create(&db.Karyawan{Nama: namaKaryawan, Umur: umurKaryawan, Asal: asalKaryawan})
	err := result.Error
	if err != nil {
		panic(err.Error())
	}
}
