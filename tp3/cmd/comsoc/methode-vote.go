package comsoc

type Alternative int
type Profile [][]Alternative
type Count map[Alternative]int

// renvoie l'indice ou se trouve alt dans prefs
func rank(alt Alternative, prefs []Alternative) int {
	for pos, val := range prefs {
		if alt == val {
			return pos
		}
	}
	return -1 // Si alt n'est pas dans prefs
}
