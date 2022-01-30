package model

type Application struct {
	App        App        `json:"app"`
	Server     Server     `json:"servers"`
}

type App struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Version     string `json:"version"`
}

type Server struct {
	Port int `json:"port"`
}