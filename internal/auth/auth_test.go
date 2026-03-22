package auth

import (
	"net/http"
	"reflect"
	"testing"
)

func TestGetAPIKey(t *testing.T) {

	tests := []struct {
		name  string
		key   string
		value string
		want  string
	}{
		{name: "no auth header", key: "", value: "", want: ""},
		{name: "empty auth header", key: "Authorization", value: "", want: ""},
		{name: "wrong auth key", key: "Authorization", value: "api-key thisisakey", want: ""},
		{name: "valid auth", key: "Authorization", value: "ApiKey validkey", want: "validkey"},
	}

	for _, tc := range tests {
		headers := http.Header{}
		headers.Add(tc.key, tc.value)
		got, _ := GetAPIKey(headers)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("%s: expected: %v, got: %v", tc.name, tc.want, got)
		}
	}
}
