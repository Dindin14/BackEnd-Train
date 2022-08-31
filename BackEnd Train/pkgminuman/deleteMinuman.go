package pkgminuman

import "fmt"

var idMinuman string

func KurangMinuman() string {
	fmt.Println("ID Minuman: ")
	fmt.Scanln(&idMinuman)

	return idMinuman
}
