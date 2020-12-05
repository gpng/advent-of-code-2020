package passport

import (
	"regexp"
	"strconv"
	"strings"
)

// constants
const (
	HeightCm     = "cm"
	HeightInches = "in"
)

// Height model
type Height struct {
	Scale string
	Value int
}

// Passport model
type Passport struct {
	BirthYear      int
	IssueYear      int
	ExpirationYear int
	Height         Height
	HairColor      string
	EyeColor       string
	PassportID     string
	CountryID      string
}

// Parse Passports from input
func Parse(input [][]string) []Passport {
	passports := make([]Passport, 0)

	isValid := true
	passport := Passport{}
	requiredFields := getRequiredFields()

	for i := 0; i < len(input); i++ {
		row := input[i]

		if row[0] == "" {
			if isValid && len(requiredFields) == 0 {
				passports = append(passports, passport)
			}
			isValid = true
			passport = Passport{}
			requiredFields = getRequiredFields()
			continue
		}

		if !isValid {
			continue
		}

		for j := 0; j < len(row); j++ {
			if !isValid {
				break
			}
			split := strings.Split(row[j], ":")
			key := split[0]
			val := split[1]
			switch key {
			case "byr":
				{
					passport.BirthYear, isValid = parseBirthYear(val)
				}
			case "iyr":
				{
					passport.IssueYear, isValid = parseIssueYear(val)
				}
			case "eyr":
				{
					passport.ExpirationYear, isValid = parseExpiryYear(val)
				}
			case "hgt":
				{
					passport.Height, isValid = parseHeight(val)
				}
			case "hcl":
				{
					passport.HairColor, isValid = parseHairColor(val)
				}
			case "ecl":
				{
					passport.EyeColor, isValid = parseEyeColor(val)
				}
			case "pid":
				{
					passport.PassportID, isValid = parsePassportID(val)
				}
			case "cid":
				{
					passport.CountryID, _ = parseCountryID(val)
				}
			}
			if isValid {
				delete(requiredFields, key)
			}
		}
	}

	if isValid && len(requiredFields) == 0 {
		passports = append(passports, passport)
	}
	return passports
}

func parseBirthYear(byr string) (int, bool) {
	year, err := strconv.Atoi(byr)
	if err != nil || year < 1920 || year > 2002 {
		return 0, false
	}
	return year, true
}

func parseIssueYear(iyr string) (int, bool) {
	year, err := strconv.Atoi(iyr)
	if err != nil || year < 2010 || year > 2020 {
		return 0, false
	}
	return year, true
}

func parseExpiryYear(eyr string) (int, bool) {
	year, err := strconv.Atoi(eyr)
	if err != nil || year < 2020 || year > 2030 {
		return 0, false
	}
	return year, true
}

func parseHeight(hgt string) (Height, bool) {
	scale := hgt[len(hgt)-2:]
	height, err := strconv.Atoi(hgt[:len(hgt)-2])
	if err != nil {
		return Height{}, false
	}
	switch scale {
	case HeightCm:
		{
			return Height{scale, height}, height >= 150 && height <= 193
		}
	case HeightInches:
		{
			return Height{scale, height}, height >= 59 && height <= 76
		}
	default:
		{
			return Height{}, false
		}
	}
}

func parseHairColor(hcl string) (string, bool) {
	if len(hcl) != 7 || hcl[:1] != "#" {
		return "", false
	}
	color := hcl[1:]
	re := regexp.MustCompile("^[a-zA-Z0-9]*$")
	return color, re.MatchString(color)
}

func parseEyeColor(ecl string) (string, bool) {
	validColors := map[string]bool{
		"amb": true,
		"blu": true,
		"brn": true,
		"gry": true,
		"grn": true,
		"hzl": true,
		"oth": true,
	}
	return ecl, validColors[ecl]
}

func parsePassportID(pid string) (string, bool) {
	if len(pid) != 9 {
		return pid, false
	}
	re := regexp.MustCompile("^[0-9]*$")
	return pid, re.MatchString(pid)
}

func parseCountryID(cid string) (string, bool) {
	return cid, true
}

func getRequiredFields() map[string]bool {
	return map[string]bool{"byr": true, "iyr": true, "eyr": true, "hgt": true, "hcl": true, "ecl": true, "pid": true}
}
