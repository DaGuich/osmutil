package geotypes

import (
	"math"
	"reflect"
	"testing"
)

func TestVector2D_Length(t *testing.T) {
	t.Parallel()
	type fields struct {
		X float64
		Y float64
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{
			"Easy",
			fields{1, 0},
			1.0,
		},
		{
			"2",
			fields{1, 1},
			math.Sqrt(2),
		},
		{
			"3",
			fields{2, 0},
			2,
		},
		{
			"4",
			fields{0, 2},
			2,
		},
		{
			"5",
			fields{4, 4},
			math.Sqrt(32),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &Vector2D{
				X: tt.fields.X,
				Y: tt.fields.Y,
			}
			if got := v.Length(); got != tt.want {
				t.Errorf("Vector2D.Length() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVector3D_Length(t *testing.T) {
	t.Parallel()
	type fields struct {
		X float64
		Y float64
		Z float64
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{
			"Easy",
			fields{X: 1, Y: 0, Z: 0},
			1.0,
		},
		{
			"2",
			fields{X: 1, Y: 1, Z: 0},
			math.Sqrt(2),
		},
		{
			"3",
			fields{X: 1, Y: 1, Z: 1},
			math.Sqrt(3),
		},
		{
			"4",
			fields{X: 4, Y: 4, Z: 2},
			6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &Vector3D{
				X: tt.fields.X,
				Y: tt.fields.Y,
				Z: tt.fields.Z,
			}
			if got := v.Length(); got != tt.want {
				t.Errorf("Vector3D.Length() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVector2D_Normalize(t *testing.T) {
	t.Parallel()
	type fields struct {
		X float64
		Y float64
	}
	tests := []struct {
		name   string
		fields fields
		want   *Vector2D
	}{
		{
			"1",
			fields{1, 0},
			&Vector2D{1, 0},
		},
		{
			"2",
			fields{0, 1},
			&Vector2D{0, 1},
		},
		{
			"3",
			fields{-1, 0},
			&Vector2D{-1, 0},
		},
		{
			"4",
			fields{0, -1},
			&Vector2D{0, -1},
		},
		{
			"5",
			fields{0, 5},
			&Vector2D{0, 1},
		},
		{
			"6",
			fields{1, 1},
			&Vector2D{1 / math.Sqrt(2), 1 / math.Sqrt(2)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &Vector2D{
				X: tt.fields.X,
				Y: tt.fields.Y,
			}
			if got := v.Normalize(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Vector2D.Normalize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVector3D_Normalize(t *testing.T) {
	t.Parallel()
	type fields struct {
		X float64
		Y float64
		Z float64
	}
	tests := []struct {
		name   string
		fields fields
		want   *Vector3D
	}{
		{
			"1",
			fields{1, 0, 0},
			&Vector3D{1, 0, 0},
		},
		{
			"2",
			fields{0, 1, 0},
			&Vector3D{0, 1, 0},
		},
		{
			"3",
			fields{0, 0, 1},
			&Vector3D{0, 0, 1},
		},
		{
			"4",
			fields{1, 0, 0},
			&Vector3D{1, 0, 0},
		},
		{
			"5",
			fields{1, 0, 0},
			&Vector3D{1, 0, 0},
		},
		{
			"6",
			fields{1, 0, 0},
			&Vector3D{1, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &Vector3D{
				X: tt.fields.X,
				Y: tt.fields.Y,
				Z: tt.fields.Z,
			}
			if got := v.Normalize(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Vector3D.Normalize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVector2D_Add(t *testing.T) {
	t.Parallel()
	type fields struct {
		X float64
		Y float64
	}
	type args struct {
		other *Vector2D
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Vector2D
	}{
		// TODO: Add test cases.
		{
			"1",
			fields{1, 0},
			args{&Vector2D{1, 0}},
			&Vector2D{2, 0},
		},
		{
			"2",
			fields{1, 1},
			args{&Vector2D{1, 2}},
			&Vector2D{2, 3},
		},
		{
			"3",
			fields{5, 8},
			args{&Vector2D{2, 12}},
			&Vector2D{7, 20},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &Vector2D{
				X: tt.fields.X,
				Y: tt.fields.Y,
			}
			if got := v.Add(tt.args.other); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Vector2D.Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVector3D_Add(t *testing.T) {
	t.Parallel()
	type fields struct {
		X float64
		Y float64
		Z float64
	}
	type args struct {
		other *Vector3D
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Vector3D
	}{
		// TODO: Add test cases.
		{
			"1",
			fields{1, 0, 0},
			args{&Vector3D{1, 0, 0}},
			&Vector3D{2, 0, 0},
		},
		{
			"2",
			fields{1, 0, 0},
			args{&Vector3D{1, 2, 3}},
			&Vector3D{2, 2, 3},
		},
		{
			"3",
			fields{5, 8, 11},
			args{&Vector3D{2, 12, -12}},
			&Vector3D{7, 20, -1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &Vector3D{
				X: tt.fields.X,
				Y: tt.fields.Y,
				Z: tt.fields.Z,
			}
			if got := v.Add(tt.args.other); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Vector3D.Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVector2D_Sub(t *testing.T) {
	t.Parallel()
	type fields struct {
		X float64
		Y float64
	}
	type args struct {
		other *Vector2D
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Vector2D
	}{
		// TODO: Add test cases.
		{
			"1",
			fields{1, 0},
			args{&Vector2D{1, 0}},
			&Vector2D{0, 0},
		},
		{
			"2",
			fields{1, 1},
			args{&Vector2D{1, 2}},
			&Vector2D{0, -1},
		},
		{
			"3",
			fields{5, 8},
			args{&Vector2D{2, 12}},
			&Vector2D{3, -4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &Vector2D{
				X: tt.fields.X,
				Y: tt.fields.Y,
			}
			if got := v.Sub(tt.args.other); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Vector2D.Sub() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVector3D_Sub(t *testing.T) {
	t.Parallel()
	type fields struct {
		X float64
		Y float64
		Z float64
	}
	type args struct {
		other *Vector3D
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Vector3D
	}{
		// TODO: Add test cases.
		{
			"1",
			fields{1, 0, 0},
			args{&Vector3D{1, 0, 0}},
			&Vector3D{0, 0, 0},
		},
		{
			"2",
			fields{1, 0, 0},
			args{&Vector3D{1, 2, 3}},
			&Vector3D{0, -2, -3},
		},
		{
			"3",
			fields{5, 8, 11},
			args{&Vector3D{2, 12, -12}},
			&Vector3D{3, -4, 23},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &Vector3D{
				X: tt.fields.X,
				Y: tt.fields.Y,
				Z: tt.fields.Z,
			}
			if got := v.Sub(tt.args.other); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Vector3D.Sub() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVector2D_Mult(t *testing.T) {
	t.Parallel()
	type fields struct {
		X float64
		Y float64
	}
	type args struct {
		other *Vector2D
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   float64
	}{
		// TODO: Add test cases.
		{
			"1",
			fields{1, 0},
			args{&Vector2D{1, 0}},
			1,
		},
		{
			"2",
			fields{1, 1},
			args{&Vector2D{1, 0}},
			1,
		},
		{
			"3",
			fields{1, 1},
			args{&Vector2D{1, 1}},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &Vector2D{
				X: tt.fields.X,
				Y: tt.fields.Y,
			}
			if got := v.Mult(tt.args.other); got != tt.want {
				t.Errorf("Vector2D.Mult() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVector3D_Mult(t *testing.T) {
	t.Parallel()
	type fields struct {
		X float64
		Y float64
		Z float64
	}
	type args struct {
		other *Vector3D
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   float64
	}{
		// TODO: Add test cases.
		{
			"1",
			fields{1, 0, 0},
			args{&Vector3D{1, 0, 0}},
			1,
		},
		{
			"2",
			fields{1, 1, 0},
			args{&Vector3D{1, 0, 0}},
			1,
		},
		{
			"3",
			fields{1, 1, 1},
			args{&Vector3D{1, 0, 0}},
			1,
		},
		{
			"4",
			fields{1, 1, 1},
			args{&Vector3D{1, 1, 0}},
			2,
		},
		{
			"5",
			fields{1, 1, 1},
			args{&Vector3D{1, 1, 1}},
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &Vector3D{
				X: tt.fields.X,
				Y: tt.fields.Y,
				Z: tt.fields.Z,
			}
			if got := v.Mult(tt.args.other); got != tt.want {
				t.Errorf("Vector3D.Mult() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVector2D_Angle(t *testing.T) {
	t.Parallel()
	type fields struct {
		X float64
		Y float64
	}
	type args struct {
		other *Vector2D
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   float64
	}{
		// TODO: Add test cases.
		{
			"1",
			fields{1, 0},
			args{&Vector2D{0, 1}},
			90,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &Vector2D{
				X: tt.fields.X,
				Y: tt.fields.Y,
			}
			if got := v.Angle(tt.args.other); got != tt.want {
				t.Errorf("Vector2D.Angle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVector3D_Angle(t *testing.T) {
	t.Parallel()
	type fields struct {
		X float64
		Y float64
		Z float64
	}
	type args struct {
		other *Vector3D
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   float64
	}{
		{
			"1",
			fields{1, 0, 0},
			args{&Vector3D{0, 1, 0}},
			90,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &Vector3D{
				X: tt.fields.X,
				Y: tt.fields.Y,
				Z: tt.fields.Z,
			}
			if got := v.Angle(tt.args.other); got != tt.want {
				t.Errorf("Vector3D.Angle() = %v, want %v", got, tt.want)
			}
		})
	}
}
