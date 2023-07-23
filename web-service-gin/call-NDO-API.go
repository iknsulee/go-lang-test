package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"kslee/cisco"
	_ "kslee/cisco"
	_ "kslee/ksleelogin"
	_ "kslee/ksleeutility"
	"net/http"
)

func main() {

	var statusCode, responseString, err = GetAllSitesInfo()
	if err != nil {
		panic("GetAllSitesInfo")
	}

	//var aciErrorResponse = cisco.ACIErrorResponse{}
	if statusCode == 401 {
		fmt.Printf("[" + responseString + "]\n")

		// 응답코드가 401 이면 유효기간 만료, 빈토큰 값 등 다시 로그인해서 토큰을 받아와야 하는 상황이다.
		cisco.LoginDCloud()
		statusCode, responseString, err = GetAllSitesInfo()
		if statusCode != 200 {
			panic("Fail to get all sites information")
		}
	} else if statusCode != 200 {
		fmt.Println("http 에러!!")
		fmt.Println("응답 코드 :", statusCode)
		return
	}

	// 결과 출력

	// token 유효 기간 확인

	// 요휴기간이 지난 경우 다시 로그인해서 토큰 받아온 다음 전역변수에 담아주기

	// NDO API Test
}

func GetAllSitesInfo() (int, string, error) {

	fmt.Printf("------------------------------------------------------------\n")
	fmt.Printf("Get all sites information\n")
	fmt.Printf("------------------------------------------------------------\n")
	req, err := http.NewRequest("GET", "https://198.18.133.100/mso/api/v2/sites", nil)
	if err != nil {
		panic(err)
	}
	req.Header.Add("Authorization", "Bearer "+cisco.AciNdoLogin.Jwttoken)

	transCfg := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // ignore expired SSL certificates
	}
	client := &http.Client{Transport: transCfg}
	//client := &http.Client{}

	// Client 객체에서 Request 실행
	httpResponse, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer httpResponse.Body.Close()

	bytes, _ := io.ReadAll(httpResponse.Body)
	httpResponseString := string(bytes) //바이트를 문자열로

	return httpResponse.StatusCode, httpResponseString, nil
}
