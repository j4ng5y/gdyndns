# GDYNDNS

This is a simple little application that can be set to run via whatever means you want, but all it does is simple update a synthetic dynamic dns record for a google domains domain.

Please see [This URL](https://support.google.com/domains/answer/6147083?hl=en) if you don't know what I'm talking about.

## Usage

```bash
Usage of ./gdyndns:
  -hostname string
        your google domains hostname to update
  -password string
        your google domains password
  -username string
        your google domains username
  -ipinfo-token string
        your IPINFO.io API Token
```

## Installation
### Prereqs:

* Go >= 1.14.0
* make

### Installing:

1) clone this repo
2) `cd` into this repo
3) `make && make install`

If you use systemd (and there is a good chance you do), you can install a service for this as well using `make install-service`, but you will need to modify the `/etc/systemd/system/gdyndns.service` file with your actual information. Then, you should do a `systemctl daemon-reload && systemctl restart gdyndns.timer`.

## Note

This uses the ipinfo.io backend to get your current IP address, so if you don't set the `-ipinfo-token` flag, you will be limited to 1000 requests per day (which is probably ok).

If you need up to 50,000/mo requests (which might be the case if you use this across a bunch of servers), you will need to sign up at [https://ipinfo.io](https://ipinfo.io/signup).

If you need more than 50,000/mo requests, you are going to have to use a paid ipinfo account :smile:

## Disclaimer

I am not affiliated with ipinfo.io at all, i just like their service.