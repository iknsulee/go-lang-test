package schema

import (
	"encoding/json"
	"ksleemodule/cisco"
	"ksleemodule/kslee/http"
)

type Template struct {
	Name         string `json:"name"`
	DisplayName  string `json:"displayName"`
	TenantID     string `json:"tenantId"`
	TemplateType string `json:"templateType"`
	Anps         []Anp  `json:"anps"`
	Vrfs         []Vrf  `json:"vrfs"`
}

type Vrf struct {
	Name                  string                  `json:"name"`
	DisplayName           string                  `json:"displayName"`
	AutoRouteTargetImport []AutoRouteTargetImport `json:"autoRouteTargetImport"`
	AutoRouteTargetExport []AutoRouteTargetExport `json:"autoRouteTargetExport"`
}

type AutoRouteTargetImport struct {
	Aci string `json:"aci"`
}

type AutoRouteTargetExport struct {
	Aci string `json:"aci"`
}

type Anp struct {
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
}
type Site struct {
	SiteID       string `json:"siteId"`
	TemplateName string `json:"templateName"`
}

type Schema struct {
	DisplayName string     `json:"displayName"`
	Templates   []Template `json:"templates"`
	Sites       []Site     `json:"sites"`
}

//type Schema struct {
//	DisplayName string `json:"displayName"`
//	Templates   []struct {
//		Name         string `json:"name"`
//		DisplayName  string `json:"displayName"`
//		TenantID     string `json:"tenantId"`
//		TemplateType string `json:"templateType"`
//		Anps         []struct {
//			Name        string `json:"name"`
//			DisplayName string `json:"displayName"`
//		} `json:"anps"`
//		Vrfs []struct {
//			Name                  string `json:"name"`
//			DisplayName           string `json:"displayName"`
//			AutoRouteTargetImport []struct {
//				Aci string `json:"aci"`
//			} `json:"autoRouteTargetImport"`
//			AutoRouteTargetExport []struct {
//				Aci string `json:"aci"`
//			} `json:"autoRouteTargetExport"`
//		} `json:"vrfs"`
//	} `json:"templates"`
//	Sites []struct {
//		SiteID       string `json:"siteId"`
//		TemplateName string `json:"templateName"`
//	} `json:"sites"`
//}

func CreateSchema(schema Schema) (int, string, error) {
	var statusCode, httpResponse = _createSchema(schema)

	if cisco.IsSessionExpired(statusCode) {
		statusCode, httpResponse = _createSchema(schema)
	}

	return statusCode, httpResponse, nil
}

func _createSchema(schema Schema) (int, string) {

	schemaByte, err := json.Marshal(schema)
	if err != nil {
		panic(err)
	}

	return http.PostHttp(cisco.AciNdoLogin.Jwttoken, "https://198.18.133.100/mso/api/v1/schemas", string(schemaByte))

}
