package main

import "testing"

func BenchmarkPortScanner(b *testing.B) {
	for b.Loop() {
		ScanIp("127.0.0.1") //Benchmarking with localhost
	}
}
