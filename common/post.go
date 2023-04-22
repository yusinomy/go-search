package common

import (
	"bytes"
	"crypto/tls"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"time"
)

func Post(target string, data []byte) []string {
	s := []string{}
	var tr *http.Transport
	if HttpProxy != "" {
		uri, _ := url.Parse(HttpProxy)
		tr = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			Proxy:           http.ProxyURL(uri),
		}
	} else {
		tr = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
	}
	client := &http.Client{
		Timeout:   time.Duration(Timeout) * time.Second,
		Transport: tr,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse //不允许跳转
		}}
	req, err := http.NewRequest("POST", target, bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json;charset=utf8;")
	if err != nil {
		log.Println(err)
	}
	get, err := client.Do(req)
	exp, _ := regexp.Compile(`http[s]{0,1}://(([a-zA-Z0-9\._-]+\.[a-zA-Z]{2,6})|([0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}))(:[0-9]{1,4})*(/[a-zA-Z0-9\&%_\./-~-]*)?`)
	match := exp.FindAllString(conetent(get), -1)
	for z, _ := range match {
		a := match[z]
		s = append(s, a)
	}
	return s
}
