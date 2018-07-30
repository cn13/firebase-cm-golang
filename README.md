# firebase-cm-golang
FireBase Cloude Message


# Установка

```go get github.com/cn13/firebase-cm-golang```


# Использование

```go
package main

import (
	"fmt"
	"encoding/json"
	"github.com/cn13/firebase-cm-golang"
)

type Data struct {
	Name  string
	Value string
}

func main() {

	firebase_cm_golang.SetUrl("https://you-link.firebaseio.com/");
	request := Data{Name: "test1", Value: "1233"}
	firebase_cm_golang.Send("my_test/first", request)

	var response = firebase_cm_golang.Get("my_test/first")
	data := Data{}
	json.Unmarshal(response.Body.Bytes, &data)

	fmt.Println(data.Name)
	fmt.Println(data.Value)
}
 ```
