package main

// Barrowed from https://github.com/mpolney/clext
var palette []uint16 = []uint16{
	0xff, 0xff, 0xff, 0xff, 0xff, 0xcc, 0xff, 0xff, 0x99, 0xff, 0xff, 0x66, 0xff, 0xff, 0x33,
	0xff, 0xff, 0x00, 0xff, 0xcc, 0xff, 0xff, 0xcc, 0xcc, 0xff, 0xcc, 0x99, 0xff, 0xcc, 0x66,
	0xff, 0xcc, 0x33, 0xff, 0xcc, 0x00, 0xff, 0x99, 0xff, 0xff, 0x99, 0xcc, 0xff, 0x99, 0x99,
	0xff, 0x99, 0x66, 0xff, 0x99, 0x33, 0xff, 0x99, 0x00, 0xff, 0x66, 0xff, 0xff, 0x66, 0xcc,
	0xff, 0x66, 0x99, 0xff, 0x66, 0x66, 0xff, 0x66, 0x33, 0xff, 0x66, 0x00, 0xff, 0x33, 0xff,
	0xff, 0x33, 0xcc, 0xff, 0x33, 0x99, 0xff, 0x33, 0x66, 0xff, 0x33, 0x33, 0xff, 0x33, 0x00,
	0xff, 0x00, 0xff, 0xff, 0x00, 0xcc, 0xff, 0x00, 0x99, 0xff, 0x00, 0x66, 0xff, 0x00, 0x33,
	0xff, 0x00, 0x00, 0xcc, 0xff, 0xff, 0xcc, 0xff, 0xcc, 0xcc, 0xff, 0x99, 0xcc, 0xff, 0x66,
	0xcc, 0xff, 0x33, 0xcc, 0xff, 0x00, 0xcc, 0xcc, 0xff, 0xcc, 0xcc, 0xcc, 0xcc, 0xcc, 0x99,
	0xcc, 0xcc, 0x66, 0xcc, 0xcc, 0x33, 0xcc, 0xcc, 0x00, 0xcc, 0x99, 0xff, 0xcc, 0x99, 0xcc,
	0xcc, 0x99, 0x99, 0xcc, 0x99, 0x66, 0xcc, 0x99, 0x33, 0xcc, 0x99, 0x00, 0xcc, 0x66, 0xff,
	0xcc, 0x66, 0xcc, 0xcc, 0x66, 0x99, 0xcc, 0x66, 0x66, 0xcc, 0x66, 0x33, 0xcc, 0x66, 0x00,
	0xcc, 0x33, 0xff, 0xcc, 0x33, 0xcc, 0xcc, 0x33, 0x99, 0xcc, 0x33, 0x66, 0xcc, 0x33, 0x33,
	0xcc, 0x33, 0x00, 0xcc, 0x00, 0xff, 0xcc, 0x00, 0xcc, 0xcc, 0x00, 0x99, 0xcc, 0x00, 0x66,
	0xcc, 0x00, 0x33, 0xcc, 0x00, 0x00, 0x99, 0xff, 0xff, 0x99, 0xff, 0xcc, 0x99, 0xff, 0x99,
	0x99, 0xff, 0x66, 0x99, 0xff, 0x33, 0x99, 0xff, 0x00, 0x99, 0xcc, 0xff, 0x99, 0xcc, 0xcc,
	0x99, 0xcc, 0x99, 0x99, 0xcc, 0x66, 0x99, 0xcc, 0x33, 0x99, 0xcc, 0x00, 0x99, 0x99, 0xff,
	0x99, 0x99, 0xcc, 0x99, 0x99, 0x99, 0x99, 0x99, 0x66, 0x99, 0x99, 0x33, 0x99, 0x99, 0x00,
	0x99, 0x66, 0xff, 0x99, 0x66, 0xcc, 0x99, 0x66, 0x99, 0x99, 0x66, 0x66, 0x99, 0x66, 0x33,
	0x99, 0x66, 0x00, 0x99, 0x33, 0xff, 0x99, 0x33, 0xcc, 0x99, 0x33, 0x99, 0x99, 0x33, 0x66,
	0x99, 0x33, 0x33, 0x99, 0x33, 0x00, 0x99, 0x00, 0xff, 0x99, 0x00, 0xcc, 0x99, 0x00, 0x99,
	0x99, 0x00, 0x66, 0x99, 0x00, 0x33, 0x99, 0x00, 0x00, 0x66, 0xff, 0xff, 0x66, 0xff, 0xcc,
	0x66, 0xff, 0x99, 0x66, 0xff, 0x66, 0x66, 0xff, 0x33, 0x66, 0xff, 0x00, 0x66, 0xcc, 0xff,
	0x66, 0xcc, 0xcc, 0x66, 0xcc, 0x99, 0x66, 0xcc, 0x66, 0x66, 0xcc, 0x33, 0x66, 0xcc, 0x00,
	0x66, 0x99, 0xff, 0x66, 0x99, 0xcc, 0x66, 0x99, 0x99, 0x66, 0x99, 0x66, 0x66, 0x99, 0x33,
	0x66, 0x99, 0x00, 0x66, 0x66, 0xff, 0x66, 0x66, 0xcc, 0x66, 0x66, 0x99, 0x66, 0x66, 0x66,
	0x66, 0x66, 0x33, 0x66, 0x66, 0x00, 0x66, 0x33, 0xff, 0x66, 0x33, 0xcc, 0x66, 0x33, 0x99,
	0x66, 0x33, 0x66, 0x66, 0x33, 0x33, 0x66, 0x33, 0x00, 0x66, 0x00, 0xff, 0x66, 0x00, 0xcc,
	0x66, 0x00, 0x99, 0x66, 0x00, 0x66, 0x66, 0x00, 0x33, 0x66, 0x00, 0x00, 0x33, 0xff, 0xff,
	0x33, 0xff, 0xcc, 0x33, 0xff, 0x99, 0x33, 0xff, 0x66, 0x33, 0xff, 0x33, 0x33, 0xff, 0x00,
	0x33, 0xcc, 0xff, 0x33, 0xcc, 0xcc, 0x33, 0xcc, 0x99, 0x33, 0xcc, 0x66, 0x33, 0xcc, 0x33,
	0x33, 0xcc, 0x00, 0x33, 0x99, 0xff, 0x33, 0x99, 0xcc, 0x33, 0x99, 0x99, 0x33, 0x99, 0x66,
	0x33, 0x99, 0x33, 0x33, 0x99, 0x00, 0x33, 0x66, 0xff, 0x33, 0x66, 0xcc, 0x33, 0x66, 0x99,
	0x33, 0x66, 0x66, 0x33, 0x66, 0x33, 0x33, 0x66, 0x00, 0x33, 0x33, 0xff, 0x33, 0x33, 0xcc,
	0x33, 0x33, 0x99, 0x33, 0x33, 0x66, 0x33, 0x33, 0x33, 0x33, 0x33, 0x00, 0x33, 0x00, 0xff,
	0x33, 0x00, 0xcc, 0x33, 0x00, 0x99, 0x33, 0x00, 0x66, 0x33, 0x00, 0x33, 0x33, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x00, 0xff, 0xcc, 0x00, 0xff, 0x99, 0x00, 0xff, 0x66, 0x00, 0xff, 0x33,
	0x00, 0xff, 0x00, 0x00, 0xcc, 0xff, 0x00, 0xcc, 0xcc, 0x00, 0xcc, 0x99, 0x00, 0xcc, 0x66,
	0x00, 0xcc, 0x33, 0x00, 0xcc, 0x00, 0x00, 0x99, 0xff, 0x00, 0x99, 0xcc, 0x00, 0x99, 0x99,
	0x00, 0x99, 0x66, 0x00, 0x99, 0x33, 0x00, 0x99, 0x00, 0x00, 0x66, 0xff, 0x00, 0x66, 0xcc,
	0x00, 0x66, 0x99, 0x00, 0x66, 0x66, 0x00, 0x66, 0x33, 0x00, 0x66, 0x00, 0x00, 0x33, 0xff,
	0x00, 0x33, 0xcc, 0x00, 0x33, 0x99, 0x00, 0x33, 0x66, 0x00, 0x33, 0x33, 0x00, 0x33, 0x00,
	0x00, 0x00, 0xff, 0x00, 0x00, 0xcc, 0x00, 0x00, 0x99, 0x00, 0x00, 0x66, 0x00, 0x00, 0x33,
	0x00, 0x00, 0x00, 0xee, 0x00, 0x00, 0xdd, 0x00, 0x00, 0xbb, 0x00, 0x00, 0xaa, 0x00, 0x00,
	0x88, 0x00, 0x00, 0x77, 0x00, 0x00, 0x55, 0x00, 0x00, 0x44, 0x00, 0x00, 0x22, 0x00, 0x00,
	0x11, 0x00, 0x00, 0x00, 0xee, 0x00, 0x00, 0xdd, 0x00, 0x00, 0xbb, 0x00, 0x00, 0xaa, 0x00,
	0x00, 0x88, 0x00, 0x00, 0x77, 0x00, 0x00, 0x55, 0x00, 0x00, 0x44, 0x00, 0x00, 0x22, 0x00,
	0x00, 0x11, 0x00, 0x00, 0x00, 0xee, 0x00, 0x00, 0xdd, 0x00, 0x00, 0xbb, 0x00, 0x00, 0xaa,
	0x00, 0x00, 0x88, 0x00, 0x00, 0x77, 0x00, 0x00, 0x55, 0x00, 0x00, 0x44, 0x00, 0x00, 0x22,
	0xff, 0xff, 0xff, 0xee, 0xee, 0xee, 0xdd, 0xdd, 0xdd, 0xbb, 0xbb, 0xbb, 0xaa, 0xaa, 0xaa,
	0x88, 0x88, 0x88, 0x77, 0x77, 0x77, 0x55, 0x55, 0x55, 0x44, 0x44, 0x44, 0x22, 0x22, 0x22,
	0x11, 0x11, 0x11,
}
