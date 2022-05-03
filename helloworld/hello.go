package main

import (
	"net/http"
	_ "net/http/pprof"
)

func main() {
	//request, err := http.NewRequest(http.MethodGet, "http://www.zgtfintech.com", nil)
	//if err != nil {
	//	panic(err)
	//}
	//request.Header.Add("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 13_2_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.3 Mobile/15E148 Safari/604.1")
	//client := http.Client{
	//	CheckRedirect: func(req *http.Request, via []*http.Request) error {
	//		fmt.Println("redirect: ", req)
	//		return nil
	//	},
	//}
	//response, err := client.Do(request)
	//defer response.Body.Close()
	//if s, err := httputil.DumpResponse(response, true); err == nil {
	//	fmt.Println(string(s))
	//}
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("hello"))
	})

	http.ListenAndServe(":8888", nil)

}
