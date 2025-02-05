package cmd

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
