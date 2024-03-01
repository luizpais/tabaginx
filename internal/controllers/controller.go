package controllers

import (
	"fmt"
	"github.com/luizpais/tabaginx/internal/models"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync/atomic"
)

type Controller struct {
	destinations []*url.URL
	counter      uint32
}

func (c *Controller) Init(config models.Config) {
	mux := http.NewServeMux()

	// Initialize the list of destination URLs
	c.destinations = make([]*url.URL, len(config.Destinations))
	for i, destination := range config.Destinations {
		c.destinations[i] = mustParseURL(destination)
	}

	debugRequest := config.Debug
	debugRequestBody := config.DebugReqBody

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Get the current destination and increment the counter
		destination := c.destinations[c.counter%uint32(len(c.destinations))]
		atomic.AddUint32(&c.counter, 1)

		proxy := httputil.NewSingleHostReverseProxy(destination)

		if debugRequest {
			requestDump, err := httputil.DumpRequest(r, debugRequestBody)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(string(requestDump))
			}
		}

		proxy.ServeHTTP(w, r)
	})

	fmt.Printf("Listening on port %s\n", config.Port)
	port := fmt.Sprintf(":%d", config.Port)
	if err := http.ListenAndServe(port, mux); err != nil {
		panic(err)
	}

}

func mustParseURL(s string) *url.URL {
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}
	return u
}
