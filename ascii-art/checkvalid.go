package ascii_art

//IsValidSym - function for checking the validity of characters
func IsValidSym(s string) bool {
	for _, s := range s {
		if ((s < 32 && s != 10) || s > 126) && s != '\r' {
			return false
		}
	}
	return true
}
