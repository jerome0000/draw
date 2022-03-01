package process

import (
	"testing"

	"github.com/jerome0000/draw/conf"
	"github.com/stretchr/testify/assert"
)

func TestRuleHandler(t *testing.T) {
}

func Test_getRuleByRate(t *testing.T) {
	rules := []*conf.Rule{
		{ID: 1, Rate: 1},
		{ID: 2, Rate: 1},
		{ID: 3, Rate: 1},
		{ID: 4, Rate: 0},
	}
	assert.Equal(t, true, getRuleByRate(rules) != nil)
}
