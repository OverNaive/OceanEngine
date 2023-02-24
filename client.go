package oceanengine

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/pkg/errors"
)

type Client struct {
	accessToken string
}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) SetAccessToken(accessToken string) {
	c.accessToken = accessToken
}

func (c *Client) GetList(ctx context.Context, gw string, req map[string]interface{}) (*ListResponse, error) {
	// build query
	query := EncodeQuery(req)
	apiUrl := c.getApiUrl(gw, query)

	// build httpReq
	httpReq, err := http.NewRequestWithContext(ctx, "GET", apiUrl, nil)
	if err != nil {
		return nil, err
	}
	if c.accessToken != "" {
		httpReq.Header.Add("Access-Token", c.accessToken)
	}

	// build httpResp
	httpResp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpResp.Body.Close()

	// build resp
	resp := new(ListResponse)
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(resp); err != nil {
		return nil, err
	}
	if resp.IsError() {
		return resp, errors.New(resp.ErrorMessage())
	}
	return resp, nil
}

func (c *Client) Get(ctx context.Context, gw string, req map[string]interface{}) (*DataResponse, error) {
	// build query
	query := EncodeQuery(req)
	apiUrl := c.getApiUrl(gw, query)

	// build httpReq
	httpReq, err := http.NewRequestWithContext(ctx, "GET", apiUrl, nil)
	if err != nil {
		return nil, err
	}
	if c.accessToken != "" {
		httpReq.Header.Add("Access-Token", c.accessToken)
	}

	// build httpResp
	httpResp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpResp.Body.Close()

	// build resp
	resp := new(DataResponse)
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(resp); err != nil {
		return nil, err
	}
	if resp.IsError() {
		return resp, errors.New(resp.ErrorMessage())
	}
	return resp, nil
}

func (c *Client) Post(ctx context.Context, gw string, req map[string]interface{}) (*DataResponse, error) {
	// build body
	body := EncodeBody(req)
	apiUrl := c.getApiUrl(gw, "")

	// build httpReq
	httpReq, err := http.NewRequestWithContext(ctx, "POST", apiUrl, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	httpReq.Header.Add("Content-Type", "application/json")
	if c.accessToken != "" {
		httpReq.Header.Add("Access-Token", c.accessToken)
	}

	// build httpResp
	httpResp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpResp.Body.Close()

	// build resp
	resp := new(DataResponse)
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(resp); err != nil {
		return nil, err
	}
	if resp.IsError() {
		return resp, errors.New(resp.ErrorMessage())
	}
	return resp, nil
}

func (c *Client) getApiUrl(gw string, query string) (apiUrl string) {
	apiUrl = gw
	if !strings.Contains(gw, "https") {
		apiUrl = fmt.Sprintf("%s%s", BaseUrl, gw)
	}
	if query != "" {
		apiUrl = fmt.Sprintf("%s?%s", apiUrl, query)
	}
	return
}
