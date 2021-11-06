package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	// "go_playground/helpers"
	// "go_playground/leetcode"
	"go_playground/educative"
	"go_playground/models"
	"go_playground/polygons"
	"go_playground/utils"
)

func main() {

	//
	c := &polygons.Circle{Rad: 3}
	polygons.Calculate(c)

	// hackerrank.HCRunFuncs()
	educative.BSRunTest()

	fmt.Println("--------consuming rest service------")
	//get request
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	//we read teh response body in line below
	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// //convert the boyd to string
	// sb := string(body)
	// fmt.Println(sb)

	var users []models.User

	if err := json.NewDecoder(resp.Body).Decode(&users); err != nil {
		fmt.Printf("Error occured %v", err)
	}
	fmt.Printf("Users count %d\n", len(users))

	//post request
	payload, _ := json.Marshal(map[string]string{
		"name":  "Tobdy",
		"email": "toby@example.com",
	})

	paylodBuff := bytes.NewBuffer(payload)

	res, err := http.Post("https://postman-echo.com/post", "application/json", paylodBuff)
	if err != nil {
		fmt.Println(err.Error())
	}

	defer res.Body.Close()

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err.Error())
	}

	strResp := string(resBody)

	fmt.Printf("Post response: %s\n", strResp)

	fmt.Println("---------------")
	fmt.Println("---read/write file demo")

	lines, isError := utils.ReadFile("Logs.txt")
	if isError {
		return
	}
	fmt.Println(len(lines))

	set := map[string]bool{}
	for _, line := range lines {
		words := strings.Split(line, " ")
		if len(words) < 5 {
			continue
		}
		httpMethod, fileName, statusCode := words[5], words[6], words[8]
		if httpMethod == "\"GET" && statusCode == "200" && (strings.Contains(fileName, ".gif") || strings.Contains(fileName, ".GIF")) {
			// fmt.Println(fileName)
			fn := strings.Split(fileName, "/")
			last := fn[len(fn)-1]
			// fmt.Print(last[0:len(last)-4] + "\n")
			if !set[last] {
				set[last] = true
			}
		}
	}

	errorInWrite := utils.WriteFile("UniqueGif.txt", set)
	if errorInWrite {
		return
	}
	fmt.Println("---------------------------------------------------------")

	// fmt.Println(len(words))
	// for i, w := range words {
	// 	fmt.Printf("index %d | word %s\n", i, w)
	// }

	// arr := rand.Perm(5)
	// fmt.Printf("Original arr: %v\n", arr)
	// qs := helpers.QuickSort(arr, 0, len(arr)-1)
	// fmt.Printf("Quick sorted arr: %v\n", qs)
	// ms := helpers.MergeSort(arr)
	// fmt.Printf("merge sorted arr: %v\n", ms)
	// inssort := helpers.InsertionSort(arr)
	// fmt.Printf("merge sorted arr: %v\n", inssort)
	// leetcode.RunTest1()
	// helpers.RunTest(20)
}
