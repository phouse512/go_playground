package main

import "fmt"

type HttpResponse struct {
	status_code int
	headers     map[string]string
	body        string
}

func New(status_code int) (*HttpResponse, error) {
	r := new(HttpResponse)

	r.headers = make(map[string]string)
	r.status_code = status_code
	r.body = ""
	return r, nil
}

func (r HttpResponse) validResponse() bool {
	// a value receiver that reads the HttpResponse copy
	if r.status_code < 300 {
		return true
	}
	return false
}

func (r HttpResponse) add_header(key string, value string) {
	// a value receiver that modifies the original map
	r.headers[key] = value
}

func (r *HttpResponse) updateStatus(new_status int) {
	// correct use of a pointer receiver for updating an int field
	r.status_code = new_status
}

func (r HttpResponse) updateStatusFail(new_status int) {
	// this is a bad use of a value receiver, nothing will happen
	r.status_code = new_status
}

func main() {
	response, _ := New(230)

	response.updateStatusFail(300)
	fmt.Println(response.status_code) // 230, the original response object wasn't updated

	response.updateStatus(300)
	fmt.Println(response.status_code) // 300, correct use of a pointer receiver

	response.add_header("Content-Type", "text/javascript")
	fmt.Println(response.headers) // map[Content-Type:text/javascript
}
