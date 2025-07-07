package config

var AllowedOrigins = []string{
	"http://localhost:3000",
	"http://localhost:5173",
	"http://localhost:6969",
	"https://ubiquitous-belekoy-512d48.netlify.app/",
}

func GetAllowedOrigins() []string {
	return AllowedOrigins
}