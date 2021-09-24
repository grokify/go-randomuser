package randomuser // import "github.com/grokify/go-randomuser/v1.3"

import (
	"math/rand"
	"strings"
	"time"
)

const (
	CountryAustralia     = "AU"
	CountryBrazil        = "BR"
	CountryCanada        = "CA"
	CountryDenmark       = "DK"
	CountryFinland       = "FI"
	CountryFrance        = "FR"
	CountryGermany       = "DE"
	CountryIran          = "IR"
	CountryIreland       = "IE"
	CountryNetherlands   = "NL"
	CountryNewZealand    = "NZ"
	CountryNorway        = "NO"
	CountrySpain         = "ES"
	CountrySwitzerland   = "CH"
	CountryTurkey        = "TR"
	CountryUnitedKingdom = "GB"
	CountryUnitedStates  = "US"
)

func RandomCountry() string {
	countries := Countries()
	rand.Seed(time.Now().Unix())
	return countries[rand.Intn(len(countries))]
}

func Countries() []string {
	return []string{
		CountryAustralia,
		CountryBrazil,
		CountryCanada,
		CountryDenmark,
		CountryFinland,
		CountryFrance,
		CountryGermany,
		CountryIran,
		CountryIreland,
		CountryNetherlands,
		CountryNewZealand,
		CountryNorway,
		CountrySpain,
		CountrySwitzerland,
		CountryTurkey,
		CountryUnitedKingdom,
		CountryUnitedStates}
}

func IsCountry(c string) bool {
	c = strings.ToUpper(c)
	for _, try := range Countries() {
		if try == c {
			return true
		}
	}
	return false
}
