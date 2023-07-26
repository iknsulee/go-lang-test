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

type CampSchemaV struct {
	Schemas []struct {
		ID          string `json:"id"`
		DisplayName string `json:"displayName"`
		Templates   []struct {
			Name            string `json:"name"`
			DisplayName     string `json:"displayName"`
			TenantID        string `json:"tenantId"`
			TemplateID      string `json:"templateID"`
			Anps            []any  `json:"anps"`
			Bds             []any  `json:"bds"`
			Contracts       []any  `json:"contracts"`
			ExternalEpgs    []any  `json:"externalEpgs"`
			Filters         []any  `json:"filters"`
			IntersiteL3Outs []any  `json:"intersiteL3outs"`
			Networks        []any  `json:"networks"`
			ServiceGraphs   []any  `json:"serviceGraphs"`
			Vrfs            []any  `json:"vrfs"`
			TemplateType    string `json:"templateType"`
			Description     string `json:"description"`
			Version         int    `json:"version"`
		} `json:"templates"`
		Sites []struct {
			SiteID          string `json:"siteId"`
			TemplateName    string `json:"templateName"`
			TemplateID      string `json:"templateID"`
			Anps            []any  `json:"anps"`
			Vrfs            []any  `json:"vrfs"`
			Bds             []any  `json:"bds"`
			Contracts       []any  `json:"contracts"`
			ExternalEpgs    []any  `json:"externalEpgs"`
			ServiceGraphs   []any  `json:"serviceGraphs"`
			IntersiteL3Outs []any  `json:"intersiteL3outs"`
			Networks        []any  `json:"networks"`
		} `json:"sites"`
		Summary struct {
			NumOfPolicies int `json:"numOfPolicies"`
			Templates     []struct {
				Anps struct {
					Count int `json:"count"`
				} `json:"anps"`
				Epgs struct {
					Count int `json:"count"`
				} `json:"epgs"`
				Vrfs struct {
					Count int `json:"count"`
				} `json:"vrfs"`
				Bds struct {
					Count int `json:"count"`
				} `json:"bds"`
				Contracts struct {
					Count int `json:"count"`
				} `json:"contracts"`
				Filters struct {
					Count int `json:"count"`
				} `json:"filters"`
				ExternalEpgs struct {
					Count int `json:"count"`
				} `json:"externalEpgs"`
				ServiceGraphs struct {
					Count int `json:"count"`
				} `json:"serviceGraphs"`
				IntersiteL3Outs struct {
					Count int `json:"count"`
				} `json:"intersiteL3outs"`
				Networks struct {
					Count int `json:"count"`
				} `json:"networks"`
			} `json:"templates"`
		} `json:"summary"`
		Description string `json:"description"`
	} `json:"schemas"`
}

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

func GetAllSchemas() (int, string, error) {
	var statusCode, httpResponse = _getAllSchemas()

	if cisco.IsSessionExpired(statusCode) {
		statusCode, httpResponse = _getAllSchemas()
	}

	return statusCode, httpResponse, nil
}

func _getAllSchemas() (int, string) {
	return http.GetHttp(cisco.AciNdoLogin.Jwttoken, "https://198.18.133.100/mso/api/v1/schemas/list-identity")
}
