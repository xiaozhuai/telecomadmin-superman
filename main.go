package main

import (
	"flag"
	"fmt"
	"os"
	"telecomadmin-superman/exp"
	"time"
)

var (
	host string
)

func main() {
	flag.StringVar(&host, "host", "192.168.1.1", "set modem host")
	flag.Parse()
	fmt.Printf("- Modem host: %s\n", host)
	suc := false
	var result *exp.Result = nil
	var err error = nil

	allExp := []*exp.Exp{
		exp.ALCL_E_140W_P(),
		// TODO Add more exp...
	}
	for _, e := range allExp {
		fmt.Printf("  - [%s] Trying...\n", e.Name)
		result, err = e.Handle(host)
		if err != nil {
			fmt.Printf("  - [%s] Failed! %v\n", e.Name, err)
		} else {
			fmt.Printf("  - [%s] Succeed!\n", e.Name)
			suc = true
			break
		}
	}
	if suc {
		fmt.Printf("- Done!\n")
		fmt.Printf("  - Username: %s\n", result.Username)
		fmt.Printf("  - Password: %s\n", result.Password)
		now := time.Now()
		year, mon, day := now.Date()
		hour, min, sec := now.Clock()
		zone, _ := now.Zone()
		configFile := fmt.Sprintf("modem_%04d%02d%02d%02d%02d%02d%s.xml", year, mon, day, hour, min, sec, zone)
		err = os.WriteFile(configFile, result.Config, 0644)
		if err != nil {
			fmt.Printf("  - Error write modem config file\n")
		} else {
			fmt.Printf("  - Modem Config: %s\n", configFile)
		}
	} else {
		fmt.Printf("- Failed, no suitable exp found\n")
	}

	fmt.Printf("- Press Enter to continue...\n")
	_, _ = fmt.Scanln()
}
