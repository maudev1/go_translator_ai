package models

type Config struct {
	ID        int
	Language  string
	InputFile string
}

func GetConfig() Config {
	return Config{
		ID:        1,
		Language:  "pt_br",
		InputFile: "files/input/base.file",
	}
}
