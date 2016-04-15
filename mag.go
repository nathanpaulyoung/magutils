package magutils

type mags int

//Constant identifiers for each mag for use in the SetSpecies function.
const (
	BasicMag mags = iota
	Agastya
	Andhaka
	Apsaras
	Ashvinau
	Bana
	Bhima
	Bhirava
	Chao
	Chuchu
	Churel
	Deva
	Diwari
	Durga
	Gaelgill
	Garuda
	Geungsi
	Ila
	Kabanda
	Kaitabha
	Kalki
	Kama
	Kapukapu
	Kumara
	Madhu
	Marica
	Marutah
	Mitra
	Moro
	Naga
	Namuci
	Nandin
	Naraka
	Nidra
	Opaopa
	Pian
	Pioneer2
	Pitri
	Preta
	Pushan
	Puyo
	Rappy
	Rati
	Ravana
	Ribhava
	Robochao
	Rudra
	Rukmin
	Sato
	Savitri
	Sita
	Soma
	Soniti
	Strikerunit
	Sumba
	Surya
	Tapas
	Tellusis
	Ushasu
	Varaha
	Varuna
	Vayu
	Vritra
	Yahoo
	Yaksa
)

//Mag defines all parameters of a Mag incrementally during the simulation
type Mag struct {
	species magSpecies
	cost    int
	sync    int
	iq      int
	def     float64
	pow     float64
	dex     float64
	mind    float64
	pb1     blast
	pb2     blast
	pb3     blast
}
