package thousands

import "strings"

func Seperate(n string, lang string) string {

	c := false

	if strings.ContainsAny(n, "abcdefghijklmnopqrstuvwxyz°^§$%&/()=?+#;:_-<>¡“¶¢[]|{}≠¿'!¬”#£ﬁ^˚’—÷˛@ƒ©ªº∆@œæ‘≤¥≈ç√∫~µ∞…–") || strings.ContainsAny(n, "`") || strings.ContainsAny(n, `"`) || strings.ContainsAny(n, "'") {
		c = true
	}

	if c {
		return n
	}

	switch lang {

	case "de":

		n = strings.ReplaceAll(n, ",", ".")

		dec := ""

		if strings.Index(n, ".") != -1 {
			dec = n[strings.Index(n, ".")+1 : len(n)]
			n = n[0:strings.Index(n, ".")]

		}

		for i := 0; i <= len(n); i = i + 4 {

			a := n[0 : len(n)-i]
			b := n[len(n)-i : len(n)]
			n = a + "." + b

		}

		if n[0:1] == "." {
			n = n[1:len(n)]
		}

		if n[len(n)-1:len(n)] == "." {
			n = n[0 : len(n)-1]
		}

		if dec != "" {
			n = n + "," + dec
		}

		return n

	case "en":

		n = strings.ReplaceAll(n, ",", "")

		dec := ""

		if strings.Index(n, ".") != -1 {
			dec = n[strings.Index(n, ".")+1 : len(n)]
			n = n[0:strings.Index(n, ".")]

		}

		for i := 0; i <= len(n); i = i + 4 {

			a := n[0 : len(n)-i]
			b := n[len(n)-i : len(n)]
			n = a + "," + b

		}

		if n[0:1] == "," {
			n = n[1:len(n)]
		}

		if n[len(n)-1:len(n)] == "," {
			n = n[0 : len(n)-1]
		}

		if dec != "" {
			n = n + "." + dec
		}

		return n

	}

	return n

}
