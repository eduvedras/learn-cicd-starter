package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	type test struct{
		input http.Header
		want string
		err error
	}

	tests := []test{
		{input: createHeader("ApiKey key1"), want: "key1"},
		{input: createHeader("ApiKey key2"), want: "key2"},
		{input: http.Header{}, err: errors.New("no authorization header included")},
		{input: createHeader("ApiKy key3"), err: errors.New("malformed authorization header")},
	}

	for _, tc := range tests {
        got, err := GetAPIKey(tc.input)
        if tc.want != got{
			t.Fatalf("expected: %v, got: %v", tc.want, got)
        }
		if tc.err != nil && err != nil && tc.err.Error() != err.Error(){
			t.Fatalf("expected err: %v, got err: %v", tc.err, err)
		}
    }
}

func createHeader(text string) http.Header{
	req := new(http.Request)
	req.Header = make(http.Header)
	req.Header.Add("Authorization", text)
	return req.Header
}