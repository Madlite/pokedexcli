package pokeapi

import(
	"fmt"
	"net/http"
)

type Client struct {
	httpClient *http.Client
}

func NewClient() Client {
	return Client{
		httpClient: http.DefaultClient,
	}
}

func (c *Client) Get(url string) (*http.Response, error) {
	return c.httpClient.Get(url)
}

func getApiRequest()  {
	fmt.Println("Fetching api...")
}