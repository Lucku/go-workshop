package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSplitHostPortFromURL(t *testing.T) {

	// in: "192.168.0.1:8080" 	out1: "192.168.0.1" 	out2: 8080
	// in: "192.168.0.1:hey"	out1: "192.168.0.1" 	out2: -1
	// in: ""					out1: "" 				out2: -1
	// in: ":8080"				out1: ""				out2: 8080

	tests := []struct {
		giveURL  string
		wantHost string
		wantPort int
	}{
		{
			giveURL: "192.168.0.1:8080",
			wantHost: "192.168.0.1",
			wantPort: 8080,
		},
		{
			giveURL: "192.168.0.1:hey",
			wantHost: "192.168.0.1",
			wantPort: -1,
		},
		{
			giveURL: "",
			wantHost: "",
			wantPort: -1,
		},
		{
			giveURL: ":8080",
			wantHost: "",
			wantPort: 8080,
		},
	}

	for _, tt := range tests {

		t.Run(tt.giveURL, func(t *testing.T) {

			actualHost, actualPort := splitHostPortFromURL(tt.giveURL)

			assert.Equal(t, tt.wantHost, actualHost)
			assert.Equal(t, tt.wantPort, actualPort)
		})
	}
}

func BenchmarkSplitHostPortFromURL(b *testing.B) {

	url := "192.168.100:8080"

	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		splitHostPortFromURL(url)
	}
}