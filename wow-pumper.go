package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"golang.org/x/text/number"
)

var logfile string
var player string
var totaldamage int

func main() {

	flag.StringVar(&logfile, "logfile", "default", "usage: --logfile '<logfile>'")
	flag.StringVar(&player, "player", "default", "<character name>")
	flag.Parse()

	// Open the file.
	file, err := os.Open(logfile)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file
	scanner := bufio.NewScanner(file)

	// count files
	lineCount := 0

	// Loop through each line of the file.
	for scanner.Scan() {
		lineCount++
		line := scanner.Text()

		if strings.Contains(line, "SPELL_DAMAGE") && strings.Contains(line, player) || strings.Contains(line, "SPELL_DAMAGE_SUPPORT") && strings.Contains(line, "Prescience") || strings.Contains(line, "SPELL_DAMAGE_SUPPORT") && strings.Contains(line, "Ebon Might") {
			fields := strings.Split(line, ",")
			//ability := string([]rune(fields[10])[1 : len([]rune(fields[10]))-1])
			//fmt.Println(fields[0], ability, fields[30])

			damage, err := strconv.Atoi(fields[30])
			if err != nil {
				fmt.Println("Error converting string to int:", err)
				return
			}
			totaldamage = totaldamage + damage
		}
	}

	// Check for errors during Scan.
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from file:", err)
	}
	fmt.Printf("lines processed: %d \n", lineCount)
	p := message.NewPrinter(language.English)
	p.Printf("total damage: %v", number.Decimal(totaldamage))
}
