package main

import (
	"flag"
	"fmt"
	"os"
	"telecomadmin-superman/exp"
	"time"
)

func main() {
	var (
		host         string
		username     string
		password     string
		autoContinue bool
	)
	flag.StringVar(&host, "host", "192.168.1.1", "光猫地址")
	flag.StringVar(&username, "username", "", "普通管理员账号")
	flag.StringVar(&password, "password", "", "普通管理员密码")
	flag.BoolVar(&autoContinue, "auto-continue", false, "完成后自动退出进程")
	flag.Parse()
	context := exp.NewContext(host, username, password)

	fmt.Printf("- 光猫地址: %s\n", context.GetHost())
	suc := false
	var result *exp.Result = nil
	var err error = nil

	allExp := []*exp.Exp{
		exp.ALCL_E_140W_P(),
		// TODO Add more exp...
	}
	for _, e := range allExp {
		fmt.Printf("  - [%s] 正在尝试...\n", e.Name)
		result, err = e.Handle(context)
		if err != nil {
			fmt.Printf("  - [%s] 失败! %v\n", e.Name, err)
		} else {
			fmt.Printf("  - [%s] 成功!\n", e.Name)
			suc = true
			break
		}
	}
	if suc {
		fmt.Printf("- 完成!\n")
		fmt.Printf("  - 超级管理员账号: %s\n", result.Username)
		fmt.Printf("  - 超级管理员密码: %s\n", result.Password)
		now := time.Now()
		year, mon, day := now.Date()
		hour, min, sec := now.Clock()
		zone, _ := now.Zone()
		configFile := fmt.Sprintf("modem_%04d%02d%02d%02d%02d%02d%s.xml", year, mon, day, hour, min, sec, zone)
		err = os.WriteFile(configFile, result.Config, 0644)
		if err != nil {
			fmt.Printf("  - 写入光猫配置失败! %v\n", err)
		} else {
			fmt.Printf("  - 光猫配置文件: %s\n", configFile)
		}
	} else {
		fmt.Printf("- 失败, 没有找到适用于你的光猫设备的方案\n")
	}

	if !autoContinue {
		fmt.Printf("- 按回车键退出...\n")
		_, _ = fmt.Scanln()
	}
}
