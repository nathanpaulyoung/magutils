package magutils

type magSpecies struct {
	name       string
	group      magGroup
	pb         string
	actBarFill string
	actTenthHP string
	actDeath   string
	actBoss    string
}

var basicMag = magSpecies{
	name:       "Mag",
	group:      magGroupMag,
	pb:         "none",
	actBarFill: "none",
	actTenthHP: "none",
	actDeath:   "none",
	actBoss:    "none",
}

var agastya = magSpecies{
	name:       "Agastya",
	group:      magGroupCell,
	pb:         "none",
	actBarFill: "Invulnerability",
	actTenthHP: "Resta",
	actDeath:   "Reverser",
	actBoss:    "Invulnerability",
}

var andhaka = magSpecies{
	name:       "Andhaka",
	group:      magGroupE,
	pb:         "Estlla",
	actBarFill: "Invulnerability",
	actTenthHP: "Resta",
	actDeath:   "Reverser",
	actBoss:    "Shifta/Deband",
}

var apsaras = magSpecies{
	name:       "Apsaras",
	group:      magGroupC,
	pb:         "Estlla",
	actBarFill: "Invulnerability",
	actTenthHP: "Invulnerability",
	actDeath:   "Reverser",
	actBoss:    "Invulnerability",
}

var ashvinau = magSpecies{
	name:       "Ashvinau",
	group:      magGroupA,
	pb:         "Pilla",
	actBarFill: "Invulnerability",
	actTenthHP: "Shifta/Deband",
	actDeath:   "none",
	actBoss:    "none",
}

var bana = magSpecies{
	name:       "Bana",
	group:      magGroupE,
	pb:         "Estlla",
	actBarFill: "Invulnerability",
	actTenthHP: "Invulnerability",
	actDeath:   "Reverser",
	actBoss:    "Invulnerability",
}

var bhima = magSpecies{
	name:       "Bhima",
	group:      magGroup4thA,
	pb:         "none",
	actBarFill: "Shifta/Deband",
	actTenthHP: "Invulnerability",
	actDeath:   "Reverser",
	actBoss:    "Invulnerability",
}

var bhirava = magSpecies{
	name:       "Bhirava",
	group:      magGroupC,
	pb:         "Pilla",
	actBarFill: "Invulnerability",
	actTenthHP: "Resta",
	actDeath:   "Reverser",
	actBoss:    "Shifta/Deband",
}

var chao = magSpecies{
	name:       "Chao",
	group:      magGroupCell,
	pb:         "none",
	actBarFill: "Invulnerability",
	actTenthHP: "Resta",
	actDeath:   "Reverser",
	actBoss:    "Shifta/Deband",
}

var chuchu = magSpecies{
	name:       "Chu Chu",
	group:      magGroupCell,
	pb:         "none",
	actBarFill: "Shifta/Deband",
	actTenthHP: "Invulnerability",
	actDeath:   "Reverser",
	actBoss:    "Shifta/Deband",
}

var churel = magSpecies{
	name:       "Churel",
	group:      magGroupCell,
	pb:         "none",
	actBarFill: "Resta",
	actTenthHP: "Resta",
	actDeath:   "Reverser",
	actBoss:    "Resta",
}

var deva = magSpecies{
	name:       "Deva",
	group:      magGroup4thB,
	pb:         "none",
	actBarFill: "Shifta/Deband",
	actTenthHP: "Resta",
	actDeath:   "Reverser",
	actBoss:    "Shifta/Deband",
}

var diwari = magSpecies{
	name:       "Diwari",
	group:      magGroup4thC,
	pb:         "none",
	actBarFill: "Shifta/Deband",
	actTenthHP: "Resta",
	actDeath:   "Reverser",
	actBoss:    "Shifta/Deband",
}

var durga = magSpecies{
	name:       "Durga",
	group:      magGroupD,
	pb:         "Estlla",
	actBarFill: "Invulnerability",
	actTenthHP: "Invulnerability",
	actDeath:   "Reverser",
	actBoss:    "Invulnerability",
}

var gaelgill = magSpecies{
	name:       "Gael-Gill",
	group:      magGroupCell,
	pb:         "none",
	actBarFill: "Invulnerability",
	actTenthHP: "Invulnerability",
	actDeath:   "none",
	actBoss:    "Shifta/Deband",
}

var garuda = magSpecies{
	name:       "Garuda",
	group:      magGroupD,
	pb:         "Pilla",
	actBarFill: "Resta",
	actTenthHP: "Resta",
	actDeath:   "Reverser",
	actBoss:    "Resta",
}

