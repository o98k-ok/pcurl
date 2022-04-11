package main

import (
	"fmt"
	"github.com/o98k-ok/pcurl/format"
	"os"
	"strings"
)

func main() {
	r := strings.NewReader(`{
"code": 200,
    "message": "操作成功",
    "timestamp": 1648883705955,
    "requestId": "8cef9adb6f",
    "serverIp": "172.168",
    "handleTime": 10,
    "data": [
        {
            "longitude": "113.935349",
            "name": "安盛行名车门口",
            "serverAllday": 0,
            "status": 3,
            "code": "B04232002",
            "id": 33191,
            "latitude": "22.526247",
            "idCommunity1": 37767,
            "serverPeopleInt": 1,
            "workerNumber": 4,
            "areaName": "南山区",
            "communityName": "粤桂社区",
            "distance": 0,
            "address": "安盛行名车门口（城中村抽样）",
            "phone": "",
            "serverPeople": "绿码普通人员",
            "streetName": "粤海街道",
            "updateTime": 1648883704000,
            "createTime": 1648047903000,
            "serverNight": 1,
            "serverTime": "09:00-21:00"
        }]}`)
	err := format.NewFormatter(os.Stdout).WithoutColor().Jsonify(r)
	if err != nil {
		fmt.Println(err)
	}
}
