package cisco

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"io"
	"kslee/ksleeutility"
	"net/http"
)

var token = "eyJhbGciOiJSUzI1NiIsImtpZCI6IjlrMm5iZXN4MTBoMG9sdWdrNnlja2I5bzJ3YnBicDBwIiwidHlwIjoiSldUIn0.eyJhdnBhaXIiOiJzaGVsbDpkb21haW5zPWFsbC9hZG1pbi8iLCJjbHVzdGVyIjoiNmU2NDZmMDAtMDAwMC0wMDAwLTAwMDAtMDAwMDAwMDAwMDAwIiwiZXhwIjoxNjg5OTUxNDE2LCJpYXQiOjE2ODk5NTAyMTYsImlkIjoiNDhkMTA1YmRmYmM0OWE1ZmNmMzlhMTBiOTYxMzg2ZTYxZGZlNDAwODVjYjAzMTVkODE4Yjc2MWM1NzM1ZGFmYSIsImlzcyI6Im5kIiwiaXNzLWhvc3QiOiIxOTguMTguMTMzLjEwMCIsInJiYWMiOlt7ImRvbWFpbiI6ImFsbCIsInJvbGVzUiI6MTY3NzcyMTYsInJvbGVzVyI6MSwicm9sZXMiOltbImFkbWluIiwiV3JpdGVQcml2Il0sWyJhcHAtdXNlciIsIlJlYWRQcml2Il1dfV0sInNlc3Npb25pZCI6InhCaHVsR3hVTjhUSkxmeExMPUUyeUNVeiIsInVzZXJmbGFncyI6MCwidXNlcmlkIjoyNTAwMiwidXNlcm5hbWUiOiJhZG1pbiIsInVzZXJ0eXBlIjoibG9jYWwifQ.6Qmt-_i6FLpK-CzGNQcKYFyaYhblGDR2idtWCNy7dIcMlu9bdQLbYb_GVUw3F-uDu_vGQrwLMA5bS0K61zuL-GQdDas0DjFo4dh1IrRP6g0FY_YJGUvffs0xigEauGVThnfTRWfkub-pw6UhAPd5C13ErK_eOCSDXq93I0ON6FwO9TGbUgxgYeayLifH_KfY8wCaMSl4wzRLXfEK2kC83x0x0vdiW10Xj372cEZZkyL9v9cloUAi1yLUeSS-OuIoqc2vVBLGEU2BrjsQajv0CxlxQ-e0xtFIlfPPKZIBUFS1iRtKwr01wYQgQ1LwhkVUT8uqMbp5Izd24FH72DFypA"

type ACIErrorResponse struct {
	Error string `json:"error"`
}

type ACINDOLogin struct {
	Jwttoken   string `json:"jwttoken"`
	Username   string `json:"username"`
	Usertype   string
	Rbac       string
	StatusCode int
	Token      string
}

var AciNdoLogin = ACINDOLogin{}

func LoginDCloud() {
	reqBody := bytes.NewBufferString("{\"username\": \"admin\", \"password\": \"C1sco12345\"}")

	// https://www.socketloop.com/tutorials/golang-disable-security-check-for-http-ssl-with-bad-or-expired-certificate
	transCfg := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // ignore expired SSL certificates
	}
	client := &http.Client{Transport: transCfg}

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

	// parse JSON string to ksleelogin response struct
	err = json.Unmarshal(respBody, &AciNdoLogin)
	if err != nil {
		panic(err)
		return
	}

	// 구조체를 예쁘게 JSON 형식으로 출력하기 위한 함수
	ksleeutility.PrintPrettyStruct("set global ACINDOLogin variable: ", AciNdoLogin)

}
