package service

import (
	"html/template"
	"net/http"

	ascii "ascii/ascii-art"
)

type page struct {
	Fs     string
	Input  string
	Output string
}

var tmp, err = template.ParseFiles("template/index.html")

func MainPageHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 "+http.StatusText(404), 404)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, "405 "+http.StatusText(405), 405)
		return
	}
	myPage := page{}
	tmp.Execute(w, myPage)
}

func PostPageHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/ascii-art/" {
		http.Error(w, "404 "+http.StatusText(404), 404)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "405 "+http.StatusText(405), 405)
		return
	}
	var site = page{}
	site.Fs = r.FormValue("fs")
	site.Input = r.FormValue("input")
	if site.Input == "" || !correctFs(site.Fs) {
		http.Error(w, "400 "+http.StatusText(400), 400)
		return
	}

	var err int
	site.Output, err = ascii.ASCII(site.Input, site.Fs)
	if err != http.StatusOK {
		if err == 400 {
			http.Error(w, "400 "+http.StatusText(400), 400)
			return
		}
		http.Error(w, "500 "+http.StatusText(500), 500)
		return
	}

	tmp.Execute(w, site)
}

func correctFs(s string) bool {
	if s == "standard" || s == "shadow" || s == "thinkertoy" || s == "zigzag" {
		return true
	}
	return false
}
