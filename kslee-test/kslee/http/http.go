package http

import (
	"bytes"
	"crypto/tls"
	"io"
	"net/http"
)

func GetHttp(token string, url string) (int, string) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Add("Authorization", "Bearer "+token)

	transCfg := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // ignore expired SSL certificates
	}
	client := &http.Client{Transport: transCfg}

	// Client 객체에서 Request 실행
	httpResponse, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer httpResponse.Body.Close()

	var bodyBytes, _ = io.ReadAll(httpResponse.Body)
	var httpResponseString = string(bodyBytes) //바이트를 문자열로
	return httpResponse.StatusCode, httpResponseString
}

func PostHttp(token string, url string, reqBody string) (int, string) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(reqBody)))
	if err != nil {
		panic(err)
	}
	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Add("Content-Type", "application/json")

	transCfg := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // ignore expired SSL certificates
	}
	client := &http.Client{Transport: transCfg}

	// Client 객체에서 Request 실행
	httpResponse, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer httpResponse.Body.Close()

	var bodyBytes, _ = io.ReadAll(httpResponse.Body)
	var httpResponseString = string(bodyBytes) //바이트를 문자열로
	//return httpResponse, bodyBytes, httpResponseString
	return httpResponse.StatusCode, httpResponseString
}

func DeleteHttp(token string, url string, reqBody string) (int, string) {
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Add("Authorization", "Bearer "+token)

	transCfg := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // ignore expired SSL certificates
	}
	client := &http.Client{Transport: transCfg}

	// Client 객체에서 Request 실행
	httpResponse, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer httpResponse.Body.Close()

	var bodyBytes, _ = io.ReadAll(httpResponse.Body)
	var httpResponseString = string(bodyBytes) //바이트를 문자열로
	//return httpResponse, bodyBytes, httpResponseString
	return httpResponse.StatusCode, httpResponseString
}
