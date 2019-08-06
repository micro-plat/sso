package sso

import (
	"fmt"
)

var client *SSOClient

func init() {
	clientT, err := New("http://192.168.106.226:6689", "sso", "B128F779D5741E701923346F7FA9F95C")
	if err != nil {
		fmt.Println(err)
	}
	client = clientT
}

func TestMenu() {
	data, err := client.GetUserMenu(10123)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data)
}

func TestSystemInfo() {
	data, err := client.GetSystemInfo()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data)

}

func TestGetUserInfoByName() {
	data, err := client.GetUserInfoByName("programer")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data)
}
