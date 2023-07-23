package main

import (
	"encoding/json"
	"fmt"
	"ksleemodule/cisco"
	"ksleemodule/cisco/schema"
	"ksleemodule/cisco/tenant"
	"ksleemodule/ksleeutility"
	"strings"
)

type Tenant struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
	//SiteAssociation   string `json:"id"`
	Description      string `json:"description"`
	UpdateVersion    int    `json:"_updateVersion"`
	VersionDefaulted bool   `json:"_versionDefaulted"`
}

type TenantInfo struct {
	Tenants []Tenant `json:"tenants"`
}

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

	statusCode, responseString, err = tenant.GetAllTenants()
	if err != nil {
		panic("GetAllTenants")
	}
	fmt.Printf("[%d][%s]\n", statusCode, ksleeutility.GetPrettyStringFromJSONString(responseString))

	statusCode, responseString, err = tenant.CreateTenant()
	if err != nil {
		panic("CreateTenant")
	}
	fmt.Printf("[%d][%s]\n", statusCode, ksleeutility.GetPrettyStringFromJSONString(responseString))

	statusCode, responseString, err = tenant.GetAllTenants()
	if err != nil {
		panic("GetAllTenants")
	}
	fmt.Printf("[%d][%s]\n", statusCode, ksleeutility.GetPrettyStringFromJSONString(responseString))
	var tenantInfo TenantInfo
	err = json.Unmarshal([]byte(responseString), &tenantInfo)
	if err != nil {
		return
	}

	var targetTenant Tenant
	for _, _tenant := range tenantInfo.Tenants {
		if strings.Compare(_tenant.Name, "kslee-test") == 0 {
			targetTenant = _tenant
			break
		}
	}
	ksleeutility.PrintPrettyStruct("Target Tenant", targetTenant)

	var _schema = schema.Schema{}
	_schema.DisplayName = "kslee-test"

	_schema.Templates = make([]schema.Template, 2)

	_schema.Templates[0].Name = "kslee-test-AC"
	_schema.Templates[0].DisplayName = "kslee-test-AC"
	_schema.Templates[0].TenantID = targetTenant.Id
	_schema.Templates[0].TemplateType = "stretched-template"

	_schema.Templates[0].Anps = make([]schema.Anp, 1)
	_schema.Templates[0].Anps[0].Name = "kslee-test"
	_schema.Templates[0].Anps[0].DisplayName = "kslee-test"

	_schema.Templates[0].Vrfs = make([]schema.Vrf, 1)
	_schema.Templates[0].Vrfs[0].Name = "kslee-test-AC_VRF"
	_schema.Templates[0].Vrfs[0].DisplayName = "kslee-test-AC_VRF"

	_schema.Templates[0].Vrfs[0].AutoRouteTargetImport = make([]schema.AutoRouteTargetImport, 1)
	_schema.Templates[0].Vrfs[0].AutoRouteTargetImport[0].Aci = "route-target:as4-nn2:16777227:30780"
	_schema.Templates[0].Vrfs[0].AutoRouteTargetExport = make([]schema.AutoRouteTargetExport, 1)
	_schema.Templates[0].Vrfs[0].AutoRouteTargetExport[0].Aci = "route-target:as4-nn2:16777226:30780"

	_schema.Templates[1].Name = "kslee-test-BC"
	_schema.Templates[1].DisplayName = "kslee-test-BC"
	_schema.Templates[1].TenantID = targetTenant.Id
	_schema.Templates[1].TemplateType = "stretched-template"

	_schema.Templates[1].Anps = make([]schema.Anp, 1)
	_schema.Templates[1].Anps[0].Name = "kslee-test"
	_schema.Templates[1].Anps[0].DisplayName = "kslee-test"

	_schema.Templates[1].Vrfs = make([]schema.Vrf, 1)
	_schema.Templates[1].Vrfs[0].Name = "kslee-test-BC_VRF"
	_schema.Templates[1].Vrfs[0].DisplayName = "kslee-test-BC_VRF"

	_schema.Templates[1].Vrfs[0].AutoRouteTargetImport = make([]schema.AutoRouteTargetImport, 1)
	_schema.Templates[1].Vrfs[0].AutoRouteTargetImport[0].Aci = "route-target:as4-nn2:16777229:30780"
	_schema.Templates[1].Vrfs[0].AutoRouteTargetExport = make([]schema.AutoRouteTargetExport, 1)
	_schema.Templates[1].Vrfs[0].AutoRouteTargetExport[0].Aci = "route-target:as4-nn2:16777228:30780"

	_schema.Sites = make([]schema.Site, 2)
	_schema.Sites[0].SiteID = "64b8f2a376fa8d974ea1238a"
	_schema.Sites[0].TemplateName = "kslee-test-AC"

	_schema.Sites[1].SiteID = "64b8f2b576fa8d974ea1238b"
	_schema.Sites[1].TemplateName = "kslee-test-BC"

	statusCode, responseString, err = schema.CreateSchema(_schema)
	if err != nil {
		panic("CreateSchema")
	}
	fmt.Printf("[%d][%s]\n", statusCode, ksleeutility.GetPrettyStringFromJSONString(responseString))

	statusCode, responseString, err = tenant.DeleteTenant(targetTenant.Id)
	if err != nil {
		panic("DeleteTenant")
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
