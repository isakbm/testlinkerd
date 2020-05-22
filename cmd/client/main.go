package main

import (
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
)

var (
	cacert = flag.String("cacert", ".ca/tls.crt", "path to cacert file")
)

func main() {

	rootCAs, err := x509.SystemCertPool()
	if err != nil {
		panic(err)
	}

	pem, err := ioutil.ReadFile(*cacert)
	if err != nil {
		panic(err)
	}
	if !rootCAs.AppendCertsFromPEM(pem) {
		panic("unable to append cert from pem")
	}

	https := http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{RootCAs: rootCAs},
		},
	}

	resp, err := https.Get("https://web/sayHello")
	if err != nil {
		fmt.Println(err)
	}
	if resp != nil {
		fmt.Println(resp)
	}

}
