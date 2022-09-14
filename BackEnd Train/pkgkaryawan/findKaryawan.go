package pkgkaryawan

import "gormstudy/db"

func TemukanKaryawan() []db.Karyawan {

	var karyawans []db.Karyawan

	result := db.DB.Find(&karyawans)
	err := result.Error
	if err != nil {
		panic(err.Error())
	}
	return karyawans
}
