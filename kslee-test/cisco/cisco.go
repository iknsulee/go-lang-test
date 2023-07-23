package cisco

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"io"
	"ksleemodule/ksleeutility"
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

func getHttp(url string) (int, string) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Add("Authorization", "Bearer "+AciNdoLogin.Jwttoken)

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

func postHttp(url string, reqBody string) (int, string) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(reqBody)))
	if err != nil {
		panic(err)
	}
	req.Header.Add("Authorization", "Bearer "+AciNdoLogin.Jwttoken)
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

func isSessionExpired(statusCode int) bool {
	if statusCode == 401 {
		// 응답코드가 401 이면 유효기간 만료, 빈토큰 값 등 다시 로그인해서 토큰을 받아와야 하는 상황이다.
		LoginDCloud()
		return true
	} else {
		return false
	}
}

func CreateTenant() (int, string, error) {
	var statusCode, httpResponse = _createTenant()

	if isSessionExpired(statusCode) {
		statusCode, httpResponse = _createTenant()
	}

	return statusCode, httpResponse, nil
}

func _createTenant() (int, string) {
	var reqTenant = `{
		"name": "kslee-test",
		"displayName": "kslee-test",
		"siteAssociations": [
			{
				"siteId": "64b8f2a376fa8d974ea1238a"
		   },
			{
				"siteId": "64b8f2b576fa8d974ea1238b"
			}
		],
		"userAssociations": [
			{
				"userId": "48d105bdfbc49a5fcf39a10b961386e61dfe40085cb0315d818b761c5735dafa"
			}
		],
		"_updateVersion": 0
	}`

	return postHttp("https://198.18.133.100/mso/api/v1/tenants", reqTenant)

}

func GetAllSitesInfo() (int, string, error) {
	var statusCode, httpResponse = _getAllSitesInfo()

	if isSessionExpired(statusCode) {
		statusCode, httpResponse = _getAllSitesInfo()
	}

	return statusCode, httpResponse, nil
}

func GetAllUserInfo() (int, string, error) {
	var statusCode, httpResponse = _getAllUserInfo()

	if isSessionExpired(statusCode) {
		statusCode, httpResponse = _getAllUserInfo()
	}

	return statusCode, httpResponse, nil
}

func GetAllTenants() (int, string, error) {
	var statusCode, httpResponse = _getAllTenants()

	if isSessionExpired(statusCode) {
		statusCode, httpResponse = _getAllTenants()
	}

	return statusCode, httpResponse, nil
}

func _getAllSitesInfo() (int, string) {
	return getHttp("https://198.18.133.100/mso/api/v2/sites")
}

func _getAllUserInfo() (int, string) {
	return getHttp("https://198.18.133.100/mso/api/v1/users")
}

func _getAllTenants() (int, string) {
	return getHttp("https://198.18.133.100/mso/api/v1/tenants")
}
