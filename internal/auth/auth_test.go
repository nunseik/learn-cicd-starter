package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestAuth(t *testing.T) {
	tests := map[string]struct {
		input http.Header
		want  string
		err   error
	}{
		"simple": {
			input: http.Header{
				"Authorization": {"ApiKey your_jwt_token_here"},
			},
			want: "your_jwt_token_here",
			err:  nil,
		},
		"no auth": {
			input: http.Header{},
			want:  "",
			err:   ErrNoAuthHeaderIncluded,
		},
		"malformed": {
			input: http.Header{
				"Authorization": {"ApiKey"},
			},
			want: "",
			err:  ErrMalformedAuthHeader,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, gotErr := GetAPIKey(tc.input)

			if got != tc.want {
				t.Fatalf("%s: want value %q, got %q (err=%v)", name, tc.want, got, gotErr)
			}

			if !errors.Is(gotErr, tc.err) {
				t.Fatalf("%s: want error %v, got %v", name, tc.err, gotErr)
			}
		})
	}
}
