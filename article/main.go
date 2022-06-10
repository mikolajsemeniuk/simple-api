package article

// struct { "items": []Article, "total": 10000 }
// struct { "items": []Article + something }
// []Article
// { "description": "one" }
// { "description": null }
//

// type JSON interface {
// 	JSON() (string, error)
// }

// type Bytes interface {
// 	Bytes() ([]byte, error)
// }
// func (a *Articles) Bytes(input []byte) error {
// 	return json.Unmarshal(input, &a)
// }

type Articles []Article

type Article struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// func (a *Article) Bytes(input []byte) error {
// 	return json.Unmarshal(input, &a)
// }

// func (a *Article) JSON() (string, error) {
// 	bytes, err := json.Marshal(a)
// 	if err != nil {
// 		return "", err
// 	}
// 	return string(bytes), nil
// }

// func (a *Article) Bytes() ([]byte, error) {
// 	return json.Marshal(a)
// }
