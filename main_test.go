package main

import "testing"

func BenchmarkDedicatedW100R100Buf0(b *testing.B) {
	for n := 0; n < b.N; n++ {
		dedicated(100, 100, 0)
	}
}

func BenchmarkDedicatedW100R100Buf1000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		dedicated(100, 100, 1000)
	}
}

func BenchmarkSynchroW100R100(b *testing.B) {
	for n := 0; n < b.N; n++ {
		synchro(100, 100)
	}
}

func BenchmarkDedicatedW10R1000Buf0(b *testing.B) {
	for n := 0; n < b.N; n++ {
		dedicated(10, 1000, 0)
	}
}
func BenchmarkDedicatedW10R1000Buf1000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		dedicated(10, 1000, 1000)
	}
}

func BenchmarkSynchroW10R1000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		synchro(10, 1000)
	}
}
