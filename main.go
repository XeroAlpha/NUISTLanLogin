package main

import (
	"fmt"
	"os"
)

const defaultHost = "http://10.255.255.46/"

func main() {
	args := os.Args[1:]
	host := os.Getenv("NUISTLAN_HOST")
	if host == "" {
		host = defaultHost
	}
	if len(args) >= 3 {
		switch args[0] {
		case "list":
			listChannel(host, args[1], args[2])
			return
		case "login":
			if len(args) >= 4 {
				login(host, args[1], args[2], args[3])
				return
			}
			break
		case "status":
			onlineInfo(host, args[1], args[2])
			return
		case "logout":
			logout(host, args[1], args[2])
			return
		}
	}
	if len(args) > 0 {
		fmt.Println("Require more arguments")
		fmt.Println("")
	}
	fmt.Println("NUIST LAN Login Helper V1.0")
	fmt.Println("https://github.com/XeroAlpha/NUISTLanLogin")
	fmt.Println("")
	fmt.Println("NUISTLanLogin.exe {list|login|status|logout} <user> <pass> [<channel name>]")
	fmt.Println("")
	fmt.Println("list   - List all available channels")
	fmt.Println("login  - Login with specified channel, should provide <channel name>")
	fmt.Println("status - Print online status")
	fmt.Println("logout - Logout")
	if len(args) > 0 {
		os.Exit(1)
	}
}

func listChannel(host string, userName string, password string) {
	ip := GetIP(host)
	res := Login(host, LoginRequest{
		UserName:  userName,
		Password:  password,
		AutoLogin: 1,
		ChannelId: ChannelIdList,
		Action:    ActionFirstAuth,
		IPAddress: ip,
	})
	fmt.Println("Channels:")
	for _, channel := range res.Data.Channels {
		fmt.Printf("%s - %s\n", channel.Id, channel.Name)
	}
}

func checkLoginRes(res LoginResponse) bool {
	if res.Data.PromptText != "" {
		fmt.Println(res.Data.PromptText)
	}
	if res.Data.PromptURL != "" {
		fmt.Println(res.Data.PromptURL)
	}
	return res.Code == 200
}

func printOnlineInfo(res LoginResponse) {
	sec := res.Data.OnlineDuration
	hr := sec / 3600
	min := sec / 60 % 60
	sec = sec % 60
	fmt.Printf("User name: %s\n", res.Data.UserName)
	fmt.Printf("Balance:   %g\n", res.Data.Balance)
	fmt.Printf("Online:    %dh%02dm%02ds\n", hr, min, sec)
	fmt.Printf("Port:      %s\n", res.Data.CurrentPort)
	fmt.Printf("IP:        %s\n", res.Data.IPAddress)
}

func login(host string, userName string, password string, channelName string) {
	ip := GetIP(host)
	firstRes := Login(host, LoginRequest{
		UserName:  userName,
		Password:  password,
		AutoLogin: 1,
		ChannelId: ChannelIdList,
		Action:    ActionFirstAuth,
		IPAddress: ip,
	})
	if !checkLoginRes(firstRes) {
		fmt.Println("Login failed")
		return
	}
	var channelId string = ""
	for _, channel := range firstRes.Data.Channels {
		if channel.Name == channelName || channel.Id == channelName {
			channelId = channel.Id
			break
		}
	}
	if channelId == "" {
		fmt.Printf("Cannot find channel with name: %s", channelName)
		return
	}
	secondRes := Login(host, LoginRequest{
		UserName:  userName,
		Password:  password,
		AutoLogin: 1,
		ChannelId: channelId,
		Action:    ActionSecondAuth,
		IPAddress: ip,
	})
	if !checkLoginRes(secondRes) {
		fmt.Println("Login failed")
		return
	}
	fmt.Println("Login successfully!")
	printOnlineInfo(secondRes)
}

func onlineInfo(host string, userName string, password string) {
	ip := GetIP(host)
	res := Login(host, LoginRequest{
		UserName:  userName,
		Password:  password,
		AutoLogin: 1,
		ChannelId: ChannelIdOnlineInfo,
		Action:    ActionThirdAuth,
		IPAddress: ip,
	})
	if !checkLoginRes(res) {
		fmt.Println("Fetch failed")
		return
	}
	printOnlineInfo(res)
}

func logout(host string, userName string, password string) {
	ip := GetIP(host)
	res := Login(host, LoginRequest{
		UserName:  userName,
		Password:  password,
		AutoLogin: 1,
		ChannelId: ChannelIdLogout,
		Action:    ActionThirdAuth,
		IPAddress: ip,
	})
	if !checkLoginRes(res) {
		fmt.Println("Logout failed")
		return
	}
	fmt.Println("Logout successfully!")
}
