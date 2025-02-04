package cmd

import (
	"fmt"
	"math"
)

/*
*
*  Making my vector library bc why not, matrices later although the C version of this only used vectors I think?
*  Prob will need them for perspective projection translations.
*
 */
type Vec3 [3]float64

type Vec2 [2]float64

type Vec4 [4]float64

/*
*
* All methods for Vec2
*
 */

func newVec2(x, y float64) Vec2 {
	return Vec2{x, y}
}

func (v Vec2) add(v2 Vec2) Vec2 {
	return Vec2{
		v[0] + v2[0],
		v[1] + v2[1],
	}
}

func (v Vec2) sub(v2 Vec2) Vec2 {
	return Vec2{
		v[0] - v2[0],
		v[1] - v2[1],
	}
}

func (v Vec2) mult(v2 Vec2) Vec2 {
	return Vec2{
		v[0] * v2[0],
		v[1] * v2[1],
	}
}

func (v Vec2) len() float64 {
	return math.Sqrt(v[0]*v[0] + v[1]*v[1])
}

func (v Vec2) addVec2(v1, v2 Vec2) Vec2 {
	return Vec2{
		v1[0] + v2[0],
		v1[1] + v2[1],
	}
}

func (v Vec2) subVec2(v1, v2 Vec2) Vec2 {
	return Vec2{
		v1[0] - v2[0],
		v1[1] - v2[1],
	}
}

/*
*
* All methods for Vec3
*
*  Any methods such as mulVec, addVec etc are meant to reduce copies of Vecs, no need to create a vec for the sole reason of adding it
 */

func newVec3(x, y, z float64) Vec3 {
	return Vec3{x, y, z}
}

func (v Vec3) add(v2 Vec3) Vec3 {
	return Vec3{
		v[0] + v2[0],
		v[1] + v2[1],
		v[2] + v2[2],
	}
}

func (v Vec3) sub(v2 Vec3) Vec3 {
	return Vec3{
		v[0] - v2[0],
		v[1] - v2[1],
		v[2] - v2[2],
	}
}

func (v Vec3) mult(v2 Vec3) Vec3 {
	return Vec3{
		v[0] * v2[0],
		v[1] * v2[1],
		v[2] * v2[2],
	}
}

func addVec3(v1, v2 Vec3) Vec3 {
	return Vec3{
		v1[0] + v2[0],
		v1[1] + v2[1],
		v1[2] + v2[2],
	}
}

func subVec3(v1, v2 Vec3) Vec3 {
	return Vec3{
		v1[0] - v2[0],
		v1[1] - v2[1],
		v1[2] - v2[2],
	}
}

func mulVec3(v1, v2 Vec3) Vec3 {
	return Vec3{
		v1[0] * v2[0],
		v1[1] * v2[1],
		v1[2] * v2[2],
	}
}

func (v Vec3) len() float64 {
	return math.Sqrt(v[0]*v[0] + v[1]*v[1] + v[2]*v[2])
}

func (v Vec3) dot(v2 Vec3) float64 {
	return v[0]*v2[0] + v[1]*v2[1] + v[2]*v2[2]
}

func (v Vec3) cross(v2 Vec3) Vec3 {
	return Vec3{
		v[1]*v2[2] - v[2]*v2[1],
		-(v[0]*v2[2] - v[2]*v2[0]),
		v[0]*v2[1] - v[1]*v2[0],
	}
}

/*
*
*  All methods for Vec4
 */

func newVec4(x, y, z, j float64) Vec4 {
	return Vec4{x, y, z, j}
}

func (v Vec4) add(v2 Vec4) Vec4 {
	return Vec4{
		v[0] + v2[0],
		v[1] + v2[1],
		v[2] + v2[2],
		v[3] + v2[3],
	}
}

func (v Vec4) sub(v2 Vec4) Vec4 {
	return Vec4{
		v[0] - v2[0],
		v[1] - v2[1],
		v[2] - v2[2],
		v[3] - v2[3],
	}
}

func (v Vec4) Mul(v2 Vec4) Vec4 {
	return Vec4{
		v[0] * v2[0],
		v[1] * v2[1],
		v[2] * v2[2],
		v[3] * v2[3],
	}
}

func (v Vec4) Len() float64 {
	return math.Sqrt(v[0]*v[0] + v[1]*v[1] + v[2]*v[2])
}

func (v Vec4) Dot(v2 Vec4) float64 {
	return 32.
}

func AddVec4(v1, v2 Vec4) Vec4 {
	return Vec4{
		v1[0] + v2[0],
		v1[1] + v2[1],
		v1[2] + v2[2],
		v1[3] + v2[3],
	}
}

func SubVec4(v1, v2 Vec4) Vec4 {
	return Vec4{
		v1[0] - v2[0],
		v1[1] - v2[1],
		v1[2] - v2[2],
		v1[3] - v2[3],
	}
}

func MulVec4(v1, v2 Vec4) Vec4 {
	return Vec4{
		v1[0] * v2[0],
		v1[1] * v2[1],
		v1[2] * v2[2],
		v1[3] * v2[3],
	}
}

// temp test for vectors

func TestVecs() {
	v := Vec3{1, 2, 3}
	v2 := Vec3{4, 5, 6}
	fmt.Print(v.cross(v2))
}
