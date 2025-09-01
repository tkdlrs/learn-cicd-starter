package auth

import (
	"errors"
	"net/http"
	"reflect"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	//
	type test struct {
		name         string
		InputHeaders http.Header
		wantErr      bool
		want         error
	}
	//
	tests := []test{
		{
			name:    "No included header",
			wantErr: true,
			want:    errors.New("no authorization header included"),
		},
		{
			name:    "Bad auth header",
			wantErr: true,
			want:    errors.New("malformed authorization header"),
		},
		{
			name:    "Doesn't start with 'ApiKey'",
			wantErr: true,
			want:    errors.New("malformed authorization header"),
		},
	}
	//
	if tests[1].InputHeaders == nil {
		tests[1].InputHeaders = make(http.Header)
		tests[1].InputHeaders.Add("Content-Type", "application/json")
		tests[1].InputHeaders.Add("Authorization", "no")
	}
	//
	if tests[2].InputHeaders == nil {
		tests[2].InputHeaders = make(http.Header)
		tests[2].InputHeaders.Add("Content-Type", "application/json")
		tests[2].InputHeaders.Add("Authorization", "this is noise")
	}
	//
	for _, tc := range tests {
		_, err := GetAPIKey(tc.InputHeaders)
		if tc.wantErr {
			if !reflect.DeepEqual(tc.want, err) {
				t.Fatalf("test: %v, \n expected error of: -%v-, instead got: -%v-", tc.name, tc.want, err)
			}
		}
		//  else {
		// 	if tc.want == got {
		// 		t.Fatalf("expected value: %v, got: %v", tc.want, got)
		// 	}
		// }

	}
}
