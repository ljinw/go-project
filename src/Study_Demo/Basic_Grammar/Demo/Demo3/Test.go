package main

func Iota_Demo()(i any...){
	const (
		a = iota
		b = iota
		c = iota
		d = 0
		e = iota
		f = iota + 1
	)
	return a, b, c, d, e, f
}
