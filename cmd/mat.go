package cmd

import "math"

/*
Mat4
1D array representation of a 2D matrix, [col][row] can be accessed with
Mat[col + row * 4]
*/
type Mat3 [3 * 3]float64
type Mat4 [4 * 4]float64

/*

All methods for Mat4

*/

func NewMat4() Mat4 {
	return Mat4{}
}

func (m Mat4) Mul(m2 Mat4) Mat4 {
	var newMat Mat4
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			sum := 0.0
			for k := 0; k < 4; k++ {
				sum += m[i*4+k] * m2[k*4+j]
			}
			newMat[i*4+j] = sum
		}
	}
	return newMat
}

func (m Mat4) Add(m2 Mat4) Mat4 {
	var newMat Mat4
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			newMat[i*4+j] = m[i*4+j] + m2[i*4+j]
		}
	}
	return newMat
}
func (m Mat4) Sub(m2 Mat4) Mat4 {
	var newMat Mat4
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			newMat[i*4+j] = m[i*4+j] - m2[i*4+j]
		}
	}
	return newMat
}

func AddMat4(a, b Mat4) Mat4 {
	newMat := a.Add(b)
	return newMat
}

func SubMat4(a, b Mat4) Mat4 {
	newMat := a.Sub(b)
	return newMat
}

func MulMat4(a, b Mat4) Mat4 {
	newMat := a.Mul(b)
	return newMat
}

/*
Common graphics matrices
*/

func NewPerspectiveMatrix(aspect, fov, far, near float64) Mat4 {
	return Mat4{
		1 / (aspect * math.Tan(fov/2)), 0, 0, 0, // Row one
		0, 1 / (math.Tan(fov / 2)), 0, 0, // Row two
		0, 0, -(far + near/far - near), -((2 * far * near) / (far - near)), // Row three
		0, 0, -1, 0, // Row four
	}
}

/*
The RotateX function takes in an angle, theta and rotates by that many degrees around the x-axis
*/
func RotateX(theta float64) Mat4 {
	return Mat4{}
}

func (m Mat4) Translate(Tx, Ty, Tz float64) Mat4 {
	return Mat4{}
}
