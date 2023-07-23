package main

import (
	"fmt"
	"ksleemodule/cisco"
	"ksleemodule/ksleeutility"
)

func main() {

	var statusCode, responseString, err = cisco.GetAllSitesInfo()
	if err != nil {
		panic("GetAllSitesInfo")
	}
	fmt.Printf("[%d][%s]\n", statusCode, ksleeutility.GetPrettyStringFromJSONString(responseString))

	statusCode, responseString, err = cisco.GetAllUserInfo()
	if err != nil {
		panic("GetAllUserInfo")
	}
	fmt.Printf("[%d][%s]\n", statusCode, ksleeutility.GetPrettyStringFromJSONString(responseString))

	statusCode, responseString, err = cisco.GetAllTenants()
	if err != nil {
		panic("GetAllTenants")
	}
	fmt.Printf("[%d][%s]\n", statusCode, ksleeutility.GetPrettyStringFromJSONString(responseString))

	statusCode, responseString, err = cisco.CreateTenant()
	if err != nil {
		panic("CreateTenant")
	}
	fmt.Printf("[%d][%s]\n", statusCode, ksleeutility.GetPrettyStringFromJSONString(responseString))

	//var aciErrorResponse = cisco.ACIErrorResponse{}
	//if statusCode == 401 {
	//	fmt.Printf("[" + responseString + "]\n")
	//
	//	// 응답코드가 401 이면 유효기간 만료, 빈토큰 값 등 다시 로그인해서 토큰을 받아와야 하는 상황이다.
	//	cisco.LoginDCloud()
	//		statusCode, responseString, err = cisco.GetAllSitesInfo()
	//	if statusCode != 200 {
	//		panic("Fail to get all sites information")
	//	}
	//	fmt.Printf("[%s]\n", responseString)
	//	fmt.Printf("[%s]\n", ksleeutility.GetPrettyStringFromJSONString(responseString))
	//} else if statusCode != 200 {
	//	fmt.Println("http 에러!!")
	//	fmt.Println("응답 코드 :", statusCode)
	//	return
	//}

	// 결과 출력

	// token 유효 기간 확인

	// 요휴기간이 지난 경우 다시 로그인해서 토큰 받아온 다음 전역변수에 담아주기

	// NDO API Test
}
