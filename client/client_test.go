package client

import (
	"fmt"
	"sync"
	"testing"
)

func BenchmarkDefaultClientRequest(b *testing.B) {
	table := []struct {
		input int
	}{
		{input: 100},
		{input: 1000},
	}

	for _, v := range table {
		b.Run(fmt.Sprintf("input_size_%d", v.input), func(b *testing.B) {
			var wg sync.WaitGroup
			wg.Add(v.input)
			for i := 0; i < v.input; i++ {
				go func() {
					DefaultClientRequest()
					wg.Done()
				}()
			}
			wg.Wait()
		})
	}

}

func BenchmarkOverridingDefaultClientRequest(b *testing.B) {
	table := []struct {
		input int
	}{
		{input: 100},
		{input: 1000},
	}

	for _, v := range table {
		b.Run(fmt.Sprintf("input_size_%d", v.input), func(b *testing.B) {
			var wg sync.WaitGroup
			wg.Add(v.input)
			for i := 0; i < v.input; i++ {
				go func() {
					OverridingDefaultClientRequest()
					wg.Done()
				}()
			}
			wg.Wait()
		})
	}

}

func BenchmarkOverridingDefaultClientRequestWithIdleConnTimeout(b *testing.B) {
	table := []struct {
		input int
	}{
		{input: 100},
		{input: 1000},
	}

	for _, v := range table {
		b.Run(fmt.Sprintf("input_size_%d", v.input), func(b *testing.B) {
			var wg sync.WaitGroup
			wg.Add(v.input)
			for i := 0; i < v.input; i++ {
				go func() {
					OverridingDefaultClientRequestWithIdleConnTimeout()
					wg.Done()
				}()
			}
			wg.Wait()
		})
	}

}
