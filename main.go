package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	hostname, err := os.Hostname()
	if err != nil || hostname == "" {
		hostname = "unknown-container"
	}

	mux := http.NewServeMux()

	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		serverTime := time.Now().UTC().Format("2006-01-02 15:04:05 UTC")
		processingDuration := time.Since(start)

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		
		fmt.Fprintf(w, `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>K3S Demo</title>
    <style>
        body {
            background-color: #000000;
            color: #ededed;
            font-family: ui-monospace, SFMono-Regular, Menlo, Monaco, Consolas, monospace;
            display: flex;
            justify-content: center;
            align-items: center;
            height: 100vh;
            margin: 0;
        }
        .container {
            border: 1px solid #222222;
            padding: 2rem;
            max-width: 420px;
            width: 100%%;
            background-color: #050505;
        }
        h2 {
            margin: 0 0 0.5rem 0;
            font-size: 0.95rem;
            text-transform: uppercase;
            letter-spacing: 0.15em;
            color: #ffffff;
        }
        .subtitle {
            color: #666666;
            font-size: 0.75rem;
            margin-bottom: 1.5rem;
            letter-spacing: 0.05em;
        }
        ul {
            list-style: none;
            padding: 0;
            margin: 0;
        }
        li {
            display: flex;
            justify-content: space-between;
            margin: 0.75rem 0;
            font-size: 0.8rem;
            border-bottom: 1px solid #111111;
            padding-bottom: 0.5rem;
        }
        .label {
            color: #777777;
        }
        .value {
            color: #00ff66;
        }
    </style>
</head>
<body>
    <div class="container">
        <h2>K3S Demo</h2>
        <div class="subtitle">This app is running in a container.</div>
        <ul>
            <li><span class="label">Container ID:</span> <span class="value">%s</span></li>
            <li><span class="label">Server Time:</span> <span class="value">%s</span></li>
            <li><span class="label">Processing:</span> <span class="value">%s</span></li>
        </ul>
    </div>
</body>
</html>`, hostname, serverTime, processingDuration)
	})

	port := ":8080"
	fmt.Printf("Server started on port %s...\n", port)
	if err := http.ListenAndServe(port, mux); err != nil {
		panic(err)
	}
}
