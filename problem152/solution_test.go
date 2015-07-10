package main

import "testing"

func TestSolution(t *testing.T) {
	for age, want := range cases {
		got := category(age)
		if got != want {
			t.Errorf("category(%q) == %q, want %q", age, got, want)
		}
	}
}

var cases = map[string]string{
	"-1":  "This program is for humans",
	"0":   "Still in Mama's arms",
	"1":   "Still in Mama's arms",
	"2":   "Still in Mama's arms",
	"3":   "Preschool Maniac",
	"4":   "Preschool Maniac",
	"5":   "Elementary school",
	"6":   "Elementary school",
	"7":   "Elementary school",
	"8":   "Elementary school",
	"9":   "Elementary school",
	"10":  "Elementary school",
	"11":  "Elementary school",
	"12":  "Middle school",
	"13":  "Middle school",
	"14":  "Middle school",
	"15":  "High school",
	"16":  "High school",
	"17":  "High school",
	"18":  "High school",
	"19":  "College",
	"20":  "College",
	"21":  "College",
	"22":  "College",
	"23":  "Working for the man",
	"24":  "Working for the man",
	"25":  "Working for the man",
	"26":  "Working for the man",
	"27":  "Working for the man",
	"28":  "Working for the man",
	"29":  "Working for the man",
	"30":  "Working for the man",
	"31":  "Working for the man",
	"32":  "Working for the man",
	"33":  "Working for the man",
	"34":  "Working for the man",
	"35":  "Working for the man",
	"36":  "Working for the man",
	"37":  "Working for the man",
	"38":  "Working for the man",
	"39":  "Working for the man",
	"40":  "Working for the man",
	"41":  "Working for the man",
	"42":  "Working for the man",
	"43":  "Working for the man",
	"44":  "Working for the man",
	"45":  "Working for the man",
	"46":  "Working for the man",
	"47":  "Working for the man",
	"48":  "Working for the man",
	"49":  "Working for the man",
	"50":  "Working for the man",
	"51":  "Working for the man",
	"52":  "Working for the man",
	"53":  "Working for the man",
	"54":  "Working for the man",
	"55":  "Working for the man",
	"56":  "Working for the man",
	"57":  "Working for the man",
	"58":  "Working for the man",
	"59":  "Working for the man",
	"60":  "Working for the man",
	"61":  "Working for the man",
	"62":  "Working for the man",
	"63":  "Working for the man",
	"64":  "Working for the man",
	"65":  "Working for the man",
	"66":  "The Golden Years",
	"67":  "The Golden Years",
	"68":  "The Golden Years",
	"69":  "The Golden Years",
	"70":  "The Golden Years",
	"71":  "The Golden Years",
	"72":  "The Golden Years",
	"73":  "The Golden Years",
	"74":  "The Golden Years",
	"75":  "The Golden Years",
	"76":  "The Golden Years",
	"77":  "The Golden Years",
	"78":  "The Golden Years",
	"79":  "The Golden Years",
	"80":  "The Golden Years",
	"81":  "The Golden Years",
	"82":  "The Golden Years",
	"83":  "The Golden Years",
	"84":  "The Golden Years",
	"85":  "The Golden Years",
	"86":  "The Golden Years",
	"87":  "The Golden Years",
	"88":  "The Golden Years",
	"89":  "The Golden Years",
	"90":  "The Golden Years",
	"91":  "The Golden Years",
	"92":  "The Golden Years",
	"93":  "The Golden Years",
	"94":  "The Golden Years",
	"95":  "The Golden Years",
	"96":  "The Golden Years",
	"97":  "The Golden Years",
	"98":  "The Golden Years",
	"99":  "The Golden Years",
	"100": "The Golden Years",
	"101": "This program is for humans",
}