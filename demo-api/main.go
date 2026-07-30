package main

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var requestsTotal = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "demo_api_requests_total",
		Help: "Total de requisicoes recebidas pela aplicacao.",
	},
)

func handler(w http.ResponseWriter, r *http.Request) {
	requestsTotal.Inc()

	fmt.Fprintln(w, "Hello DevOps!")
}

func main() {
	prometheus.MustRegister(requestsTotal)

	http.HandleFunc("/", handler)
	http.Handle("/metrics", promhttp.Handler())

	fmt.Println("Servidor iniciado na porta 8080")

	http.ListenAndServe(":8080", nil)
}