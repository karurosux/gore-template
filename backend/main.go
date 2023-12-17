package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	echo "github.com/labstack/echo/v4"
	do "github.com/samber/do"
	lo "github.com/samber/lo"
)

func main() {
	var requestedSeeding bool = lo.Some[string](os.Args, []string{"--seed-db"})
	var generateTsTypes bool = lo.Some[string](os.Args, []string{"--generate-ts"})

	dotEnvErr := godotenv.Load(".env")

	if dotEnvErr != nil {
		log.Fatalf("Error loading .env file")
	}

	mode := os.Getenv("MODE")

	switch {
	case requestedSeeding && mode == "development":
		RunSeeder()
	case generateTsTypes && mode == "development":
		GenerateTsTypes()
	default:
		RunHTTPServer()
	}
}

func RunSeeder() {
	Seed()
}

func RunHTTPServer() {
	app := echo.New()
	container := do.New()

	var s *Server = &Server{
		app:       app,
		container: container,
	}

	s.Start()
}
