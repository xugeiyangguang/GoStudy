package unitTestExcercise

import (
	"reflect"
	"testing"
)

func Test_displayTable(t *testing.T) {
	type args struct {
		orders [][]string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		// TODO: Add test cases.
		{"1", args{[][]string{{"Laura", "2", "Bean Burrito"}, {"Jhon", "2", "Beef Burrito"}, {"Melissa", "2", "Soda"}}}, [][]string{{"Table", "Bean Burrito", "Beef Burrito", "Soda"}, {"2", "1", "1", "1"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := displayTable(tt.args.orders); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("displayTable() = %v, want %v", got, tt.want)
			}
		})
	}
}
