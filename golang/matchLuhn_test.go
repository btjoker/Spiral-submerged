package main

import "testing"

func Test_matchLuhn(t *testing.T) {
	type item struct {
		name string
		args string
		want bool
	}

	items := []item{
		item{
			name: "无效卡号",
			args: "655002000100000328",
			want: false,
		},
		item{
			name: "有效卡号",
			args: "6550020001000003287",
			want: true,
		},
	}

	for _, tt := range items {
		t.Run(tt.name, func(t *testing.T) {
			if got := matchLuhn(tt.args); got != tt.want {
				t.Errorf("matchLuhn() = %v, want %s", got, tt.want)
			}
		})
	}
}

func Benchmark_matchLuhn(b *testing.B) {
	b.StopTimer()
	item := "6550020001000003287"
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		matchLuhn(item)
	}
}
