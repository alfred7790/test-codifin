package main

import "codifin-challenge/infrastructure/web"

// @title Codifin Challenge API
// @version 1.0
// @description Service to manage products
// @BasePath /v1

// @schemes http https
// @BasePath /v1
// @Produce json
// @Consumes json

// @contact.name   API Support
// @contact.email  alfred.7790@gmail.com
func main() {
	server := web.NewServer()
	server.Run()
}
