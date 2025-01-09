package aliU

import (
	"encoding/json"
)

func PrettyPrint[T any](r *T) string {
	response, err := json.MarshalIndent(*r, "", "  ")
	if err != nil {
		return "ERROR PRETTY PRINTING"
	}

	return string(response)
}
