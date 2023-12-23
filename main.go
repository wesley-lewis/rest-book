package main

import "github.com/joho/godotenv"


func main() {
	godotenv.Load()
	s := NewServer(":8000")
	
	s.Start()
}
