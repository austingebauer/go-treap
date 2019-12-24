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
			want: &Treap{},
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
						},
						right: &node{
							value:    "e",
							priority: 1,
						},
					},
					right: &node{
						value:    "t",
						priority: 7,
						left: &node{
							value:    "h",
							priority: 3,
						},
						right: &node{
							value:    "x",
							priority: 6,
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
						},
						right: &node{
							value:    "e",
							priority: 1,
						},
					},
					right: &node{
						value:    "t",
						priority: 7,
						left: &node{
							value:    "h",
							priority: 3,
						},
						right: &node{
							value:    "x",
							priority: 6,
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
					},
					right: &node{
						value:    "e",
						priority: 1,
					},
				},
				right: &node{
					value:    "t",
					priority: 7,
					left: &node{
						value:    "k",
						priority: 5,
						left: &node{
							value:    "h",
							priority: 3,
						},
					},
					right: &node{
						value:    "x",
						priority: 6,
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
			trp.root = insert(trp.root, tt.args.value, tt.args.priority)
			assert.Equal(t, tt.want, trp.root)
			assert.True(t, trp.Search(tt.args.value))
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
						},
						right: &node{
							value:    "e",
							priority: 1,
						},
					},
					right: &node{
						value:    "t",
						priority: 7,
						left: &node{
							value:    "h",
							priority: 3,
						},
						right: &node{
							value:    "x",
							priority: 6,
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
						},
						right: &node{
							value:    "e",
							priority: 1,
						},
					},
					right: &node{
						value:    "t",
						priority: 7,
						left: &node{
							value:    "h",
							priority: 3,
						},
						right: &node{
							value:    "x",
							priority: 6,
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

func Test_rotateRight(t *testing.T) {
	type args struct {
		root *node
	}
	tests := []struct {
		name string
		args args
		want *node
	}{
		{
			name: "rotate tree right",
			args: args{
				root: &node{
					value:    "5",
					priority: 10,
					right: &node{
						value:    "7",
						priority: 9,
					},
					left: &node{
						value:    "3",
						priority: 8,
						right: &node{
							value:    "4",
							priority: 5,
						},
						left: &node{
							value:    "2",
							priority: 3,
						},
					},
				},
			},
			want: &node{
				value:    "3",
				priority: 8,
				right: &node{
					value:    "5",
					priority: 10,
					right: &node{
						value:    "7",
						priority: 9,
					},
					left: &node{
						value:    "4",
						priority: 5,
					},
				},
				left: &node{
					value:    "2",
					priority: 3,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// assert returned value is the new root after the rotation
			assert.Equal(t, tt.want, rotateRight(tt.args.root, tt.args.root.left))
		})
	}
}

func Test_rotateLeft(t *testing.T) {
	type args struct {
		root  *node
		pivot *node
	}
	tests := []struct {
		name string
		args args
		want *node
	}{
		{
			name: "rotate tree left",
			args: args{
				root: &node{
					value:    "3",
					priority: 8,
					right: &node{
						value:    "5",
						priority: 10,
						right: &node{
							value:    "7",
							priority: 9,
						},
						left: &node{
							value:    "4",
							priority: 5,
						},
					},
					left: &node{
						value:    "2",
						priority: 3,
					},
				},
			},
			want: &node{
				value:    "5",
				priority: 10,
				right: &node{
					value:    "7",
					priority: 9,
				},
				left: &node{
					value:    "3",
					priority: 8,
					right: &node{
						value:    "4",
						priority: 5,
					},
					left: &node{
						value:    "2",
						priority: 3,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// assert returned value is the new root after the rotation
			assert.Equal(t, tt.want, rotateLeft(tt.args.root, tt.args.root.right))
		})
	}
}
