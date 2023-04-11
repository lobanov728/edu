package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

type Item struct {
	Title, SourceChannel, GUID string
}

type Fetcher interface {
	Fetch() (items []Item, next time.Time, err error)
}

func Fetch(domain string) Fetcher {
	return &fakeFetcher{
		dataChannel: domain,
	}
}

var _ Fetcher = (*fakeFetcher)(nil)

type fakeFetcher struct {
	dataChannel string
	items       []Item
}

// FakeDuplicates causes the fake fetcher to return duplicate items.
var FakeDuplicates bool

func (f *fakeFetcher) Fetch() (items []Item, next time.Time, err error) {
	now := time.Now()
	next = now.Add(time.Duration(rand.Intn(5)) * 500 * time.Millisecond)

	item := Item{
		Title:         fmt.Sprintf("Item %d", len(f.items)),
		SourceChannel: f.dataChannel,
	}
	item.GUID = item.SourceChannel + "/" + item.Title

	f.items = append(f.items, item)

	if FakeDuplicates {
		items = f.items
	} else {
		items = []Item{item}
	}

	return
}

type Subscription interface {
	Updates() <-chan Item
	Close() error
}

var _ Subscription = (*subcrition)(nil)

func Subscribe(fetcher Fetcher) Subscription {
	s := &subcrition{
		fetcher: fetcher,
		updates: make(chan Item),
		closing: make(chan chan error),
	}

	go func() {
		var pending []Item
		var next time.Time
		var err error

		for {
			var fetchDelay time.Duration
			now := time.Now()
			if next.After(now) {
				fetchDelay = next.Sub(now)
			}
			startFetch := time.After(fetchDelay)

			var first Item
			var updates chan Item
			if len(pending) > 0 {
				first = pending[0]
				updates = s.updates
			}

			select {
			case errCh := <-s.closing:
				errCh <- err
				close(s.updates)
				return
			case <-startFetch:
				var feched []Item
				feched, next, err = s.fetcher.Fetch()
				if err != nil {
					next = time.Now().Add(10 * time.Second)
					break
				}
				pending = append(pending, feched...)
			case updates <- first:
				pending = pending[1:]
			}

		}

	}()

	return s
}

type subcrition struct {
	fetcher Fetcher
	updates chan Item
	closing chan chan error
}

func (s *subcrition) Updates() <-chan Item {
	return s.updates
}

func (s *subcrition) Close() error {
	errCh := make(chan error)
	s.closing <- errCh
	return <-errCh
}

type naiveMerge struct {
	subs    []Subscription
	updates chan Item
}

var _ Subscription = (*naiveMerge)(nil)

func NaiveMerge(subs ...Subscription) Subscription {
	m := &naiveMerge{
		subs:    subs,
		updates: make(chan Item),
	}

	for _, sub := range subs {
		go func(s Subscription) {
			for item := range s.Updates() {
				m.updates <- item
			}
		}(sub)
	}

	return m
}

// Close implements Subscription
func (m *naiveMerge) Close() (err error) {
	for _, s := range m.subs {
		if subErr := s.Close(); err == nil && subErr != nil {
			err = subErr
		}
	}

	close(m.updates)

	return
}

// Updates implements Subscription
func (m *naiveMerge) Updates() <-chan Item {
	return m.updates
}

func main() {
	rand.Seed(time.Now().UnixNano())

	merged := NaiveMerge(
		Subscribe(Fetch("blog.golang.org")),
		Subscribe(Fetch("googleblog.blogspot.com")),
		Subscribe(Fetch("googledevelopers.blogspot.com")),
	)

	time.AfterFunc(3*time.Second, func() {
		fmt.Println("closed:", merged.Close())
	})

	for it := range merged.Updates() {
		fmt.Println(runtime.NumGoroutine())
		fmt.Println(it.SourceChannel, it.Title)
	}

	fmt.Println(runtime.NumGoroutine())
}
