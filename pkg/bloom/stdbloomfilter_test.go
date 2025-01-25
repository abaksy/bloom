package bloom

import (
	"testing"
)

func TestGetBitArraySize(t *testing.T) {
	type args struct {
		N int
		P float64
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "small set",
			args: args{
				N: 100,
				P: 0.01,
			},
			want:    958,
			wantErr: false,
		},
		{
			name: "medium set",
			args: args{
				N: 1000,
				P: 0.05,
			},
			want:    6235,
			wantErr: false,
		},
		{
			name: "large set",
			args: args{
				N: 10000,
				P: 0.1,
			},
			want:    47925,
			wantErr: false,
		},
		{
			name: "edge case - invalid P",
			args: args{
				N: 100,
				P: 1.1,
			},
			want:    -1,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetBitArraySize(tt.args.N, tt.args.P)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBitArraySize() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetBitArraySize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetNumHashFunctions(t *testing.T) {
	type args struct {
		N int
		M int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "standard",
			args: args{
				N: 1000,
				M: 10000,
			},
			want:    6,
			wantErr: false,
		},
		{
			name: "small N",
			args: args{
				N: 100,
				M: 1000,
			},
			want:    6,
			wantErr: false,
		},
		{
			name: "large N",
			args: args{
				N: 10000,
				M: 100000,
			},
			want:    6,
			wantErr: false,
		},
		{
			name: "edge case - minimal size",
			args: args{
				N: 100,
				M: 100,
			},
			want:    1, // Should return at least 1
			wantErr: false,
		},
		{
			name: "edge case - very large ratio",
			args: args{
				N: 10,
				M: 10000,
			},
			want:    693, // ln(2) * (m/n)
			wantErr: false,
		},
		{
			name: "edge case - N = 0",
			args: args{
				N: 0,
				M: 100,
			},
			want:    -1,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetNumHashFunctions(tt.args.N, tt.args.M)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetNumHashFunctions() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetNumHashFunctions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewStandardBloomFilter(t *testing.T) {
	type args struct {
		N int
		P float64
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "small N",
			args: args{
				N: 100,
				P: 0.01,
			},
			wantErr: false,
		},
		{
			name: "medium N",
			args: args{
				N: 1000,
				P: 0.001,
			},
			wantErr: false,
		},
		{
			name: "large N",
			args: args{
				N: 10000,
				P: 0.0001,
			},
			wantErr: false,
		},
		{
			name: "edge case - invalid N",
			args: args{
				N: 0,
				P: 0.01,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bf, err := NewStandardBloomFilter(tt.args.N, tt.args.P)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewStandardBloomFilter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if bf == nil {
				if !tt.wantErr {
					t.Fatal("Expected non-nil BloomFilter")
				} else {
					return // nil value was expected, no need to run further tests
				}
			}

			if bf.size != 0 {
				t.Errorf("Expected initial size 0, got %d", bf.size)
			}
			if len(bf.Arr) == 0 {
				t.Error("Expected non-zero length bit array")
			}
			if bf.K == 0 {
				t.Error("Expected non-zero number of hash functions")
			}
		})
	}
}

func TestStandardBloomFilter_GetHashIndex(t *testing.T) {
	type args struct {
		value string
		idx   int
	}
	tests := []struct {
		name string
		b    *StandardBloomFilter
		args args
		want uint64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.GetHashIndex(tt.args.value, tt.args.idx); got != tt.want {
				t.Errorf("StandardBloomFilter.GetHashIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStandardBloomFilter_Add(t *testing.T) {
	type args struct {
		ele interface{}
	}
	tests := []struct {
		name    string
		b       *StandardBloomFilter
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.b.Add(tt.args.ele); (err != nil) != tt.wantErr {
				t.Errorf("StandardBloomFilter.Add() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStandardBloomFilter_Contains(t *testing.T) {
	type args struct {
		ele interface{}
	}
	tests := []struct {
		name string
		b    *StandardBloomFilter
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.Contains(tt.args.ele); got != tt.want {
				t.Errorf("StandardBloomFilter.Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStandardBloomFilter_Clear(t *testing.T) {
	tests := []struct {
		name    string
		b       *StandardBloomFilter
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.b.Clear(); (err != nil) != tt.wantErr {
				t.Errorf("StandardBloomFilter.Clear() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
