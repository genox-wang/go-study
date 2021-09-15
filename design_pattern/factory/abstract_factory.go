package factory

type Dioer interface {
	Do() string
}

type HttpClient struct{}

func (c *HttpClient) Do() string {
	return "HttpClient"
}

func NewHttpClient() Dioer {
	return &HttpClient{}
}

type MockHttpClient struct{}

func (c *MockHttpClient) Do() string {
	return "MockHttpClient"
}

func NewMockHttpClient() Dioer {
	return &MockHttpClient{}
}

func dioerDo(dioer Dioer) string {
	return dioer.Do()
}
