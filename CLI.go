package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
)

func main() {
	// first is the program
	argsWithProg := os.Args
	argsWOProg := os.Args[1:]
	arg := "dummy"
	// arg := argsWOProg[2]

	fmt.Println(argsWithProg, argsWOProg, arg)

	command := exec.Command("date")
	bytes, _ := command.Output()
	fmt.Println(">date")
	fmt.Println(string(bytes))

	grepcmd := exec.Command("grep", "hello")
	grepIn, _ := grepcmd.StdinPipe()
	grepOut, _ := grepcmd.StdoutPipe()
	grepcmd.Start()
	grepIn.Write([]byte("hello this is grep test"))
	grepIn.Close()
	all, _ := ioutil.ReadAll(grepOut)
	grepcmd.Wait()
	fmt.Println(">grep hello")
	fmt.Println(string(all))

	// unix signals
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	// register signals
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		fmt.Println(sig)
		done <- true
	}()

	fmt.Println("awaiting signals")
	<-done
	fmt.Println("exiting")

	// exec others to replace the current go
	path, e := exec.LookPath("ls")
	if e != nil {
		panic(e)
	}
	// first arg should be program
	args := []string{"ls", "-l", "-a", "-h"}
	env := os.Environ()
	// go does not support fork
	e = syscall.Exec(path, args, env)
	if e != nil {
		panic(e)
	}

	// defer will not be run when using os.Exit
	defer fmt.Println("!exit")
	os.Exit(3)
}
