package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"strings"
)

func emailsFrom(in string) []string {
	r := csv.NewReader(strings.NewReader(in))
	r.Comma = ';'
	r.Comment = '#'

	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	header := records[0]
	email := -1
	for i := 0; i < len(header); i++ {
		if header[i] == "email" {
			email = i
			break
		}
	}
	size := len(records)
	raw := records[1:size]
	var emails = make([]string, size+10)
	for j, col := range raw {
		emails[j] = col[email]
	}
	return emails

}

func intersect(sect []string, orig []string, tgt []string) []string {
	for _, i := range orig {
		for _, x := range tgt {
			if i == x {
				if !exists(x, sect) {
					sect = append(sect, x)
				}
				return intersect(sect, orig[1:], tgt[1:])
			}
		}
	}
	return sect
}

func exists(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func testIntersect() {
	in := `first_name;last_name;username;email
"John";"Pike";john;john@example.com
Jônatas;Paganini;jonatasdp;jonatasdp@gmail.com
Kin;Thompson;kin;kin@example.com
Jônatas;Paganini;jonatasdp;jonatasdp@gmail.com
"John";"Pike";john;john@example.com
Jônatas;Paganini;jonatasdp;jonatasdp@gmail.com
Kin;Thompson;kin;kin@example.com
"John";"Pike";john;john@example.com
Kin;Thompson;kin;kin@example.com
"Johnert";"Griesemer";"gri";gri@example.com
Jônatas;Paganini;jonatasdp;jonatasdp@gmail.com
`
	compare := `username;email
gri;gri@example.com
jonatasdp;jonatasdp@gmail.com
jonatasdp;jonatasdp@gmail.com
jonatasdp;jonatasdp@gmail.com
john;john@example.com
jonatasdp;jonatasdp@gmail.com
jonatasdp;jonatasdp@gmail.com
`
	a := emailsFrom(in)
	b := emailsFrom(compare)
	fmt.Println(intersect([]string{}, b, a))

}
