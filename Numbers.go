package main

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"math/rand"
	"net"
	"net/url"
	"strconv"
	"time"
)

func main() {
	// Intn [0,n)
	fmt.Println(rand.Intn(100))
	fmt.Println(rand.Intn(100))
	// Float64 [0.0, 1.0)
	fmt.Println(rand.Float64())
	// special range [5.0, 10)
	fmt.Println(rand.Float64()*5 + 5)

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	// same as rand.Intn
	fmt.Println(r.Intn(100))
	fmt.Println(r.Intn(100))

	// same source generate same rand
	// to be secret, use crypto/rand
	s1 := rand.New(rand.NewSource(43))
	s2 := rand.New(rand.NewSource(43))
	fmt.Println(s1.Intn(100), s1.Intn(100))
	fmt.Println(s2.Intn(100), s2.Intn(100))

	f, _ := strconv.ParseFloat("3.14159", 64)
	fmt.Println(f)

	i, _ := strconv.ParseInt("123", 0, 64)
	fmt.Println(i)

	k, _ := strconv.Atoi("124")
	fmt.Println(k)

	_, e := strconv.Atoi("abc")
	fmt.Println(e)

	u := "postgres://user:pass@host.com:5432/path?k=v#f"
	uu, e := url.Parse(u)
	if e != nil {
		panic(e)
	}
	fmt.Println(uu.Scheme)
	fmt.Println(uu.User)
	fmt.Println(uu.User.Username())
	fmt.Println(uu.User.Password())
	fmt.Println(uu.Host)
	fmt.Println(net.SplitHostPort(uu.Host))
	fmt.Println(uu.Path)
	fmt.Println(uu.RawQuery)
	fmt.Println(url.ParseQuery(uu.RawQuery))
	fmt.Println(uu.Fragment)

	strs := "this is sha1 testing"
	fmt.Println(strs)
	h := sha1.New()
	h.Write([]byte(strs))
	// usually nil
	bs := h.Sum(nil)
	fmt.Printf("%x\n", bs)

	data := "data to be base64"
	toString := base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(toString)

	bytes, e := base64.StdEncoding.DecodeString(toString)
	fmt.Println(string(bytes))

	// URL compatible base64 encoding
	urlStr := base64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(urlStr)
}
