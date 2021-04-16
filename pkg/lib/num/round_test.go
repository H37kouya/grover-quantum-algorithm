package num

import (
	"strconv"
	"testing"
)

func TestRound(t *testing.T) {
	type args struct {
		num    float64
		places float64
	}
	tests := []struct {
		args args
		want float64
	}{
		{args: args{num: 20.0, places: 0}, want: 20},
		{args: args{num: 20.1, places: 0}, want: 20},
		{args: args{num: 20.4, places: 0}, want: 20},
		{args: args{num: 20.5, places: 0}, want: 21},
		{args: args{num: 20.6, places: 0}, want: 21},
		{args: args{num: 20.9, places: 0}, want: 21},
		{args: args{num: 10.00, places: 1}, want: 10.0},
		{args: args{num: 10.04, places: 1}, want: 10.0},
		{args: args{num: 10.05, places: 1}, want: 10.1},
		{args: args{num: 10.06, places: 1}, want: 10.1},
		{args: args{num: 1000, places: -2}, want: 1000},
		{args: args{num: 1001, places: -2}, want: 1000},
		{args: args{num: 1049, places: -2}, want: 1000},
		{args: args{num: 1050, places: -2}, want: 1100},
		{args: args{num: 1051, places: -2}, want: 1100},
		{args: args{num: 1060, places: -2}, want: 1100},
	}
	for _, tt := range tests {
		numStr := strconv.FormatFloat(tt.args.num, 'f', -1, 64)
		placesStr := strconv.FormatFloat(tt.args.places, 'f', -1, 64)
		wantStr := strconv.FormatFloat(tt.want, 'f', -1, 64)
		t.Run(
			numStr+"を"+placesStr+"桁で四捨五入すると"+wantStr+"である",
			func(t *testing.T) {
				if got := Round(tt.args.num, tt.args.places); got != tt.want {
					t.Errorf("Round() = %v, want %v", got, tt.want)
				}
			})
	}
}

func TestRoundDown(t *testing.T) {
	type args struct {
		num    float64
		places float64
	}
	tests := []struct {
		args args
		want float64
	}{
		{args: args{num: 20.0, places: 0}, want: 20},
		{args: args{num: 20.1, places: 0}, want: 20},
		{args: args{num: 20.4, places: 0}, want: 20},
		{args: args{num: 20.5, places: 0}, want: 20},
		{args: args{num: 20.6, places: 0}, want: 20},
		{args: args{num: 20.9, places: 0}, want: 20},
		{args: args{num: 10.00, places: 1}, want: 10.0},
		{args: args{num: 10.04, places: 1}, want: 10.0},
		{args: args{num: 10.05, places: 1}, want: 10.0},
		{args: args{num: 10.06, places: 1}, want: 10.0},
		{args: args{num: 1000, places: -2}, want: 1000},
		{args: args{num: 1001, places: -2}, want: 1000},
		{args: args{num: 1049, places: -2}, want: 1000},
		{args: args{num: 1050, places: -2}, want: 1000},
		{args: args{num: 1051, places: -2}, want: 1000},
		{args: args{num: 1060, places: -2}, want: 1000},
	}
	for _, tt := range tests {
		numStr := strconv.FormatFloat(tt.args.num, 'f', -1, 64)
		placesStr := strconv.FormatFloat(tt.args.places, 'f', -1, 64)
		wantStr := strconv.FormatFloat(tt.want, 'f', -1, 64)
		t.Run(
			numStr+"を"+placesStr+"桁で切り捨てすると"+wantStr+"である",
			func(t *testing.T) {
				if got := RoundDown(tt.args.num, tt.args.places); got != tt.want {
					t.Errorf("RoundDown() = %v, want %v", got, tt.want)
				}
			})
	}
}

func TestRoundUp(t *testing.T) {
	type args struct {
		num    float64
		places float64
	}
	tests := []struct {
		args args
		want float64
	}{
		{args: args{num: 20.0, places: 0}, want: 21},
		{args: args{num: 20.1, places: 0}, want: 21},
		{args: args{num: 20.4, places: 0}, want: 21},
		{args: args{num: 20.5, places: 0}, want: 21},
		{args: args{num: 20.6, places: 0}, want: 21},
		{args: args{num: 20.9, places: 0}, want: 21},
		{args: args{num: 10.00, places: 1}, want: 10.1},
		{args: args{num: 10.04, places: 1}, want: 10.1},
		{args: args{num: 10.05, places: 1}, want: 10.1},
		{args: args{num: 10.06, places: 1}, want: 10.1},
		{args: args{num: 1000, places: -2}, want: 1100},
		{args: args{num: 1001, places: -2}, want: 1100},
		{args: args{num: 1049, places: -2}, want: 1100},
		{args: args{num: 1050, places: -2}, want: 1100},
		{args: args{num: 1051, places: -2}, want: 1100},
		{args: args{num: 1060, places: -2}, want: 1100},
	}
	for _, tt := range tests {
		numStr := strconv.FormatFloat(tt.args.num, 'f', -1, 64)
		placesStr := strconv.FormatFloat(tt.args.places, 'f', -1, 64)
		wantStr := strconv.FormatFloat(tt.want, 'f', -1, 64)
		t.Run(
			numStr+"を"+placesStr+"桁で切り上げすると"+wantStr+"である",
			func(t *testing.T) {
				if got := RoundUp(tt.args.num, tt.args.places); got != tt.want {
					t.Errorf("RoundUp() = %v, want %v", got, tt.want)
				}
			})
	}
}

func Test_roundInt(t *testing.T) {
	type args struct {
		num float64
	}
	tests := []struct {
		args args
		want float64
	}{
		{args: args{num: 20.0}, want: 20},
		{args: args{num: 20.1}, want: 20},
		{args: args{num: 20.4}, want: 20},
		{args: args{num: 20.5}, want: 21},
		{args: args{num: 20.6}, want: 21},
		{args: args{num: 20.9}, want: 21},
	}
	for _, tt := range tests {
		numStr := strconv.FormatFloat(tt.args.num, 'f', -1, 64)
		wantStr := strconv.FormatFloat(tt.want, 'f', -1, 64)
		t.Run(
			numStr+"で四捨五入すると"+wantStr+"である",
			func(t *testing.T) {
				if got := roundInt(tt.args.num); got != tt.want {
					t.Errorf("roundInt() = %v, want %v", got, tt.want)
				}
			})
	}
}

func Test_roundUpInt(t *testing.T) {
	type args struct {
		num float64
	}
	tests := []struct {
		args args
		want float64
	}{
		{args: args{num: 20.0}, want: 21},
		{args: args{num: 20.1}, want: 21},
		{args: args{num: 20.4}, want: 21},
		{args: args{num: 20.5}, want: 21},
		{args: args{num: 20.6}, want: 21},
		{args: args{num: 20.9}, want: 21},
	}
	for _, tt := range tests {
		numStr := strconv.FormatFloat(tt.args.num, 'f', -1, 64)
		wantStr := strconv.FormatFloat(tt.want, 'f', -1, 64)
		t.Run(
			numStr+"で切り上げすると"+wantStr+"である",
			func(t *testing.T) {
				if got := RoundUpInt(tt.args.num); got != tt.want {
					t.Errorf("RoundUpInt() = %v, want %v", got, tt.want)
				}
			})
	}
}
