package http

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHttp(t *testing.T) {
	//if rsp, err := http.Get("https://www.imooc.com"); err != nil {
	//	defer rsp.Body.Close()
	//	t.Log(rsp)
	//	if s, err := httputil.DumpResponse(rsp, true); err != nil {
	//		t.Log(string(s))
	//	}
	//} else {
	//	t.Log(err)
	//}
	resp, err := http.Get("https://www.imooc.com")
	if err != nil {
		fmt.Println(resp)
	} else {
		panic(err)
	}
}
