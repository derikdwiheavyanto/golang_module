package bangundatar

import "math"

func (l Lingkaran) JariJari() float64 {
	return l.Diameter / 2
}

func (l Lingkaran) Luas() float64 {
	return math.Pi * math.Pow(l.Diameter, 2)
}

func (l Lingkaran) Keliling() float64 {
	return math.Pi * l.Diameter
}

func (p Persegi) Luas() float64 {
	return math.Pow(p.Sisi, 2)
}

func (p Persegi) Keliling() float64 {
	return 4 * p.Sisi
}

func (s Segitiga) Luas() float64 {
	return (1.0 / 2.0) * s.Alas * s.Tinggi
}

func (s Segitiga) SisiMiring() float64 {
	return math.Sqrt(math.Pow(s.Alas, 2) + math.Pow(s.Tinggi, 2))
}

func (s Segitiga) Keliling() float64 {
	return s.Alas + s.Tinggi + s.SisiMiring()
}

func (k Kubus) Luas() float64 {
	return math.Pow(k.Sisi, 2) * 6
}

func (k Kubus) Keliling() float64 {
	return k.Sisi * 12
}

func (k Kubus) Volume() float64 {
	return math.Pow(k.Sisi, 3)
}
