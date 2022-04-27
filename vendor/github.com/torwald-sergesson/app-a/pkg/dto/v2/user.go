package dto

type User struct {
	ID   int64    `json:"id"`
	Name string   `json:"name"`
	Age  int      `json:"age"`
	Tags []string `json:"tags"`
}
