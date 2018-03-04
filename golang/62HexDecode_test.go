package main

import (
	"testing"
)

func Test_ReverseString(t *testing.T) {
	type item struct {
		name string
		want string
		args []byte
	}

	tests := []item{
		item{
			name: "英文",
			want: "abcdefghijkmlnopqrstuvwxyz",
			args: []byte("zyxwvutsrqponlmkjihgfedcba"),
		},
		item{
			name: "数字",
			want: "0123456789",
			args: []byte("9876543210"),
		},
		item{
			name: "符号",
			want: "~!@#$%^&*()_",
			args: []byte("_)(*&^%$#@!~"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseString(tt.args); got != tt.want {
				t.Errorf("ReverseString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Base10ToBase62(t *testing.T) {
	type item struct {
		name string
		want string
		args int64
	}

	tests := []item{
		item{
			name: "恶臭数字",
			want: "1fuTIX",
			args: 1145141919,
		},
		item{
			name: "唯一qq",
			want: "mRz86",
			args: 337845818,
		},
		item{
			name: "电话",
			want: "vgAiJjUY",
			args: 110112114119120,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Base10ToBase62(tt.args); got != tt.want {
				t.Errorf("Base10ToBase62() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Base62ToBase10(t *testing.T) {
	type item struct {
		name string
		args string
		want int64
	}
	tests := []item{
		item{
			name: "golang",
			args: "golang",
			want: 15017802146,
		},
		item{
			name: "python",
			args: "python",
			want: 23412694595,
		},
		item{
			name: "acfun",
			args: "acfun",
			want: 150682839,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := Base62ToBase10(tt.args); gotResult != tt.want {
				t.Errorf("Base62ToBase10() = %v, want %v", gotResult, tt.want)
			}
		})
	}
}

func Benchmark_ReverseString(b *testing.B) {
	b.StopTimer()
	item := []byte("abcdefghijkmlnopqrstuvwxyz")
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		ReverseString(item)
	}
}

func Benchmark_Base62ToBase10(b *testing.B) {
	b.StopTimer()
	item := "golang"
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		Base62ToBase10(item)
	}
}

func Benchmark_Base10ToBase62(b *testing.B) {
	b.StopTimer()
	var item int64 = 337845818
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		Base10ToBase62(item)
	}
}
