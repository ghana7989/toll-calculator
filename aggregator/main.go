package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"

	"github.com/ghana7989/toll-calculator/types"
	"github.com/sirupsen/logrus"
)

func main() {
	listenAddr := flag.String("listen", ":3000", "The listen address of the http server")
	flag.Parse()
	store := NewMemoryStore()
	svc := NewInvoiceAggregator(store)
	svc = NewLogMiddleware(svc)
	makeHttpTransport(*listenAddr, svc)
}
func makeHttpTransport(listenAddr string, svc Aggregator) {
	fmt.Println("Starting HTTP transport on ", listenAddr)
	http.HandleFunc("/aggregate", handleAggregate(svc))
	http.ListenAndServe(listenAddr, nil)
}
func handleAggregate(svc Aggregator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var distance types.Distance
		if err := json.NewDecoder(r.Body).Decode(&distance); err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]string{
				"error": err.Error(),
			})
			return
		}
		if err := svc.AggregateDistance(distance); err != nil {
			logrus.WithError(err).Error("Failed to aggregate distance")
			writeJSON(w, http.StatusInternalServerError, map[string]string{
				"error": err.Error(),
			})
			return
		}
	}
}

func writeJSON(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}
