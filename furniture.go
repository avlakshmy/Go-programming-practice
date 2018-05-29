package furniture

type Furniture struct {
	Name string
	Material string
	Price int
}

func (f *Furniture) RetName () string {
	return f.Name
}

func (f *Furniture) RetPrice () int {
	return f.Price
}

func (f *Furniture) RetMaterial () string {
	return f.Material
}
