package main

import (
	"encoding/json"
	"net/http"

	"github.com/gocolly/colly/v2"
)

type APIResponse struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Status      int      `json:"status"`
	Icons       []string `json:"favicons"`
}

type APIRequest struct {
	Link string `json:"link"`
}

func PreviewHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		d := &APIRequest{}
		err := json.NewDecoder(r.Body).Decode(d)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("'link' property missing from request body."))
			return
		}

		res := &APIResponse{
			Icons: make([]string, 0),
		}
		c := colly.NewCollector()

		c.OnHTML("meta[name='description']", func(e *colly.HTMLElement) {
			res.Description = e.Attr("content")
		})

		c.OnHTML("title", func(h *colly.HTMLElement) {
			res.Title = h.Text
		})

		c.OnHTML("link[rel='icon']", func(h *colly.HTMLElement) {
			res.Icons = append(res.Icons, h.Request.AbsoluteURL(h.Attr("href")))
		})

		c.OnResponse(func(r *colly.Response) {
			res.Status = r.StatusCode
		})

		err = c.Visit(d.Link)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		b, err := json.Marshal(res)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(b)
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}
