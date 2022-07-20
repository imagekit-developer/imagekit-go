package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

type Tstringer struct {
	name string
}

func (ts Tstringer) String() string {
	return ts.name
}

// read closer implementation
type MockedRC struct {
	called bool
	data   io.ReadCloser
}

func (rc *MockedRC) Close() error {
	rc.called = true
	return nil
}

func (*MockedRC) Read(p []byte) (n int, err error) {
	return n, err
}

type MockedResponse struct {
	ResponseMetaData
}

func (resp *MockedResponse) SetMeta(meta ResponseMetaData) {
	resp.ResponseMetaData = meta
}

func (resp *MockedResponse) ParseError() error {
	return nil
}

func Test_SetMeta(t *testing.T) {
	resp := &Response{}
	meta := ResponseMetaData{}

	resp.SetMeta(meta)

	if !cmp.Equal(resp.ResponseMetaData, meta) {
		t.Error("ResponseMetaData set")
	}
}

func Test_Body(t *testing.T) {
	b := []byte("test")

	resp := &Response{
		ResponseMetaData{
			Body: b,
		},
	}

	if !cmp.Equal(resp.Body(), b) {
		t.Error("invalid body")
	}
}

func Test_ParseError(t *testing.T) {
	h := http.Header{"content-type": []string{"application/json"}}

	type Ctype struct {
		code   int
		result error
	}
	var cases = map[string]Ctype{
		"nill": {
			200,
			nil,
		},
		"bad request": {
			400,
			ErrBadRequest,
		},
		"unauthorized": {
			401,
			ErrUnauthorized,
		},
		"Forbidden": {
			403,
			ErrForbidden,
		},
		"not-found": {
			404,
			ErrNotFound,
		},
		"too-many-requests": {
			429,
			ErrTooManyRequests,
		},
		"undefine err": {
			600,
			ErrUndefined,
		},
	}

	for _, c := range []int{500, 502, 503, 504} {
		cases[fmt.Sprintf("server error %d", c)] = Ctype{
			c,
			ErrServer,
		}
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			resp := &Response{
				ResponseMetaData{
					Header:     h,
					Body:       []byte(`{"message":"test error","reason":"test reason"}`),
					StatusCode: tc.code,
				},
			}

			err := resp.ParseError()

			if !errors.Is(err, tc.result) {
				t.Error("expected: " + tc.result.Error() + "\n got: " + err.Error())
			}
		})
	}

}

func Test_StructtoParams(t *testing.T) {
	var cases = map[string]struct {
		input  any
		result string
		err    error
	}{
		"test list": {
			input: struct {
				List []string
			}{[]string{"one", "two"}},
			result: `{"List[0]":["one"],"List[1]":["two"]}`,
		},
		"test array": {
			input: struct {
				List [2]string
			}{[2]string{"one", "two"}},
			result: `{"List[0]":["one"],"List[1]":["two"]}`,
		},
		"scalar": {
			input: struct {
				Rank int
				Name string
			}{2, "test"},
			result: `{"Name":["test"],"Rank":["2"]}`,
		},
		"should fail": {
			input:  "abc",
			result: "",
			err:    errors.New("json: cannot unmarshal string into Go value of type map[string]interface {}"),
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			res, err := StructToParams(tc.input)

			if err != nil && tc.err == nil {
				t.Fatal(err)
			}

			if tc.err != nil {
				if err == nil {
					t.Error("err is nil")
				}
				if err.Error() != tc.err.Error() {
					t.Error("invalid error")
				}
				return
			}
			str, err := json.Marshal(res)

			if string(str) != tc.result {
				t.Error("expected: " + tc.result + "\ngot: " + string(str))
			}
		})
	}
}

func Test_BuildPath(t *testing.T) {

	parts := []any{
		1, "two", Tstringer{"test"},
	}

	result := BuildPath(parts...)

	if result != "1/two/test" {
		t.Error("invalid")
	}
}

func Test_DeferredBodyClose(t *testing.T) {
	var rc = &MockedRC{}

	var resp = &http.Response{
		Body: rc,
	}

	DeferredBodyClose(resp)
	if rc.called == false {
		t.Error("deffered close not called")
	}
}

func Test_SetResponseMeta(t *testing.T) {
	h := http.Header{"content-type": []string{"application/json"}}

	var response = &MockedResponse{}
	var meta = ResponseMetaData{
		Header:     h,
		StatusCode: 200,
		Body:       []byte("hello"),
	}

	SetResponseMeta(&http.Response{
		Header:     h,
		Body:       io.NopCloser(strings.NewReader("hello")),
		StatusCode: 200,
	}, response)

	if !cmp.Equal(response.ResponseMetaData, meta) {
		t.Error("invalid header")
	}
}

func Test_Bool(t *testing.T) {
	resp := Bool(true)

	if *resp != true {
		t.Error("invalid bool(true)")
	}

	resp = Bool(false)

	if *resp != false {
		t.Error("invalid bool(false)")
	}
}
