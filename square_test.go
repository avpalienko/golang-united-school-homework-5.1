package square

import (
	"reflect"
	"testing"
)

func TestSquareEnd(t *testing.T) {
	tests := []struct {
		name string
		arg  Square
		want Point
	}{
		{"End/1", Square{start: Point{0, 0}, a: 4}, Point{4, 4}},
		{"End/2", Square{start: Point{-1, -1}, a: 4}, Point{3, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.arg.End(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Got = %v, want %v", got, tt.want)
			}
		})
	}

}

func TestSquareArea(t *testing.T) {
	tests := []struct {
		name string
		arg  Square
		want uint
	}{
		{"Area/1", Square{start: Point{0, 0}, a: 4}, 16},
		{"Area/2", Square{start: Point{0, 0}, a: 5}, 25},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.arg.Area(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Got = %v, want %v", got, tt.want)
			}
		})
	}

}

func TestSquarePerimeter(t *testing.T) {
	tests := []struct {
		name string
		arg  Square
		want uint
	}{
		{"Perimeter/1", Square{start: Point{0, 0}, a: 4}, 16},
		{"Perimeter/2", Square{start: Point{0, 0}, a: 5}, 20},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.arg.Perimeter(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Got = %v, want %v", got, tt.want)
			}
		})
	}

}
