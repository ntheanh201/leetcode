func addBinary(a string, b string) string {
	i := 1
	j := 1
	res := ""
	carry := 0
	n := len(a) - i
	m := len(b) - j
	for n >= 0 || m >= 0 {
		x := "0"
		y := "0"
		s := ""
		if n >= 0 {
			x = string(a[n])
		}
		if m >= 0 {
			y = string(b[m])
		}
		s, carry = cal(x, y, carry)
		res = s + res
		i++
		j++
	}
	if carry > 0 {
		return "1" + res
	}
	return res
}

func cal(a string, b string, carry int) (string, int) {
	x := 0
	y := 0
	if a == "0" {
		x = 0
	} else {
		x = 1
	}
	if b == "0" {
		y = 0
	} else {
		y = 1
	}
	if x+y > 1 {
		if carry == 0 {
			return "0", 1
		} else {
			return "1", 1
		}
	} else {
		s := x + y + carry
		if s > 1 {
			return "0", 1
		} else {
			if s == 1 {
				return "1", 0
			} else {
				return "0", 0
			}
		}
	}
}
