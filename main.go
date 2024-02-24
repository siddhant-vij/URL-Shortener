package main

import (
	"url-shortener/controller"
	"url-shortener/repository"
	"url-shortener/server"
	"url-shortener/view"
)

func main() {
	// repo := repository.NewInMemoryURLStore()
	repo, err := repository.NewMongoDBURLStore()
	if err != nil {
		panic(err)
	}
	ctrl := controller.NewURLShortenerController(repo)
	cliView := view.NewCLIView(ctrl)

	go server.StartHTTPServer(ctrl, "8080")

	cliView.Run()
}
