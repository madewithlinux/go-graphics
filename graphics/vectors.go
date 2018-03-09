package graphics

//go:generate go run ../Slice2D/generator/slice_2d_generator.go PackageLocalType $GOPACKAGE Vec2
//go:generate go run ../Slice2D/generator/slice_2d_generator.go PackageLocalType $GOPACKAGE Vec2i
//go:generate go run ../Slice2D/generator/slice_2d_generator.go PackageLocalType $GOPACKAGE Vec3
//go:generate go run ../Slice2D/generator/slice_2d_generator.go PackageLocalType $GOPACKAGE Vec3i
//go:generate go run ../Slice2D/generator/slice_2d_generator.go PackageLocalType $GOPACKAGE Vec4
//go:generate go run ../Slice2D/generator/slice_2d_generator.go PackageLocalType $GOPACKAGE Vec4i

func (v Vec2) Cross(u Vec2) Float {
	return v.X*u.Y - v.Y*u.X
}

func Vec2Midpoint(a, b Vec2) Vec2 { return a.AddV(b).DivS(2.0) }
func Vec3Midpoint(a, b Vec3) Vec3 { return a.AddV(b).DivS(2.0) }
func Vec4Midpoint(a, b Vec4) Vec4 { return a.AddV(b).DivS(2.0) }