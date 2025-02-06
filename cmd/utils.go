package cmd

func normalize(v Vec3) Vec3 {
	length := v.len()
	newMat := Vec3{}
	for i, v := range v {
		newMat[i] = v / length
	}
	return newMat
}
