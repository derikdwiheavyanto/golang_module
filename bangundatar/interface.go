package bangundatar

type Hitung interface {
	Hitung2d
	Hitung3d
}

type Hitung2d interface {
	Luas() float64
	Keliling() float64
}

type Hitung3d interface {
	Volume() float64
}

type Lingkaran struct {
	Diameter float64
}

type Persegi struct {
	Sisi float64
}

type Segitiga struct {
	Alas   float64
	Tinggi float64
}

type Kubus struct {
	Sisi float64
}