var geungsi = magSpecies{
	name:       "Geung-si",
	group:      magGroupCell,
	pb:         "none",
	actBarFill: "Resta",
	actTenthHP: "Invulnerability",
	actDeath:   "Reverser",
	actBoss:    "Invulnerability",
}

var ila = magSpecies{
	name:       "Ila",
	group:      magGroupD,
	pb:         "Mylla & Youlla",
	actBarFill: "none",
	actTenthHP: "Resta",
	actDeath:   "Reverser",
	actBoss:    "Shifta/Deband",
}

var kabanda = magSpecies{
	name:       "Kabanda",
	group:      magGroupE,
	pb:         "Mylla & Youlla",
	actBarFill: "Invulnerability",
	actTenthHP: "Shifta/Deband",
	actDeath:   "Reverser",
	actBoss:    "Shifta/Deband",
}

var kaitabha = magSpecies{
	name:       "Kaitabha",
	group:      magGroupC,
	pb:         "Mylla & Youlla",
	actBarFill: "Invulnerability",
	actTenthHP: "Resta",
	actDeath:   "Reverser",
	actBoss:    "Shifta/Deband",
}

var kalki = magSpecies{
	name:       "Kalki",
	group:      magGroupT1,
	pb:         "Estlla",
	actBarFill: "Invulnerability",
	actTenthHP: "none",
	actDeath:   "none",
	actBoss:    "Shifta/Deband",
}

var kama = magSpecies{
	name:       "Kama",
	group:      magGroupC,
	pb:         "Pilla",
	actBarFill: "none",
	actTenthHP: "Shifta/Deband",
	actDeath:   "Reverser",
	actBoss:    "Shifta/Deband",
}

var kapukapu = magSpecies{
	name:       "Kapu Kapu",
	group:      magGroupCell,
	pb:         "none",
	actBarFill: "Invulnerability",
	actTenthHP: "Resta",
	actDeath:   "Reverser",
	actBoss:    "Shifta/Deband",
}

var kumara = magSpecies{
	name:       "Kumara",
	group:      magGroupC,
	pb:         "Pilla",
	actBarFill: "Resta",
	actTenthHP: "Resta",
	actDeath:   "Reverser",
	actBoss:    "Resta",
}

var madhu = magSpecies{
	name:       "Madhu",
	group:      magGroupC,
	pb:         "Mylla & Youlla",
	actBarFill: "Invulnerability",
	actTenthHP: "Resta",
	actDeath:   "Reverser",
	actBoss:    "Shifta/Deband",
}

var marica = magSpecies{
	name:       "Marica",
	group:      magGroupE,
	pb:         "Pilla",
	actBarFill: "Invulnerability",
	actTenthHP: "Resta",
	actDeath:   "Reverser",
	actBoss:    "Shifta/Deband",
}

var marutah = magSpecies{
	name:       "Marutah",
	group:      magGroupA,
	pb:         "Pilla",
	actBarFill: "Invulnerability",
	actTenthHP: "Shifta/Deband",
	actDeath:   "none",
	actBoss:    "none",
}

var mitra = magSpecies{
	name:       "Mitra",
	group:      magGroupB,
	pb:         "Pilla",
	actBarFill: "Invulnerability",
	actTenthHP: "Shifta/Deband",
	actDeath:   "none",
	actBoss:    "none",
}

var moro = magSpecies{
	name:       "Moro",
	group:      magGroupCell,
	pb:         "none",
	actBarFill: "Invulnerability",
	actTenthHP: "Resta",
	actDeath:   "Reverser",
	actBoss:    "Resta",
}

var naga = magSpecies{
	name:       "Naga",
	group:      magGroupE,
	pb:         "Mylla & Youlla",
	actBarFill: "none",
	actTenthHP: "Shifta/Deband",
	actDeath:   "Reverser",
	actBoss:    "Shifta/Deband",
}

var namuci = magSpecies{
	name:       "Namuci",
	group:      magGroupA,
	pb:         "Mylla & Youlla",
	actBarFill: "Invulnerability",
	actTenthHP: "Resta",
	actDeath:   "none",
	actBoss:    "Resta",
}

var nandin = magSpecies{
	name:       "Nandin",
	group:      magGroupD,
	pb:         "Estlla",
	actBarFill: "Invulnerability",
	actTenthHP: "Invulnerability",
	actDeath:   "Reverser",
	actBoss:    "Invulnerability",
}

