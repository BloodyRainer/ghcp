package ghcp

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSortArticles(t *testing.T) {
	// Arrange
	m := make(map[string]int)
	m["three"] = 3
	m["seven"] = 7
	m["two"] = 2
	m["one"] = 1
	m["five"] = 5
	m["four"] = 4
	m["eight"] = 8
	m["zero"] = 0
	m["nine"] = 9
	m["d0"] = 2
	m["d1"] = 1
	m["d2"] = 3
	m["eight2"] = 8

	// Act
	articles := getHighestNArticles(5, m)

	// Assert
	assert.Equal(t, 5, len(articles))
	assert.Equal(t, 9, articles[0].Count)
	assert.Equal(t, 8, articles[1].Count)
	assert.Equal(t, 8, articles[2].Count)
	assert.Equal(t, 7, articles[3].Count)
	assert.Equal(t, 5, articles[4].Count)

}
