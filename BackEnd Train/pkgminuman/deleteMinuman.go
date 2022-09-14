package pkgminuman

import (
	"fmt"
	"gormstudy/db"
)

var idMinuman string

func KurangMinuman() {
	fmt.Println("ID Minuman: ")
	fmt.Scanln(&idMinuman)

	result := db.DB.Delete(&db.Minuman{}, idMinuman)
	err := result.Error
	if err != nil {
		panic(err.Error())
	}
}
