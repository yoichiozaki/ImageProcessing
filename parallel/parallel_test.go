package parallel

import "testing"

func TestParallelExecution(t *testing.T) {
	for n := 0; n < 1024; n++ {
		data := make([]bool, n)

		Line(len(data), func(start, end int) {
			for i := start; i < end; i++ {
				data[i] = !data[i]
			}
		})

		for i := range data {
			if data[i] != true {
				t.Errorf("TestParallelExecution failed. Failure at n = %v", n)
			}
		}
	}
}