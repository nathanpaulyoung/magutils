package magutils

//NewMag returns a pointer to a fully initialized basic Mag
func NewMag() *Mag {
	mag := new(Mag)
	mag.species = basicMag
	mag.cost = 0
	mag.sync = 20
	mag.iq = 0
	mag.def = 5.0
	mag.pow = 0
	mag.dex = 0
	mag.mind = 0
	mag.pb1 = "none"
	mag.pb2 = "none"
	mag.pb3 = "none"
	return mag
}
