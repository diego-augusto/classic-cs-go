package compress

import "testing"

func TestCompress(t *testing.T) {

	cases := []struct {
		i string
		e uint
	}{
		{"A", 0b0100},
		{"AC", 0b010001},
		{"ACG", 0b01000110},
		{"ACGT", 0b0100011011},
	}

	for _, c := range cases {

		r := compress(c.i)

		if r != c.e {
			t.Errorf("\ncase: %s\ngot: %d\nwant: %d", c.i, r, c.e)
		}
	}
}

func TestDescompress(t *testing.T) {

	cases := []struct {
		i uint
		e string
	}{
		{0b0100, "A"},
		{0b010001, "AC"},
		{0b01000110, "ACG"},
		{0b0100011011, "ACGT"},
	}

	for _, c := range cases {

		r := decompress(c.i)

		if r != c.e {
			t.Errorf("\ncase: %d\ngot: %s\nwant: %s", c.i, r, c.e)
		}
	}
}
