- Get json content and print it out from url

```
func get_content_json(){
	url_check := "https://127.0.0.1:8444/upstreams/test/health"

	res, err := http.Get(url_check)

	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}

	var data Data_Json
	json.Unmarshal(body, &data)
	fmt.Printf("Result: %v\n", data)
	os.Exit(0)

}


func main() {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	get_content_json()
}


```

Require data struct,you can get it from jsongen in here : https://github.com/bemasher/JSONGen
