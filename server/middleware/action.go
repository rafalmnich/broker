package middleware

import (
	"io/ioutil"
	"net/http"

	jsoniter "github.com/json-iterator/go"
	"github.com/rafalmnich/broker"
)

// SendAction deserializes the message from request and publishes it.
func SendAction(p broker.Publisher) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Body == nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		a, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		var actions broker.Actions
		err = jsoniter.Unmarshal(a, &actions)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = p.Publish(actions)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
