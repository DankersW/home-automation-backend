package models

type Config struct {
	Mongo struct {
		Host string `yaml:"host"`
		Port uint16 `yaml:"port"`
	}

	Rest struct {
		Port1 int `yaml:"port1"`
	}
	Weather struct {
		Key string  `yaml:"key"`
		Lat float64 `yaml:"lat"`
		Lon float64 `yaml:"lon"`
	}
}

type Route struct {
	Method string `json:"method"`
	Uri    string `json:"uri"`
}
type Routes []Route
