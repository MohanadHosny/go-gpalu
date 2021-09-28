package gpalu

import (
	"fmt"
	"net/http"
	"regexp"
)

type Mail struct {
	content string
}

type Client struct {
	client *http.Client
	baseUrl string
}

var (
	baseUrl = "https://gpa.lu/%s"
	emailRegex = regexp.MustCompile(`class="text-truncate">(.*?)<`)
	inboxRegex = regexp.MustCompile("/.*/(.*)\\?noheader")
)

// Create new Gpalu Client.
func NewClient(userClient *http.Client) *Client {
	if userClient != nil{
		return &Client{client: userClient, baseUrl: baseUrl}
	}

	return &Client{client: http.DefaultClient, baseUrl: baseUrl}
}

// Get new random email address.
func (client *Client) GetAddress() string {
	route := fmt.Sprintf(client.baseUrl, "?random")

	resp, err := client.client.Get(route)
	body := ReadBody(resp, err)

	email := emailRegex.FindStringSubmatch(string(body))
	return email[1]
}

// Get mail content by id.
func (client *Client) GetContent(email string,  messageId string) string {
	data := fmt.Sprintf("%s/%s?noheader", email, messageId)
	route := fmt.Sprintf(client.baseUrl, data)

	resp, err := client.client.Get(route)
	body := ReadBody(resp, err)

	return string(body)
}

// Get inbox.
func (client *Client) GetInbox(email string, limit int) []Mail {
	route := fmt.Sprintf(client.baseUrl, email)

	resp, err := client.client.Get(route)
	body := ReadBody(resp, err)

	allMails := inboxRegex.FindAllStringSubmatch(string(body), limit)
	var mails []Mail

	for _, match := range allMails{
		content := client.GetContent(email, match[1])
		mails = append(mails, Mail{content: content})
	}

	return mails
}
