package ascii_art

//ConsoleOutput - Outputting characters to the therminal
func ConsoleOutput(art ASCIIArt) string {
	res := ""
	for _, s := range art.Str {
		// if there is empty line, prints newline
		if s != "" {
			for i := 0; i < 8; i++ {
				for j := 0; j < len(s); j++ {
					res += art.Fileinfo[int(s[j]-32)*9+1+i]
				}
				res += "\n"
			}
		} else {
			res += "\n"
		}
	}
	return res
}
