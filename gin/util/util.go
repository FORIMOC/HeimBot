package util

import (
	"github.com/tidwall/gjson"
	url2 "net/url"
	"os/exec"
)

func Send(sender string, group string, content string) {
	var port string
	switch sender {
	default:
	case "Kalina":
		port = "5700"
		break
	case "Hoshino":
		port = "5701"
		break
	case "Komari":
		port = "5702"
		break
	}
	url := "http://127.0.0.1:" + port + "/send_group_msg"
	url = url + "?group_id=" + group + "&message=" + url2.QueryEscape(content)
	exec.Command("curl", url).Run()
}

func GetGroupInfo(groupID string) (string, int, int) {
	url := "http://127.0.0.1:5701/get_group_info?group_id=" + groupID
	output, _ := exec.Command("curl", url).CombinedOutput()
	groupName := gjson.Get(string(output), "data.group_name").String()
	member := gjson.Get(string(output), "data.member_count").Int()
	maxMember := gjson.Get(string(output), "data.max_member_count").Int()
	return groupName, int(member), int(maxMember)
}
