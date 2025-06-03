package main

// import (
// 	"context"
// 	"fmt"
// 	"io"
// 	"net"
// 	"net/http"
// 	"time"
// )

// func main() {
// 	// Create a proper HTTP client with comprhesive and timeout settings
// 	client := &http.Client{
// 		//! Overall timeot for the entire request lifecycle
// 		Timeout: 30 * time.Second,

// 		//! Transport settings for connection management
// 		Transport: &http.Transport{
// 			// !Timeout for establishing a  TCP connection
// 			DialContext: (&net.Dialer{
// 				Timeout:   10 * time.Second, //! Timeout for establishing a TCP connection
// 				KeepAlive: 30 * time.Second, //! // TCP health checks every 30s,like hearthbeat
// 			}).DialContext,

// 			//! Timeout for TLS handshake (important for HTTPS requests)
// 			TLSHandshakeTimeout: 10 * time.Second,

// 			//! Timeout for reading the response headers
// 			ResponseHeaderTimeout: 10 * time.Second,

// 			//! Timeout for waiting on 100-continue response
// 			ExpectContinueTimeout: 1 * time.Second,

// 			//! How long idle connections stay in pool
// 			IdleConnTimeout: 90 * time.Second,

// 			//! Connection pool settings
// 			MaxIdleConns:        100,
// 			MaxIdleConnsPerHost: 10,
// 		},
// 	}

// 	// !First API call
// 	fmt.Println("=== Making request to JSONPlaceholder ===")
// 	makeRequest(client, "https://jsonplaceholder.typicode.com/posts/1")

// 	fmt.Println("\n=== Making request to Dog API ===")
// 	makeRequest(client, "https://dogapi.dog/api/v2/breeds")
// }

// func makeRequest(client *http.Client, url string) {
// 	// Create a context with a timeout for the request
// 	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
// 	defer cancel() // Always cancel the context to avoid leaks

// 	// Create request with context
// 	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)

// 	if err != nil {
// 		fmt.Printf("Error creating request: %v\n", err)
// 		return
// 	}

// 	// Set headers if needed
// 	req.Header.Set("User-Agent", "MyCustomClient/1.0")

// 	// Execute the request
// 	fmt.Printf("Making request to %s\n", url)
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		fmt.Printf("Error making request: %v\n", err)
// 		return
// 	}

// 	// Always close the response body to free resources
// 	defer func() {
// 		if closeErr := resp.Body.Close(); closeErr != nil {
// 			fmt.Printf("Error closing response body: %v\n", closeErr)
// 		}
// 	}()

// 	//Check HTTP status code
// 	if resp.StatusCode != http.StatusOK {
// 		fmt.Printf("Received non-200 response: %d %s\n", resp.StatusCode, http.StatusText(resp.StatusCode))
// 		return
// 	}

// 	//Read response body with size limit to prevent memory issues
// 	const maxResponseBodySize = 1024 * 1024 // 1 MB limit
// 	limitedReader := io.LimitReader(resp.Body, maxResponseBodySize)

// 	body, err := io.ReadAll(limitedReader)

// 	if err != nil {
// 		fmt.Printf("Error reading response body: %v\n", err)
// 		return
// 	}

// 	fmt.Printf("Response Status: %d %s\n", resp.StatusCode, http.StatusText(resp.StatusCode))
// 	fmt.Printf("Content-Type : %s\n", resp.Header.Get("Content-Type"))
// 	fmt.Printf("Content-Length: %d bytes\n", len(body))
// 	fmt.Printf("Response body preview (first 200 chars):\n%s\n", truncateString(string(body), 200))
// }

// // !Helper function to truncate strings for logging
// func truncateString(s string, maxLen int) string {
// 	if len(s) <= maxLen {
// 		return s
// 	}
// 	return s[:maxLen] + "..."
// }


