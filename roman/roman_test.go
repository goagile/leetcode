package roman

import (
	"testing"
)

// func Test_D(t *testing.T) {
// 	got := ToInt("MDCXCV")

// 	t.Fatal(got)
// }

func Test_Examples(t *testing.T) {
	examples := []struct {
		want  int
		roman string
	}{
		{1, "I"},
		{2, "II"},
		{3, "III"},
		{4, "IV"},
		{5, "V"},
		{6, "VI"},
		{7, "VII"},
		{8, "VIII"},
		{9, "IX"},
		{10, "X"},
		{14, "XIV"},
		{19, "XIX"},
		{30, "XXX"},
		{34, "XXXIV"},
		{38, "XXXVIII"},
		{39, "XXXIX"},
		{40, "XL"},
		{44, "XLIV"},
		{49, "XLIX"},
		{50, "L"},
		{54, "LIV"},
		{59, "LIX"},
		{60, "LX"},
		{61, "LXI"},
		{64, "LXIV"},
		{69, "LXIX"},
		{89, "LXXXIX"},
		{90, "XC"},
		{100, "C"},
		{104, "CIV"},
		{109, "CIX"},
		{114, "CXIV"},
		{119, "CXIX"},
		{149, "CXLIX"},
		{150, "CL"},
		{189, "CLXXXIX"},
		{199, "CXCIX"},
		{399, "CCCXCIX"},
		{400, "CD"},
		{406, "CDVI"},
		{490, "CDXC"},
		{500, "D"},
		{900, "CM"},
		{1000, "M"},
		{1695, "MDCXCV"},
		{2000, "MM"},
		{2059, "MMLIX"},
		{3999, "MMMDCDXCIX"},
	}

	for i, e := range examples {
		got := ToInt(e.roman)
		if e.want != got {
			t.Fatalf(
				"\n%v example '%v' -> want:%v got:%v\n",
				i, e.roman, e.want, got,
			)
		}
	}
}
