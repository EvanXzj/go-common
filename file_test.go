package common

import "testing"

func TestIsFile(t *testing.T) {
	cases := []struct {
		in   string
		want bool
	}{
		{"file.go", true},
		{"file_not_exists", false},
	}

	for _, c := range cases {
		got := IsFile(c.in)
		if got != c.want {
			t.Errorf("IsFile(%s) == %v, want %v", c.in, got, c.want)
		}
	}
}

func BenchmarkIsFile(b *testing.B) { // 感觉好慢啊！
	for i := 0; i < b.N; i++ {
		IsFile("file_not_exists")
	}
}
