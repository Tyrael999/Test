package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Проверка наличия других символов, которые не из английского и русского алфавитов
func strValid(s string) bool {
	return regexp.MustCompile(`^[A-Za-zА-Яа-я]*$`).MatchString(s)
}

// Сворачивание строки
func strSort(s string) (ss string) {
	sr := []rune(s)
	for sym := 'A'; sym <= 'я'; sym++ {
		if sym == '{' { //в unicode "{" идёт сразу после "z"
			sym = 'А' //русское А!
		}
		for i := range sr {
			if sr[i] == sym {
				count := strconv.Itoa(strings.Count(s, string(sym)))
				ss += string(sr[i]) + count
				break
			}
		}
	}
	return
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	str := scanner.Text()
	valid := strValid(str)
	if !valid {
		fmt.Println("Только латиница или кириллица без пробелов!")
	} else {
		fmt.Println(strSort(str))
	}
}
