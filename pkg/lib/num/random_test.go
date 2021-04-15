package num

import (
	"fmt"
	"math/big"
	"reflect"
	"testing"
)

func TestRandomInt(t *testing.T) {
	type args struct {
		num int64
	}
	tests := []struct {
		name     string
		args     args
		want     *big.Int
		wantErr  bool
		isNotVal bool
		loop     int
	}{
		{name: "ランダムな値を取得できる", args: args{num: 2}, want: big.NewInt(1), wantErr: false, isNotVal: false, loop: 100},
		{name: "ランダムな値を取得できる", args: args{num: 8}, want: big.NewInt(4), wantErr: false, isNotVal: false, loop: 100},
		{name: "範囲外の値は取得できない", args: args{num: 1}, want: big.NewInt(1), wantErr: false, isNotVal: true, loop: 100},
		{name: "範囲外の値は取得できない", args: args{num: 2}, want: big.NewInt(2), wantErr: false, isNotVal: true, loop: 100},
		{name: "範囲外の値は取得できない", args: args{num: 2}, want: big.NewInt(3), wantErr: false, isNotVal: true, loop: 100},
		{name: "負の値はエラーが起きる", args: args{num: -1}, want: nil, wantErr: true, isNotVal: false, loop: 100},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for i := 0; i < tt.loop; i++ {
				got, err := RandomInt(tt.args.num)
				if (err != nil) != tt.wantErr {
					t.Errorf("RandomInt() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if reflect.DeepEqual(got, tt.want) {
					return
				}
				fmt.Println(got)
			}

			if !tt.isNotVal {
				t.Errorf("RandomInt() want %v, loop %v", tt.want, tt.loop)
			}
		})
	}
}
