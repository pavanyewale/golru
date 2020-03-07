package lru

import (
	"container/list"
	"reflect"
	"testing"
)

func TestNewLRU(t *testing.T) {
	type args struct {
		size int
	}
	tests := []struct {
		name    string
		args    args
		want    *LRU
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "first",
			args: args{size: 4},
			want: &LRU{
				size:          4,
				keyIndexMap:   make(map[interface{}]*list.Element, 4),
				itemList:      list.New(),
				leastUsedItem: nil,
				worstUsedItem: nil,
			},
			wantErr: false,
		},
		{
			name:    "second",
			args:    args{size: 1},
			want:    nil,
			wantErr: true,
		}, {
			name:    "third",
			args:    args{size: 0},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewLRU(tt.args.size)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewLRU() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLRU() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLRU_SetItem(t *testing.T) {
	type fields struct {
		keyIndexMap   map[interface{}]*list.Element
		size          int
		itemList      *list.List
		leastUsedItem *list.Element
		worstUsedItem *list.Element
		lruInterface  lruInterface
	}
	type args struct {
		key   interface{}
		value interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lru := &LRU{
				keyIndexMap:   tt.fields.keyIndexMap,
				size:          tt.fields.size,
				itemList:      tt.fields.itemList,
				leastUsedItem: tt.fields.leastUsedItem,
				worstUsedItem: tt.fields.worstUsedItem,
				lruInterface:  tt.fields.lruInterface,
			}
			lru.SetItem(tt.args.key, tt.args.value)
		})
	}
}

func TestLRU_GetItem(t *testing.T) {
	type fields struct {
		keyIndexMap   map[interface{}]*list.Element
		size          int
		itemList      *list.List
		leastUsedItem *list.Element
		worstUsedItem *list.Element
		lruInterface  lruInterface
	}
	type args struct {
		key interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lru := &LRU{
				keyIndexMap:   tt.fields.keyIndexMap,
				size:          tt.fields.size,
				itemList:      tt.fields.itemList,
				leastUsedItem: tt.fields.leastUsedItem,
				worstUsedItem: tt.fields.worstUsedItem,
				lruInterface:  tt.fields.lruInterface,
			}
			got, err := lru.GetItem(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("LRU.GetItem() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LRU.GetItem() = %v, want %v", got, tt.want)
			}
		})
	}
}
