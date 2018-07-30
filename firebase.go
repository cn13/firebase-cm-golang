package firebase_cm_golang

/**
	type Data struct {
		Name   string
		Value  string
	}

	firebase-cm-golang.SetUrl("https://you-link.firebaseio.com/");

	Example Send:
		request := Data{Name: "test1", Value: "1233"}
		firebase-cm-golang.Send("my_class/first", request)

	Example Get:
		var response = firebase-cm-golang.Get("my_class/first")
		data := Data{}
		json.Unmarshal(response.Body.Bytes, &data)
		fmt.Println(data.Name)
 */

import (
	"io"
	"io/ioutil"
	"net/http"
	"fmt"
	"encoding/json"
	"strings"
)

var urlFireBase string;

type Request struct {
	BaseUrl     string
	Endpoint    string
	QueryString string
	Headers     map[string][]string
}

type Response struct {
	RawBody  []byte
	Body     Body
	Headers  map[string][]string
	Status   int
	Protocol string
}

type Body struct {
	Bytes  []byte
	String string
}

func SetUrl(url string) {
	urlFireBase = url
}

func Send(path string, message interface{}) {
	reader := strings.NewReader(prepareMessageSend(message))
	request, err := http.NewRequest("PATCH", urlFireBase+path+".json", reader)
	client := &http.Client{}
	resp, err := client.Do(request)
	if (err != nil) {
		fmt.Println(err)
	} else {
		fmt.Println(resp)
	}
}

func Get(path string) Response {
	// Build the Request struct
	req := Request{
		BaseUrl:     urlFireBase,
		Endpoint:    path + ".json",
		QueryString: "?print=pretty",
		Headers:     nil,
	}

	// Send the request
	res, err := getUrl(req)

	if err != nil {
		fmt.Println(err)
	}

	return res
}

func getUrl(req Request) (Response, error) {

	// Build the URL
	var url string
	url += req.BaseUrl
	url += req.Endpoint
	url += req.QueryString

	// Create an HTTP client
	c := &http.Client{}

	// Create an HTTP request
	r, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Response{}, err
	}

	// Add any defined headers
	if req.Headers != nil {
		r.Header = http.Header(req.Headers)
	}

	// Add User-Agent if none is given
	if r.Header["User-Agent"] == nil {
		r.Header.Set("User-Agent", "Golang easyget")
	}

	// Send the request
	res, err := c.Do(r)

	// Check for error
	if err != nil {
		return Response{}, err
	}

	// Make sure to close after reading
	defer res.Body.Close()

	// Limit response body to 1mb
	lr := &io.LimitedReader{res.Body, 1000000}

	// Read all the response body
	rb, err := ioutil.ReadAll(lr)

	// Check for error
	if err != nil {
		return Response{}, err
	}

	// Build the output
	responseOutput := Response{
		Body: Body{
			Bytes:  rb,
			String: string(rb),
		},
		Headers:  res.Header,
		Status:   res.StatusCode,
		Protocol: res.Proto,
	}

	// Send it along
	return responseOutput, nil

}

func prepareMessageSend(message interface{}) string {
	result, _ := json.Marshal(message)
	return string(result)
}
