package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Define network-related metrics
var (
	networkTrafficReceived = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "network_traffic_bytes_received_total",
			Help: "Total bytes received on the network interface",
		},
		[]string{"interface"},
	)
	networkTrafficTransmitted = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "network_traffic_bytes_transmitted_total",
			Help: "Total bytes transmitted on the network interface",
		},
		[]string{"interface"},
	)
	networkPacketsReceived = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "network_packets_received_total",
			Help: "Total packets received on the network interface",
		},
		[]string{"interface"},
	)
	networkPacketsTransmitted = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "network_packets_transmitted_total",
			Help: "Total packets transmitted on the network interface",
		},
		[]string{"interface"},
	)
	networkPacketsReceivedErrors = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "network_packets_received_errors_total",
			Help: "Total packets received errors on the network interface",
		},
		[]string{"interface"},
	)
	networkPacketsTransmittedErrors = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "network_packets_transmitted_errors_total",
			Help: "Total packets transmitted errors on the network interface",
		},
		[]string{"interface"},
	)
	networkPacketsDroppedReceived = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "network_packets_dropped_received_total",
			Help: "Total dropped packets received on the network interface",
		},
		[]string{"interface"},
	)
	networkPacketsDroppedTransmitted = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "network_packets_dropped_transmitted_total",
			Help: "Total dropped packets transmitted on the network interface",
		},
		[]string{"interface"},
	)
	networkTcpConnectionsEstablished = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "network_tcp_connections_established",
			Help: "Number of established TCP connections",
		},
	)
	networkPingLatency = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "network_ping_latency_seconds",
			Help: "Ping round-trip time in seconds",
		},
	)
	networkInterfaceSpeed = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "network_interface_speed_bits_per_second",
			Help: "Speed of network interface in bits per second",
		},
		[]string{"interface"},
	)
)

func init() {
	// Register metrics
	prometheus.MustRegister(networkTrafficReceived)
	prometheus.MustRegister(networkTrafficTransmitted)
	prometheus.MustRegister(networkPacketsReceived)
	prometheus.MustRegister(networkPacketsTransmitted)
	prometheus.MustRegister(networkPacketsReceivedErrors)
	prometheus.MustRegister(networkPacketsTransmittedErrors)
	prometheus.MustRegister(networkPacketsDroppedReceived)
	prometheus.MustRegister(networkPacketsDroppedTransmitted)
	prometheus.MustRegister(networkTcpConnectionsEstablished)
	prometheus.MustRegister(networkPingLatency)
	prometheus.MustRegister(networkInterfaceSpeed)
}

func main() {
	// Expose metrics at /metrics endpoint
	http.Handle("/metrics", promhttp.Handler())
	go func() {
		// Bind to localhost explicitly (ensure it's only accessible via localhost)
		if err := http.ListenAndServe("localhost:8080", nil); err != nil {
			fmt.Println("Error starting server: ", err)
			os.Exit(1)
		}
	}()
	
		// Collect network metrics every 10 seconds
	for {
		collectNetworkMetrics()
		time.Sleep(10 * time.Second)
	}
}

func collectNetworkMetrics() {
	// Get network interfaces
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("Error fetching network interfaces: ", err)
		return
	}

	for _, iface := range interfaces {
		// Skip loopback interface
		if strings.HasPrefix(iface.Name, "lo") {
			continue
		}

		// Fetch and collect the network stats for the interface
		collectInterfaceStats(iface.Name)
	}
}

func collectInterfaceStats(interfaceName string) {
	// Here, you can use `sysctl`, `netstat`, or `procfs` for fetching stats
	// For simplicity, I will simulate fetching these values.

	// Simulate network traffic (bytes)
	networkTrafficReceived.WithLabelValues(interfaceName).Add(1024)
	networkTrafficTransmitted.WithLabelValues(interfaceName).Add(2048)

	// Simulate packet stats
	networkPacketsReceived.WithLabelValues(interfaceName).Add(100)
	networkPacketsTransmitted.WithLabelValues(interfaceName).Add(200)
	networkPacketsReceivedErrors.WithLabelValues(interfaceName).Add(5)
	networkPacketsTransmittedErrors.WithLabelValues(interfaceName).Add(10)
	networkPacketsDroppedReceived.WithLabelValues(interfaceName).Add(1)
	networkPacketsDroppedTransmitted.WithLabelValues(interfaceName).Add(2)

	// Simulate TCP connections
	networkTcpConnectionsEstablished.Add(50)

	// Simulate Ping latency
	networkPingLatency.Set(0.05)

	// Simulate interface speed
	networkInterfaceSpeed.WithLabelValues(interfaceName).Set(1000000000) // 1 Gbps
}

