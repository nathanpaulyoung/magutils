package magutils

import "math"

//SetStats sets a mag's stats
func (mag *Mag) SetStats(sync int, iq int, def float64, pow float64, dex float64, mind float64) (syncerr bool, iqerr bool, staterr bool) {
	if sync >= 0 && sync <= 120 {
		mag.sync = sync
		syncerr = true
	} else {
		syncerr = false
	}

	if iq >= 0 && iq <= 200 {
		mag.iq = iq
		iqerr = true
	} else {
		iqerr = false
	}

	level := math.Floor(def) + math.Floor(pow) + math.Floor(dex) + math.Floor(mind)
	if level >= 0 && level <= 200 {
		mag.def = def
		mag.pow = pow
		mag.dex = dex
		mag.mind = mind
		staterr = true
	} else {
		staterr = false
	}

	return
}

//SetPB sets a mag's Photon Blast
func (mag *Mag) SetPB(pb blast, pos int) {
	switch {
	case pos == 1:
		mag.pb1 = pb
	case pos == 2:
		mag.pb2 = pb
	case pos == 3:
		mag.pb3 = pb
	}
}

//SetSpecies sets a mag's species
func (mag *Mag) SetSpecies(species mags) {
	switch {
	case species == BasicMag:
		mag.species = basicMag
	case species == Agastya:
		mag.species = agastya
	case species == Andhaka:
		mag.species = andhaka
	case species == Apsaras:
		mag.species = apsaras
	case species == Ashvinau:
		mag.species = ashvinau
	case species == Bana:
		mag.species = bana
	case species == Bhima:
		mag.species = bhima
	case species == Bhirava:
		mag.species = bhirava
	case species == Chao:
		mag.species = chao
	case species == Chuchu:
		mag.species = chuchu
	case species == Churel:
		mag.species = churel
	case species == Deva:
		mag.species = deva
	case species == Diwari:
		mag.species = diwari
	case species == Durga:
		mag.species = durga
	case species == Gaelgill:
		mag.species = gaelgill
	case species == Garuda:
		mag.species = garuda
	case species == Ila:
		mag.species = ila
	case species == Kabanda:
		mag.species = kabanda
	case species == Kaitabha:
		mag.species = kaitabha
	case species == Kalki:
		mag.species = kalki
	case species == Kama:
		mag.species = kama
	case species == Kapukapu:
		mag.species = kapukapu
	case species == Kumara:
		mag.species = kumara
	case species == Madhu:
		mag.species = madhu
	case species == Marica:
		mag.species = marica
	case species == Marutah:
		mag.species = marutah
	case species == Mitra:
		mag.species = mitra
	case species == Moro:
		mag.species = moro
	case species == Naga:
		mag.species = naga
	case species == Namuci:
		mag.species = namuci
	case species == Nandin:
		mag.species = nandin
	case species == Naraka:
		mag.species = naraka
	case species == Nidra:
		mag.species = nidra
	case species == Opaopa:
		mag.species = opaopa
	case species == Pian:
		mag.species = pian
	case species == Pioneer2:
		mag.species = pioneer2
	case species == Pitri:
		mag.species = pitri
	case species == Preta:
		mag.species = preta
	case species == Pushan:
		mag.species = pushan
	case species == Puyo:
		mag.species = puyo
	case species == Rappy:
		mag.species = rappy
	case species == Rati:
		mag.species = rati
	case species == Ravana:
		mag.species = ravana
	case species == Ribhava:
		mag.species = ribhava
	case species == Robochao:
		mag.species = robochao
	case species == Rudra:
		mag.species = rudra
	case species == Rukmin:
		mag.species = rukmin
	case species == Sato:
		mag.species = sato
	case species == Savitri:
		mag.species = savitri
	case species == Sita:
		mag.species = sita
	case species == Soma:
		mag.species = soma
	case species == Soniti:
		mag.species = soniti
	case species == Strikerunit:
		mag.species = strikerunit
	case species == Sumba:
		mag.species = sumba
	case species == Surya:
		mag.species = surya
	case species == Tapas:
		mag.species = tapas
	case species == Tellusis:
		mag.species = tellusis
	case species == Ushasu:
		mag.species = ushasu
	case species == Varaha:
		mag.species = varaha
	case species == Varuna:
		mag.species = varuna
	case species == Vayu:
		mag.species = vayu
	case species == Vritra:
		mag.species = vritra
	case species == Yahoo:
		mag.species = yahoo
	case species == Yaksa:
		mag.species = yaksa
	}
}
