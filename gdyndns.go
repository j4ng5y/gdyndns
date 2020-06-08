package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/ipinfo/go-ipinfo/ipinfo"
)

const (
	RESPONSE_GOOD = "good"
	RESPONSE_NOCHANGE = "nochg"
	RESPONSE_HOHOST = "nohost"
	RESPONSE_BADAUTH = "badauth"
	RESPONSE_NOTFQDN = "notfqdn"
	RESPONSE_BADAGENT = "badagent"
	RESPONSE_ABUSE = "abuse"
	RESPONSE_911 = "911"
	RESPONSE_CONFLICT = "conflict"
)

type GoogleDomainsResponse http.Response

func (G GoogleDomainsResponse) CheckError() (err error) {
	var b []byte
	if G.Body != nil {
		b, err = ioutil.ReadAll(G.Body)
		if err != nil {
			return err
		}
	} else {
		return fmt.Errorf("body can not be nil")
	}

	r := strings.Split(string(b), " ")
	switch r[0] {
	case RESPONSE_GOOD:
		log.Printf("success: updated DNS to point to %s", r[1])
		return nil
	case RESPONSE_NOCHANGE:
		log.Printf("success: no change required")
		return nil
	case RESPONSE_HOHOST:
		return fmt.Errorf("error: the hostname does not exist, or does not have Dynamic DNS enabled")
	case RESPONSE_BADAUTH:
		return fmt.Errorf("error: the username / password combination is not valid for the specified host")
	case RESPONSE_NOTFQDN:
		return fmt.Errorf("error: the supplied hostname is not a valid fully-qualified domain name")
	case RESPONSE_BADAGENT:
		return fmt.Errorf("error: your dynamic DNS client is making bad requests. ensure the user agent is set in teh request")
	case RESPONSE_ABUSE:
		return fmt.Errorf("error: dynamic DNS access for the hostname has been blocked due to failure to interpret previous responses correctly")
	case RESPONSE_911:
		return fmt.Errorf("error: an error happened on our end. Wait 5 minute and try again")
	case RESPONSE_CONFLICT:
		return fmt.Errorf("error: a custom %s resource record conflicts with the update, delete the indicated resource record within the DNS settings page and try the update again", r[1])
	default:
		return fmt.Errorf("error: response string did not conform: %s", b)
	}
}

func GetIP(token string) (ip string, err error) {
	var c *ipinfo.Client
	if token != "" {
		auth := ipinfo.AuthTransport{Token: token}
		authc := auth.Client()
		c = ipinfo.NewClient(authc)
	} else {
		c = ipinfo.NewClient(http.DefaultClient)
	}

	info, err := c.GetIP(nil)
	if err != nil {
		return "", err
	}
	return info.String(), err
}

func SetIP(u, p, h, i string) (err error) {
	c := http.DefaultClient
	URL := fmt.Sprintf("https://domains.google.com/nic/update?hostname=%s&myip=%s", h, i)

	req, err := http.NewRequest(http.MethodPost, URL, nil)
	if err != nil {
		return err
	}

	req.Header.Set("User-Agent", "gdyndns/0.1.0")
	req.Header.Set("Authorization", fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", u, p)))))

	resp, err := c.Do(req)
	if err != nil {
		return err
	}

	gDomainsResp := GoogleDomainsResponse(*resp)

	if err := gDomainsResp.CheckError(); err != nil {
		return err
	}

	return err
}

func main() {
	var googleDomainsUsername, googleDomainsPassword, googleDomainsHostname, ipinfoAPIToken string
	flag.StringVar(&googleDomainsUsername, "username", "", "your google domains username")
	flag.StringVar(&googleDomainsPassword, "password", "", "your google domains password")
	flag.StringVar(&googleDomainsHostname, "hostname", "", "your google domains hostname to update")
	flag.StringVar(&ipinfoAPIToken, "ipinfo-token", "", "Your ipinfo API Token")
	flag.Parse()

	switch {
	case googleDomainsHostname == "":
		log.Fatal("--hostname must not be blank")
	case googleDomainsUsername == "":
		log.Fatal("--username must not be blank")
	case googleDomainsPassword == "":
		log.Fatal("--password must not be blank")
	}

	ip, err := GetIP(ipinfoAPIToken)
	if err != nil {
		log.Fatal(err)
	}

	if err := SetIP(googleDomainsUsername, googleDomainsPassword, googleDomainsHostname, ip); err != nil {
		log.Fatal(err)
	}
}