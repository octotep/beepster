package main

import (
	"github.com/octotep/beepster"
)

var melody []beepster.Note
var rest []beepster.Note
var tag []beepster.Note

func init() {
	// Initialize the main melody
	melody = make([]beepster.Note, 56)
	melody[0] = beepster.Note{261.0, 500, 5}
	melody[1] = beepster.Note{261.0, 500, 5}
	melody[2] = beepster.Note{261.0, 333, 5}
	melody[3] = beepster.Note{293.0, 166, 5}
	melody[4] = beepster.Note{329.0, 500, 5}
	melody[5] = beepster.Note{329.0, 333, 5}
	melody[6] = beepster.Note{293.0, 166, 5}
	melody[7] = beepster.Note{329.0, 333, 5}
	melody[8] = beepster.Note{349.0, 166, 5}
	melody[9] = beepster.Note{392.0, 1000, 5}
	melody[10] = beepster.Note{523.0, 166, 5}
	melody[11] = beepster.Note{523.0, 166, 5}
	melody[12] = beepster.Note{523.0, 166, 5}
	melody[13] = beepster.Note{392.0, 166, 5}
	melody[14] = beepster.Note{392.0, 166, 5}
	melody[15] = beepster.Note{392.0, 166, 5}
	melody[16] = beepster.Note{329.0, 166, 5}
	melody[17] = beepster.Note{329.0, 166, 5}
	melody[18] = beepster.Note{329.0, 166, 5}
	melody[19] = beepster.Note{261.0, 166, 5}
	melody[20] = beepster.Note{261.0, 166, 5}
	melody[21] = beepster.Note{261.0, 166, 5}
	melody[22] = beepster.Note{392.0, 333, 5}
	melody[24] = beepster.Note{349.0, 166, 5}
	melody[25] = beepster.Note{329.0, 333, 5}
	melody[26] = beepster.Note{293.0, 166, 5}
	melody[28] = beepster.Note{261.0, 1000, 5}
	melody[29] = beepster.Note{261.0, 500, 5}
	melody[30] = beepster.Note{261.0, 500, 5}
	melody[31] = beepster.Note{261.0, 333, 5}
	melody[32] = beepster.Note{293.0, 166, 5}
	melody[33] = beepster.Note{329.0, 500, 5}
	melody[34] = beepster.Note{329.0, 333, 5}
	melody[35] = beepster.Note{293.0, 166, 5}
	melody[36] = beepster.Note{329.0, 333, 5}
	melody[37] = beepster.Note{349.0, 166, 5}
	melody[38] = beepster.Note{392.0, 1000, 5}
	melody[39] = beepster.Note{523.0, 166, 5}
	melody[40] = beepster.Note{523.0, 166, 5}
	melody[41] = beepster.Note{523.0, 166, 5}
	melody[42] = beepster.Note{392.0, 166, 5}
	melody[43] = beepster.Note{392.0, 166, 5}
	melody[44] = beepster.Note{392.0, 166, 5}
	melody[45] = beepster.Note{329.0, 166, 5}
	melody[46] = beepster.Note{329.0, 166, 5}
	melody[47] = beepster.Note{329.0, 166, 5}
	melody[48] = beepster.Note{261.0, 166, 5}
	melody[49] = beepster.Note{261.0, 166, 5}
	melody[50] = beepster.Note{261.0, 166, 5}
	melody[51] = beepster.Note{392.0, 333, 5}
	melody[52] = beepster.Note{349.0, 166, 5}
	melody[53] = beepster.Note{329.0, 333, 5}
	melody[54] = beepster.Note{293.0, 166, 5}
	melody[55] = beepster.Note{261.0, 1000, 5}

	// Initialize one measure of rest
	rest = make([]beepster.Note, 5)
	rest[0] = beepster.Note{0, 500, 5}
	rest[1] = beepster.Note{0, 500, 5}
	rest[2] = beepster.Note{0, 333, 5}
	rest[3] = beepster.Note{0, 166, 5}
	rest[4] = beepster.Note{0, 500, 5}

	// Initialize the ending tag
	tag = make([]beepster.Note, 5)
	tag[0] = beepster.Note{392.0, 333, 5}
	tag[1] = beepster.Note{349.0, 166, 5}
	tag[2] = beepster.Note{329.0, 333, 5}
	tag[3] = beepster.Note{293.0, 166, 5}
	tag[4] = beepster.Note{261.0, 1000, 5}
}
