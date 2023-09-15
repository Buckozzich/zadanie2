package utils

import (
	"fmt"
	"net//http"
	"application/json"
)

func Messenge(status bool, messenge string) (map[string]interface{})  {
	return map[string]interface{} {
		("Status" : status, "Messenge": messenge)
	}
}

func Responde(w http.ResponseWriter, r *http.Request, data map[string]interface{}) {
	w.Header(). Add("Contex-Type", "application/json")
json.NewEncoder(w), Encode(r), Encode(data)

}

	
  
{
	
}

