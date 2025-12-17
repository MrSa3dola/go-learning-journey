package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://www.google.com:443/search?q=youtube&newwindow=1&sca_esv=07b4a2215fc3982c&sxsrf=AE3TifMrOCnig6L44sXOutPLE_E8eJkIwg%3A1766001696750&ei=IAxDadbALfnOxc8Pxa7z-As&gs_ssp=eJzj4tTP1TewzEouKzZg9GKvzC8tKU1KBQA_-AaN&oq=you&gs_lp=Egxnd3Mtd2l6LXNlcnAiA3lvdSoCCAAyEBAuGEMYgwEYsQMYgAQYigUyERAAGIAEGIoFGJECGLEDGIMBMhAQABiABBiKBRhDGLEDGIMBMgoQABiABBiKBRhDMhAQABiABBiKBRhDGLEDGIMBMhAQABiABBiKBRhDGLEDGIMBMggQABiABBixAzILEAAYgAQYsQMYgwEyDhAuGIAEGIoFGLEDGIMBMggQABiABBixA0id-w9Q6_APWP7xD3ADeAGQAQCYAZkBoAH0AqoBAzAuM7gBA8gBAPgBAZgCBqACjAPCAgoQABhHGNYEGLADwgIKECMYgAQYigUYJ8ICChAjGPAFGMkCGCfCAhEQLhiDARiRAhixAxiABBiKBcICChAjGMkCGPAFGCeYAwCIBgGQBgiSBwMzLjOgB4MksgcDMC4zuAeCA8IHBTAuMi40yAcSgAgB&sclient=gws-wiz-serp"
 
func main() {
	fmt.Println("urls")
	fmt.Println(myurl)

	// prasing URL (extract information)
	result, err := url.Parse(myurl)
	checkNil(err)
	// fmt.Println(result.Scheme) // https
	// fmt.Println(result.Host) // google.com:443
	// fmt.Println(result.Path)  // for example (/search) or (/dev)
	// fmt.Println(result.Port()) // 443
	// fmt.Println(result.RawQuery) // query example (coruse=ai&id=1)

	qparams := result.Query()
	fmt.Printf("type of query params are: %T\n", qparams)
	
	fmt.Println(qparams["q"])


	for k, val := range qparams {
		fmt.Printf("key is {%v} value is {%v}\n", k, val)
	}

	partsOfUrl := &url.URL{
		Scheme: "https",
		Host: "google.com",
		Path: "/search",
		RawPath: "q=youtube",
	}
	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl) // https://google.com/search
}


func checkNil(err error){
	if err != nil{
		panic(err)
	}
}