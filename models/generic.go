package models

type Route struct {
	Method string `json:"method"`
	Uri    string `json:"uri"`
}
type Routes []Route
