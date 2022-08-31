package pkgkaryawan

import "fmt"

var namaKaryawan string
var umurKaryawan int
var asalKaryawan string

func TambahKaryawan() (string, int, string) {
	fmt.Println("Nama Karyawan: ")
	fmt.Scanln(&namaKaryawan)
	fmt.Println("Umur: ")
	fmt.Scanln(&umurKaryawan)
	fmt.Println("Asal: ")
	fmt.Scanln(&asalKaryawan)

	return namaKaryawan, umurKaryawan, asalKaryawan
}
