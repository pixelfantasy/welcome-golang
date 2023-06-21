package helper

import "fmt"
import "encoding/json"

func OutputRawJSON(contents []byte) {
	var result map[string]any
	jsonError := json.Unmarshal(contents, &result)
	if jsonError == nil {
		fmt.Println("== OutputRawJSON >> ==========================================")
		fmt.Println(result)
		fmt.Println("== OutputRawJSON << ==========================================")
	}
}
