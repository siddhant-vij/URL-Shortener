package main

import (
	"url-shortener/controller"
	"url-shortener/repository"
	"url-shortener/server"
	"url-shortener/view"
)

func main() {
	repo := repository.NewInMemoryURLStore()
	ctrl := controller.NewURLShortenerController(repo)
	cliView := view.NewCLIView(ctrl)

	go server.StartHTTPServer(ctrl, "8080")

	cliView.Run()
}
