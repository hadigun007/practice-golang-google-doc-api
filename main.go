package main

import (
	"fmt"
	core "quickstart/core/docs"
)

func main() {

	data, err := core.CreateBlackDocument("test-ketiga")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Berhasil tambah dokumen baru: " + data.Title)

	// data, err := core.GetFiles()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println()
	// for _, d := range data.Files {
	// 	fmt.Println(d.Id)
	// 	fmt.Println(d.Name)

	// }
}
