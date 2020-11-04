package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"io/ioutil"
)

type Coupon struct {
	Code string
}

type Coupons struct {
	Coupon []Coupon
}

func (c Coupons) Check(code string) string {
	for _, item := range c.Coupon {
		if code == item.Code {
			return "valid"
		}
	}
	return "invalid"
}

type Result struct {
	Status string
}

var coupons Coupons

func main() {
	coupon := Coupon{
		Code: "abc",
	}

	coupons.Coupon = append(coupons.Coupon, coupon)


	http.HandleFunc("/", home)
	http.ListenAndServe(":9092", nil)
}

func home(w http.ResponseWriter, r *http.Request) {

	
	serviceD := makeHttpCall("http://localhost:9094");

	result := Result{Status: "invalid"}

	if serviceD.Status == "Service D is out!"{
		result.Status =  "Service D is out!"
	}

	coupon := r.PostFormValue("coupon")
	valid := coupons.Check(coupon)

	result = Result{Status: valid}

	 
	jsonResult, err := json.Marshal(result)
	if err != nil {
		log.Fatal("Error converting json")
	}

	fmt.Fprintf(w, string(jsonResult))
	
}

func makeHttpCall(urlMicroservice string) Result {

	values := url.Values{}

	res, err := http.PostForm(urlMicroservice, values)
	if err != nil{
		result := Result{Status: "Service D is out!"}
		return result
	}

	defer res.Body.Close()


	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal("Error processing result")
	}

	result := Result{}

	json.Unmarshal(data, &result)

	return result
}