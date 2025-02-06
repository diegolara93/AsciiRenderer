package cmd

import "math"

/*
Mat4
1D array representation of a 2D matrix, [col][row] can be accessed with
Mat[col + row * 4]
*/
type (
	Mat3 [3 * 3]float64
	Mat4 [4 * 4]float64
)

/*

All methods for Mat4

*/

func NewMat4(a, b, c, d, e, f, h, i, j, k, l, m, n, o, p, q float64) Mat4 {
	return Mat4{
		a, b, c, d,
		e, f, h, i,
		j, k, l, m,
		n, o, p, q,
	}
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

func NewViewMatrix() Mat4 {
	return Mat4{}
}

/*

In computer graphics, contrary to how dimensions are portrayed in Calculus, we have the z axis pointing towards us, like this:
						y-axis
					|
					|
					|
					|
					|
					/--------------------- x-axis
		          /
			    /
		       / z-axis
*/

/*
The RotateX function takes in an angle, theta and rotates by that many degrees around the x-axis
*/
func (m Mat4) RotateX(theta float64) Mat4 {
	rotationXMatrix := NewMat4(
		1, 0, 0, 0,
		0, math.Cos(theta), math.Sin(theta), 0,
		0, -(math.Sin(theta)), math.Cos(theta), 0,
		0, 0, 0, 1,
	)
	return m.Mul(rotationXMatrix)
}

/*
The RotateY function takes in an angle, theta and rotates by that many degrees around the y-axis
*/
func (m Mat4) RotateY(theta float64) Mat4 {
	rotationYMatrix := NewMat4(
		1, 0, 0, 0,
		0, math.Cos(theta), math.Sin(theta), 0,
		0, -(math.Sin(theta)), math.Cos(theta), 0,
		0, 0, 0, 1,
	)
	return m.Mul(rotationYMatrix)
}

/*
The RotateZ function takes in an angle, theta and rotates by that many degrees around the z-axis
*/
func (m Mat4) RotateZ(theta float64) Mat4 {
	rotationZMatrix := NewMat4(
		math.Cos(theta), -(math.Sin(theta)), 0, 0,
		math.Sin(theta), math.Cos(theta), 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	)
	return m.Mul(rotationZMatrix)
}

/*
Translates the matrix m by Tx, Ty, Tz
*/
func (m Mat4) Translate(Tx, Ty, Tz float64) Mat4 {
	translationMatrix := NewMat4(
		1, 0, 0, Tx,
		0, 1, 0, Ty,
		0, 0, 1, Tz,
		0, 0, 0, 1,
	)
	return m.Mul(translationMatrix)
}

/*
Scales the matrix m by sX, sY, sZ
*/

func (m Mat4) Scale(sX, sY, sZ float64) Mat4 {
	scalingMatrix := NewMat4(
		sX, 0, 0, 0,
		0, sY, 0, 0,
		0, 0, sZ, 0,
		0, 0, 0, 1,
	)
	return m.Mul(scalingMatrix)
}

func LookAtRH(eye, target, up Vec3) Mat4 {
	zAxis := normalize(eye.sub(target))
	xAxis := normalize(up.cross(zAxis))
	yaxis := zAxis.cross(xAxis)

	orientation := NewMat4(
		zAxis[0], yaxis[0], zAxis[0], 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		-eye[0], -eye[1], -eye[2], 1,
	)

	translation := NewMat4(
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		-eye[0], -eye[1], -eye[2], 1,
	)

	return MulMat4(orientation, translation)
}
