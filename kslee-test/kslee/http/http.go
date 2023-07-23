package http

import (
	"bytes"
	"crypto/tls"
	"io"
	"net/http"
)

func GetHttp(token string, url string) (int, string) {
	return _http("GET", token, url, "")

}

func PostHttp(token string, url string, reqBody string) (int, string) {
	return _http("POST", token, url, reqBody)
}

func PatchHttp(token string, url string, reqBody string) (int, string) {
	return _http("PATCH", token, url, reqBody)
}

func DeleteHttp(token string, url string, reqBody string) (int, string) {
	return _http("DELETE", token, url, "")
}

func _http(method string, token string, url string, reqBody string) (int, string) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer([]byte(reqBody)))
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
