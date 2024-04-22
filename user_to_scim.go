package randomuser // import "github.com/grokify/go-randomuser/v1.3"

import (
	"github.com/grokify/goauth/scim"
)

func UserToScim(rUser User) scim.User {
	scimUser := scim.User{
		Schemas:  []string{"urn:ietf:params:scim:schemas:core:2.0:User"},
		ID:       rUser.Login.UUID,
		UserName: rUser.Login.Username,
		Name: scim.Name{
			HonorificPrefix: rUser.Name.Title,
			GivenName:       rUser.Name.First,
			FamilyName:      rUser.Name.Last},
		Emails: []scim.Item{{
			Value:   rUser.Email,
			Type:    "work",
			Primary: true}},
		PhoneNumbers: []scim.Item{{
			Value:   rUser.PhoneFormatted(),
			Type:    "work",
			Primary: true}, {
			Value:   rUser.CellFormatted(),
			Type:    "mobile",
			Primary: false}},
		Addresses: []scim.Address{{
			Type:          "work",
			StreetAddress: rUser.Location.Street.String(),
			Locality:      rUser.Location.City,
			Region:        rUser.Location.State,
			PostalCode:    string(rUser.Location.Postcode),
			Country:       rUser.Nationality,
			Primary:       true}},
	}
	scimUser.Addresses[0].InflateStreetAddress(true)
	return scimUser
}
