package main 

import "net/http"
import "fmt"
import "bytes"
import "io/ioutil"

func main() {
	url := "http://127.0.0.1:8000/users/new"
    fmt.Println("URL:>", url)

    var jsonStr = []byte(`{ "username": "eduardo78d", "first_name": "Eduardo Ismael", "last_name": "Garcia"}`)

    req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
    req.Header.Set("X-Custom-Header", "myvalue")
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    fmt.Println("response Status:", resp.Status)
    fmt.Println("response Headers:", resp.Header)
    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("response Body:", string(body))
}