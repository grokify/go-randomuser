package main

import (
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"

	randomuser "github.com/grokify/go-randomuser/v1.3"
	"github.com/grokify/goauth/hubspot"
	"github.com/grokify/goauth/scim"
	"github.com/grokify/mogo/encoding/jsonutil"
	flags "github.com/jessevdk/go-flags"
)

type Options struct {
	Number    int      `short:"N" long:"number" description:"Number of users to create"`
	Countries []string `short:"C" long:"country" description:"Countries"`
	JSONFile  string   `short:"J" long:"jsonfile" description:"Create XLSX file"`
	XSLXFile  string   `short:"X" long:"xlsxfile" description:"Create XLSX file"`
	Seed      string   `short:"S" long:"seed" description:"Seed"`
}

func (opts *Options) OneCountry() (string, error) {
	if len(opts.Countries) == 0 {
		return randomuser.RandomCountry(), nil
	} else if len(opts.Countries) == 1 {
		if randomuser.IsCountry(opts.Countries[0]) {
			return strings.ToUpper(opts.Countries[0]), nil
		}
		return "", fmt.Errorf("not a valid country [%s]", opts.Countries[0])
	}

	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)

	c := opts.Countries[rng.Intn(len(opts.Countries))]
	if !randomuser.IsCountry(c) {
		return "", fmt.Errorf("not a valid country [%s]", c)
	}
	return strings.ToUpper(c), nil
}

func main() {
	opts := Options{}
	_, err := flags.Parse(&opts)
	if err != nil {
		log.Fatal(err)
	}

	c, err := opts.OneCountry()
	if err != nil {
		log.Fatal(err)
	}
	if opts.Number < 1 {
		opts.Number = 1
	}

	users, _, err := randomuser.GetUsers(&randomuser.Request{
		Count:   uint16(opts.Number),
		Seed:    opts.Seed,
		Country: c})
	if err != nil {
		log.Fatal(err)
	}

	scimSet := scim.UserSet{Users: []scim.User{}}

	for _, usr := range users.Results {
		sc := usr.Scim()
		sc.InflateDisplayName(false, false, true)
		scimSet.Users = append(scimSet.Users, sc)
	}

	if len(opts.JSONFile) > 0 {
		err := jsonutil.WriteFile(opts.JSONFile, scimSet, "", "  ", 0644)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("WROTE [%s]\n", opts.JSONFile)
	}
	if len(opts.XSLXFile) > 0 {
		err = hubspot.WriteContactsXLSX(opts.XSLXFile, scimSet.Users)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("WROTE [%s]\n", opts.XSLXFile)
	}
	fmt.Println("DONE")
}
