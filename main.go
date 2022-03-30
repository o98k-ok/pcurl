package main

import (
	"fmt"
	hp "github.com/o98k-ok/pcurl/http"
	"net/http"
)

func main() {
	cli := &http.Client{}
	resp, err := hp.NewRequest(cli, "www.baidu.com").Do()
	defer resp.Body.Close()
	fmt.Println(err)
	fmt.Println(resp.String())
}
