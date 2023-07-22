package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

var token = "eyJhbGciOiJSUzI1NiIsImtpZCI6IjlrMm5iZXN4MTBoMG9sdWdrNnlja2I5bzJ3YnBicDBwIiwidHlwIjoiSldUIn0.eyJhdnBhaXIiOiJzaGVsbDpkb21haW5zPWFsbC9hZG1pbi8iLCJjbHVzdGVyIjoiNmU2NDZmMDAtMDAwMC0wMDAwLTAwMDAtMDAwMDAwMDAwMDAwIiwiZXhwIjoxNjg5OTUxNDE2LCJpYXQiOjE2ODk5NTAyMTYsImlkIjoiNDhkMTA1YmRmYmM0OWE1ZmNmMzlhMTBiOTYxMzg2ZTYxZGZlNDAwODVjYjAzMTVkODE4Yjc2MWM1NzM1ZGFmYSIsImlzcyI6Im5kIiwiaXNzLWhvc3QiOiIxOTguMTguMTMzLjEwMCIsInJiYWMiOlt7ImRvbWFpbiI6ImFsbCIsInJvbGVzUiI6MTY3NzcyMTYsInJvbGVzVyI6MSwicm9sZXMiOltbImFkbWluIiwiV3JpdGVQcml2Il0sWyJhcHAtdXNlciIsIlJlYWRQcml2Il1dfV0sInNlc3Npb25pZCI6InhCaHVsR3hVTjhUSkxmeExMPUUyeUNVeiIsInVzZXJmbGFncyI6MCwidXNlcmlkIjoyNTAwMiwidXNlcm5hbWUiOiJhZG1pbiIsInVzZXJ0eXBlIjoibG9jYWwifQ.6Qmt-_i6FLpK-CzGNQcKYFyaYhblGDR2idtWCNy7dIcMlu9bdQLbYb_GVUw3F-uDu_vGQrwLMA5bS0K61zuL-GQdDas0DjFo4dh1IrRP6g0FY_YJGUvffs0xigEauGVThnfTRWfkub-pw6UhAPd5C13ErK_eOCSDXq93I0ON6FwO9TGbUgxgYeayLifH_KfY8wCaMSl4wzRLXfEK2kC83x0x0vdiW10Xj372cEZZkyL9v9cloUAi1yLUeSS-OuIoqc2vVBLGEU2BrjsQajv0CxlxQ-e0xtFIlfPPKZIBUFS1iRtKwr01wYQgQ1LwhkVUT8uqMbp5Izd24FH72DFypA"

type ErrorResponse struct {
	Error string `json:"error"`
}

type NdoLogin2 struct {
	Jwttoken   string `json:"jwttoken"`
	Username   string `json:"username"`
	Usertype   string
	Rbac       string
	StatusCode int
	Token      string
}

func main() {

	fmt.Println("Start of Call NDO API")
	fmt.Printf("token[%s]\n", token)

	// token 값 가지고 NDO API 호출
	// Request 객체 생성
	req, err := http.NewRequest("GET", "https://198.18.133.100/mso/api/v2/sites", nil)
	if err != nil {
		panic(err)
	}
	req.Header.Add("Authorization", "Bearer "+token)

	transCfg := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // ignore expired SSL certificates
	}
	client := &http.Client{Transport: transCfg}
	//client := &http.Client{}

	// Client 객체에서 Request 실행
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	bytes, _ := io.ReadAll(resp.Body)
	responseString := string(bytes) //바이트를 문자열로
	fmt.Println(responseString)

	var errorResponse = ErrorResponse{}
	if resp.StatusCode == 401 {
		json.Unmarshal(bytes, &errorResponse)
		fmt.Printf("%s\n", errorResponse)
		if strings.Contains(errorResponse.Error, "expired") {
			fmt.Printf("Token has been expired. We're loging again!")

			newNdoLogin := login()
			fmt.Printf("[%s]", newNdoLogin)

		} else {
			return
		}
	} else if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		fmt.Println("http 에러!!")
		fmt.Println("응답 코드 :", resp.StatusCode)
		return
	}

	// 결과 출력

	// token 유효 기간 확인

	// 요휴기간이 지난 경우 다시 로그인해서 토큰 받아온 다음 전역변수에 담아주기

	// NDO API Test
}

func login() NdoLogin2 {

	transCfg := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // ignore expired SSL certificates
	}
	client := &http.Client{Transport: transCfg}

	reqBody := bytes.NewBufferString("{\"username\": \"admin\", \"password\": \"C1sco12345\"}")
	resp, err :=
		client.Post(
			"https://198.18.133.100/api/v1/auth/login",
			"",
			reqBody)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err == nil {
		str := string(respBody)
		println(str)
	}

	var ndoLogin = NdoLogin2{}

	// parse JSON string to login response struct
	err = json.Unmarshal(respBody, &ndoLogin)
	if err != nil {
		panic(err)
		return ndoLogin
	}
	fmt.Printf("[%s]", ndoLogin)

	return ndoLogin
}
