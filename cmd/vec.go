package cmd

type Vec3 [3]float64

type Vec2 [2]float64

type Vec4 [4]float64

func newVec3(x, y, z float64) Vec3 {
	return Vec3{x, y, z}
}

func newVec2(x, y float64) Vec2 {
	return Vec2{x, y}
}

func newVec4(x, y, z, j float64) Vec4 {
	return Vec4{x, y, z, j}
}

func (v Vec2) add(v2 Vec2) Vec2 {
	return Vec2{
		v[0] + v2[0],
		v[1] + v2[1],
	}
}

func (v Vec3) add(v2 Vec3) Vec3 {
	return Vec3{
		v[0] + v2[0],
		v[1] + v2[1],
		v[2] + v2[2],
	}
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
