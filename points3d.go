package vu

type Points3d []Point3d

func (pp Points3d) Bounds() (min, max Point3d) {
	min = pp[0]
	max = pp[0]
	for _, p := range pp {
		min = min.Min(p)
		max = max.Max(p)
	}
	return
}
