package utild

import "net/http"

func GetRequestDomain(r *http.Request) string {
	url := ""
	if r.TLS == nil {
		url = "http://"
	} else {
		url = "https://"
	}
	return url + r.Host
}

func GetRequestLang(r *http.Request) string {
	return r.FormValue("lang")
}
