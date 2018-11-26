package node

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
)

func TestNewAddress(t *testing.T) {
	cases := []struct {
		host       string
		port       string
		difficulty int
	}{
		{"host1", "port1", 1},
		{"host2", "port2", 2},
		{"host3", "port3", 3},
	}

	for i, cc := range cases {
		c := cc
		t.Run(fmt.Sprintf("%d-th case", i), func(t *testing.T) {
			na, nonce, err := newAddress(c.host, c.port, nil, c.difficulty)
			assert.Equal(t, nil, err)
			assert.Equal(t, true, verifyAddress(na, c.host, c.port, nil, nonce, c.difficulty))
		})
	}
}

func TestVerifyAddress(t *testing.T) {
	cases := []struct {
		da         doogleAddress
		host       string
		port       string
		pk         []byte
		nonce      []byte
		difficulty int
		expected   bool
	}{
		{
			doogleAddress{},
			"",
			"",
			nil,
			nil,
			10,
			false,
		},
		{
			doogleAddress{137, 247, 252, 74, 101, 232, 49, 193, 122, 237, 123, 84, 199, 94, 78, 176, 92, 104, 69, 253},
			"ab",
			"80",
			[]byte("pk"),
			[]byte{124, 101, 169, 225, 58, 47, 235, 38, 179, 1},
			1,
			true,
		},
		{
			doogleAddress{137, 247, 252, 74, 101, 232, 49, 193, 122, 237, 123, 84, 199, 94, 78, 176, 92, 104, 69, 253},
			"ab",
			"80",
			[]byte("pk"),
			[]byte{172, 171, 254, 98, 171, 6, 169, 186, 105, 145},
			2,
			true,
		},
	}

	for i, cc := range cases {
		c := cc
		t.Run(fmt.Sprintf("%d-th case", i), func(t *testing.T) {
			actual := verifyAddress(c.da, c.host, c.port, c.pk, c.nonce, c.difficulty)
			assert.Equal(t, actual, c.expected)
		})
	}
}

func TestGetNonce(t *testing.T) {
	for i := 0; i < 10e3; i++ {
		_, err := getNonce()
		assert.Equal(t, nil, err)
	}
}

func TestLessThan(t *testing.T) {
	cases := []struct {
		da       doogleAddress
		a        doogleAddress
		expected bool
	}{
		{
			doogleAddress{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0},
			doogleAddress{1, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0},
			true,
		},
		{
			doogleAddress{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0},
			doogleAddress{0, 0, 0, 0, 1, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0},
			false,
		},
	}

	for i, cc := range cases {
		c := cc
		t.Run(fmt.Sprintf("%d-th case", i), func(t *testing.T) {
			actual := c.da.lessThanEqual(c.a)
			assert.Equal(t, actual, c.expected)
		})
	}
}
