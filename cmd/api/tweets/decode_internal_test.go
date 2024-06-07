package tweets

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGlobalToLocalXPath_success(t *testing.T) {
	xpath := "div/div/div[2]/div[2]/div[1]/div/div[1]/div/div/div[2]/div/div[3]/a/time"

	want := "div/div/div[position()=2]/div[position()=2]/div[position()=1]/div/div[position()=1]/div/div/div[position()=2]/div/div[position()=3]/a/time"
	got := globalToLocalXPath(xpath)

	assert.Equal(t, want, got)
}
