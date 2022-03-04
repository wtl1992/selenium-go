package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"selenium-go/ljxwtl/lang"
	"selenium-go/ljxwtl/selenium/httpcores"
	"strings"
	"sync"
)

func main() {
	projectDir, _ := os.Getwd()

	var driverAbsPath = projectDir + "/binary/chromedriver"

	fmt.Println(driverAbsPath)

	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		process := exec.Command(driverAbsPath, "--port=8080")
		err := process.Start()
		if err != nil {
			return
		}

		var httpClient = http.Client{}

		var params = make(map[string]interface{})
		params["capabilities"] = map[string]interface{}{}

		paramBytes, _ := json.Marshal(params)

		request, _ := http.NewRequest(http.MethodPost, "http://localhost:8080/session", strings.NewReader(string(paramBytes)))

		var response *http.Response = nil
		for {
			response, _ = httpClient.Do(request)

			if response != nil {
				break
			}
		}

		log.Println(response)

		var httpEntity = new(httpcores.HttpEntity)
		httpEntity.R = response.Body

		bytes, err := httpEntity.HandleResponse()

		if err != nil {
			panic(err)
		}

		var contentMap = make(map[string]interface{})

		_ = json.Unmarshal(bytes, &contentMap)

		fmt.Println(contentMap)

	}()

	var obj = lang.Object{}

	code := obj.HashCode()

	fmt.Println(code)

	wg.Wait()
}
