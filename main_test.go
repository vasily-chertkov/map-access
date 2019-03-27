package main

import "testing"

func BenchmarkDedicatedW100R100(b *testing.B) {
	for n := 0; n < b.N; n++ {
		dedicated(100, 100)
	}
}

func BenchmarkSynchroW100R100(b *testing.B) {
	for n := 0; n < b.N; n++ {
		synchro(100, 100)
	}
}

func BenchmarkDedicatedW10R1000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		dedicated(100, 100)
	}
}

func BenchmarkSynchroW10R1000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		synchro(100, 100)
	}
}
