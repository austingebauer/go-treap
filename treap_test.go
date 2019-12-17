package go_treap

import (
	"reflect"
	"testing"
)

func TestNewTreap(t *testing.T) {
	tests := []struct {
		name string
		want *Treap
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTreap(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTreap() = %v, want %v", got, tt.want)
			}
		})
	}
}
