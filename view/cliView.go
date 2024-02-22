package view

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"url-shortener/controller"
)

type CLIView struct {
	controller *controller.URLShortenerController
}

func NewCLIView(controller *controller.URLShortenerController) *CLIView {
	return &CLIView{controller: controller}
}

func (v *CLIView) Run() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("URL Shortener CLI")
	fmt.Print("Enter '<url>' to shorten a URL: ")

	for scanner.Scan() {
		input := scanner.Text()
		args := strings.Fields(input)

		if len(args) < 1 {
			fmt.Println("Invalid command. Please try again.")
			continue
		}
		v.controller.ShortenURL(args[0])
	}
}
