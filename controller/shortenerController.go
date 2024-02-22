package controller

import (
	"fmt"
	"url-shortener/model"
	"url-shortener/repository"
)

type URLShortenerController struct {
	repo repository.URLRepository
}

func NewURLShortenerController(repo repository.URLRepository) *URLShortenerController {
	return &URLShortenerController{repo: repo}
}

func (c *URLShortenerController) ShortenURL(originalURL string) {
	urlData := model.URLData{OriginalURL: originalURL}
	err := urlData.Validate()
	if err != nil {
		fmt.Println("Error validating URL:", err)
		return
	}

	counter := int64(0)
	for {
		shortened := model.GenerateHash(urlData.OriginalURL, counter)
		if !c.repo.Exists(shortened) {
			urlData.ShortenedURL = shortened
			break
		}
		counter++
	}

	err = c.repo.Save(urlData)
	if err != nil {
		fmt.Println("Error saving URL:", err)
		return
	}

	fmt.Printf("Shortened URL: http://localhost:8080/%s\n", urlData.ShortenedURL)
}

func (c *URLShortenerController) RetrieveOriginalURL(shortURL string) (string, error) {
	urlData, err := c.repo.FindByShortURL(shortURL)
	if err != nil {
		fmt.Println("Error retrieving URL:", err)
		return "", err
	}

	return urlData.OriginalURL, nil
}
