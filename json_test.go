package dobaos

import (
	"reflect"
	"testing"
)

func Test_jsonWrapper_Marshal(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name    string
		jw      *jsonWrapper
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name:    "test 1",
			jw:      &jsonWrapper{},
			args:    args{v: struct{}{}},
			want:    []byte("{}"),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jw := &jsonWrapper{}
			got, err := jw.Marshal(tt.args.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("jsonWrapper.Marshal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("jsonWrapper.Marshal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_jsonWrapper_Unmarshal(t *testing.T) {
	var placeholder struct{}

	type args struct {
		b []byte
		v interface{}
	}
	tests := []struct {
		name    string
		jw      *jsonWrapper
		args    args
		wantErr bool
	}{
		{
			name:    "test 1",
			jw:      &jsonWrapper{},
			args:    args{b: []byte("{}"), v: &placeholder},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jw := &jsonWrapper{}
			if err := jw.Unmarshal(tt.args.b, tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("jsonWrapper.Unmarshal() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
