type DynamoAccessor interface {
	Do(*Request) (*Response, error)
}

func (c *Client) Do(request *Request) (*Response, error) {
	switch request.Method {
	case BatchWrite:
		resp, err := c.batchWriteItem(request)
		break
	case Update:
		resp, err := c.updateItem(request)
		break

		// More dynamo functions...

		return resp, nil
	}
}

func (c *FakeClient) Do(request *Request) (*Response, error) {
	return what_we_want
}
