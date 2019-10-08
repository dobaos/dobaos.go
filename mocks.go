package dobaos

import "github.com/go-redis/redis"

type mockPublisher struct {
	IntCmdToReturn *redis.IntCmd
}

func (m *mockPublisher) Publish(string, interface{}) *redis.IntCmd {
	return m.IntCmdToReturn
}

type mockReceiver struct {
	MessageToReturn *redis.Message
	ErrorToReturn   error
}

func (m mockReceiver) ReceiveMessage() (*redis.Message, error) {
	return m.MessageToReturn, m.ErrorToReturn
}

type mockMarshalUnmarshaller struct {
	BytesToReturn  []byte
	UErrorToReturn error
	MErrorToReturn error
}

func (m *mockMarshalUnmarshaller) Marshal(v interface{}) ([]byte, error) {
	return m.BytesToReturn, m.MErrorToReturn
}

func (m *mockMarshalUnmarshaller) Unmarshal(b []byte, v interface{}) error {
	return m.UErrorToReturn
}
