package models

type Config struct {
	Mongo struct {
		Addr string `yaml:"addr"`
		Port int    `yaml:"port"`
		User string `yaml:"user"`
		Pass string `yaml:"pass"`
	}

	Rest struct {
		Port int `yaml:"port"`
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
