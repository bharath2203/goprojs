package prometheus

import "github.com/prometheus/client_golang/prometheus"

var messageReceiveCounter = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "apache_pulsar_message_processed_counter",
		Help: "Counts the number of messages processed from the apache pulsar",
	},
	[]string{"topic", "method"},
)

func init() {
	prometheus.MustRegister(messageReceiveCounter)
}

func IncrementCounter(topic string, method string) {
	messageReceiveCounter.WithLabelValues(topic, method).Inc()
}