var naraka = magSpecies{
	name:       "Naraka",
	group:      magGroupE,
	pb:         "Golla",
	actBarFill: "Resta",
	actTenthHP: "Resta",
	actDeath:   "Reverser",
	actBoss:    "Resta",
}

var nidra = magSpecies{
	name:       "Nidra",
	group:      magGroup4thC,
	pb:         "none",
	actBarFill: "Invulnerability",
	actTenthHP: "Invulnerability",
	actDeath:   "Reverser",
	actBoss:    "Invulnerability",
}

var opaopa = magSpecies{
	name:       "Opa Opa",
	group:      magGroupCell,
	pb:         "none",
	actBarFill: "Invulnerability",
	actTenthHP: "Resta",
	actDeath:   "Reverser",
	actBoss:    "Shifta/Deband",
}

var pian = magSpecies{
	name:       "Pian",
	group:      magGroupCell,
	pb:         "none",
	actBarFill: "Invulnerability",
	actTenthHP: "Resta",
	actDeath:   "Reverser",
	actBoss:    "Shifta/Deband",
}

var pioneer2 = magSpecies{
	name:       "Pioneer 2",
	group:      magGroupCell,
	pb:         "none",
	actBarFill: "Invulnerability",
	actTenthHP: "Resta",
	actDeath:   "Reverser",
	actBoss:    "Invulnerability",
}

var pitri = magSpecies{
	name:       "Pitri",
	group:      magGroupCell,
	pb:         "none",
	actBarFill: "Resta",
	actTenthHP: "Resta",
	actDeath:   "Reverser",
	actBoss:    "Resta",
}

var preta = magSpecies{
	name:       "Preta",
	group:      magGroupCell,
	pb:         "none",
	actBarFill: "Shifta/Deband",
	actTenthHP: "Shifta/Deband",
	actDeath:   "Reverser",
	actBoss:    "Shifta/Deband",
}

var pushan = magSpecies{
	name:       "Pushan",
	group:      magGroup4thA,
	pb:         "none",
	actBarFill: "Shifta/Deband",
	actTenthHP: "Resta",
	actDeath:   "Reverser",
	actBoss:    "Shifta/Deband",
}

var puyo = magSpecies{
	name:       "Puyo",
	group:      magGroupCell,
	pb:         "none",
	actBarFill: "Shifta/Deband",
	actTenthHP: "Shifta/Deband",
	actDeath:   "Reverser",
	actBoss:    "Invulnerability",
}

var rappy = magSpecies{
	name:       "Rappy",
	group:      magGroupCell,
	pb:         "none",
	actBarFill: "Resta",
	actTenthHP: "Invulnerability",
	actDeath:   "Reverser",
	actBoss:    "Resta",
}

var rati = magSpecies{
	name:       "Rati",
	group:      magGroup4thA,
	pb:         "none",
	actBarFill: "Shifta/Deband",
	actTenthHP: "Resta",
	actDeath:   "Reverser",
	actBoss:    "Shifta/Deband",
}

var ravana = magSpecies{
	name:       "Ravana",
	group:      magGroupE,
	pb:         "Farlla",
	actBarFill: "Invulnerability",
	actTenthHP: "Shifta/Deband",
	actDeath:   "Reverser",
	actBoss:    "Shifta/Deband",
}

var ribhava = magSpecies{
	name:       "Ravana",
	group:      magGroupD,
	pb:         "Farlla",
	actBarFill: "Invulnerability",
	actTenthHP: "Invulnerability",
	actDeath:   "Reverser",
	actBoss:    "Invulnerability",
}

var robochao = magSpecies{
	name:       "Robochao",
	group:      magGroupCell,
	pb:         "none",
	actBarFill: "Invulnerability",
	actTenthHP: "Resta",
	actDeath:   "Reverser",
	actBoss:    "Shifta/Deband",
}

var rudra = magSpecies{
	name:       "Rudra",
	group:      magGroupA,
	pb:         "Golla",
	actBarFill: "Invulnerability",
	actTenthHP: "Resta",
	actDeath:   "none",
	actBoss:    "Shifta/Deband",
}

var rukmin = magSpecies{
	name:       "Rukmin",
	group:      magGroup4thB,
	pb:         "none",
	actBarFill: "Shifta/Deband",
	actTenthHP: "Resta",
	actDeath:   "Reverser",
	actBoss:    "Shifta/Deband",
}

