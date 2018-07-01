package ghcp

import (
	"time"
)

type Counter struct {
	acMaps       []acMap  // the process frequently switches between two article count maps. At a time one is to be written to and one is to be read out of.
	wsw          chan int // wsw stands for write-switch
	rsw          chan int // rsw stands for read-switch
	ticker       *time.Ticker
	switchChan   chan bool // chan for the ticker that emits the signal to switch between maps
	topN         int       // collect the top n articles with the highest Count
	topNDone chan bool
}

type acMap map[string]int // acMap stands for article count map

func MakeCounter(interval time.Duration, numberOfTopArticles int) *Counter {
	c := &Counter{
		acMaps:     make([]acMap, 2),
		switchChan: make(chan bool),
		wsw:        make(chan int),
		rsw:        make(chan int),
		ticker:     time.NewTicker(interval),
		topN:       numberOfTopArticles,
		topNDone: 	make(chan bool),
	}

	c.acMaps[0] = make(map[string]int)
	c.acMaps[1] = make(map[string]int)

	go c.startTicker()

	go c.switchMap()

	return c
}

func (rcv *Counter) CountArticleChan() chan<- string {
	countArticleNr := make(chan string)

	go func() {

		var w int

		// prioritize the switch to the write-map, before writing to it
		for {
			select {
			case w = <-rcv.wsw:
				//log.Printf("received write switch top, map nr. %v \n", w)
			default:
			}

			select {
			case w = <-rcv.wsw:
				//log.Printf("received write switch bottom, map nr. %v \n", w)

			case aNr := <-countArticleNr:

				readMap := rcv.acMaps[w]

				//log.Printf("writing to map nr. %v, len %v", w, len(readMap))

				// increasing article-Number Count value by one
				readMap[aNr]++
			}
		}
	}()

	return countArticleNr
}

func (rcv *Counter) TopNArticlesChan() <-chan []Article {

	collection := make(chan []Article)

	go func() {
		for {
			select {
			case r := <-rcv.rsw:

				readMap := rcv.acMaps[r]

				//log.Printf("reading from map nr. %d, length of map %v", r, len(readMap))
				collection <- getHighestNArticles(rcv.topN, readMap)

				// reset map
				rcv.acMaps[r] = make(map[string]int)
				rcv.topNDone <- true
			}
		}
	}()

	return collection
}

func (rcv *Counter) switchMap() {
	writeMap := 0
	rcv.wsw <- 0

	for {
		select {
		case <-rcv.switchChan:
			if writeMap == 0 {
				writeMap = 1

				//log.Println("sending write switch map1")
				rcv.wsw <- 1

				//log.Println("sending read switch map0")
				rcv.rsw <- 0
			} else {
				writeMap = 0

				//log.Println("sending write switch map0")
				rcv.wsw <- 0

				//log.Println("sending read switch map1")
				rcv.rsw <- 1
			}
		}
	}
}

func (rcv *Counter) startTicker() {

	var proceed = true

	for {
		select {
		case proceed = <- rcv.topNDone:

		case <-rcv.ticker.C:

			// waiting for the topN-Articles-Counter to finish
			if !proceed {
				continue
			}

			//log.Println("send switch signal")
			rcv.switchChan <- true
			proceed = false
		}
	}
}
