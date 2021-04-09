package service

import (
	"reflect"
	"testing"
)

func TestShorFactorization(t *testing.T) {
	type args struct {
		M int
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		wantErr bool
	}{
		{name: "-1は自然数ではないので、エラーをはく", args: args{M: -1}, want: []int{}, wantErr: true},
		{name: "0は自然数ではないので、エラーをはく", args: args{M: 0}, want: []int{}, wantErr: true},
		{name: "1の素因数分解", args: args{M: 1}, want: []int{1}, wantErr: false},
		{name: "2の素因数分解", args: args{M: 2}, want: []int{2}, wantErr: false},
		{name: "3の素因数分解", args: args{M: 3}, want: []int{3}, wantErr: false},
		{name: "4の素因数分解", args: args{M: 4}, want: []int{2, 2}, wantErr: false},
		{name: "5の素因数分解", args: args{M: 5}, want: []int{5}, wantErr: false},
		{name: "6の素因数分解", args: args{M: 6}, want: []int{2, 3}, wantErr: false},
		{name: "7の素因数分解", args: args{M: 7}, want: []int{7}, wantErr: false},
		{name: "8の素因数分解", args: args{M: 8}, want: []int{2, 2, 2}, wantErr: false},
		{name: "9の素因数分解", args: args{M: 9}, want: []int{3, 3}, wantErr: false},
		{name: "10の素因数分解", args: args{M: 10}, want: []int{2, 5}, wantErr: false},
		{name: "11の素因数分解", args: args{M: 11}, want: []int{11}, wantErr: false},
		{name: "12の素因数分解", args: args{M: 12}, want: []int{2, 2, 3}, wantErr: false},
		{name: "13の素因数分解", args: args{M: 13}, want: []int{13}, wantErr: false},
		{name: "14の素因数分解", args: args{M: 14}, want: []int{2, 7}, wantErr: false},
		{name: "15の素因数分解", args: args{M: 15}, want: []int{3, 5}, wantErr: false},
		{name: "16の素因数分解", args: args{M: 16}, want: []int{2, 2, 2, 2}, wantErr: false},
		{name: "17の素因数分解", args: args{M: 17}, want: []int{17}, wantErr: false},
		{name: "18の素因数分解", args: args{M: 18}, want: []int{2, 3, 3}, wantErr: false},
		{name: "19の素因数分解", args: args{M: 19}, want: []int{19}, wantErr: false},
		{name: "20の素因数分解", args: args{M: 20}, want: []int{2, 2, 5}, wantErr: false},
		{name: "21の素因数分解", args: args{M: 21}, want: []int{3, 7}, wantErr: false},
		{name: "22の素因数分解", args: args{M: 22}, want: []int{2, 11}, wantErr: false},
		{name: "23の素因数分解", args: args{M: 23}, want: []int{23}, wantErr: false},
		{name: "24の素因数分解", args: args{M: 24}, want: []int{2, 2, 2, 3}, wantErr: false},
		{name: "25の素因数分解", args: args{M: 25}, want: []int{5, 5}, wantErr: false},
		{name: "26の素因数分解", args: args{M: 26}, want: []int{2, 13}, wantErr: false},
		{name: "27の素因数分解", args: args{M: 27}, want: []int{3, 3, 3}, wantErr: false},
		{name: "28の素因数分解", args: args{M: 28}, want: []int{2, 2, 7}, wantErr: false},
		{name: "29の素因数分解", args: args{M: 29}, want: []int{29}, wantErr: false},
		{name: "30の素因数分解", args: args{M: 30}, want: []int{2, 3, 5}, wantErr: false},
		{name: "31の素因数分解", args: args{M: 31}, want: []int{31}, wantErr: false},
		{name: "32の素因数分解", args: args{M: 32}, want: []int{2, 2, 2, 2, 2}, wantErr: false},
		{name: "33の素因数分解", args: args{M: 33}, want: []int{3, 11}, wantErr: false},
		{name: "34の素因数分解", args: args{M: 34}, want: []int{2, 17}, wantErr: false},
		{name: "35の素因数分解", args: args{M: 35}, want: []int{5, 7}, wantErr: false},
		{name: "36の素因数分解", args: args{M: 36}, want: []int{2, 2, 3, 3}, wantErr: false},
		{name: "37の素因数分解", args: args{M: 37}, want: []int{37}, wantErr: false},
		{name: "38の素因数分解", args: args{M: 38}, want: []int{2, 19}, wantErr: false},
		{name: "39の素因数分解", args: args{M: 39}, want: []int{3, 13}, wantErr: false},
		{name: "40の素因数分解", args: args{M: 40}, want: []int{2, 2, 2, 5}, wantErr: false},
		{name: "41の素因数分解", args: args{M: 41}, want: []int{41}, wantErr: false},
		{name: "42の素因数分解", args: args{M: 42}, want: []int{2, 3, 7}, wantErr: false},
		{name: "43の素因数分解", args: args{M: 43}, want: []int{43}, wantErr: false},
		{name: "44の素因数分解", args: args{M: 44}, want: []int{2, 2, 11}, wantErr: false},
		{name: "45の素因数分解", args: args{M: 45}, want: []int{3, 3, 5}, wantErr: false},
		{name: "46の素因数分解", args: args{M: 46}, want: []int{2, 23}, wantErr: false},
		{name: "47の素因数分解", args: args{M: 47}, want: []int{47}, wantErr: false},
		{name: "48の素因数分解", args: args{M: 48}, want: []int{2, 2, 2, 2, 3}, wantErr: false},
		{name: "49の素因数分解", args: args{M: 49}, want: []int{7, 7}, wantErr: false},
		{name: "50の素因数分解", args: args{M: 50}, want: []int{2, 5, 5}, wantErr: false},
		{name: "51の素因数分解", args: args{M: 51}, want: []int{3, 17}, wantErr: false},
		{name: "52の素因数分解", args: args{M: 52}, want: []int{2, 2, 13}, wantErr: false},
		{name: "53の素因数分解", args: args{M: 53}, want: []int{53}, wantErr: false},
		{name: "54の素因数分解", args: args{M: 54}, want: []int{2, 3, 3, 3}, wantErr: false},
		{name: "55の素因数分解", args: args{M: 55}, want: []int{5, 11}, wantErr: false},
		{name: "56の素因数分解", args: args{M: 56}, want: []int{2, 2, 2, 7}, wantErr: false},
		{name: "57の素因数分解", args: args{M: 57}, want: []int{3, 19}, wantErr: false},
		{name: "58の素因数分解", args: args{M: 58}, want: []int{2, 29}, wantErr: false},
		{name: "59の素因数分解", args: args{M: 59}, want: []int{59}, wantErr: false},
		{name: "60の素因数分解", args: args{M: 60}, want: []int{2, 2, 3, 5}, wantErr: false},
		{name: "61の素因数分解", args: args{M: 61}, want: []int{61}, wantErr: false},
		{name: "62の素因数分解", args: args{M: 62}, want: []int{2, 31}, wantErr: false},
		{name: "63の素因数分解", args: args{M: 63}, want: []int{3, 3, 7}, wantErr: false},
		{name: "64の素因数分解", args: args{M: 64}, want: []int{2, 2, 2, 2, 2, 2}, wantErr: false},
		{name: "65の素因数分解", args: args{M: 65}, want: []int{5, 13}, wantErr: false},
		{name: "66の素因数分解", args: args{M: 66}, want: []int{2, 3, 11}, wantErr: false},
		{name: "67の素因数分解", args: args{M: 67}, want: []int{67}, wantErr: false},
		{name: "68の素因数分解", args: args{M: 68}, want: []int{2, 2, 17}, wantErr: false},
		{name: "69の素因数分解", args: args{M: 69}, want: []int{3, 23}, wantErr: false},
		{name: "70の素因数分解", args: args{M: 70}, want: []int{2, 5, 7}, wantErr: false},
		{name: "119の素因数分解", args: args{M: 119}, want: []int{7, 17}, wantErr: false},
		{name: "3331の素因数分解", args: args{M: 3331}, want: []int{3331}, wantErr: false},
		{name: "33331の素因数分解", args: args{M: 33331}, want: []int{33331}, wantErr: false},
		{name: "333331の素因数分解", args: args{M: 333331}, want: []int{333331}, wantErr: false},
		{name: "5*7*11*13*17*19の素因数分解", args: args{M: 5 * 7 * 11 * 13 * 17 * 19}, want: []int{5, 7, 11, 13, 17, 19}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ShorFactorization(tt.args.M)
			if (err != nil) != tt.wantErr {
				t.Errorf("ShorFactorization() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ShorFactorization() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_divisionTwoAsFarAsPossible(t *testing.T) {
	type args struct {
		M int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := divisionTwoAsFarAsPossible(tt.args.M)
			if got != tt.want {
				t.Errorf("divisionTwoAsFarAsPossible() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("divisionTwoAsFarAsPossible() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
