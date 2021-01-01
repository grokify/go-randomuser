package main

import (
	"fmt"
	"log"

	randomuser "github.com/grokify/go-randomuser/v1.3"
	"github.com/grokify/oauth2more/hubspot"
	"github.com/grokify/oauth2more/scim"
	"github.com/grokify/simplego/fmt/fmtutil"
)

func main() {
	users, _, err := randomuser.GetUsers(&randomuser.Request{
		Count:   1,
		Seed:    "abc",
		Country: "us"})
	if err != nil {
		log.Fatal(err)
	}
	// fmtutil.PrintJSON(users)
	scims := []scim.User{}
	for _, usr := range users.Results {
		sc := usr.Scim()
		fmtutil.PrintJSON(sc)
		scims = append(scims, sc)
	}
	outfile := "_contacts_x.xlsx"
	err = hubspot.WriteContactsXLSX(outfile, scims)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("WROTE [%s]\n", outfile)
	fmt.Println("DONE")
}
