package globals

import (
	"encoding/json"
	"log"

	"github.com/deadcheat/twister/types"
)

// ConvertInterfaceToMatch convert match data in session attribute
func ConvertInterfaceToMatch(i interface{}) *types.Match {
	d, err := json.Marshal(i)
	if err != nil {
		log.Println(err)
		return nil
	}
	var m types.Match
	err = json.Unmarshal(d, &m)
	if err != nil {
		log.Println(err)
		return nil
	}
	return &m
}
