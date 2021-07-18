package rippling

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/workspace/rippling-data-fetch/logger"
)

type Response map[string]string;

type Client struct {
	baseURL *url.URL
	httpClient *http.Client
	token string
}

func (c *Client) MakeRequest(method, requestURL string, body io.Reader) (*http.Response, error) {

	req, err := http.NewRequest("POST", c.baseURL.String() + requestURL, body)
	if err != nil {
		logger.Log.Error("Could not make new request" )
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", "CrowdCowRipClient-1.0")
	req.Header.Add("authorization", "Bearer "+ c.token)
	return c.httpClient.Do(req)
}

func NewProdClient(token string) Client {
	return NewRipClient("https://api.rippling.com/platform/api", token)
}

func NewSandBoxClient(token string) Client {
	return NewRipClient("https://sandbox.rippling.com/api/platform/api" , token)
}

func NewRipClient(baseURLStr,token string) Client {
	baseURL, err := url.Parse(baseURLStr)
	if err != nil {
		logger.Log.Error(err.Error())
		os.Exit(1)
	}
	return Client {
		baseURL: baseURL,
		httpClient: &http.Client{
			Timeout: time.Second * 10,
		},
		token: token,
	}
}

func (c *Client) GetReportForDate(token, report, date string) (Response, error) { 
	c.MakeRequest("POST", "/reports/report_data", strings.NewReader(fmt.Sprintf(`{"report_name":"%s"}`, report)))
	
	
	// return Client{
	// 	httpClient: client,
	// }
	
	// c.Do(req)
	return map[string]string{"hello": "world"}, nil
}