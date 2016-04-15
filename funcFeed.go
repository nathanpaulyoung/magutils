package magutils

import "math"

//Feed adds the parameters of a given item i a set number of n times,
//and applies evolutions when applicable based on sectionid s and
//class c.
func (mag *Mag) Feed(n int, i items, s sectionid, c class) {
	for iter := 0; iter < n; iter++ {
		mag.feed(i)
		mag.Evolve(s, c)
	}
}

func (mag *Mag) feed(feedItem items) {
	var i item
	switch {
	case feedItem == Monomate:
		i = mag.species.group.monomate
	case feedItem == Dimate:
		i = mag.species.group.dimate
	case feedItem == Trimate:
		i = mag.species.group.trimate
	case feedItem == Monofluid:
		i = mag.species.group.monofluid
	case feedItem == Difluid:
		i = mag.species.group.difluid
	case feedItem == Trifluid:
		i = mag.species.group.trifluid
	case feedItem == Antidote:
		i = mag.species.group.antidote
	case feedItem == Antiparalysis:
		i = mag.species.group.antiparalysis
	case feedItem == Solatomizer:
		i = mag.species.group.solatomizer
	case feedItem == Moonatomizer:
		i = mag.species.group.moonatomizer
	case feedItem == Staratomizer:
		i = mag.species.group.staratomizer
	}

	mag.cost += i.cost

	if mag.sync+i.sync <= 120 && mag.sync+i.sync >= 0 {
		mag.sync += i.sync
	} else if mag.sync+i.sync > 120 {
		mag.sync = 120
	} else if mag.sync+i.sync < 0 {
		mag.sync = 0
	}

	if mag.iq+i.iq <= 200 && mag.iq+i.iq >= 0 {
		mag.iq += i.iq
	} else if mag.iq+i.iq > 200 {
		mag.iq = 200
	} else if mag.iq+i.iq < 0 {
		mag.iq = 0
	}

	if mag.def+i.def+math.Floor(mag.pow)+math.Floor(mag.dex)+math.Floor(mag.mind) > 200 {
		mag.def = math.Ceil(mag.def)
	} else if mag.def+i.def >= math.Floor(mag.def) {
		mag.def += i.def
	} else if mag.def+i.def < math.Floor(mag.def) {
		mag.def = math.Floor(mag.def)
	}

	if math.Floor(mag.def)+mag.pow+i.pow+math.Floor(mag.dex)+math.Floor(mag.mind) > 200 {
		mag.pow = math.Ceil(mag.pow)
	} else if mag.pow+i.pow >= math.Floor(mag.pow) {
		mag.pow += i.pow
	} else if mag.pow+i.pow < math.Floor(mag.pow) {
		mag.pow = math.Floor(mag.pow)
	}

	if math.Floor(mag.def)+math.Floor(mag.pow)+mag.dex+i.dex+math.Floor(mag.mind) > 200 {
		mag.dex = math.Ceil(mag.dex)
	} else if mag.dex+i.dex >= math.Floor(mag.dex) {
		mag.dex += i.dex
	} else if mag.dex+i.dex < math.Floor(mag.dex) {
		mag.dex = math.Floor(mag.dex)
	}

	if math.Floor(mag.def)+math.Floor(mag.pow)+math.Floor(mag.dex)+mag.mind+i.mind > 200 {
		mag.mind = math.Ceil(mag.mind)
	} else if mag.mind+i.mind >= math.Floor(mag.mind) {
		mag.mind += i.mind
	} else if mag.mind+i.mind < math.Floor(mag.mind) {
		mag.mind = math.Floor(mag.mind)
	}
}
