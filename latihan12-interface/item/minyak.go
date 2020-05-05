package item

type Minyak struct{
	Nama string
	Volume int
	Harga float32
}

func (m *Minyak)Oil()(*Minyak){
	m.Nama = "Minyak"
	m.Volume = 10
	m.Harga = 120000.0

	return m
}