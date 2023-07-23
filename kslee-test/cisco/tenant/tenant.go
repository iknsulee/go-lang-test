package tenant

import (
	"ksleemodule/cisco"
	"ksleemodule/kslee/http"
)

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
