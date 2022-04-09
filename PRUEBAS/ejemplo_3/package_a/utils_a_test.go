package package_a

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHello(t *testing.T) {

	//assert := assert.New(T)
	assert.Equal(t, "Helo, World", Hello())

}
