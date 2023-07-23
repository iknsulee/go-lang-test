package ksleelogin

//import (
//	"bytes"
//	"crypto/tls"
//	"encoding/json"
//	"fmt"
//	"io"
//	"log"
//	"net/http"
//)
//
//type NdoLogin struct {
//	Jwttoken   string `json:"jwttoken"`
//	Username   string `json:"username"`
//	Usertype   string
//	Rbac       string
//	StatusCode int
//	Token      string
//}
//
//func loginDCloud() {
//	reqBody := bytes.NewBufferString("{\"username\": \"admin\", \"password\": \"C1sco12345\"}")
//
//	// https://www.socketloop.com/tutorials/golang-disable-security-check-for-http-ssl-with-bad-or-expired-certificate
//	transCfg := &http.Transport{
//		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // ignore expired SSL certificates
//	}
//	client := &http.Client{Transport: transCfg}
//
//	resp, err :=
//		client.Post(
//			"https://198.18.133.100/api/v1/auth/login",
//			"",
//			reqBody)
//	if err != nil {
//		panic(err)
//	}
//
//	defer resp.Body.Close()
//
//	respBody, err := io.ReadAll(resp.Body)
//	if err == nil {
//		str := string(respBody)
//		println(str)
//	}
//
//	var ndoLogin = NdoLogin{}
//
//	// parse JSON string to ksleelogin response struct
//	err = json.Unmarshal(respBody, &ndoLogin)
//	if err != nil {
//		panic(err)
//		return
//	}
//	//fmt.Printf("[%s]", ndoLogin)
//
//	// 구조체를 예쁘게 JSON 형식으로 출력하기 위한 함수
//	ndoLoginString, err := json.MarshalIndent(ndoLogin, "", "  ")
//	if err != nil {
//		log.Fatalf(err.Error())
//	}
//	fmt.Printf("MarshalIndent funnction output %s\n", string(ndoLoginString))
//
//}
