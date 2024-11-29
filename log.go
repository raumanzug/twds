package main

import (
	"golang.org/x/net/webdav"
	"iter"
	"log"
	"maps"
	"net/http"
	"slices"
	"strings"
)

func getHeaderStringIter(pHeader *http.Header) iter.Seq[string] {
	return func(yield func(string) bool) {
		for _, k := range slices.Sorted(maps.Keys(*pHeader)) {
			vs := (*pHeader)[k]
			headerItemString := k + ":" + strings.Join(vs, "; ")
			if !yield(headerItemString) {
				return
			}
		}
	}
}

func access_logger(pH *webdav.Handler, pReq *http.Request, err error) {
	errString := ""
	if err != nil {
		errString += "; ERROR: " + err.Error()
	}

	headerStringIter := getHeaderStringIter(&pReq.Header)
	headerString := strings.Join(slices.Collect(headerStringIter), " ")

	log.Printf("%s %s %s; %s => %s; %s %s\n",
		pReq.Method,
		pReq.URL.String(),
		pReq.Proto,
		pReq.RemoteAddr,
		pReq.Host,
		headerString,
		errString)
}
