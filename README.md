#**MagUtils** - a library for computing PSOBB mag growth and evolution

##**Functions:**

###`func NewMag() *Mag`
`NewMag` returns a pointer to a fully initialized basic (or "baby") Mag, with Defense set to 0, Synchro set to 20, and all other stats set at zero.

###`func (mag *Mag) Feed(n int, i items, s sectionid, c class)`
`Feed` feeds a food item `i` to an initialized Mag and processes evolution based on SectionID `s` and Class `c` when correct conditions are met. Supports feeding `n` multiples of an item in one function call.

**Values for `i`:**  
`Monomate`, `Dimate`, `Trimate`, `Monofluid`, `Difluid`, `Trifluid`, `Antidote`, `Antiparalysis`, `Solatomizer`, `Moonatomizer`, `Staratomizer`

**Values for `s`:**  
`Viridia`, `Greenill`, `Skyly`, `Bluefull`, `Purplenum`, `Pinkal`, `Redria`, `Oran`, `Yellowboze`, `Whitill`

**Values for `c`:**  
`Humar`, `Hunewearl`, `Hucast`, `Hucaseal`, `Ramar`, `Ramarl`, `Racast`, `Racaseal`, `Fomar`, `Fomarl`, `Fonewm`, `Fonewearl`

###`func (mag *Mag) Evolve(s sectionid, c class)`
`Evolve` is called by `Feed` automatically, but may be called independently on a properly initialized mag to process evolution based on SectionID `s` and Class `c` when correct conditions are met.

**Values for `s`:**  
`Viridia`, `Greenill`, `Skyly`, `Bluefull`, `Purplenum`, `Pinkal`, `Redria`, `Oran`, `Yellowboze`, `Whitill`

**Values for `c`:**  
`Humar`, `Hunewearl`, `Hucast`, `Hucaseal`, `Ramar`, `Ramarl`, `Racast`, `Racaseal`, `Fomar`, `Fomarl`, `Fonewm`, `Fonewearl`

###`func (mag *Mag) SetStats(sync int, iq int, def float64, pow float64, dex float64, mind float64)`
`SetStats` should never NEED to be used, unless you want to allow sandbox testing of mags, or the ability to save/import mags.

##**Mag Object Properties**
* `mag.species.name`  
*The name of the mag's current species.*
* `mag.species.group.name`  
*The name of the mag's current mag feeding group.*
* `mag.species.group.<item>.name`  
*The name of \<item\>.*
* `mag.species.group.<item>.cost`  
*The cost in Meseta of \<item\>.*
* `mag.species.group.<item>.sync`  
*The contribution to Synchro of \<item\>.*
* `mag.species.group.<item>.iq`  
*The contribution to IQ of \<item\>.*
* `mag.species.group.<item>.def`  
*The contribution to Def of \<item\>, represented as a float64.*
* `mag.species.group.<item>.pow`  
*The contribution to Pow of \<item\>, represented as a float64.*
* `mag.species.group.<item>.dex`  
*The contribution to Dex of \<item\>, represented as a float64.*
* `mag.species.group.<item>.mind`  
*The contribution to Mind of \<item\>, represented as a float64.*
* `mag.species.pb`  
*The contribution to the mag's list of Photon Blasts, granted by the species upon evolution.*
* `mag.species.actBarFill`  
*The mag's activation effect upon filling the PB Meter, granted by the species upon evolution.*
* `mag.species.actTenthHP`  
*The mag's activation effect upon falling to 1/10th HP, granted by the species upon evolution.*
* `mag.species.actDeath`  
*The mag's activation effect upon player death, granted by the species upon evolution.*
* `mag.species.actBoss`  
*The mag's activation effect upon entering a boss room, granted by the species upon evolution.*
* `mag.cost`  
*The mag's running total cost, incremented by mag.species.group.<item>.cost in Feed function.*
* `mag.sync`  
*The mag's current Synchro stat, represented as a float64.*
* `mag.iq`  
*The mag's current IQ stat, represented as a float64.*
* `mag.def`  
*The mag's current Def stat, represented as a float64.*
* `mag.pow`  
*The mag's current Pow stat, represented as a float64.*
* `mag.dex`  
*The mag's current Dex stat, represented as a float64.*
* `mag.mind`  
*The mag's current Mind stat, represented as a float64.*
* `mag.pb1`  
*The mag's first (green) Photon Blast, granted by mag.species.pb1 in Evolve function.*
* `mag.pb2`  
*The mag's second (blue) Photon Blast, granted by mag.species.pb2 in Evolve function.*
* `mag.pb3`  
*The mag's third (red) Photon Blast, granted by mag.species.pb3 in Evolve function.*

##**Usage Example:**
`mag := NewMag()`  
`mag.Feed(13, Antidote, Purplenum, Racast)`  
`println(mag.species.name)`

Will output:  
`Kalki`

##**To do:**
* Finish funcSetMagParams.go
* Make `Feed` return `true` on evolution and `false` on no evolution.
* Make `Feed` break on evolution and return text indicating how many items were fed.
* Implement feed/evolution history.
* Implement Undo/stepback function.
