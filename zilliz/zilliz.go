package zilliz

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"time"
)

const (
	apiTemplateUrl string = "https://controller.api.%s.zillizcloud.com/v1/"
)

type Client struct {
	CloudRegionId string
	HTTPClient    *http.Client
	baseUrl       string
	apiKey        string
	userAgent     string
}

// NewClient - creates new Pinecone client.
func NewClient(apiKey string, cloudRegionId string) *Client {
	c := &Client{
		CloudRegionId: cloudRegionId,
		HTTPClient:    &http.Client{Timeout: 30 * time.Second},
		baseUrl:       apiTemplateUrl,
		apiKey:        apiKey,
		userAgent:     "skyscrapr/terraform-provider-zilliz",
	}
	return c
}

type zillizResponse[T any] struct {
	Code int `json:"code"`
	Data T   `json:"data"`
}

type zillizPage struct {
	Count       int `json:"count"`
	CurrentPage int `json:"currentPage"`
	PageSize    int `json:"pageSize"`
}

type CloudProvider struct {
	CloudId     string `json:"cloudId"`
	Description string `json:"description"`
}

func (c *Client) ListCloudProviders() ([]CloudProvider, error) {
	var cloudProviders zillizResponse[[]CloudProvider]
	err := c.do("GET", "clouds", nil, nil, &cloudProviders)
	return cloudProviders.Data, err
}

type CloudRegion struct {
	ApiBaseUrl string `json:"apiBaseUrl"`
	CloudId    string `json:"cloudId"`
	RegionId   string `json:"regionId"`
}

func (c *Client) ListCloudRegions(cloudId string) ([]CloudRegion, error) {
	var cloudRegions zillizResponse[[]CloudRegion]
	err := c.do("GET", "clouds", nil, nil, &cloudRegions)
	return cloudRegions.Data, err
}

type Clusters struct {
	zillizPage
	Clusters []Cluster `json:"clusters"`
}

type Cluster struct {
	ClusterId          string `json:"clusterId"`
	ClusterName        string `json:"clusterName"`
	Description        string `json:"description"`
	RegionId           string `json:"regionId"`
	ClusterType        string `json:"clusterType"`
	CuSize             int64  `json:"cuSize"`
	Status             string `json:"status"`
	ConnectAddress     string `json:"connectAddress"`
	PrivateLinkAddress string `json:"privateLinkAddress"`
	CreateTime         string `json:"createTime"`
}

func (c *Client) ListClusters() (Clusters, error) {
	var clusters zillizResponse[Clusters]
	err := c.do("GET", "clusters", nil, nil, &clusters)
	return clusters.Data, err
}

func (c *Client) DescribeCluster(clusterId string) (Cluster, error) {
	var cluster zillizResponse[Cluster]
	err := c.do("GET", "clusters/"+clusterId, nil, nil, &cluster)
	return cluster.Data, err
}

type CreateClusterParams struct {
	CreateCollection        bool   `json:"createCollection"`
	CreateExampleCollection bool   `json:"createExampleCollection"`
	InstanceName            string `json:"instanceName"`
	ProjectId               int    `json:"projectId"`
	RegionId                string `json:"regionId"`
}

func (c *Client) CreateCluster(params CreateClusterParams) error {
	err := c.do("POST", "clusters", params, nil, nil)
	return err
}

func (c *Client) do(method string, path string, body interface{}, values url.Values, result interface{}) error {
	u, err := c.buildURL(path)
	if err != nil {
		return err
	}
	req, err := c.newRequest(method, u, body)
	if err != nil {
		return err
	}
	if values != nil {
		req.URL.RawQuery = values.Encode()
	}
	return c.doRequest(req, result)
}

func (c *Client) newRequest(method string, u *url.URL, body interface{}) (*http.Request, error) {
	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}
	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}
	// req.Header.Set("User-Agent", c.UserAgent)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	return req, nil
}

func (c *Client) doRequest(req *http.Request, v any) error {
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		return c.handleHTTPErrorResp(res)
	}

	return decodeResponse(res.Body, v)
}

func decodeResponse(body io.Reader, v any) error {
	if v == nil {
		return nil
	}
	b, err := io.ReadAll(body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(b, v)
	return err
	// return json.NewDecoder(body).Decode(v)
}

func (c *Client) buildURL(endpointPath string) (*url.URL, error) {
	u, err := url.Parse(endpointPath)
	if err != nil {
		return nil, err
	}
	sBaseUrl := c.baseUrl
	if c.CloudRegionId != "" {
		sBaseUrl = fmt.Sprintf(apiTemplateUrl, c.CloudRegionId)
	}
	baseUrl, err := url.Parse(sBaseUrl)
	if err != nil {
		return nil, err
	}
	u.Path = path.Join(baseUrl.Path, u.Path)
	return baseUrl.ResolveReference(u), err
}

func (c *Client) handleHTTPErrorResp(resp *http.Response) error {
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	reqErr := &HTTPError{
		StatusCode: resp.StatusCode,
		Status:     resp.Status,
		Message:    string(data),
	}
	return reqErr
}

// HTTPError provides informations about generic HTTP errors.
type HTTPError struct {
	StatusCode int
	Status     string
	Message    string
}

func (e *HTTPError) Error() string {
	return fmt.Sprintf("error, status code: %d, message: %s", e.StatusCode, e.Message)
}
