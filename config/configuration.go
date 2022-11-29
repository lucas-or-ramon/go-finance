package config

type Web struct {
	Address string
}

type Database struct {
	URL          string
	DatabaseName string
}

type Configuration struct {
	Web      Web
	Database Database
}
