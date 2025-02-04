package cmd

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

/*
*
* All methods for Vec3
*
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

func (v Vec4) mult(v2 Vec4) Vec4 {
	return Vec4{
		v[0] * v2[0],
		v[1] * v2[1],
		v[2] * v2[2],
		v[3] * v2[3],
	}
}

func addVec4(v1, v2 Vec4) Vec4 {
	return Vec4{
		v1[0] * v2[0],
		v1[1] * v2[1],
		v1[2] * v2[2],
		v1[3] * v2[3],
	}
}
