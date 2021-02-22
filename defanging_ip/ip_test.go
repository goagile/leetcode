package defangingip

import "testing"

var examples = []struct {
	want string
	ip   string
}{
	{
		want: "1[.]1[.]1[.]1",
		ip:   "1.1.1.1",
	},
	{
		want: "255[.]100[.]50[.]0",
		ip:   "255.100.50.0",
	},
}

func Test_examples(t *testing.T) {
	for i, e := range examples {
		got := defangIPaddr(e.ip)
		if e.want != got {
			t.Fatalf("\ni:%v want:%v got: %v", i, e.want, got)
		}
	}
}
