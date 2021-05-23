package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	json_string = `{"title":"Cloud Native Go","author":"M.-L. Reimer","isbn":"0123456789"}`
)

func TestBookToJson(t *testing.T) {
	book := Book{Title: "Cloud Native Go", Author: "M.-L. Reimer", ISBN: "0123456789"}
	json := book.ToJSON()

	assert.Equal(t,
		json_string,
		string(json),
		"Book JSON marshalling wrong.",
	)
}

func TestBookFromJson(t *testing.T) {
	json := []byte(json_string)
	book := FromJSON(json)
	assert.Equal(t,
		Book{Title: "Cloud Native Go", Author: "M.-L. Reimer", ISBN: "0123456789"},
		book,
		"Book JSON unmarshalling wrong",
	)
}
