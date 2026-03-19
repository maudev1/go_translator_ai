package models

type Config struct {
	ID        int
	Language  string
	Menu      [3]string
	InputFile string
}

func GetConfig() Config {
	return Config{
		ID:        1,
		Language:  "pt_br",
		Menu:      [3]string{"Go", "Java", "PHP"},
		InputFile: "files/input/base.file",
	}
}
