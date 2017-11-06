package httpreq

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
)

const (
	charsetUTF8 = "charset=UTF-8"
)

type Header struct {
	ContentType string
	Token       string
}

const (
	MIMEApplicationJSON                  = "application/json"
	MIMEApplicationJSONCharsetUTF8       = "application/json" + "; " + charsetUTF8
	MIMEApplicationJavaScript            = "application/javascript"
	MIMEApplicationJavaScriptCharsetUTF8 = "application/javascript" + "; " + charsetUTF8
	MIMEApplicationXML                   = "application/xml"
	MIMEApplicationXMLCharsetUTF8        = "application/xml" + "; " + charsetUTF8
	MIMETextXML                          = "text/xml"
	MIMETextXMLCharsetUTF8               = "text/xml" + "; " + charsetUTF8
	MIMEApplicationForm                  = "application/x-www-form-urlencoded"
	MIMEApplicationFormUTF8              = "application/x-www-form-urlencoded" + "; " + charsetUTF8
	MIMEApplicationProtobuf              = "application/protobuf"
	MIMEApplicationMsgpack               = "application/msgpack"
	MIMETextHTML                         = "text/html"
	MIMETextHTMLCharsetUTF8              = "text/html" + "; " + charsetUTF8
	MIMETextPlain                        = "text/plain"
	MIMETextPlainCharsetUTF8             = "text/plain" + "; " + charsetUTF8
	MIMEMultipartForm                    = "multipart/form-data"
	MIMEOctetStream                      = "application/octet-stream"
)

func POST(token, url string, param interface{}, v interface{}) (resp *http.Response, err error) {
	b, err := json.Marshal(param)
	if err != nil {
		return
	}

	httpReq, err := http.NewRequest("POST", url, bytes.NewBuffer(b))
	if err != nil {
		err = fmt.Errorf("HTTP New Request Error: %s", err)
		return
	}
	httpReq.Header.Set("Content-Type", "application/json")
	if token != "" {
		httpReq.Header.Set("Authorization", "Bearer "+token)
	}

	resp, err = (&http.Client{}).Do(httpReq)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		b, _ := ioutil.ReadAll(resp.Body)
		err = fmt.Errorf("[%d %s]%s", resp.StatusCode, resp.Status, string(b))
		return
	}
	if v != nil {
		dec := json.NewDecoder(resp.Body)
		if err = dec.Decode(&v); err != nil {
			return
		}
	}

	return
}

func GET(token, url string, v interface{}) (resp *http.Response, err error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		err = fmt.Errorf("HTTP New Request Error: %s", err)
		return
	}
	if token != "" {
		req.Header.Set("Authorization", "Bearer "+token)

	}
	client := &http.Client{}
	resp, err = client.Do(req)
	if err != nil {
		err = fmt.Errorf("HTTP Request Error: %s", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		b, _ := ioutil.ReadAll(resp.Body)
		err = fmt.Errorf("[%d %s]%s", resp.StatusCode, resp.Status, string(b))
		return
	}

	if v != nil {
		dec := json.NewDecoder(resp.Body)
		if err = dec.Decode(&v); err != nil {
			return
		}
	}

	return
}

func NewPost(url string, param []byte, header *Header, transport *http.Transport) (resp *http.Response, body []byte, err error) {

	httpReq, err := http.NewRequest("POST", url, bytes.NewBuffer(param))
	if header != nil {
		httpReq.Header.Set("Content-Type", header.ContentType)
		if header.Token != "" {
			httpReq.Header.Set("Authorization", "Bearer "+header.Token)
		}
	} else {
		httpReq.Header.Set("Content-Type", MIMEApplicationFormUTF8)
	}
	var client *http.Client
	if transport != nil {
		client = &http.Client{Transport: transport}
	} else {
		client = &http.Client{}
	}
	resp, err = client.Do(httpReq)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		b, _ := ioutil.ReadAll(resp.Body)
		err = fmt.Errorf("[%d %s]%s", resp.StatusCode, resp.Status, string(b))
		return
	}
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	return
}

func NewGet(url string, header *Header, transport *http.Transport) (resp *http.Response, body []byte, err error) {

	httpReq, err := http.NewRequest("GET", url, nil)
	if header != nil {
		if header.ContentType != "" {
			httpReq.Header.Set("Content-Type", header.ContentType)
		}
		if header.Token != "" {
			httpReq.Header.Set("Authorization", "Bearer "+header.Token)
		}
	} else {
		httpReq.Header.Set("Content-Type", MIMEApplicationFormUTF8)
	}

	var client *http.Client
	if transport != nil {
		client = &http.Client{Transport: transport}
	} else {
		client = &http.Client{}
	}
	resp, err = client.Do(httpReq)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		b, _ := ioutil.ReadAll(resp.Body)
		err = fmt.Errorf("[%d %s]%s", resp.StatusCode, resp.Status, string(b))
		return
	}
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	return

}

type FileItem struct {
	Key      string //image_content
	FileName string //test.jpg
	Content  []byte //[]byte
}

func NewPostFile(url string, paramTexts map[string]interface{}, paramFile FileItem) (resp *http.Response, body []byte, err error) {
	// if paramFiles ==nil {
	// 	return NewPost(url,paramTexts,header,transport)
	// }

	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	for k, v := range paramTexts {
		bodyWriter.WriteField(k, v.(string))
	}
	fileWriter, err := bodyWriter.CreateFormFile(paramFile.Key, paramFile.FileName)
	if err != nil {
		return
	}
	fileWriter.Write(paramFile.Content)
	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	resp, err = http.Post(url, contentType, bodyBuf)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		b, _ := ioutil.ReadAll(resp.Body)
		err = fmt.Errorf("[%d %s]%s", resp.StatusCode, resp.Status, string(b))
		return
	}
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	return
}

func POSTForm(url string, param []byte, header *Header, transport *http.Transport, v interface{}) (resp *http.Response, err error) {

	httpReq, err := http.NewRequest("POST", url, bytes.NewBuffer(param))
	if header != nil {
		httpReq.Header.Set("Content-Type", header.ContentType)
		if header.Token != "" {
			httpReq.Header.Set("Authorization", "Bearer "+header.Token)
		}
	} else {
		httpReq.Header.Set("Content-Type", MIMEApplicationFormUTF8)
	}
	var client *http.Client
	if transport != nil {
		client = &http.Client{Transport: transport}
	} else {
		client = &http.Client{}
	}
	resp, err = client.Do(httpReq)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		b, _ := ioutil.ReadAll(resp.Body)
		err = fmt.Errorf("[%d %s]%s", resp.StatusCode, resp.Status, string(b))
		return
	}
	if v != nil {
		dec := json.NewDecoder(resp.Body)
		if err = dec.Decode(&v); err != nil {
			return
		}
	}

	return
}
