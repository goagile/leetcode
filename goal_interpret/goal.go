package goal

import "strings"

func interpret(cmd string) string {
	z := make([]byte, 0)
	i := 0
	for i < len(cmd) {
		switch cmd[i] {
		case 'G':
			z = append(z, 'G')
			i++
		case '(':
			switch cmd[i+1] {
			case ')':
				z = append(z, 'o')
				i += 2
			case 'a':
				z = append(z, "al"...)
				i += 4
			}
		}
	}
	return string(z)
}

func interpretBuiltin(cmd string) string {
	cmd = strings.Replace(cmd, "()", "o", -1)
	cmd = strings.Replace(cmd, "(al)", "al", -1)
	return cmd
}

func interpretMap(cmd string) string {
	z := ""
	m := map[string]string{
		"G":    "G",
		"()":   "o",
		"(al)": "al",
	}
	key := ""
	for _, c := range cmd {
		key += string(c)
		if v, ok := m[key]; ok {
			z += v
			key = ""
		}
	}
	return z
}
