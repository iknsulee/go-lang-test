package ksleeutility

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

func PrintPrettyStruct(title string, v any) {
	// 구조체를 예쁘게 JSON 형식으로 출력하기 위한 함수
	anyString, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		log.Fatalf(err.Error())
		return
	}
	fmt.Printf("%s[%s]\n", title, string(anyString))

}

func GetPrettyStringFromJSONString(str string) string {
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, []byte(str), "", "    "); err != nil {
		return ""
	}
	return prettyJSON.String()
}
