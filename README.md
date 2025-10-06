# 🌐 Network Metrics Exporter

A **Go-based network metrics exporter** that collects and exposes real-time network performance data — including packet statistics, connection status, throughput, and error rates — through a local HTTP endpoint.  
This tool enables easy integration with monitoring systems like **Prometheus** and **Node Exporter**, providing deep insights into network health and performance.

---

## 🚀 Features

### 📊 Real-Time Network Metrics
- Collects key network metrics such as:
  - Packet counts (sent, received, dropped)
  - Active connections
  - Throughput (upload/download rates)
  - Error and retransmission rates

### ⚙️ Lightweight & Efficient
- Built in **Go**, leveraging **system APIs** (`netstat`, `ifstat`, and `/proc/net/`) for low-overhead data collection.  
- Optimized for minimal CPU and memory usage.

### 🌍 HTTP Metrics Endpoint
- Exposes collected metrics through a **local HTTP server**, compatible with:
  - **Prometheus** scrapers
  - **Node Exporter text format**
- Easy integration with existing monitoring setups.

### 🧩 Prometheus & Node Exporter Compatibility
- Designed to complement **Node Exporter**, allowing seamless collection of network metrics alongside CPU, memory, and disk stats.

### 💻 Local Web Interface
- Simple HTTP endpoint for quick visualization of current network statistics.  
- Provides a straightforward, developer-friendly way to monitor system network performance.

---


## 🧰 Tech Stack

- **Language:** Go (Golang)  
- **APIs Used:** netstat, ifstat, procfs  
- **Output Format:** Prometheus / Node Exporter-compatible text  
- **Interface:** HTTP server (localhost)

---

## ⚡ Getting Started

### Prerequisites
- Go 1.21 or higher
- Linux or macOS environment with `/proc` filesystem

### Installation
```bash
- git clone https://github.com/sarthak21-negi/network-metrics-exporter.git

- cd network-metrics-exporter

- go mod tidy

```

---

### Run the Exporter
```bash
go run main.go
```

---

### Access Metrics
```bash
http://localhost:9090/metrics
```
## 🧠 Architecture Overview

