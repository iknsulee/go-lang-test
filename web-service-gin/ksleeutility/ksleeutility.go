package ksleeutility

import (
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
