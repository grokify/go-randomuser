package randomuser // import "github.com/grokify/go-randomuser/v1.3"

type Results struct {
	Results []User `json:"results,omitempty"`
	Info    Info   `json:"info,omitempty"`
	Error   string `json:"error,omitempty`
}

type Info struct {
	Seed    string `json:"seed"`
	Results int    `json:"results"`
	Page    int    `json:"page"`
	Version string `json:"version"`
}
