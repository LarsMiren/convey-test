package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGet(t *testing.T) {
	c := http.DefaultClient
	Convey("Test GET method", t, func() {
		resp, err := c.Get("https://jsonplaceholder.typicode.com/todos/1")
		Convey("No errors occurs while requesting", func() {
			So(err, ShouldBeNil)
		})
		Convey("Status code must be 200", func() {
			So(resp.StatusCode, ShouldEqual, 200)
		})

		b, err := ioutil.ReadAll(resp.Body)
		Convey("No errors occurs while reading responce body", func() {
			So(err, ShouldBeNil)
		})

		var todo ToDo
		err = json.Unmarshal(b, &todo)
		Convey("No errors occurs while unmarshalling body", func() {
			So(err, ShouldBeNil)
		})
	})
}

func TestPost(t *testing.T) {
	c := http.DefaultClient
	Convey("Test POST method", t, func() {
		todo, err := json.Marshal(&ToDo{
			255,
			1,
			"New ToDo",
			false,
		})
		Convey("No errors occurs while marsalling", func() {
			So(err, ShouldBeNil)
		})
		resp, err := c.Post(
			"https://jsonplaceholder.typicode.com/todos",
			"appliction/json",
			bytes.NewBuffer(todo),
		)
		Convey("No errors occurs while requesting", func() {
			So(err, ShouldBeNil)
		})
		Convey("Status code must be 201", func() {
			So(resp.StatusCode, ShouldEqual, 201)
		})

	})
}

type ToDo struct {
	ID        int64  `json:"id"`
	UserID    int64  `json:"userId"`
	Title     string `json:"title"`
	Complited bool   `json:"complited"`
}
