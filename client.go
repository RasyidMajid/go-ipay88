package go_ipay88

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"time"
)

type Client struct {
	Env          Env
	MerchantCode string
	MerchantKey  string
	Log          *log.Logger
}

var defHTTPTimeout = 80 * time.Second
var httpClient = &http.Client{Timeout: defHTTPTimeout}

func (c *Client) DoProses(method, url string, body interface{}, response interface{}) error {
	bytesJson, err := json.Marshal(body)
	if nil != err {
		return err
	}
	fmt.Println(string(bytesJson))
	req, err := http.NewRequest(method, url, bytes.NewBuffer(bytesJson))
	if nil != err {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	fmt.Println("url = ", req.URL)
	fmt.Println("header = ", req.Header)
	fmt.Println("body = ", req.Body)

	res, err := httpClient.Do(req)
	if nil != err {
		return err
	}

	defer func() {
		_ = res.Body.Close()
	}()

	if http.StatusOK != res.StatusCode {
		errMessage := fmt.Sprintf("Cannot send request and got status code %d from Ipay88 V2", res.StatusCode)
		return errors.New(errMessage)
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if res.Header.Get("Content-Type") == "text/html" {
		reflect.ValueOf(response).Elem().SetString(string(resBody))
	} else {
		err = json.Unmarshal(resBody, response)
		if err != nil {
			return err
		}
	}

	return nil
}
