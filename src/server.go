package src

import (
	"fmt"
	"net/http"

	"github.com/containrrr/shoutrrr"
)

func Send(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		req.ParseForm()
		serviceName := req.Form.Get("service")
		message := req.Form.Get("message")
		serviceUrl, serviceUrlOk := GetServiceUrl(serviceName)
		if !serviceUrlOk {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		err := shoutrrr.Send(serviceUrl, message)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			fmt.Printf("Internal server error while request: %v", err)
		} else {
			fmt.Printf("Message %q sent to service %q.", message, serviceName)
			w.WriteHeader(http.StatusNoContent)
		}
	} else {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}
