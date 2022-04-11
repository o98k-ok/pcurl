package main

import (
	"fmt"
	"github.com/o98k-ok/pcurl/format"
	hp "github.com/o98k-ok/pcurl/http"
	"jaytaylor.com/html2text"
	"net/http"
	"os"
)

func main() {
	cli := &http.Client{}

	url := "https://szwj.borycloud.com/ilhapi/wjw/checkpoint/list"
	resp, err := hp.NewRequest(cli, url).WithMethod(http.MethodPost).
		AddParam("areaId", "136").
		AddParam("streetId", "867").
		AddParam("page", "1").Do()
	if err != nil {
		fmt.Println(err)
		return
	}

	if err := format.NewFormatter(os.Stdout).Jsonify(resp.Body); err != nil {
		fmt.Println(err)
		return
	}
}

func help() {
	cli := &http.Client{}

	url := "https://articles.zsxq.com/id_f1ui373tzrzk.html"
	if err != nil {
		fmt.Println(err)
		return
	}

	// fmt.Println(resp.String())
	text, err := html2text.FromString(resp.String(), html2text.Options{PrettyTables: true})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(text)
}
