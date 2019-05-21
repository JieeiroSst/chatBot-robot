package function

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func handle(w http.ResponseWriter, r *http.Request) {
	var input []byte

	if r.Body != nil {
		defer r.Body.Close()

		body, _ := ioutil.ReadAll(r.Body)

		input = body
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("hello world ,input was:%s", string(input))))
}
