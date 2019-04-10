package main

import (
	"fmt"
	"os"
)

func main() {
	// will terminate the main
	//panic("dummy problem")

	//_, err := os.Create("/tmp/fileByGo")
	//if err != nil {
	//	panic(err)
	//}

	// defer for cleanup
	f := createFile("/tmp/go-file.txt")
	defer closeFile(f)
	writeFile(f)
	panic("test panic before defer")
}

func createFile(path string) *os.File {
	fmt.Println("creating...")
	f, e := os.Create(path)
	if e != nil {
		panic(e)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing...")
	_, _ = fmt.Fprintln(f, "this is the data")
}

func closeFile(f *os.File) {
	fmt.Println("closing...")
	f.Close()
}
