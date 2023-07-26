package tenant

import (
	"ksleemodule/cisco"
	"ksleemodule/kslee/http"
)

type CampTenant struct {
	Tenants []struct {
		ID               string `json:"id"`
		Name             string `json:"name"`
		DisplayName      string `json:"displayName"`
		SiteAssociations []struct {
			SiteID          string `json:"siteId"`
			SecurityDomains []any  `json:"securityDomains"`
			AzureAccount    []any  `json:"azureAccount"`
			AwsAccount      []any  `json:"awsAccount"`
			GcpAccount      []any  `json:"gcpAccount"`
			GatewayRouter   []any  `json:"gatewayRouter"`
		} `json:"siteAssociations"`
		UserAssociations []struct {
			UserID string `json:"userId"`
		} `json:"userAssociations"`
		Description      string `json:"description"`
		UpdateVersion    int    `json:"_updateVersion"`
		VersionDefaulted bool   `json:"_versionDefaulted"`
	} `json:"tenants"`
}

func CreateTenant() (int, string, error) {
	var statusCode, httpResponse = _createTenant()

	if cisco.IsSessionExpired(statusCode) {
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
				"siteId": "64bf870b51da8d2bf8e7ed33"
		   },
			{
				"siteId": "64bf874551da8d2bf8e7ed34"
			}
		],
		"userAssociations": [
			{
				"userId": "48d105bdfbc49a5fcf39a10b961386e61dfe40085cb0315d818b761c5735dafa"
			}
		],
		"_updateVersion": 0
	}`

	return http.PostHttp(cisco.AciNdoLogin.Jwttoken, "https://198.18.133.100/mso/api/v1/tenants", reqTenant)

}

func DeleteTenant(id string) (int, string, error) {
	var statusCode, httpResponse = _deleteTenant(id)

	if cisco.IsSessionExpired(statusCode) {
		statusCode, httpResponse = _deleteTenant(id)
	}

	return statusCode, httpResponse, nil
}

func _deleteTenant(id string) (int, string) {
	return http.DeleteHttp(cisco.AciNdoLogin.Jwttoken, "https://198.18.133.100/mso/api/v1/tenants/"+id, "")
}

func GetAllTenants() (int, string, error) {
	var statusCode, httpResponse = _getAllTenants()

	if cisco.IsSessionExpired(statusCode) {
		statusCode, httpResponse = _getAllTenants()
	}

	return statusCode, httpResponse, nil
}

func _getAllTenants() (int, string) {
	return http.GetHttp(cisco.AciNdoLogin.Jwttoken, "https://198.18.133.100/mso/api/v1/tenants")
}
