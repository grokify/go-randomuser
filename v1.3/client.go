package randomuser // import "github.com/grokify/go-randomuser/v1.3"

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

const (
	ApiURL = "https://randomuser.me/api/"
)

type Request struct {
	Count            uint16
	Country          string
	Gender           string
	PasswordSettings PasswordSettings
	Seed             string
}

type PasswordSettings struct {
	Special   bool
	Upper     bool
	Lower     bool
	Number    bool
	MinLength uint8
	MaxLength uint8
}

func GetUsers(qry *Request) (Results, *http.Response, error) {
	apiURL := ApiURL
	if qry != nil {
		query := url.Values{}
		if qry.Count > 0 {
			query.Add("results", strconv.Itoa(int(qry.Count)))
		}
		qry.Country = strings.ToUpper(strings.TrimSpace(qry.Country))
		if len(qry.Country) > 0 {
			query.Add("nat", qry.Country)
		}
		qry.Gender = strings.ToLower(strings.TrimSpace(qry.Gender))
		if qry.Gender == "male" || qry.Gender == "female" {
			query.Add("gender", qry.Gender)
		}

		if len(qry.Seed) > 0 {
			query.Add("seed", qry.Seed)
		}

		passwordParts := []string{}
		if qry.PasswordSettings.Special {
			passwordParts = append(passwordParts, "special")
		}
		if qry.PasswordSettings.Upper {
			passwordParts = append(passwordParts, "upper")
		}
		if qry.PasswordSettings.Lower {
			passwordParts = append(passwordParts, "lower")
		}
		if qry.PasswordSettings.Number {
			passwordParts = append(passwordParts, "number")
		}
		min := uint8(0)
		max := uint8(0)

		if qry.PasswordSettings.MinLength > 0 {
			min = qry.PasswordSettings.MinLength
		}
		if qry.PasswordSettings.MaxLength > 0 {
			min = qry.PasswordSettings.MaxLength
		}
		if min == 0 && max != 0 {
			passwordParts = append(passwordParts, strconv.Itoa(int(max)))
		} else if max == 0 && min != 0 {
			passwordParts = append(passwordParts, strconv.Itoa(int(min)))
		} else if min != 0 && max != 0 {
			if min == max {
				passwordParts = append(passwordParts, strconv.Itoa(int(min)))
			} else if min > max {
				passwordParts = append(passwordParts, fmt.Sprintf("%d-%d", max, min))
			} else {
				passwordParts = append(passwordParts, fmt.Sprintf("%d-%d", min, max))
			}
		}
		if len(passwordParts) > 0 {
			query.Add("password", strings.Join(passwordParts, ","))
		}
		qryStr := query.Encode()
		if len(qryStr) > 0 {
			apiURL += "?" + qryStr
		}
	}

	res := Results{}
	resp, err := http.Get(apiURL)
	if err != nil {
		return res, resp, err
	} else if resp.StatusCode >= 300 {
		return res, resp, fmt.Errorf("RandomUser API Response Status Code [%d]", resp.StatusCode)
	}

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return res, resp, err
	}

	// fmt.Println(string(bytes))

	err = json.Unmarshal(bytes, &res)
	return res, resp, err
}
