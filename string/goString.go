package main

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"
)

func main() {
	p := fmt.Println

	s := "golang"

	p(strings.Contains(s, "go"))
	p(strings.ContainsAny(s, "o"))
	p(strings.Count(s, "g"))
	p(strings.HasPrefix(s, "go"))
	p(strings.HasSuffix(s, "lang"))
	p(strings.Index(s, "go"))
	p(strings.Join([]string{"a", "b"}, "-"))
	p(strings.Repeat(s, 2))
	p(strings.Replace(s, "go", "java", 1))
	p(strings.Replace(s, "g", "G", -1))
	p(strings.Split("good,bad", ","))
	p(strings.ToLower(s))
	p(strings.ToUpper(s))
	// note the multi-byte characters
	p(len(s))
	p(s[1])

	// lots of string formats

	// regexp
	matched, _ := regexp.MatchString("[a-z]", "G")
	p(matched)

	r, _ := regexp.Compile("\\d{4}-(\\d{2})-\\d{2}")
	p(r.MatchString("2010-01-01"))
	println(r.FindString("2010-01-01 12:34:56"))
	p(r.FindStringSubmatch("2010-01-01"))
	p(r.FindAllString("2010-01-01 to 2010-01-02", 1))
	p(r.FindAllString("2010-01-01 to 2010-01-02", 2))

	fmt.Println(r.Match([]byte("2010-01-01")))

	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println(r)

	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))
}
