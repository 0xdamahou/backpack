package authenticated

import (
	"bytes"
	"crypto/ed25519"
	"encoding/base64"
	"io"

	//"encoding/json"
	"fmt"
	json "github.com/bytedance/sonic"
	"log"
	"net/http"
	"sort"
	"strings"
	"time"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func (c *BackpackClient) SignRequest(instruction string, params string, window string) (string, string, error) {
	timestamp := fmt.Sprintf("%d", time.Now().UTC().UnixMilli())
	if window == "" {
		window = "5000"
	}
	var signStr strings.Builder
	signStr.WriteString("instruction=")
	signStr.WriteString(instruction)
	signStr.WriteString("&")
	if len(params) > 0 {
		signStr.WriteString(params)
		signStr.WriteString("&")
	}
	signStr.WriteString("timestamp=")
	signStr.WriteString(timestamp)
	signStr.WriteString("&window=")
	signStr.WriteString(window)
	//log.Printf("SignRequest: %s", signStr.String())
	signature := ed25519.Sign(c.PrivateKey, []byte(signStr.String()))
	return timestamp, base64.StdEncoding.EncodeToString(signature), nil
}

// MapToQueryString 将 map[string]string 转换为查询字符串，按键的字母顺序排列
func MapToQueryString(params map[string]interface{}) string {
	//fmt.Sprintf("--%+v", params)
	if len(params) == 0 {
		return ""
	}

	// 获取所有键并排序
	var keys []string
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var queryStr strings.Builder
	for i, k := range keys {
		if i > 0 {
			queryStr.WriteString("&")
		}
		switch v := params[k].(type) {
		case string:
			queryStr.WriteString(fmt.Sprintf("%s=%s", k, v))
		case bool:
			queryStr.WriteString(fmt.Sprintf("%s=%t", k, v)) // %t 用于 bool 类型
		case int:
			queryStr.WriteString(fmt.Sprintf("%s=%d", k, v)) // %d 用于 int 类型
		case int64:
			queryStr.WriteString(fmt.Sprintf("%s=%d", k, v)) // %d 用于 int 类型
		case float64:
			queryStr.WriteString(fmt.Sprintf("%s=%f", k, v)) // %f 用于 float64 类型
		default:
			// 对于其他类型，使用 fmt.Sprint 转换为字符串
			queryStr.WriteString(fmt.Sprintf("%s=%s", k, fmt.Sprint(v)))
		}
	}

	return queryStr.String()
}

func (c *BackpackClient) DoPost(endpoint string, instruction string, postBody []byte, result interface{}) error {
	return c.DoWithBody(POST, endpoint, instruction, postBody, result)
}

func (c *BackpackClient) DoDelete(endpoint string, instruction string, postBody []byte, result interface{}) error {
	return c.DoWithBody(DELETE, endpoint, instruction, postBody, result)
}

func (c *BackpackClient) DoPatch(endpoint string, instruction string, postBody []byte, result interface{}) error {
	return c.DoWithBody(PATCH, endpoint, instruction, postBody, result)
}

func (c *BackpackClient) DoWithBody(method string, endpoint string, instruction string, postBody []byte, result interface{}) error {
	var params map[string]interface{}
	if postBody != nil {
		err := json.Unmarshal(postBody, &params)
		if err != nil {
			return err
		}
	}
	//log.Printf("Body:%s", string(postBody))

	q := MapToQueryString(params)
	//log.Printf("params:%s", q)
	return c.DoRequest(method, endpoint, instruction, postBody, q, result)
}

func (c *BackpackClient) DoGet(endpoint string, instruction string, params string, result interface{}) error {
	return c.DoRequest(GET, endpoint, instruction, nil, params, result)
}

func (c *BackpackClient) DoRequest(method, endpoint string, instruction string, reqBody []byte, params string, result interface{}) error {
	url := baseURL + endpoint
	req, err := http.NewRequest(method, url, bytes.NewBuffer(reqBody))
	if err != nil {
		log.Println(err)
		return err
	}

	if len(params) > 0 {
		req.URL.RawQuery = params
	}
	//log.Println(req.URL.RawQuery)

	window := req.Header.Get("X-Window")
	if window == "" {
		window = "5000"
	}

	timestamp, signature, err := c.SignRequest(instruction, params, window)
	if err != nil {
		log.Println(err)
		return err
	}

	req.Header.Set("X-Api-Key", c.APIKey)
	req.Header.Set("X-Timestamp", timestamp)
	req.Header.Set("X-Window", window)
	req.Header.Set("X-Signature", signature)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.Client.GetClient().Do(req)
	if err != nil {
		log.Println(err)
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("request failed with status: %s\n", resp.Status)
	}
	//use decoder in production
	bs, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return err
	}
	//	log.Println(resp.StatusCode, string(bs))
	if result != nil {
		err = json.Unmarshal(bs, result)
	}

	//if result != nil {
	//	err = json.NewDecoder(resp.Body).Decode(result)
	//	if err != nil {
	//		log.Println(err)
	//		return err
	//	}
	//}

	return err
}
