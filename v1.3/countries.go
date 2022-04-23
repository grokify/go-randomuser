package randomuser // import "github.com/grokify/go-randomuser/v1.3"

import (
	"strings"

	"github.com/grokify/mogo/crypto/randutil"
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

func RandomCountry() (string, error) {
	countries := Countries()
	rnd, err := randutil.CryptoRandInt64(nil, int64(len(countries)))
	if err != nil {
		return "", err
	}
	return countries[rnd], nil
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
