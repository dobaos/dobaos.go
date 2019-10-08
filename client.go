package dobaos

import (
	"github.com/go-redis/redis"
)

// Client is used to interact with dobaos. It is
// initialised with the redis pub and pubsub
type Client struct {
	pub    publisher
	pubsub receiver
	json   marshalUnmarshaller
}

// New creates a new instance of Client
func New() *Client {
	var redisOpts = redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	}

	sub := redis.NewClient(&redisOpts)
	pubsub := sub.PSubscribe("dobaosgo_*")
	pub := redis.NewClient(&redisOpts)
	return &Client{
		pub:    pub,
		pubsub: pubsub,
		json:   &jsonWrapper{},
	}
}

func (c *Client) commonRequest(channel string, method string, payload interface{}) (map[string]interface{}, error) {
	req := make(map[string]interface{})
	req["method"] = method
	req["payload"] = payload
	req["response_channel"] = "dobaosgo_42"

	res := make(map[string]interface{})
	b, err := c.json.Marshal(&req)
	if err != nil {
		return nil, err
	}

	c.pub.Publish(channel, b)
	msgi, err := c.pubsub.ReceiveMessage()
	if err != nil {
		return nil, err
	}

	err = c.json.Unmarshal([]byte(msgi.Payload), &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GetValue is used to get a value from the
// redis store
func (c *Client) GetValue(id interface{}) (map[string]interface{}, error) {
	return c.commonRequest("dobaos_req", "get value", id)
}

// SetValue is used to set a value to the
// redis store
func (c *Client) SetValue(payload interface{}) (map[string]interface{}, error) {
	return c.commonRequest("dobaos_req", "set value", payload)
}

type publisher interface {
	Publish(string, interface{}) *redis.IntCmd
}

type receiver interface {
	ReceiveMessage() (*redis.Message, error)
}

type marshalUnmarshaller interface {
	Marshal(interface{}) ([]byte, error)
	Unmarshal([]byte, interface{}) error
}
