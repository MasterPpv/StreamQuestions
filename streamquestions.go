package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	//"sync"
	//"strings"
	"github.com/garyburd/go-oauth/oauth"
	"github.com/garyburd/twitterstream"
	"github.com/ChimeraCoder/anaconda"
	"github.com/nsf/termbox-go"
)

var (
	configPath = flag.String("config", "config.json", "Path to configuration file containing the application's credentials.")

	accessToken oauth.Credentials

	oauthClient = oauth.Client {
		TemporaryCredentialRequestURI: "https://api.twitter.com/oauth/request_token",
		ResourceOwnerAuthorizationURI: "https://api.twitter.com/oauth/authorize",
		TokenRequestURI:               "https://api.twitter.com/oauth/access_token",
	}
)

func readConfig() error {
	b, err := ioutil.ReadFile(*configPath)
	if err != nil {
		return err
	}
	var config = struct {
		Consumer, Access *oauth.Credentials
	}{
		&oauthClient.Credentials, &accessToken,
	}
	return json.Unmarshal(b, &config)
}

func usage() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n %s keyword ...\n", os.Args[0], os.Args[0])
	flag.PrintDefaults()
	os.Exit(2)
}

func main() {
	//queryMutex := new(sync.Mutex)
	//queryMutex.Lock()
	termbox_err := termbox.Init()
	if termbox_err != nil {
		panic(termbox_err)
	}
	defer termbox.Close()
	termbox.SetInputMode(termbox.InputEsc)
	edit_box.InsertRune('#')

	redraw_all()
	query := ""
	stringSlices := make([]rune, 0)
	stringSlices = append(stringSlices, '#')
	queryloop: for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				break queryloop
			case termbox.KeySpace:
				edit_box.InsertRune(' ')
				stringSlices = append(stringSlices, ' ')
			case termbox.KeyEnter:
				query = string(stringSlices)
				break queryloop
			case termbox.KeyBackspace, termbox.KeyBackspace2:
				edit_box.DeleteRuneBackward()
				if len(stringSlices) > 0 {
					tempSlices := make([]rune, len(stringSlices) - 1)
					copy(tempSlices, stringSlices[0:])
					stringSlices[len(stringSlices) - 1] = 0
					stringSlices = tempSlices
				}
			default:
				if ev.Ch != 0 {
					edit_box.InsertRune(ev.Ch)
					stringSlices = append(stringSlices, ev.Ch)
				}
			}
		case termbox.EventError:
			panic(ev.Err)
		}
		redraw_all()
	}
	if query == "" || query == "#" || query == " " || query == "# " {
		fmt.Println("No query given. Nothing to search for. Program exiting...")
		//queryMutex.Unlock()
		return
	} else {
		termbox.Close()
		fmt.Print("Searching for: ", query, "\n")
		fmt.Println()
		//queryMutex.Unlock()
	}

	displayQueue := new(Queue)
	displayQueue = displayQueue.Init()
	newTweetsQueue := new(Queue)
	newTweetsQueue = newTweetsQueue.Init()
	oldTweetsQueue := new(Queue)
	oldTweetsQueue = oldTweetsQueue.Init()

	flag.Usage = usage
	flag.Parse()
	if err := readConfig(); err != nil {
		log.Fatalf("Error reading configuration, %v", err)
	}

	ts, err := twitterstream.Open(
		&oauthClient,
		&accessToken,
		"https://stream.twitter.com/1.1/statuses/filter.json",
		//url.Values{"track": {strings.Join(flag.Args(), ", ")}})
		url.Values{"track": {query}})
	if err != nil {
		log.Fatal(err)
	}
	defer ts.Close()

	// Loop until stream has a permanent error.
	for ts.Err() == nil {
		var t anaconda.Tweet
		if err := ts.UnmarshalNext(&t); err != nil {
			log.Fatal(err)
		}
		displaySize := displayQueue.GetSize()
		if displaySize < 10 {
			displayQueue.Enqueue(t)
			fmt.Print("@", t.User.ScreenName, ": ", t.Text, " (https://twitter.com/", t.User.ScreenName, "/status/", t.Id,")")
			fmt.Println()
			fmt.Println()
		} else {
			newTweetsSize := newTweetsQueue.GetSize()
			if newTweetsSize == 0 {
				fmt.Println("Display queue filled.")
				fmt.Println("Now filling new Tweets queue.")
				fmt.Println()
			}
			newTweetsQueue.Enqueue(t)
			fmt.Print("@", t.User.ScreenName, ": ", t.Text, " (https://twitter.com/", t.User.ScreenName, "/status/", t.Id,")")
			fmt.Println()
			fmt.Println()
		}
	}
	log.Print(ts.Err)
}