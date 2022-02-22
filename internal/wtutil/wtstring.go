package wtutil

import (
	"fmt"
	"strings"
)

func StringFunction() bool {
	res := true
	var trimed string

	//Test Contains
	res = true
	if !strings.Contains("Tralalilo", "la") {
		res = false
	}
	fmt.Printf("Contains of strings.Contains(\"Tralalilo\", \"la\"): >>>%v<<<%v\n", strings.Contains("Tralalilo", "la"), res)

	res = true
	if !strings.ContainsAny("Tralalilo", "zli") {
		res = false
	}
	fmt.Printf("Contains of strings.ContainsAny(\"Tralalilo\", \"zli\"): >>>%v<<<%v\n", strings.ContainsAny("Tralalilo", "zli"), res)

	res = true
	if strings.Count("Tralalalala", "la") == 0 {
		res = false
	}
	fmt.Printf("Count of strings.Count(\"Tralalalala\", \"la\"): >>>%v<<<%v\n", strings.Count("Tralalalala", "la"), res)

	//Test Has Contains
	res = true
	if !strings.HasPrefix("Tralalilo", "Tra") {
		res = false
	}
	fmt.Printf("HasPrefix of strings.HasPrefix(\"Tralalilo\", \"Tra\"): >>>%v<<<%v\n", strings.HasPrefix("Tralalilo", "Tra"), res)

	res = true
	if !strings.HasSuffix("Tralalilo", "lo") {
		res = false
	}
	fmt.Printf("HasSuffix of strings.HasSuffix(\"Tralalilo\", \"lo\"): >>>%v<<<%v\n", strings.HasSuffix("Tralalilo", "lo"), res)

	//Test Index
	res = true
	if strings.Index("Tralalalala", "la") < 1 {
		res = false
	}
	fmt.Printf("Index of strings.Index(\"Tralalalala\", \"la\"): >>>%v<<<%v\n", strings.Index("Tralalalala", "la"), res)

	res = true
	if strings.IndexAny("Tralalalala", "la") < 1 {
		res = false
	}
	fmt.Printf("IndexAny of strings.IndexAny(\"Tralalalala\", \"la\"): >>>%v<<<%v\n", strings.IndexAny("Tralalalala", "la"), res)

	res = true
	if strings.LastIndex("Tralalalala", "la") < 1 {
		res = false
	}
	fmt.Printf("LastIndex of strings.LastIndex(\"Tralalalala\", \"la\"): >>>%v<<<%v\n", strings.LastIndex("Tralalalala", "la"), res)

	res = true
	if strings.LastIndexAny("Tralalalala", "la") < 1 {
		res = false
	}
	fmt.Printf("LastIndexAny of strings.LastIndexAny(\"Tralalalala\", \"la\"): >>>%v<<<%v\n", strings.LastIndexAny("Tralalalala", "la"), res)

	//Test Replace
	replaced := ""
	res = true
	if replaced = strings.Replace("Tralalalala", "la", ".", 2); replaced != "Tra..lala" {
		res = false
	}
	fmt.Printf("Replace of strings.Replace(\"Tralalalala\", \"la\", \".\",2): >>>%v<<<%v\n", strings.Replace("Tralalalala", "la", ".", 2), res)

	res = true
	f := func(r rune) rune {
		return r + 1
	}

	if replaced = strings.Map(f, "hahiho"); replaced != "ibijip" {
		res = false
	}
	fmt.Printf("Apply function of strings.Map(f, \"hahiho\"): >>>%v<<<%v\n", strings.Map(f, "hahiho"), res)

	res = true
	if replaced = strings.ToUpper("hOhoHo"); replaced != "HOHOHO" {
		res = false
	}
	fmt.Printf("ToUpper of strings.ToUpper(\"hOhoHo\"): >>>%v<<<%v\n", strings.ToUpper("hOhoHo"), res)

	res = true
	if replaced = strings.ToLower("hOhoHo"); replaced != "hohoho" {
		res = false
	}
	fmt.Printf("ToLower of strings.ToLower(\"hOhoHo\"): >>>%v<<<%v\n", strings.ToLower("hOhoHo"), res)

	res = true
	if replaced = strings.Title("ha Ha hi Hi Ho ho"); replaced != "Ha Ha Hi Hi Ho Ho" {
		res = false
	}
	fmt.Printf("Title of strings.Title(\"ha Ha hi Hi Ho ho\"): >>>%v<<<%v\n", strings.Title("ha Ha hi Hi Ho ho"), res)

	//Test Trim
	res = true
	if trimed = strings.TrimSpace(" foo "); trimed != "foo" {
		res = false
	}
	fmt.Printf("TrimSpace of strings.TrimSpace(\" foo \"): >>>%v<<<%v\n", strings.TrimSpace(" foo "), res)

	res = true
	if trimed = strings.TrimSpace(" foo "); trimed != "foo" {
		res = false
	}
	fmt.Printf("TrimSpace of strings.TrimSpace(\" foo \"): >>>%v<<<%v\n", strings.TrimSpace(" foo "), res)

	res = true
	if trimed = strings.TrimSpace(" foo zz "); trimed != "foo zz" {
		res = false
	}
	fmt.Printf("TrimSpace of strings.TrimSpace(\" foo zz \"): >>>%v<<<%v\n", strings.TrimSpace(" foo zz "), res)

	res = true
	if trimed = strings.TrimLeft(" foo ", " "); trimed != "foo " {
		res = false
	}
	fmt.Printf("TrimLeft of strings.TrimLeft(\" foo \", \" \"): >>>%v<<<%v\n", strings.TrimLeft(" foo ", " "), res)

	res = true
	if trimed = strings.TrimRight(" foo ", " "); trimed != " foo" {
		res = false
	}
	fmt.Printf("TrimRight of strings.TrimRight(\" foo \", \" \"): >>>%v<<<%v\n", strings.TrimRight(" foo ", " "), res)

	res = true
	if trimed = strings.Trim("foosooofdfff", "of"); trimed != "sooofd" {
		res = false
	}
	fmt.Printf("Trim of strings.Trim(\"foosooofdfff\",\"of\"): >>>%v<<<%v\n", strings.Trim("foosooofdfff", "of"), res)

	res = true
	if trimed = strings.TrimPrefix("fofoo", "fo"); trimed != "foo" {
		res = false
	}
	fmt.Printf("TrimPrefix of strings.TrimPrefix(\"fofoo\",\"fo\"): >>>%v<<<%v\n", strings.TrimPrefix("fofoo", "fo"), res)

	res = true
	if trimed = strings.TrimSuffix("fofozozo", "zo"); trimed != "fofozo" {
		res = false
	}
	fmt.Printf("TrimSuffix of strings.TrimSuffix(\"fofozozo\",\"zo\"): >>>%v<<<%v\n", strings.TrimSuffix("fofozozo", "zo"), res)

	res = true
	eval := []string{"a", "b"}
	// var trimed2 []string
	trimed2 := strings.Fields(" a \t b \n")
	if len(trimed2) != len(eval) {
		res = false
	}
	for i, v := range trimed2 {
		if v != eval[i] {
			res = false
		}
	}
	fmt.Printf("Fields of strings.Fields(\" a\t b\n\"): >>>%v<<<%v\n", strings.Fields(" a\t b\n"), res)

	res = true
	eval = []string{"a", "b", "c"}
	trimed2 = strings.Split("a,b,c", ",")
	if len(trimed2) != len(eval) {
		res = false
	}
	for i, v := range trimed2 {
		if v != eval[i] {
			res = false
		}
	}
	fmt.Printf("Split of strings.Split(\"a,b,c\", \",\"): >>>%v<<<%v\n", strings.Split("a,b,c", ","), res)

	res = true
	eval = []string{"a,", "b,", "c"}
	trimed2 = strings.SplitAfter("a,b,c", ",")
	if len(trimed2) != len(eval) {
		res = false
	}
	for i, v := range trimed2 {
		if v != eval[i] {
			res = false
		}
	}
	fmt.Printf("Split of strings.SplitAfter(\"a,b,c\", \",\"): >>>%v<<<%v\n", strings.SplitAfter("a,b,c", ","), res)

	res = true
	if trimed = strings.Join([]string{"a", "b", ":c"}, ":"); trimed != "a:b::c" {
		res = false
	}
	fmt.Printf("Join of strings.Join([]string{\"a\", \"b\",\":c\"}, \":\"): >>>%v<<<%v\n", strings.Join([]string{"a", "b", ":c"}, ":"), res)

	res = true
	if trimed = strings.Repeat("bla", 3); trimed != "blablabla" {
		res = false
	}
	fmt.Printf("Repeat of strings.Repeat(\"bla\", 3): >>>%v<<<%v\n", strings.Repeat("bla", 3), res)
	return res
}
