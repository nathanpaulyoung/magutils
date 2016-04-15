package magutils

type items int

//Constant identifiers for each item for use in the Feed function.
const (
	Monomate items = iota
	Dimate
	Trimate
	Monofluid
	Difluid
	Trifluid
	Antidote
	Antiparalysis
	Solatomizer
	Moonatomizer
	Staratomizer
)

type item struct {
	name string
	cost int
	sync int
	iq   int
	def  float64
	pow  float64
	dex  float64
	mind float64
}
