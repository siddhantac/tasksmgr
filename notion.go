package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/sidc9/gotion"
)

func notion() {
	b, err := ioutil.ReadFile(".env")
	if err != nil {
		log.Fatal(err)
	}

	apiKey := strings.TrimSuffix(string(b), "\n")

	gotion.Init(apiKey, gotion.DefaultURL)
	c := gotion.GetClient()

	// pg, err := c.GetPage("a0e3feca-85c9-440f-91cc-8c367d6aa9f4")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(pg.Title())

	db, err := c.GetDatabase("539a391b9f83427f933518f5dc2b6c83")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(db.Title[0].PlainText)
	pages, err := db.Query(nil)
	if err != nil {
		log.Fatal(err)
	}

	for _, p := range pages.Results {
		var done bool
		d, ok := p.Properties["Done"]
		if ok {
			done = d.Checkbox
		}
		fmt.Printf(" - [%t] %s %s\n", done, p.Title(), p.ID)
	}
}
