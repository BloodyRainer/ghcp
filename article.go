package ghcp

import (
	"sort"
	"fmt"
	"log"
)

type Article struct {
	Number string
	Count  int
}

func (rcv Article) String() string {
	return fmt.Sprintf("Article[nr: %v, Count: %v]", rcv.Number, rcv.Count)
}

type articleList []Article

func (al articleList) Len() int {
	return len(al)
}

func (al articleList) Less(i, j int) bool {
	return al[i].Count < al[j].Count
}

func (al articleList) Swap(i, j int) {
	al[i], al[j] = al[j], al[i]
}

func getHighestNArticles(n int, acMap map[string]int) []Article {

	if len(acMap) == 0 {
		return nil
		log.Println("no articles in acMap")
	}

	al := make(articleList, len(acMap))

	i := 0

	for k, v := range acMap {
		al[i] = Article{k, v}
		i ++
	}

	sort.Sort(sort.Reverse(al))

	firstN := make([]Article, n, n)

	for i := 0; i < n; i ++ {
		firstN[i] = al[i]
	}

	return firstN

}
