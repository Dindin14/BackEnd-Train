package pkgkaryawan

import "fmt"

var idKaryawan string

func KurangKaryawan() string {
	fmt.Println("ID Karyawan: ")
	fmt.Scanln(&idKaryawan)

	return idKaryawan
}
