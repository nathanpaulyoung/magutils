# **MagUtils** - *a library for computing PSOBB Mag growth and evolution*

## **History**

Sonic Team's *Phantasy Star Online* is a game beloved by many across many different platforms, including Dreamcast (twice), Gamecube, Xbox, and PC (also twice). The latest iteration of the game, *Phantasy Star Online: Blue Burst* was released on PC during the Summer of 2005, and although SEGA shut down the official servers in early 2008, players have kept the game and community alive with a number of [Fan Servers](http://ephinea.pioneer2.net) running open source server software, and playing home to thousands of players who still love this game to this day.

One aspect of PSOBB that is often quite challenging to handle is the highly mathematical and complex system behind Mags, small companion pets that float behind your character providing stat bonuses and powerful abilities during battle. Mags can be fed items from the player's inventory and will gain stats and abilities, level up, and evolve over the course of their growth. Figuring out how your Mag will grow from a given item or at a given stage of their development can be *very* challenging, with many players taking the time to scour the internet for elaborate tables of stat gains and trees detailing evolutionary conditions.

**MagUtils is a library designed to supplement or replace this process, making it clear and easy how Mag growth and evolution works and will progress, and guiding players toward a perfect Mag.**

*If you are interested in joining the PSOBB community, head over to [Ephinea PSOBB](http://ephinea.pioneer2.net) and check them out. They never accept money for any reason, and are very true to the official SEGA servers!*

***
## **Functions:**

#### `func NewMag() *Mag`
* `NewMag` returns a pointer to a fully initialized basic (or "baby") Mag, with Defense set to 5, Synchro set to 20, and all other stats set at zero.

#### `func (mag *Mag) Feed(n int, i items, s sectionid, c class)`
* `Feed` feeds a food item `i` to an initialized Mag and processes evolution based on SectionID `s` and Class `c` when correct conditions are met. Supports feeding `n` multiples of an item in one function call.

* Values for `i`:  
`Monomate`, `Dimate`, `Trimate`, `Monofluid`, `Difluid`, `Trifluid`, `Antidote`, `Antiparalysis`, `Solatomizer`, `Moonatomizer`, `Staratomizer`

* Values for `s`:  
`Viridia`, `Greenill`, `Skyly`, `Bluefull`, `Purplenum`, `Pinkal`, `Redria`, `Oran`, `Yellowboze`, `Whitill`

* Values for `c`:  
`Humar`, `Hunewearl`, `Hucast`, `Hucaseal`, `Ramar`, `Ramarl`, `Racast`, `Racaseal`, `Fomar`, `Fomarl`, `Fonewm`, `Fonewearl`

#### `func (mag *Mag) Evolve(s sectionid, c class)`
* `Evolve` is called by `Feed` automatically, but may be called independently on a properly initialized mag to process evolution based on SectionID `s` and Class `c` when correct conditions are met.

* Values for `s`:  
`Viridia`, `Greenill`, `Skyly`, `Bluefull`, `Purplenum`, `Pinkal`, `Redria`, `Oran`, `Yellowboze`, `Whitill`

* Values for `c`:  
`Humar`, `Hunewearl`, `Hucast`, `Hucaseal`, `Ramar`, `Ramarl`, `Racast`, `Racaseal`, `Fomar`, `Fomarl`, `Fonewm`, `Fonewearl`

#### `func (mag *Mag) SetStats(sync int, iq int, def float64, pow float64, dex float64, mind float64) (syncerr bool, iqerr bool, staterr bool)`
* `SetStats` should never NEED to be used, unless you want to allow sandbox testing of mags, or the ability to save/import mags. Returns three bools representing success or failure.

#### `func (mag *Mag) SetPB(pb blast, pos int)`
* `SetPB` is, again, just for sandboxing or importing mags.

* Values for `pb`:  
`None`, `Estlla`, `Leilla`, `Farlla`, `Pilla`, `Golla`, `Myllayoulla`

#### `func (mag *Mag) SetSpecies(species mags)`
* `SetSpecies` is still just for sandboxing or importing mags.

* Values for `species`:  
`BasicMag`, `Agastya`, `Andhaka`, `Apsaras`, `Ashvinau`, `Bana`, `Bhima`, `Bhirava`, `Chao`, `Chuchu`, `Churel`, `Deva`, `Diwari`, `Durga`, `Gaelgill`, `Garuda`, `Geungsi`, `Ila`, `Kabanda`, `Kaitabha`, `Kalki`, `Kama`, `Kapukapu`, `Kumara`, `Madhu`, `Marica`, `Marutah`, `Mitra`, `Moro`, `Naga`, `Namuci`, `Nandin`, `Naraka`, `Nidra`, `Opaopa`, `Pian`, `Pioneer2`, `Pitri`, `Preta`, `Pushan`, `Puyo`, `Rappy`, `Rati`, `Ravana`, `Ribhava`, `Robochao`, `Rudra`, `Rukmin`, `Sato`, `Savitri`, `Sita`, `Soma`, `Soniti`, `Strikerunit`, `Sumba`, `Surya`, `Tapas`, `Tellusis`, `Ushasu`, `Varaha`, `Varuna`, `Vayu`, `Vritra`, `Yahoo`, `Yaksa`

## **Mag Object Properties**
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

## **Usage Example:**
`mag := NewMag()`  
`mag.Feed(13, Antidote, Purplenum, Racast)`  
`println(mag.species.name)`

Will output:  
`Kalki`

## **To do:**
* ~~Finish funcSetMagParams.go~~
* Make `Feed` return `true` on evolution and `false` on no evolution.
* Make `Feed` break on evolution and return text indicating how many items were fed.
* Implement feed/evolution history.
* Implement Undo/stepback function.
