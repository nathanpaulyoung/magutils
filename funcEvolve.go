package magutils

import "math"

//Evolve applies evolutions when appropriate conditions are met
func (mag *Mag) Evolve(s sectionid, c class) {
	def := int(math.Floor(mag.def))
	pow := int(math.Floor(mag.pow))
	dex := int(math.Floor(mag.dex))
	mind := int(math.Floor(mag.mind))
	level := int(def + pow + dex + mind)

	//Level 10 basic evolution
	if level == 10 && mag.species == basicMag {
		switch {
		case c >= 0 && c <= 3:
			mag.species = varuna
		case c >= 4 && c <= 7:
			mag.species = kalki
		case c >= 8 && c <= 11:
			mag.species = vritra
		}
		mag.pb1 = mag.species.pb
	}

	//Level 35 Hunter evolution
	if level == 35 && mag.species == varuna {
		switch {
		case pow > dex && pow > mind:
			mag.species = rudra
		case dex > pow && dex > mind:
			mag.species = marutah
		case mind > pow && mind > dex:
			mag.species = vayu
		}
		mag.pb2 = mag.species.pb
	}

	//Level 35 Ranger evolution
	if level == 35 && mag.species == kalki {
		switch {
		case pow > dex && pow > mind:
			mag.species = surya
		case dex > pow && dex > mind:
			mag.species = mitra
		case mind > pow && mind > dex:
			mag.species = tapas
		}
		mag.pb2 = mag.species.pb
	}

	//Level 35 Force evolution
	if level == 35 && mag.species == vritra {
		switch {
		case pow > dex && pow > mind:
			mag.species = rudra
		case dex > pow && dex > mind:
			mag.species = marutah
		case mind > pow && mind > dex:
			mag.species = vayu
		}
		mag.pb2 = mag.species.pb
	}

	if level >= 50 && level < 100 && level%5 == 0 {
		switch {
		case c >= 8 && c <= 11 && dex >= 45:
			switch {
			case pow > dex+mind:
				mag.species = andhaka
			case mind > pow+dex || dex > pow+mind:
				mag.species = bana
			}
		case pow > dex && dex > mind:
			switch {
			case (c >= 0 || c <= 3) && s%2 == 0:
				mag.species = varaha
			case (c >= 0 || c <= 3) && s%2 == 1:
				mag.species = kama
			case (c >= 4 || c <= 7) && s%2 == 0:
				mag.species = kama
			case (c >= 4 || c <= 7) && s%2 == 1:
				mag.species = madhu
			case (c >= 8 || c <= 11) && s%2 == 0:
				mag.species = naraka
			case (c >= 8 || c <= 11) && s%2 == 1:
				mag.species = marica
			}
		case pow > mind && mind > dex:
			switch {
			case (c >= 0 || c <= 3) && s%2 == 0:
				mag.species = bhirava
			case (c >= 0 || c <= 3) && s%2 == 1:
				mag.species = apsaras
			case (c >= 4 || c <= 7) && s%2 == 0:
				mag.species = bhirava
			case (c >= 4 || c <= 7) && s%2 == 1:
				mag.species = kaitabha
			case (c >= 8 || c <= 11) && s%2 == 0:
				mag.species = ravana
			case (c >= 8 || c <= 11) && s%2 == 1:
				mag.species = naga
			}
		case dex > pow && pow > mind:
			switch {
			case (c >= 0 || c <= 3) && s%2 == 0:
				mag.species = ila
			case (c >= 0 || c <= 3) && s%2 == 1:
				mag.species = garuda
			case (c >= 4 || c <= 7) && s%2 == 0:
				mag.species = bhirava
			case (c >= 4 || c <= 7) && s%2 == 1:
				mag.species = kaitabha
			case (c >= 8 || c <= 11) && s%2 == 0:
				mag.species = ribhava
			case (c >= 8 || c <= 11) && s%2 == 1:
				mag.species = garuda
			}
		case dex > mind && mind > pow:
			switch {
			case (c >= 0 || c <= 3) && s%2 == 0:
				mag.species = nandin
			case (c >= 0 || c <= 3) && s%2 == 1:
				mag.species = yaksa
			case (c >= 4 || c <= 7) && s%2 == 0:
				mag.species = kama
			case (c >= 4 || c <= 7) && s%2 == 1:
				mag.species = varaha
			case (c >= 8 || c <= 11) && s%2 == 0:
				mag.species = sita
			case (c >= 8 || c <= 11) && s%2 == 1:
				mag.species = bhirava
			}
		case mind > pow && pow > dex:
			switch {
			case (c >= 0 || c <= 3) && s%2 == 0:
				mag.species = kabanda
			case (c >= 0 || c <= 3) && s%2 == 1:
				mag.species = bana
			case (c >= 4 || c <= 7) && s%2 == 0:
				mag.species = varaha
			case (c >= 4 || c <= 7) && s%2 == 1:
				mag.species = kabanda
			case (c >= 8 || c <= 11) && s%2 == 0:
				mag.species = naga
			case (c >= 8 || c <= 11) && s%2 == 1:
				mag.species = kumara
			}
		case mind > dex && dex > pow:
			switch {
			case (c >= 0 || c <= 3) && s%2 == 0:
				mag.species = ushasu
			case (c >= 0 || c <= 3) && s%2 == 1:
				mag.species = soma
			case (c >= 4 || c <= 7) && s%2 == 0:
				mag.species = apsaras
			case (c >= 4 || c <= 7) && s%2 == 1:
				mag.species = durga
			case (c >= 8 || c <= 11) && s%2 == 0:
				mag.species = kabanda
			case (c >= 8 || c <= 11) && s%2 == 1:
				mag.species = ila
			}
		case pow == dex && dex > mind:
			switch {
			case (c >= 0 || c <= 3) && s%2 == 0:
				mag.species = varaha
			case (c >= 0 || c <= 3) && s%2 == 1:
				mag.species = kama
			case (c >= 4 || c <= 7) && s%2 == 0:
				mag.species = bhirava
			case (c >= 4 || c <= 7) && s%2 == 1:
				mag.species = kaitabha
			case (c >= 8 || c <= 11) && s%2 == 0:
				mag.species = naga
			case (c >= 8 || c <= 11) && s%2 == 1:
				mag.species = kumara
			}
		case pow == mind && mind > dex:
			switch {
			case (c >= 0 || c <= 3) && s%2 == 0:
				mag.species = bhirava
			case (c >= 0 || c <= 3) && s%2 == 1:
				mag.species = apsaras
			case (c >= 4 || c <= 7) && s%2 == 0:
				mag.species = kama
			case (c >= 4 || c <= 7) && s%2 == 1:
				mag.species = varaha
			case (c >= 8 || c <= 11) && s%2 == 0:
				mag.species = naga
			case (c >= 8 || c <= 11) && s%2 == 1:
				mag.species = kumara
			}
		case dex == mind && mind > pow:
			switch {
			case (c >= 0 || c <= 3) && s%2 == 0:
				mag.species = varaha
			case (c >= 0 || c <= 3) && s%2 == 1:
				mag.species = kama
			case (c >= 4 || c <= 7) && s%2 == 0:
				mag.species = kama
			case (c >= 4 || c <= 7) && s%2 == 1:
				mag.species = varaha
			case (c >= 8 || c <= 11) && s%2 == 0:
				mag.species = kabanda
			case (c >= 8 || c <= 11) && s%2 == 1:
				mag.species = ila
			}
		}
		if mag.pb1 != mag.species.pb && mag.pb2 != mag.species.pb && mag.pb3 == "none" {
			mag.pb3 = mag.species.pb
		}
	}

	if level >= 100 && level%10 == 0 {
		switch {
		case c == 0 || c == 2:
			switch {
			case def+dex == pow+mind && (s == 0 || s == 6 || s == 3 || s == 9):
				mag.species = deva
			case def+pow == dex+mind && (s == 2 || s == 8 || s == 5):
				mag.species = rati
			case def+mind == pow+dex && (s == 4 || s == 1 || s == 7):
				mag.species = rati
			}
		case c == 1 || c == 3:
			switch {
			case def+dex == pow+mind && (s == 0 || s == 6 || s == 3 || s == 9):
				mag.species = savitri
			case def+pow == dex+mind && (s == 2 || s == 8 || s == 5):
				mag.species = savitri
			case def+mind == pow+dex && (s == 4 || s == 1 || s == 7):
				mag.species = savitri
			}
		case c == 4 || c == 6:
			switch {
			case def+dex == pow+mind && (s == 0 || s == 6 || s == 3 || s == 9):
				mag.species = pushan
			case def+pow == dex+mind && (s == 2 || s == 8 || s == 5):
				mag.species = pushan
			case def+mind == pow+dex && (s == 4 || s == 1 || s == 7):
				mag.species = pushan
			}
		case c == 5 || c == 7:
			switch {
			case def+dex == pow+mind && (s == 0 || s == 6 || s == 3 || s == 9):
				mag.species = rukmin
			case def+pow == dex+mind && (s == 2 || s == 8 || s == 5):
				mag.species = diwari
			case def+mind == pow+dex && (s == 4 || s == 1 || s == 7):
				mag.species = diwari
			}
		case c == 8 || c == 10:
			switch {
			case def+dex == pow+mind && (s == 0 || s == 6 || s == 3 || s == 9):
				mag.species = nidra
			case def+pow == dex+mind && (s == 2 || s == 8 || s == 5):
				mag.species = nidra
			case def+mind == pow+dex && (s == 4 || s == 1 || s == 7):
				mag.species = nidra
			}
		case c == 9 || c == 11:
			switch {
			case def+dex == pow+mind && (s == 0 || s == 6 || s == 3 || s == 9):
				mag.species = sato
			case def+pow == dex+mind && (s == 2 || s == 8 || s == 5):
				mag.species = bhima
			case def+mind == pow+dex && (s == 4 || s == 1 || s == 7):
				mag.species = bhima
			}
		}
	}
}
