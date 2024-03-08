package draw

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewCanvas(t *testing.T) {
	type args struct {
		width  int
		height int
	}
	tests := []struct {
		name string
		args args
		want *Canvas
	}{
		{
			name: "Success: returns valid Canvas instance.",
			args: args{
				width:  16,
				height: 8,
			},
			want: &Canvas{
				width:  16,
				height: 8,
				pixels: make([]byte, 16*8),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewCanvas(tt.args.width, tt.args.height)
			require.NotNil(t, got)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestCanvas_SetPixel(t *testing.T) {
	type args struct {
		x     int
		y     int
		color byte
	}
	tests := []struct {
		name   string
		target *Canvas
		args   args
	}{
		{
			name: "Success: sets ",
			target: NewCanvas(16, 8),
			args: args{
				x: 2,
				y: 3,
				color: 4,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.target.SetPixel(tt.args.x, tt.args.y, tt.args.color)
		})
	}
}
