package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"golang.org/x/text/number"
)

var totaldamage int

type Config struct {
	player  string
	zone    string
	logfile string
}

func LoadConfig(filename string) (*Config, error) {
	// Open the config file
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open config file: %v", err)
	}
	defer file.Close()

	config := &Config{}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "" || strings.HasPrefix(line, "#") {
			// Skip empty lines and comments
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid config line: %s", line)
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		switch key {
		case "player":
			config.player = value
		case "zone":
			config.zone = value
		case "logfile":
			config.logfile = value
		default:
			return nil, fmt.Errorf("unknown config key: %s", key)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading config file: %v", err)
	}

	return config, nil
}

func main() {
	config, err := LoadConfig("config.conf")
	if err != nil {
		fmt.Println("Could not load config:", err)
	}

	// Open the file.
	file, err := os.Open(config.logfile)
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

		if strings.Contains(line, "SPELL_DAMAGE") && strings.Contains(line, config.player) ||
			strings.Contains(line, "SPELL_DAMAGE_SUPPORT") && strings.Contains(line, "Prescience") ||
			strings.Contains(line, "SPELL_PERIODIC_DAMAGE_SUPPORT") && strings.Contains(line, "Prescience") ||
			strings.Contains(line, "SPELL_DAMAGE_SUPPORT") && strings.Contains(line, "Ebon Might") {
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
	p.Printf("total damage: %v \n", number.Decimal(totaldamage))
}
