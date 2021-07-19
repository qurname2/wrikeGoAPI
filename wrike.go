package wrikeGoAPI

import (
	"encoding/json"
	"fmt"
	"github.com/google/go-querystring/query"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	defaultBaseURL = "https://www.wrike.com/"
	apiVersionPath = "api/v4"
	userAgent      = "Golang Wrike API client"
	foldersUrl     = "/folders"
	contactsUrl    = "/contacts"
	tasksUrl       = "/tasks"
)

type Client struct {
	client    *http.Client
	baseURL   *url.URL
	token     string
	userAgent string
	Folders   *FolderService
	Contacts  *ContactService
	Tasks     *TaskService
}

// FolderService endpoint, see Wrike API docs: https://developers.wrike.com/documentation/api/methods/get-folder-tree
type FolderService struct {
	client *Client
}

type ContactService struct {
	client *Client
}

// TaskService is Tasks endpoint, see https://developers.wrike.com/documentation/api/methods/query-tasks
type TaskService struct {
	client *Client
}

//NewClient create a new http client
func NewClient(httpClient *http.Client, token string) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	client := &Client{client: httpClient, userAgent: userAgent}
	client.token = token
	client.Folders = &FolderService{client: client}
	client.Contacts = &ContactService{client: client}
	client.Tasks = &TaskService{client: client}
	if err := client.SetBaseURL(defaultBaseURL); err != nil {
		panic(err)
	}
	return client
}

//GetUserID find the user ID to which to assign the task
func (cl *ContactService) GetUserID(contacts *Contact, username, emailCompany string) string {
	userList := contacts.Data
	currentUser := new(Profiles)
	currentUser.Email = fmt.Sprintf("%v@%s", username, emailCompany)

	for _, userInfo := range userList {
		if userInfo.Profiles != nil {
			if userInfo.Profiles[0].Email == currentUser.Email {
				return userInfo.ID
			}
		}
	}
	return ""
}

// NewRequest build request
func (c *Client) NewRequest(method string, path string, params interface{}) (*http.Request, error) {
	u := *c.baseURL // https://www.wrike.com/api/v4

	bearerToken := fmt.Sprintf("bearer %s", c.token)
	unescaped, err := url.PathUnescape(path)

	if err != nil {
		return nil, err
	}

	u.RawPath = c.baseURL.Path + path
	u.Path = c.baseURL.Path + unescaped

	req := &http.Request{
		Method:     method,
		URL:        &u,
		Proto:      "HTTP/1.1",
		ProtoMajor: 1,
		ProtoMinor: 1,
		Header:     make(http.Header),
		Host:       u.Host,
	}

	if params != nil {
		q, err := query.Values(params)
		if err != nil {
			return nil, err
		}
		u.RawQuery = q.Encode()
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", bearerToken)

	if c.userAgent != "" {
		req.Header.Set("User-Agent", c.userAgent)
	}
	return req, nil
}

// SetBaseURL sets the base URL for API requests to a custom endpoint. urlStr
// should always be specified with a trailing slash.
func (c *Client) SetBaseURL(urlStr string) error {
	// Make sure the given URL end with a slash
	if !strings.HasSuffix(urlStr, "/") {
		urlStr += "/"
	}

	baseURL, err := url.Parse(urlStr)
	if err != nil {
		return err
	}

	if !strings.HasSuffix(baseURL.Path, apiVersionPath) {
		baseURL.Path += apiVersionPath
	}

	// Update the base URL of the client.
	c.baseURL = baseURL

	return nil
}

// Do execute a request
func (c *Client) Do(req *http.Request, v interface{}) (*Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	response := newResponse(resp)
	err = CheckResponse(resp)
	if err != nil {
		// even though there was an error, we still return the response
		// in case the caller wants to inspect it further
		return response, err
	}

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			_, err = io.Copy(w, resp.Body)
		} else {
			err = json.NewDecoder(resp.Body).Decode(v)
		}
	}

	return response, err
}

// CheckResponse check response returned
func CheckResponse(resp *http.Response) error {
	switch resp.StatusCode {
	case 200:
		return nil
	}
	return &HTTPError{}
}

type HTTPError struct{}

func (m *HTTPError) Error() string {
	return "error with making HTTP request occurred"
}

func newResponse(resp *http.Response) *Response {
	response := &Response{Response: resp}
	return response
}

// Response struct
type Response struct {
	*http.Response
}

// CustomField struct
type CustomField struct {
	ID    string `json:"id"`
	Value string `json:"value"`
}

// Metadata struct
type Metadata struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// Project struct
type Project struct {
	AuthorID      string   `json:"authorId"`
	OwnerIds      []string `json:"ownerIds"`
	Status        string   `json:"status"`
	StartDate     string   `json:"startDate"`
	EndDate       string   `json:"endDate"`
	CreatedDate   string   `json:"createdDate"`
	CompletedDate string   `json:"completedDate"`
}

// Time represents time
type Time struct {
	time.Time
}
