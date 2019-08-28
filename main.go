package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

const matchExp = "(^(N|V|AD|PHRA|CON|PREP)+|→)+"

func main() {
	result := ""
	scanner := bufio.NewScanner(os.Stdin)
	var text string
	no := 1
	for true {
		scanner.Scan()
		text = scanner.Text()
		if len(text) == 0 {
			break
		}
		ok, _ := regexp.MatchString(matchExp, text)
		if ok {
			text = fmt.Sprintf("\n%d. %s", no, text)
			no++
		}
		result += text + "\n"
	}
	fmt.Print(result[1 : len(result)-1])
}
