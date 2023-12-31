package main

import (
	"github.com/joho/godotenv"
	"rest-book/api"
)

func main() {
	godotenv.Load()
	s := api.NewServer(":8000")
	
	s.Start()
}
