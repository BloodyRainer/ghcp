package ghcp

import (
	"testing"
	"time"
	"github.com/stretchr/testify/assert"
)

func TestCounter(t *testing.T) {

	c := MakeCounter(1*time.Second, 2)
	caChan := c.CountArticleChan()
	articlesChan := c.TopNArticlesChan()

	caChan <- "one"
	caChan <- "one"
	caChan <- "two"
	caChan <- "one"
	caChan <- "three"
	caChan <- "two"
	caChan <- "four"

	articles := <-articlesChan

	assert.Equal(t, 3, articles[0].Count)
	assert.Equal(t, "one", articles[0].Number)
	assert.Equal(t, 2, articles[1].Count)
	assert.Equal(t, "two", articles[1].Number)

	caChan <- "three"
	caChan <- "one"
	caChan <- "two"
	caChan <- "two"
	caChan <- "one"
	caChan <- "two"
	caChan <- "two"

	articles = <-articlesChan

	assert.Equal(t, 4, articles[0].Count)
	assert.Equal(t, "two", articles[0].Number)
	assert.Equal(t, 2, articles[1].Count)
	assert.Equal(t, "one", articles[1].Number)

	caChan <- "three"
	caChan <- "three"
	caChan <- "three"
	caChan <- "one"
	caChan <- "one"
	caChan <- "two"

	articles = <-articlesChan

	assert.Equal(t, 3, articles[0].Count)
	assert.Equal(t, "three", articles[0].Number)
	assert.Equal(t, 2, articles[1].Count)
	assert.Equal(t, "one", articles[1].Number)

	caChan <- "four"
	caChan <- "four"
	caChan <- "four"
	caChan <- "one"
	caChan <- "one"
	caChan <- "four"

	articles = <-articlesChan

	assert.Equal(t, 4, articles[0].Count)
	assert.Equal(t, "four", articles[0].Number)
	assert.Equal(t, 2, articles[1].Count)
	assert.Equal(t, "one", articles[1].Number)

}
