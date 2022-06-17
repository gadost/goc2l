package main

var c2l = map[string]string{
	"а": "a",
	"б": "b",
	"в": "v",
	"г": "g",
	"д": "d",
	"е": "e",
	"ё": "e",
	"ж": "zh",
	"з": "z",
	"и": "i",
	"й": "y",
	"к": "k",
	"л": "l",
	"м": "m",
	"н": "n",
	"о": "o",
	"п": "p",
	"р": "r",
	"с": "s",
	"т": "t",
	"у": "y",
	"ф": "f",
	"х": "kh",
	"ц": "ts",
	"ч": "ch",
	"ш": "sh",
	"щ": "shh",
	"ъ": "",
	"ы": "y",
	"ь": "",
	"э": "e",
	"ю": "yu",
	"я": "ya",
	" ": " "}

var c2lspec = map[string]string{
	"]": "",
	"[": "",
	"!": "",
	"?": "",
	"=": "_",
	"-": "_",
	" ": "_",
	",": "",
	"'": "",
	";": "",
	":": "",
	")": "",
	"(": ""}

func Cyr2Lat(a string) string {
	if _, found := c2l[a]; found == false {
		return a
	} else {
		return c2l[a]
	}
}