var sato = magSpecies{
	name:       "Sato",
	group:      magGroup4thB,
	pb:         "none",
	actBarFill: "Invulnerability",
	actTenthHP: "Invulnerability",
	actDeath:   "Reverser",
	actBoss:    "Invulnerability",
}

var savitri = magSpecies{
	name:       "Savitri",
	group:      magGroup4thC,
	pb:         "none",
	actBarFill: "Shifta/Deband",
	actTenthHP: "Resta",
	actDeath:   "Reverser",
	actBoss:    "Shifta/Deband",
}

var sita = magSpecies{
	name:       "Sita",
	group:      magGroupD,
	pb:         "Pilla",
	actBarFill: "Invulnerability",
	actTenthHP: "Resta",
	actDeath:   "Reverser",
	actBoss:    "Shifta/Deband",
}

var soma = magSpecies{
	name:       "Soma",
	group:      magGroupD,
	pb:         "Estlla",
	actBarFill: "Invulnerability",
	actTenthHP: "Invulnerability",
	actDeath:   "Reverser",
	actBoss:    "Invulnerability",
}

var soniti = magSpecies{
	name:       "Soniti",
	group:      magGroupCell,
	pb:         "none",
	actBarFill: "Invulnerability",
	actTenthHP: "Invulnerability",
	actDeath:   "Reverser",
	actBoss:    "Invulnerability",
}

var strikerunit = magSpecies{
	name:       "Striker Unit",
	group:      magGroupCell,
	pb:         "none",
	actBarFill: "Resta",
	actTenthHP: "Resta",
	actDeath:   "Reverser",
	actBoss:    "Resta",
}

var sumba = magSpecies{
	name:       "Sumba",
	group:      magGroupA,
	pb:         "Golla",
	actBarFill: "Invulnerability",
	actTenthHP: "Resta",
	actDeath:   "none",
	actBoss:    "Shifta/Deband",
}

var surya = magSpecies{
	name:       "Surya",
	group:      magGroupB,
	pb:         "Golla",
	actBarFill: "Invulnerability",
	actTenthHP: "Resta",
	actDeath:   "none",
	actBoss:    "Shifta/Deband",
}

var tapas = magSpecies{
	name:       "Tapas",
	group:      magGroupB,
	pb:         "Mylla & Youlla",
	actBarFill: "Invulnerability",
	actTenthHP: "Resta",
	actDeath:   "none",
	actBoss:    "Resta",
}

var tellusis = magSpecies{
	name:       "Tellusis",
	group:      magGroupCell,
	pb:         "none",
	actBarFill: "Invulnerability",
	actTenthHP: "Invulnerability",
	actDeath:   "Reverser",
	actBoss:    "Invulnerability",
}

var ushasu = magSpecies{
	name:       "Ushasu",
	group:      magGroupC,
	pb:         "Golla",
	actBarFill: "Invulnerability",
	actTenthHP: "Resta",
	actDeath:   "Reverser",
	actBoss:    "Shifta/Deband",
}

var varaha = magSpecies{
	name:       "Varaha",
	group:      magGroupC,
	pb:         "Golla",
	actBarFill: "Invulnerability",
	actTenthHP: "Resta",
	actDeath:   "Reverser",
	actBoss:    "Shifta/Deband",
}

var varuna = magSpecies{
	name:       "Varuna",
	group:      magGroupC,
	pb:         "Farlla",
	actBarFill: "Invulnerability",
	actTenthHP: "none",
	actDeath:   "none",
	actBoss:    "Shifta/Deband",
}

var vayu = magSpecies{
	name:       "Vayu",
	group:      magGroupC,
	pb:         "Mylla & Youlla",
	actBarFill: "Invulnerability",
	actTenthHP: "Resta",
	actDeath:   "none",
	actBoss:    "Resta",
}

var vritra = magSpecies{
	name:       "Vritra",
	group:      magGroupT1,
	pb:         "Leilla",
	actBarFill: "Invulnerability",
	actTenthHP: "none",
	actDeath:   "none",
	actBoss:    "Resta",
}

var yahoo = magSpecies{
	name:       "Yahoo!",
	group:      magGroupCell,
	pb:         "none",
	actBarFill: "Invulnerability",
	actTenthHP: "none",
	actDeath:   "Reverser",
	actBoss:    "Invulnerability",
}

var yaksa = magSpecies{
	name:       "Yaksa",
	group:      magGroupD,
	pb:         "Golla",
	actBarFill: "Invulnerability",
	actTenthHP: "Resta",
	actDeath:   "Reverser",
	actBoss:    "Shifta/Deband",
}
