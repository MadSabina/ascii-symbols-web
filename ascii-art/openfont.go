package ascii_art

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
)

type ASCIIArt struct {
	Argument string
	Str      []string
	Fileinfo []string
	Format   string
}

//ASCII - ascii art realization, returns the finished string for html output
func ASCII(input, format string) (string, int) {
	art := ASCIIArt{
		Argument: input,
		Format:   format,
	}

	if !IsValidSym(art.Argument) {
		return "", http.StatusBadRequest
	}

	art.Str = strings.Split(art.Argument, "\r\n")

	var err int
	art.Fileinfo, err = OpenFont(art)
	if err != http.StatusOK {
		return "", http.StatusInternalServerError

	}

	return ConsoleOutput(art), http.StatusOK
}

//OpenFont - opens file and write info from file line-by-line to var
func OpenFont(art ASCIIArt) ([]string, int) {

	file, err := os.Open("fonts/" + art.Format + ".txt")
	if err != nil {
		fmt.Printf("ascii art fs: open '%s' no such file or directory\n", art.Format)
		fmt.Println("")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	str := []string{}

	//Writing file line-by-line to var str
	for scanner.Scan() {
		str = append(str, scanner.Text())
	}

	if len(str) != 855 {
		fmt.Printf("ascii art: '%s' is empty or it is not a font\n", art.Format)
		return []string{}, http.StatusInternalServerError
	}

	return str, http.StatusOK
}
