package item

type Gula struct{
	Nama string
	Volume int
	Harga float32
}

func (g *Gula) Sugar()(*Gula){
	g.Nama = "Gula"
	g.Volume = 12
	g.Harga = 50000

	return g
}