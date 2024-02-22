package server

import (
	"net/http"
	"url-shortener/controller"
)

func StartHTTPServer(ctrl *controller.URLShortenerController, port string) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path[1:]
		originalURL, err := ctrl.RetrieveOriginalURL(path)
		if err != nil {
			http.Error(w, "URL not found", http.StatusNotFound)
			return
		}
		http.Redirect(w, r, originalURL, http.StatusFound)
	})

	http.ListenAndServe(":"+port, nil)
}
