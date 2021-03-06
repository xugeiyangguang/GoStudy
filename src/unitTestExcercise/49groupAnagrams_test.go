package unitTestExcercise

import (
	"reflect"
	"testing"
)

func Test_groupAnagrams(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		// TODO: Add test cases.
		{"1", args{[]string{"eat", "tea", "tan", "ate", "nat", "bat"}},
			[][]string{{"ate", "eat", "tea"}, {"nat", "tan"}, {"bat"}}},
		{"2", args{[]string{"tea", "", "tea", ""}}, [][]string{{"tea", "tea"}, {"", ""}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := groupAnagrams(tt.args.strs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("groupAnagrams() = %v, want %v", got, tt.want)
			}
		})
	}
}
