package main

import (
	"bytes"
	"flag"
	"fmt"

	"os/exec"
	"strings"
	"sync"
	"time"
)

func main() {
	sitesFlag := flag.String("sites", "", "Comma-separated list of IP addresses or hostnames to ping")
	flag.Parse()

	if *sitesFlag == "" {
		fmt.Println("Usage: go run main.go --sites \"<ip1>,<ip2>,...\"")
		return
	}

	sites := strings.Split(*sitesFlag, ",")
	var wg sync.WaitGroup

	fmt.Printf("Pinging %d sites...\n", len(sites))

	for _, site := range sites {
		wg.Add(1)
		go func(s string) {
			defer wg.Done()
			pingSite(s)
		}(strings.TrimSpace(site))
	}

	wg.Wait()
	fmt.Println("Batch ping completed.")
}

func pingSite(site string) {
	cmd := exec.Command("ping", "-c", "4", "-W", "1", site) // -c 4: 4 packets, -W 1: 1 second timeout
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	start := time.Now()
	err := cmd.Run()
	duration := time.Since(start)

	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			fmt.Printf("🔴 %s: Failed (Exit Status: %d) - %s\n", site, exitError.ExitCode(), strings.TrimSpace(stderr.String()))
		} else {
			fmt.Printf("🔴 %s: Failed (Error: %v) - %s\n", site, err, strings.TrimSpace(stderr.String()))
		}
		return
	}

	// Basic parsing to check for successful ping, looking for "time=" in output
	if strings.Contains(stdout.String(), "time=") {
		fmt.Printf("🟢 %s: Success (Duration: %v)\n", site, duration.Round(time.Millisecond))
	} else {
		// Sometimes ping might exit 0 but still not reach the host (e.g., no route)
		fmt.Printf("🟡 %s: Unreachable or no reply (Duration: %v)\n", site, duration.Round(time.Millisecond))
	}
}
