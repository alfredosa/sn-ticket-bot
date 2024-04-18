package handlers

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadPicture(t *testing.T) {
	val, err := readPicture()
	if err != nil {
		fmt.Println(err)
	}

	assert.Nil(t, err)
	assert.NotNil(t, val)
}
