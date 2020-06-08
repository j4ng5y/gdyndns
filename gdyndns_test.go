package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestGoogleDomainsResponse_CheckError(t *testing.T) {
	nilBody := GoogleDomainsResponse(http.Response{
		Body: nil,
	})
	if err := nilBody.CheckError(); err == nil {
		t.Fail()
	} else {
		t.Log(err)
	}

	good := GoogleDomainsResponse(http.Response{
		StatusCode: http.StatusOK,
		Body: ioutil.NopCloser(bytes.NewBuffer([]byte("good 1.2.3.4"))),
	})
	if err := good.CheckError(); err != nil {
		t.Fail()
	}

	nochg := GoogleDomainsResponse(http.Response{
		StatusCode: http.StatusBadRequest,
		Body: ioutil.NopCloser(bytes.NewBuffer([]byte("nochg 1.2.3.4"))),
	})
	if err := nochg.CheckError(); err != nil {
		t.Fail()
	}

	nohost := GoogleDomainsResponse(http.Response{
		StatusCode: http.StatusBadRequest,
		Body: ioutil.NopCloser(bytes.NewBuffer([]byte("nohost"))),
	})
	if err := nohost.CheckError(); err == nil {
		t.Fail()
	} else {
		t.Log(err)
	}

	badauth := GoogleDomainsResponse(http.Response{
		StatusCode: http.StatusBadRequest,
		Body: ioutil.NopCloser(bytes.NewBuffer([]byte("badauth"))),
	})
	if err := badauth.CheckError(); err == nil {
		t.Fail()
	} else {
		t.Log(err)
	}

	notfqdn := GoogleDomainsResponse(http.Response{
		StatusCode: http.StatusBadRequest,
		Body: ioutil.NopCloser(bytes.NewBuffer([]byte("notfqdn"))),
	})
	if err := notfqdn.CheckError(); err == nil {
		t.Fail()
	} else {
		t.Log(err)
	}

	badagent := GoogleDomainsResponse(http.Response{
		StatusCode: http.StatusBadRequest,
		Body: ioutil.NopCloser(bytes.NewBuffer([]byte("badagent"))),
	})
	if err := badagent.CheckError(); err == nil {
		t.Fail()
	} else {
		t.Log(err)
	}

	abuse := GoogleDomainsResponse(http.Response{
		StatusCode: http.StatusBadRequest,
		Body: ioutil.NopCloser(bytes.NewBuffer([]byte("abuse"))),
	})
	if err := abuse.CheckError(); err == nil {
		t.Fail()
	} else {
		t.Log(err)
	}

	nineOneOne := GoogleDomainsResponse(http.Response{
		StatusCode: http.StatusBadRequest,
		Body: ioutil.NopCloser(bytes.NewBuffer([]byte("911"))),
	})
	if err := nineOneOne.CheckError(); err == nil {
		t.Fail()
	} else {
		t.Log(err)
	}

	conflictA := GoogleDomainsResponse(http.Response{
		StatusCode: http.StatusBadRequest,
		Body: ioutil.NopCloser(bytes.NewBuffer([]byte("conflict A"))),
	})
	if err := conflictA.CheckError(); err == nil {
		t.Fail()
	} else {
		t.Log(err)
	}

	conflictAAAA := GoogleDomainsResponse(http.Response{
		StatusCode: http.StatusBadRequest,
		Body: ioutil.NopCloser(bytes.NewBuffer([]byte("conflict AAAA"))),
	})
	if err := conflictAAAA.CheckError(); err == nil {
		t.Fail()
	} else {
		t.Log(err)
	}

	junk := GoogleDomainsResponse(http.Response{
		StatusCode: http.StatusOK,
		Body: ioutil.NopCloser(bytes.NewBuffer([]byte("junk"))),
	})
	if err := junk.CheckError(); err == nil {
		t.Fail()
	} else {
		t.Log(err)
	}
}