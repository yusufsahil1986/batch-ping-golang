# Go Batch Pinger

A simple command-line utility written in Go to perform concurrent pings on multiple IP addresses or hostnames.

## Features

*   **Batch Pinging:** Ping a list of sites provided as a comma-separated string.
*   **Concurrency:** Pings are performed concurrently using Go routines for efficiency.
*   **Clear Output:** Displays success or failure for each ping attempt, including duration.

## Prerequisites

*   [Go](https://golang.org/doc/install) (version 1.16 or higher recommended) installed on your system.
*   A `ping` command available in your system's PATH (common on Linux/macOS, Windows has `ping.exe`).

## How to Build

Navigate to the project directory and build the executable:

```bash
go build -o batch_ping main.go
```

This will create an executable named `batch_ping` (or `batch_ping.exe` on Windows) in the current directory.

## How to Run

Execute the `batch_ping` command, providing a comma-separated list of IP addresses or hostnames using the `--sites` flag:

```bash
./batch_ping --sites "1.2.1.3,8.8.8.8,127.0.0.1,google.com"
```

Replace `"1.2.1.3,8.8.8.8,127.0.0.1,google.com"` with the actual sites you wish to ping.

## Example Output

```
Pinging 4 sites...
🟢 google.com: Success (Duration: 3.053s)
🟢 127.0.0.1: Success (Duration: 3.06s)
🟢 8.8.8.8: Success (Duration: 4.028s)
🔴 1.2.1.3: Failed (Exit Status: 1) -
Batch ping completed.
```
