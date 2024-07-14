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

func TestCanvas_Width(t *testing.T) {
	tests := []struct {
		name   string
		target *Canvas
		want   int
	}{
		{
			name:   "Success: returns 8",
			target: NewCanvas(8, 4),
			want:   8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.target.Width()
			require.Equal(t, tt.want, got)
		})
	}
}

func TestCanvas_Height(t *testing.T) {
	tests := []struct {
		name   string
		target *Canvas
		want   int
	}{
		{
			name:   "Success: returns 4",
			target: NewCanvas(8, 4),
			want:   4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.target.Height()
			require.Equal(t, tt.want, got)
		})
	}
}

func TestCanvas_calcIndex(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name           string
		target         *Canvas
		args           args
		want           int
		wantErr        bool
		wantErrMessage string
	}{
		{
			name:   "Success: returns 0",
			target: NewCanvas(8, 4),
			args: args{
				x: 0,
				y: 0,
			},
			want: 0,
		},
		{
			name:   "Success: returns 31",
			target: NewCanvas(8, 4),
			args: args{
				x: 7,
				y: 3,
			},
			want: 31,
		},
		{
			name:   "Failure: returns 'range over' error",
			target: NewCanvas(8, 4),
			args: args{
				x: -1,
				y: 0,
			},
			wantErr:        true,
			wantErrMessage: "range over",
		},
		{
			name:   "Failure: returns 'range over' error",
			target: NewCanvas(8, 4),
			args: args{
				x: 0,
				y: -1,
			},
			wantErr:        true,
			wantErrMessage: "range over",
		},
		{
			name:   "Failure: returns 'range over' error",
			target: NewCanvas(8, 4),
			args: args{
				x: 8,
				y: 3,
			},
			wantErr:        true,
			wantErrMessage: "range over",
		},
		{
			name:   "Failure: returns 'range over' error",
			target: NewCanvas(8, 4),
			args: args{
				x: 7,
				y: 4,
			},
			wantErr:        true,
			wantErrMessage: "range over",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.target.calcIndex(tt.args.x, tt.args.y)
			if tt.wantErr {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tt.wantErrMessage)
				return
			}

			require.NoError(t, err)
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
			name:   "Success: sets ",
			target: NewCanvas(16, 8),
			args: args{
				x:     2,
				y:     3,
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

func TestCanvas_GetPixel(t *testing.T) {
	type fields struct {
		width  int
		height int
		pixels []byte
	}
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Canvas{
				width:  tt.fields.width,
				height: tt.fields.height,
				pixels: tt.fields.pixels,
			}
			got, err := c.GetPixel(tt.args.x, tt.args.y)
			if (err != nil) != tt.wantErr {
				t.Errorf("Canvas.GetPixel() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Canvas.GetPixel() = %v, want %v", got, tt.want)
			}
		})
	}
}
