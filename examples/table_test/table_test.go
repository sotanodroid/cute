//go:build example
// +build example

package table_test

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"testing"

	"github.com/ozontech/cute"
	"github.com/ozontech/cute/asserts/json"
	"github.com/ozontech/cute/errors"
)

func init() {
	os.Setenv("ALLURE_OUTPUT_PATH", "../") // custom, read Readme.md for more info
}

func TestTableExample(t *testing.T) {
	u, _ := url.Parse("https://jsonplaceholder.typicode.com/posts/1/comments")

	req, _ := http.NewRequest(http.MethodPost, u.String(), nil)
	req.Header = map[string][]string{
		"some_auth_token": []string{fmt.Sprint(11111)},
	}
	cute.NewTestBuilder().
		CreateTableTest().
		PutNewTest(
			"Execute validation 1",
			req,
			&cute.Expect{
				Code: 201,
			}).
		PutNewTest(
			"Execute validation 2",
			req,
			&cute.Expect{
				AssertBody: []cute.AssertBody{
					json.Equal("$[0].email", "Eliseo@gardner.biz"),
					json.Present("$[1].name"),
				},
			},
		).
		ExecuteTest(context.Background(), t)
}

func TestTableExample_cuteTest_one(t *testing.T) {
	test := &cute.Test{
		Name: "test_1",
		Request: &cute.Request{
			Builders: []cute.RequestBuilder{
				cute.WithURI("https://jsonplaceholder.typicode.com/posts/1/comments"),
				cute.WithMethod(http.MethodGet),
			},
		},
		Expect: nil,
	}

	test.Execute(context.Background(), t)
}

func TestTableExample_cuteTest_multi(t *testing.T) {
	tests := []*cute.Test{
		{
			Name:       "test_1",
			Middleware: nil,
			Request: &cute.Request{
				Builders: []cute.RequestBuilder{
					cute.WithURI("https://jsonplaceholder.typicode.com/posts/1/comments"),
					cute.WithMethod(http.MethodPost),
				},
			},
			Expect: &cute.Expect{
				Code: 200,
			},
		},
		{
			Name:       "test_2",
			Middleware: nil,
			Request: &cute.Request{
				Builders: []cute.RequestBuilder{
					cute.WithURI("https://jsonplaceholder.typicode.com/posts/1/comments"),
					cute.WithMethod(http.MethodGet),
				},
			},
			Expect: &cute.Expect{
				Code: 200,
				AssertBody: []cute.AssertBody{
					json.Equal("$[0].email", "Eliseo@gardner.biz"),
					json.Present("$[1].name"),
					func(body []byte) error {
						return errors.NewAssertError("example error", "example message", nil, nil)
					},
				},
			},
		},
	}

	for _, test := range tests {
		test.Execute(context.Background(), t)
	}
}
