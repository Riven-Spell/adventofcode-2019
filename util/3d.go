package util

type Point3D struct {
	X, Y, Z int64
}

func (p Point3D) Add(p2 Point3D) Point3D {
	return Point3D{
		X: p.X + p2.X,
		Y: p.Y + p2.Y,
		Z: p.Z + p2.Z,
	}
}