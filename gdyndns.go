package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type ipinfo struct {
	IP string `json:"ip"`
}

func GetIP() (ip string, err error) {
	i := &ipinfo{}
	c := http.DefaultClient
	r, err := c.Get("https://ipinfo.io")
	if err != nil {
		return "", err
	}
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return "", err
	}
	err = json.Unmarshal(b, i)
	if err != nil {
		return "", err
	}
	return i.IP, err
}

func SetIP(u, p, h, i string) (err error) {
	URL := fmt.Sprintf("https://%s:%s@domains.google.com/nic/update?hostname=%s&myip=%s", u, p, h, i)
	c := http.DefaultClient
	r, err := c.Get(URL)
	if err != nil {
		return err
	}
	if r.StatusCode != http.StatusOK {
		return fmt.Errorf("an error occurred: %s", r.Status)
	}
	return err
}

func main() {
	var googleDomainsUsername, googleDomainsPassword, googleDomainsHostname string
	flag.StringVar(&googleDomainsUsername, "username", "", "your google domains username")
	flag.StringVar(&googleDomainsPassword, "password", "", "your google domains password")
	flag.StringVar(&googleDomainsHostname, "hostname", "", "your google domains hostname to update")
	flag.Parse()

	switch {
	case googleDomainsHostname == "":
		log.Fatal("--hostname must not be blank")
	case googleDomainsUsername == "":
		log.Fatal("--username must not be blank")
	case googleDomainsPassword == "":
		log.Fatal("--password must not be blank")
	}

	ip, err := GetIP()
	if err != nil {
		log.Fatal(err)
	}

	if err := SetIP(googleDomainsUsername, googleDomainsPassword, googleDomainsHostname, ip); err != nil {
		log.Fatal(err)
	}
	log.Printf("Successfully updated %s to %s\n", googleDomainsHostname, ip)
}