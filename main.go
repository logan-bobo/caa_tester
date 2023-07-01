package main

import (
	"fmt"
	"html/template"
	"net"
	"net/http"
	"strings"

	"github.com/miekg/dns"
	"github.com/dchest/validator"
)

type pageData struct {
	Title       string
	Items       []string
	PlaceHolder string
}

func getCAA(domain string) (authorizedDomains []string, err error) {
	config, _ := dns.ClientConfigFromFile("/etc/resolv.conf")

	client := dns.Client{}
	msg := dns.Msg{}

	msg.SetQuestion(domain, dns.TypeCAA)

	response, _, err := client.Exchange(&msg, net.JoinHostPort(config.Servers[0], config.Port))

	if err != nil {
		return authorizedDomains, err
	}

	for _, answer := range response.Answer {
		result := strings.Split(answer.String(), " ")
		result[0] = result[0][len(result[0]) - 1:]
		authorizedDomains = append(authorizedDomains, strings.Join(result, " "))
	}

	return authorizedDomains, err
}

func root(w http.ResponseWriter, req *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html")

	if err != nil {
		fmt.Println(err)
	}

	data := pageData{
		Title:       "",
		PlaceHolder: "Your Domain",
	}

	tmpl.Execute(w, data)
}

func caaTester(w http.ResponseWriter, req *http.Request) {
	var data pageData

	tmpl, err := template.ParseFiles("templates/index.html")

	if err != nil {
		fmt.Println(err)
	}

	domain := req.URL.Query().Get("domain")

	if !validator.IsValidDomain(domain) {
		data = pageData{
			Title:       "Not a valid domain name",
			PlaceHolder: req.URL.Query().Get("domain"),
		}
		tmpl.Execute(w, data)
		return
	}

	// A fully qualified domain requires a dot at the end
	domain = domain + "."

	if domain == "" {
		http.ServeFile(w, req, "templates/index.html")
	}

	authorizezdDomains, err := getCAA(domain)

	if err != nil {
		fmt.Println(err)
	}

	if len(authorizezdDomains) == 0 {
		data = pageData{
			Title:       "No results, any CA can issue an SSL certificate for your domain",
			Items:       authorizezdDomains,
			PlaceHolder: req.URL.Query().Get("domain"),
		}
	} else {
		data = pageData{
			Title:       "CAs that can issue you an SSL certificate -",
			Items:       authorizezdDomains,
			PlaceHolder: req.URL.Query().Get("domain"),
		}
	}

	tmpl.Execute(w, data)
}

func main() {
	http.HandleFunc("/", root)
	http.HandleFunc("/verify_caa", caaTester)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.ListenAndServe(":8090", nil)
}
