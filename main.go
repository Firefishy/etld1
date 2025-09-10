package main

import (
	"bufio"
	"fmt"
	"net/url"
	"os"
	"strings"

	"golang.org/x/net/publicsuffix"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		if line == "" {
			continue
		}

		host := getHost(line)
		if host == "" {
			continue // skip invalid
		}

		etld1, err := publicsuffix.EffectiveTLDPlusOne(host)
		if err != nil {
			continue // skip if cannot determine
		}
		fmt.Println(etld1)
	}
	if err := sc.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "read error:", err)
	}
}

func getHost(s string) string {
	// If it looks like a URL
	if strings.Contains(s, "://") {
		u, err := url.Parse(s)
		if err == nil {
			return u.Hostname()
		}
	}
	// Try parsing as host by prefixing scheme
	u, err := url.Parse("http://" + s)
	if err == nil {
		return u.Hostname()
	}
	return ""
}
