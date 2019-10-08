package dobaos

import (
	"errors"
	"reflect"
	"testing"

	"github.com/go-redis/redis"
)

func TestClient_commonRequest(t *testing.T) {
	type fields struct {
		pub    publisher
		pubsub receiver
		json   marshalUnmarshaller
	}
	type args struct {
		channel string
		method  string
		payload interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    map[string]interface{}
		wantErr bool
	}{
		{
			name: "no errors",
			fields: fields{
				pub: &mockPublisher{
					IntCmdToReturn: &redis.IntCmd{},
				},
				pubsub: &mockReceiver{
					MessageToReturn: &redis.Message{},
					ErrorToReturn:   nil,
				},
				json: &mockMarshalUnmarshaller{
					BytesToReturn: []byte("test"),
				},
			},
			want:    map[string]interface{}{},
			wantErr: false,
		},
		{
			name: "marshal error",
			fields: fields{
				pub: &mockPublisher{
					IntCmdToReturn: &redis.IntCmd{},
				},
				pubsub: &mockReceiver{
					MessageToReturn: &redis.Message{},
					ErrorToReturn:   nil,
				},
				json: &mockMarshalUnmarshaller{
					MErrorToReturn: errors.New("dummy marshal error"),
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "receiveMessage error",
			fields: fields{
				pub: &mockPublisher{
					IntCmdToReturn: &redis.IntCmd{},
				},
				pubsub: &mockReceiver{
					MessageToReturn: nil,
					ErrorToReturn:   errors.New("dummy message error"),
				},
				json: &mockMarshalUnmarshaller{},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "unmarshal error",
			fields: fields{
				pub: &mockPublisher{
					IntCmdToReturn: &redis.IntCmd{},
				},
				pubsub: &mockReceiver{
					MessageToReturn: &redis.Message{},
					ErrorToReturn:   nil,
				},
				json: &mockMarshalUnmarshaller{
					UErrorToReturn: errors.New("dummy unmarshal error"),
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				pub:    tt.fields.pub,
				pubsub: tt.fields.pubsub,
				json:   tt.fields.json,
			}
			got, err := c.commonRequest(tt.args.channel, tt.args.method, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.commonRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.commonRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetValue(t *testing.T) {
	type fields struct {
		pub    publisher
		pubsub receiver
		json   marshalUnmarshaller
	}
	type args struct {
		id interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    map[string]interface{}
		wantErr bool
	}{
		{
			name: "no errors",
			fields: fields{
				pub: &mockPublisher{
					IntCmdToReturn: &redis.IntCmd{},
				},
				pubsub: &mockReceiver{
					MessageToReturn: &redis.Message{},
					ErrorToReturn:   nil,
				},
				json: &mockMarshalUnmarshaller{
					BytesToReturn: []byte("test"),
				},
			},
			want:    map[string]interface{}{},
			wantErr: false,
		},
		{
			name: "marshal error",
			fields: fields{
				pub: &mockPublisher{
					IntCmdToReturn: &redis.IntCmd{},
				},
				pubsub: &mockReceiver{
					MessageToReturn: &redis.Message{},
					ErrorToReturn:   nil,
				},
				json: &mockMarshalUnmarshaller{
					MErrorToReturn: errors.New("dummy marshal error"),
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "unmarshal error",
			fields: fields{
				pub: &mockPublisher{
					IntCmdToReturn: &redis.IntCmd{},
				},
				pubsub: &mockReceiver{
					MessageToReturn: &redis.Message{},
					ErrorToReturn:   nil,
				},
				json: &mockMarshalUnmarshaller{
					UErrorToReturn: errors.New("dummy unmarshal error"),
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				pub:    tt.fields.pub,
				pubsub: tt.fields.pubsub,
				json:   tt.fields.json,
			}
			got, err := c.GetValue(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.GetValue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.GetValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_SetValue(t *testing.T) {
	type fields struct {
		pub    publisher
		pubsub receiver
		json   marshalUnmarshaller
	}
	type args struct {
		payload interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    map[string]interface{}
		wantErr bool
	}{
		{
			name: "no errors",
			fields: fields{
				pub: &mockPublisher{
					IntCmdToReturn: &redis.IntCmd{},
				},
				pubsub: &mockReceiver{
					MessageToReturn: &redis.Message{},
					ErrorToReturn:   nil,
				},
				json: &mockMarshalUnmarshaller{
					BytesToReturn: []byte("test"),
				},
			},
			want:    map[string]interface{}{},
			wantErr: false,
		},
		{
			name: "marshal error",
			fields: fields{
				pub: &mockPublisher{
					IntCmdToReturn: &redis.IntCmd{},
				},
				pubsub: &mockReceiver{
					MessageToReturn: &redis.Message{},
					ErrorToReturn:   nil,
				},
				json: &mockMarshalUnmarshaller{
					MErrorToReturn: errors.New("dummy marshal error"),
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "unmarshal error",
			fields: fields{
				pub: &mockPublisher{
					IntCmdToReturn: &redis.IntCmd{},
				},
				pubsub: &mockReceiver{
					MessageToReturn: &redis.Message{},
					ErrorToReturn:   nil,
				},
				json: &mockMarshalUnmarshaller{
					UErrorToReturn: errors.New("dummy unmarshal error"),
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				pub:    tt.fields.pub,
				pubsub: tt.fields.pubsub,
				json:   tt.fields.json,
			}
			got, err := c.SetValue(tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.SetValue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.SetValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNew(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "test 1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(); got == nil {
				t.Errorf("New() returned nil")
			}
		})
	}
}
