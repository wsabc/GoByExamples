package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	name := "/tmp/go-file.txt"

	dat, e := ioutil.ReadFile(name)
	check(e)
	fmt.Println(string(dat))

	f, e := os.Open(name)
	defer f.Close()
	check(e)
	b1 := make([]byte, 5)
	n, e := f.Read(b1)
	check(e)
	fmt.Printf("%d bytes: %s\n", n, string(b1))

	o2, e := f.Seek(6, 0)
	check(e)
	b2 := make([]byte, 2)
	n2, e := f.Read(b2)
	check(e)
	fmt.Printf("%d bytes @ %d: %s\n", n2, o2, string(b2))

	o3, e := f.Seek(6, 0)
	check(e)
	b3 := make([]byte, 2)
	n3, e := io.ReadAtLeast(f, b3, 2)
	check(e)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	//rewind
	_, err := f.Seek(0, 0)
	check(err)

	r4 := bufio.NewReader(f)
	b4, e := r4.Peek(5)
	check(e)
	fmt.Printf("5 bytes: %s\n", string(b4))

	e = ioutil.WriteFile(name, []byte("good day"), 0644)
	check(e)

	ff, e := os.Create("/tmp/gofile")
	defer ff.Close()
	i, e := ff.WriteString("good good day")
	check(e)
	fmt.Printf("wrote %d bytes", i)
	ff.Sync()

	w := bufio.NewWriter(ff)
	// here append
	line, e := w.WriteString("\n again another line")
	fmt.Printf("wrote %d bytes again", line)
	w.Flush()

}
