package treap

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewTreap(t *testing.T) {
	tests := []struct {
		name string
		want *Treap
	}{
		{
			name: "new treap",
			want: &Treap{
				root: &node{
					value:    "",
					priority: 0,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, NewTreap())
		})
	}
}

func TestTreap_Delete(t *testing.T) {
	type fields struct {
		root *node
	}
	type args struct {
		value string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "delete value from the treap",
			fields: fields{
				root: &node{
					value:    "f",
					priority: 10,
					left: &node{
						value:    "d",
						priority: 8,
						left: &node{
							value:    "c",
							priority: 2,
							left:     nil,
							right:    nil,
						},
						right: &node{
							value:    "e",
							priority: 1,
							left:     nil,
							right:    nil,
						},
					},
					right: &node{
						value:    "t",
						priority: 7,
						left: &node{
							value:    "h",
							priority: 3,
							left:     nil,
							right:    nil,
						},
						right: &node{
							value:    "x",
							priority: 6,
							left:     nil,
							right:    nil,
						},
					},
				},
			},
			args: args{
				value: "d",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			trp := &Treap{
				root: tt.fields.root,
			}
			trp.Delete(tt.args.value)
			assert.False(t, trp.Search(tt.args.value))
		})
	}
}

func TestTreap_Insert(t *testing.T) {
	type fields struct {
		root *node
	}
	type args struct {
		value    string
		priority uint64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *node
	}{
		{
			name: "insert value into an empty treap",
			fields: fields{
				root: nil,
			},
			args: args{
				value:    "h",
				priority: 100,
			},
			want: &node{
				value:    "h",
				priority: 100,
			},
		},
		{
			name: "insert value into the treap",
			fields: fields{
				root: &node{
					value:    "f",
					priority: 10,
					left: &node{
						value:    "d",
						priority: 8,
						left: &node{
							value:    "c",
							priority: 2,
							left:     nil,
							right:    nil,
						},
						right: &node{
							value:    "e",
							priority: 1,
							left:     nil,
							right:    nil,
						},
					},
					right: &node{
						value:    "t",
						priority: 7,
						left: &node{
							value:    "h",
							priority: 3,
							left:     nil,
							right:    nil,
						},
						right: &node{
							value:    "x",
							priority: 6,
							left:     nil,
							right:    nil,
						},
					},
				},
			},
			args: args{
				value:    "k",
				priority: 5,
			},
			want: &node{
				value:    "f",
				priority: 10,
				left: &node{
					value:    "d",
					priority: 8,
					left: &node{
						value:    "c",
						priority: 2,
						left:     nil,
						right:    nil,
					},
					right: &node{
						value:    "e",
						priority: 1,
						left:     nil,
						right:    nil,
					},
				},
				right: &node{
					value:    "t",
					priority: 7,
					left: &node{
						value:    "h",
						priority: 3,
						left:     nil,
						right: &node{
							value:    "k",
							priority: 5,
							left:     nil,
							right:    nil,
						},
					},
					right: &node{
						value:    "x",
						priority: 6,
						left:     nil,
						right:    nil,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			trp := &Treap{
				root: tt.fields.root,
			}
			trp.insert(tt.args.value, tt.args.priority)
			assert.True(t, trp.Search(tt.args.value))
			assert.Equal(t, tt.want, trp.root)
		})
	}
}

func TestTreap_Search(t *testing.T) {
	type fields struct {
		root *node
	}
	type args struct {
		value string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "search for value in empty treap",
			fields: fields{
				root: nil,
			},
			args: args{
				value: "h",
			},
			want: false,
		},
		{
			name: "search for value in single value treap",
			fields: fields{
				root: &node{
					value:    "f",
					priority: 10,
					left:     nil,
					right:    nil,
				},
			},
			args: args{
				value: "h",
			},
			want: false,
		},
		{
			name: "search for value in single value treap",
			fields: fields{
				root: &node{
					value:    "f",
					priority: 10,
					left:     nil,
					right:    nil,
				},
			},
			args: args{
				value: "f",
			},
			want: true,
		},
		{
			name: "search for value in the treap",
			fields: fields{
				root: &node{
					value:    "f",
					priority: 10,
					left: &node{
						value:    "d",
						priority: 8,
						left: &node{
							value:    "c",
							priority: 2,
							left:     nil,
							right:    nil,
						},
						right: &node{
							value:    "e",
							priority: 1,
							left:     nil,
							right:    nil,
						},
					},
					right: &node{
						value:    "t",
						priority: 7,
						left: &node{
							value:    "h",
							priority: 3,
							left:     nil,
							right:    nil,
						},
						right: &node{
							value:    "x",
							priority: 6,
							left:     nil,
							right:    nil,
						},
					},
				},
			},
			args: args{
				value: "h",
			},
			want: true,
		},
		{
			name: "search for nonexistent value in the treap",
			fields: fields{
				root: &node{
					value:    "f",
					priority: 10,
					left: &node{
						value:    "d",
						priority: 8,
						left: &node{
							value:    "c",
							priority: 2,
							left:     nil,
							right:    nil,
						},
						right: &node{
							value:    "e",
							priority: 1,
							left:     nil,
							right:    nil,
						},
					},
					right: &node{
						value:    "t",
						priority: 7,
						left: &node{
							value:    "h",
							priority: 3,
							left:     nil,
							right:    nil,
						},
						right: &node{
							value:    "x",
							priority: 6,
							left:     nil,
							right:    nil,
						},
					},
				},
			},
			args: args{
				value: "z",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t1 *testing.T) {
			trp := &Treap{
				root: tt.fields.root,
			}
			assert.Equal(t, tt.want, trp.Search(tt.args.value))
		})
	}
}
