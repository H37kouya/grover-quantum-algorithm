package num

import (
	"reflect"
	"testing"
)

func TestCalcExponent(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name    string
		args    args
		want    Exponent
		wantErr bool
	}{
		{name: "4は指数である", args: args{num: 4}, want: Exponent{Base: 2, Exp: 2}, wantErr: false},
		{name: "8は指数である", args: args{num: 8}, want: Exponent{Base: 2, Exp: 3}, wantErr: false},
		{name: "16は指数である", args: args{num: 16}, want: Exponent{Base: 2, Exp: 4}, wantErr: false},
		{name: "64は指数である", args: args{num: 64}, want: Exponent{Base: 2, Exp: 6}, wantErr: false},
		{name: "1024は指数である", args: args{num: 1024}, want: Exponent{Base: 2, Exp: 10}, wantErr: false},
		{name: "27は指数である", args: args{num: 27}, want: Exponent{Base: 3, Exp: 3}, wantErr: false},
		{name: "100は指数である", args: args{num: 100}, want: Exponent{Base: 10, Exp: 2}, wantErr: false},
		{name: "121は指数である", args: args{num: 121}, want: Exponent{Base: 11, Exp: 2}, wantErr: false},
		{name: "3125は指数である", args: args{num: 5 * 5 * 5 * 5 * 5}, want: Exponent{Base: 5, Exp: 5}, wantErr: false},
		{name: "1e10は指数である", args: args{num: 1e10}, want: Exponent{Base: 10, Exp: 10}, wantErr: false},
		{name: "1e14は指数である", args: args{num: 1e14}, want: Exponent{Base: 10, Exp: 14}, wantErr: false},
		{name: "13^4は指数である", args: args{num: 13 * 13 * 13 * 13}, want: Exponent{Base: 13, Exp: 4}, wantErr: false},
		{name: "1は指数でない", args: args{num: 1}, want: Exponent{Base: 0, Exp: 0}, wantErr: true},
		{name: "2は指数でない", args: args{num: 2}, want: Exponent{Base: 0, Exp: 0}, wantErr: true},
		{name: "3は指数でない", args: args{num: 3}, want: Exponent{Base: 0, Exp: 0}, wantErr: true},
		{name: "7は指数でない", args: args{num: 7}, want: Exponent{Base: 0, Exp: 0}, wantErr: true},
		{name: "101は指数である", args: args{num: 101}, want: Exponent{Base: 0, Exp: 0}, wantErr: true},
		{name: "10001指数でない", args: args{num: 100001}, want: Exponent{Base: 0, Exp: 0}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CalcExponent(tt.args.num)
			if (err != nil) != tt.wantErr {
				t.Errorf("CalcExponent() got = %v, want %v", got, tt.want)
				t.Errorf("CalcExponent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CalcExponent() got = %v, want %v", got, tt.want)
			}
		})
	}
}
