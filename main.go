package main

import (
	"fmt"
	"github.com/SlyMarbo/rss"
	"github.com/bmarcelino/sec/helpers"
	"github.com/bmarcelino/sec/xbrl"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var urls = []string{} //list.New()

const (
	BASE_URL = "http://www.sec.gov/Archives/edgar/monthly/xbrlrss-{year}-{month}.xml"
)

type HttpResponse struct {
	url      string
	response *http.Response
	err      error
}

func asyncHttpGets(urls []string) []*HttpResponse {
	ch := make(chan *HttpResponse)
	responses := []*HttpResponse{}
	for _, url := range urls {
		go func(url string) {
			fmt.Printf("Fetching %s \n", url)
			resp, err := http.Get(url)
			ch <- &HttpResponse{url, resp, err}
		}(url)
	}

	for {
		select {
		case r := <-ch:
			fmt.Printf("%s was fetched\n", r.url)
			responses = append(responses, r)

			if len(responses) == len(urls) {
				return responses
			}

		case <-time.After(50 * time.Millisecond):
			fmt.Printf(".")
		}
	}
	return responses
}

func fetchFeed(url string, ch chan *rss.Feed) {
	fmt.Printf("Fetching %s \n", url)

	feed, err := rss.Fetch(url)

	if err != nil {
		fmt.Println("Failed to get feed %s\n", url)
	}

	ch <- feed
}

func parseFeed(feed *rss.Feed) {
	for _, item := range feed.Items {
		fmt.Println(item.String())
	}
}

func asyncFeedGets(urls []string) []*rss.Feed {
	ch := make(chan *rss.Feed)

	responses := []*rss.Feed{}

	for _, url := range urls {
		go fetchFeed(url, ch)
	}

	for {
		select {
		case r := <-ch:
			fmt.Printf("%s was fetched\n", r.UpdateURL)

			go parseFeed(r)

			/*
				responses = append(responses, r)

				if len(responses) == len(urls) {
					return responses
				}
			*/
		case <-time.After(50 * time.Millisecond):
			fmt.Printf(".")
		}
	}
	return responses
}

func initUrls(year, startMonth, endMonth int) {
	for month := startMonth; month <= endMonth; month++ {
		feedUrl := strings.Replace(BASE_URL, "{year}", strconv.Itoa(year), 1)
		feedUrl = strings.Replace(feedUrl, "{month}", helpers.LeftPad2Len(strconv.Itoa(month), "0", 2), 1)

		//urls.PushBack(feedUrl)
		urls = append(urls, feedUrl)
	}
}

func genUrls(year, startMonth, endMonth int) <-chan string {
	out := make(chan string)

	go func() {
		for month := startMonth; month <= endMonth; month++ {

			feedUrl := strings.Replace(BASE_URL, "{year}", strconv.Itoa(year), 1)
			feedUrl = strings.Replace(feedUrl, "{month}", helpers.LeftPad2Len(strconv.Itoa(month), "0", 2), 1)

			out <- feedUrl
		}
		close(out)
	}()

	return out
}

func getFeed(urls <-chan string) <-chan *rss.Feed {
	out := make(chan *rss.Feed)

	go func() {
		for url := range urls {

			fmt.Printf("Fetching %s \n", url)

			feed, err := rss.Fetch(url)

			if err != nil {
				fmt.Printf("Failed to get feed %s\n", url)
			} else {
				out <- feed
			}
		}
		close(out)
	}()

	return out
}

func getItems(feeds <-chan *rss.Feed) <-chan *rss.Item {
	out := make(chan *rss.Item)

	go func() {
		for feed := range feeds {
			for _, item := range feed.Items {
				if item.Content == "10-K" {
					out <- item
				}
			}
		}
		close(out)
	}()

	return out
}

func main() {
	xbrlInstance := "C:\\Users\\bmarcelino\\Downloads\\wnchvi6-20140331.xml"
	x := new(xbrl.Xbrl)
	//fin := new(xbrl.FundamentantalAccountingConcepts)

	x.Init(xbrlInstance)

	//fin.Init(x)

	// Set up the pipeline.
	/*chUrls := genUrls(2013, 1, 12)
	chFeeds := getFeed(chUrls)
	chItems := getItems(chFeeds)

	for item := range chItems {
		fmt.Println(item)
	}
	*/
	//initUrls(2013, 2013, 1, 12)
	/*
		for e := urls.Front(); e != nil; e = e.Next() {
			fmt.Println("Fetching e.Value")
		}
	*/
	/*
		feed, err := rss.Fetch("http://www.sec.gov/Archives/edgar/monthly/xbrlrss-2014-03.xml")
		if err != nil {
			fmt.Println("Failed to get feed")
		}()

		for _, item := range feed.Items {
			fmt.Println(item.String())
		}

		err = feed.Update()
		if err != nil {
			fmt.Println("Failed to update feed")
		}
	*/
	/*
		results := asyncHttpGets(urls)
		for _, result := range results {
			fmt.Printf("%s status: %s\n", result.url, result.response.Status)
		}
	*/
	//asyncFeedGets(urls)
}
