package main

import (
	"fmt"
	"net"
	"net/http"
	"strings"
	"html/template"

	"github.com/miekg/dns"
)

type pageData struct {
	Title string
	Items []string
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
		domainResult := result[len(result) - 1]
		authorizedDomains = append(authorizedDomains, domainResult[1:len(domainResult) - 1])
	}

	return authorizedDomains, err
}


func root(w http.ResponseWriter, req *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html")

	if err != nil {
		fmt.Println(err)
	}

	data := pageData{
		Title: "",
		PlaceHolder: "Your Domain",
	}

	tmpl.Execute(w, data)
}

func caaTester(w http.ResponseWriter, req *http.Request){
	var data pageData

	domain := req.URL.Query().Get("domain")

	// A fully qualified domain requires a dot at the end
	domain = domain + "."

	if domain == "" {
		http.ServeFile(w, req, "templates/index.html")
	}

	authorizezdDomains, err := getCAA(domain)

	if err != nil {
		fmt.Println(err)
	}
	
	tmpl, err := template.ParseFiles("templates/index.html")

	if err != nil {
		fmt.Println(err)
	}

	if len(authorizezdDomains) == 0 {
		data = pageData{
			Title: "No results, any CA can issue an SSL certificate for your domain",
			Items: authorizezdDomains,
			PlaceHolder: req.URL.Query().Get("domain"),
		}
	} else {
		data = pageData{
			Title: "CAs that can issue you an SSL certificate -",
			Items: authorizezdDomains,
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
