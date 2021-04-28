package valueObject

import (
	"reflect"
	"testing"
)

func TestN_ElementCount(t *testing.T) {
	type fields struct {
		value int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{name: "2の1乘である", fields: fields{value: 1}, want: 2},
		{name: "2の5乘である", fields: fields{value: 5}, want: 32},
		{name: "2の10乘である", fields: fields{value: 10}, want: 1024},
		{name: "2の20乘である", fields: fields{value: 20}, want: 1048576},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := N{
				value: tt.fields.value,
			}
			if got := n.ElementCount(); got != tt.want {
				t.Errorf("ElementCount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestN_Get(t *testing.T) {
	type fields struct {
		value int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{name: "取得できる", fields: fields{value: 1}, want: 1},
		{name: "取得できる", fields: fields{value: 5}, want: 5},
		{name: "取得できる", fields: fields{value: 10}, want: 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := N{
				value: tt.fields.value,
			}
			if got := n.Get(); got != tt.want {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewN(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name    string
		args    args
		want    N
		wantErr bool
	}{
		{name: "1は生成できる", args: args{n: 1}, want: N{value: 1}, wantErr: false},
		{name: "5は生成できる", args: args{n: 5}, want: N{value: 5}, wantErr: false},
		{name: "10は生成できる", args: args{n: 10}, want: N{value: 10}, wantErr: false},
		{name: "28は生成できる", args: args{n: 28}, want: N{value: 28}, wantErr: false},
		{name: "28より大きな値はエラーを吐く", args: args{n: 29}, want: N{value: 0}, wantErr: true},
		{name: "28より大きな値はエラーを吐く", args: args{n: 30}, want: N{value: 0}, wantErr: true},
		{name: "負の値はエラーを吐く", args: args{n: -1}, want: N{value: 0}, wantErr: true},
		{name: "0はエラーを吐く", args: args{n: 0}, want: N{value: 0}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewN(tt.args.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewN() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewN() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewNByTenDecimalNumber(t *testing.T) {
	type args struct {
		decimalNumber int
	}
	tests := []struct {
		name    string
		args    args
		want    N
		wantErr bool
	}{
		{name: "1は生成できない", args: args{decimalNumber: 1}, want: N{value: 0}, wantErr: true},
		{name: "4は生成できる", args: args{decimalNumber: 4}, want: N{value: 2}, wantErr: false},
		{name: "5は生成できる", args: args{decimalNumber: 5}, want: N{value: 3}, wantErr: false},
		{name: "16は生成できる", args: args{decimalNumber: 16}, want: N{value: 4}, wantErr: false},
		{name: "21は生成できる", args: args{decimalNumber: 21}, want: N{value: 5}, wantErr: false},
		{name: "2000は生成できる", args: args{decimalNumber: 2000}, want: N{value: 11}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewNByTenDecimalNumber(tt.args.decimalNumber)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewNByTenDecimalNumber() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewNByTenDecimalNumber() got = %v, want %v", got, tt.want)
			}
		})
	}
}
