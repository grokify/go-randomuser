package randomuser // import "github.com/grokify/go-randomuser/v1.3"

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/grokify/goauth/scim"
	"github.com/grokify/mogo/encoding/jsonutil"
	"github.com/grokify/mogo/type/stringsutil"
)

type User struct {
	Gender      string      `json:"gender,omitempty"`
	Name        Name        `json:"name,omitempty"`
	Location    Location    `json:"location,omitempty"`
	Email       string      `json:"email,omitempty"`
	Login       Login       `json:"login,omitempty"`
	DateOfBirth DateOfBirth `json:"dob,omitempty"`
	Registered  Registered  `json:"registered,omitempty"`
	Phone       string      `json:"phone,omitempty"`
	Cell        string      `json:"cell,omitempty"`
	ID          ID          `json:"id,omitempty"`
	Picture     Picture     `json:"picture,omitempty"`
	Nationality string      `json:"nat,omitempty"`
}

type Name struct {
	Title string `json:"title"`
	First string `json:"first"`
	Last  string `json:"last"`
}

var rxAZ = regexp.MustCompile(`^[A-Za-z]+$`)

func (name *Name) IsAZSimple() bool {
	name.First = strings.TrimSpace(name.First)
	name.Last = strings.TrimSpace(name.Last)
	if rxAZ.MatchString(name.First) && rxAZ.MatchString(name.Last) {
		return true
	}
	return false
}

func (user *User) PhoneFormatted() string {
	raw := strings.TrimSpace(user.Phone)
	if user.Nationality == "US" {
		try := stringsutil.DigitsOnly(raw)
		if len(try) == 10 {
			return "+1" + try
		}
	}
	return raw
}

func (user *User) CellFormatted() string {
	raw := strings.TrimSpace(user.Cell)
	if user.Nationality == "US" {
		try := stringsutil.DigitsOnly(raw)
		if len(try) == 10 {
			return "+1" + try
		}
	}
	return raw
}

func (user *User) Scim() scim.User {
	return UserToScim(*user)
}

type Location struct {
	Street      Street          `json:"street,omitempty"`
	City        string          `json:"city,omitempty"`
	State       string          `json:"state,omitempty"`
	Postcode    jsonutil.String `json:"postcode,omitempty"`
	Coordinates Coordinates     `json:"coordinates,omitempty"`
	Timezone    Timezone        `json:"timezone,omitempty"`
}

type Street struct {
	Number int    `json:"number,omitempty"`
	Name   string `json:"name,omitempty"`
}

func (street *Street) String() string {
	return fmt.Sprintf("%v %s", street.Number, street.Name)
}

type Coordinates struct {
	Latitude  string `json:"latitude,omitempty"`
	Longitude string `json:"longitude,omitempty"`
}

type Timezone struct {
	Offset      string `json:"offset,omitempty"`
	Description string `json:"description,omitempty"`
}

type Login struct {
	UUID     string `json:"uuid,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Salt     string `json:"salt,omitempty"`
	MD5      string `json:"md5,omitempty"`
	SHA1     string `json:"sha1,omitempty"`
	SHA256   string `json:"sha256,omitempty"`
}

type DateOfBirth struct {
	Date time.Time `json:"date,omitempty"`
	Age  int       `json:"age,omitempty"`
}

type Registered struct {
	Date time.Time `json:"date,omitempty"`
	Age  int       `json:"age,omitempty"`
}

type ID struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

type Picture struct {
	Large     string `json:"large,omitempty"`
	Medium    string `json:"medium,omitempty"`
	Thumbnail string `json:"thumbnail,omitempty"`
}

/*
{
        "gender": "male",
        "name": {
          "title": "mr",
          "first": "brad",
          "last": "gibson"
        },
        "location": {
          "street": "9278 new road",
          "city": "kilcoole",
          "state": "waterford",
          "postcode": "93027",
          "coordinates": {
            "latitude": "20.9267",
            "longitude": "-7.9310"
          },
          "timezone": {
            "offset": "-3:30",
            "description": "Newfoundland"
          }
        },
        "email": "brad.gibson@example.com",
        "login": {
          "uuid": "155e77ee-ba6d-486f-95ce-0e0c0fb4b919",
          "username": "silverswan131",
          "password": "firewall",
          "salt": "TQA1Gz7x",
          "md5": "dc523cb313b63dfe5be2140b0c05b3bc",
          "sha1": "7a4aa07d1bedcc6bcf4b7f8856643492c191540d",
          "sha256": "74364e96174afa7d17ee52dd2c9c7a4651fe1254f471a78bda0190135dcd3480"
        },
        "dob": {
          "date": "1993-07-20T09:44:18.674Z",
          "age": 26
        },
        "registered": {
          "date": "2002-05-21T10:59:49.966Z",
          "age": 17
        },
        "phone": "011-962-7516",
        "cell": "081-454-0666",
        "id": {
          "name": "PPS",
          "value": "0390511T"
        },
        "picture": {
          "large": "https://randomuser.me/api/portraits/men/75.jpg",
          "medium": "https://randomuser.me/api/portraits/med/men/75.jpg",
          "thumbnail": "https://randomuser.me/api/portraits/thumb/men/75.jpg"
        },
        "nat": "IE"
	  }

*/
