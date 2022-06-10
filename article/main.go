package article

import "encoding/json"

type Articles []Article

type Article struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (a Article) Valid() bool {
	if len(a.Name) < 1 || len(a.Name) > 4 {
		return false
	}

	if len(a.Description) < 1 || len(a.Description) > 4 {
		return false
	}

	return true
}

func (a Article) String() string {
	bytes, err := json.Marshal(a)
	if err != nil {
		return ""
	}
	return string(bytes)
}
