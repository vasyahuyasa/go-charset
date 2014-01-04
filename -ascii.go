package charset

import (
	"unicode/utf8"
)

// ASCII charset
var ASCII = New(table_ASCII)

var table_ASCII = [256]rune{
	0x00: 0x0000,         // NULL
	0x01: 0x0001,         // START OF HEADING
	0x02: 0x0002,         // START OF TEXT
	0x03: 0x0003,         // END OF TEXT
	0x04: 0x0004,         // END OF TRANSMISSION
	0x05: 0x0005,         // ENQUIRY
	0x06: 0x0006,         // ACKNOWLEDGE
	0x07: 0x0007,         // BELL
	0x08: 0x0008,         // BACKSPACE
	0x09: 0x0009,         // HORIZONTAL TABULATION
	0x0A: 0x000A,         // LINE FEED
	0x0B: 0x000B,         // VERTICAL TABULATION
	0x0C: 0x000C,         // FORM FEED
	0x0D: 0x000D,         // CARRIAGE RETURN
	0x0E: 0x000E,         // SHIFT OUT
	0x0F: 0x000F,         // SHIFT IN
	0x10: 0x0010,         // DATA LINK ESCAPE
	0x11: 0x0011,         // DEVICE CONTROL ONE
	0x12: 0x0012,         // DEVICE CONTROL TWO
	0x13: 0x0013,         // DEVICE CONTROL THREE
	0x14: 0x0014,         // DEVICE CONTROL FOUR
	0x15: 0x0015,         // NEGATIVE ACKNOWLEDGE
	0x16: 0x0016,         // SYNCHRONOUS IDLE
	0x17: 0x0017,         // END OF TRANSMISSION BLOCK
	0x18: 0x0018,         // CANCEL
	0x19: 0x0019,         // END OF MEDIUM
	0x1A: 0x001A,         // SUBSTITUTE
	0x1B: 0x001B,         // ESCAPE
	0x1C: 0x001C,         // FILE SEPARATOR
	0x1D: 0x001D,         // GROUP SEPARATOR
	0x1E: 0x001E,         // RECORD SEPARATOR
	0x1F: 0x001F,         // UNIT SEPARATOR
	0x20: 0x0020,         // SPACE
	0x21: 0x0021,         // EXCLAMATION MARK
	0x22: 0x0022,         // QUOTATION MARK
	0x23: 0x0023,         // NUMBER SIGN
	0x24: 0x0024,         // DOLLAR SIGN
	0x25: 0x0025,         // PERCENT SIGN
	0x26: 0x0026,         // AMPERSAND
	0x27: 0x0027,         // APOSTROPHE
	0x28: 0x0028,         // LEFT PARENTHESIS
	0x29: 0x0029,         // RIGHT PARENTHESIS
	0x2A: 0x002A,         // ASTERISK
	0x2B: 0x002B,         // PLUS SIGN
	0x2C: 0x002C,         // COMMA
	0x2D: 0x002D,         // HYPHEN-MINUS
	0x2E: 0x002E,         // FULL STOP
	0x2F: 0x002F,         // SOLIDUS
	0x30: 0x0030,         // DIGIT ZERO
	0x31: 0x0031,         // DIGIT ONE
	0x32: 0x0032,         // DIGIT TWO
	0x33: 0x0033,         // DIGIT THREE
	0x34: 0x0034,         // DIGIT FOUR
	0x35: 0x0035,         // DIGIT FIVE
	0x36: 0x0036,         // DIGIT SIX
	0x37: 0x0037,         // DIGIT SEVEN
	0x38: 0x0038,         // DIGIT EIGHT
	0x39: 0x0039,         // DIGIT NINE
	0x3A: 0x003A,         // COLON
	0x3B: 0x003B,         // SEMICOLON
	0x3C: 0x003C,         // LESS-THAN SIGN
	0x3D: 0x003D,         // EQUALS SIGN
	0x3E: 0x003E,         // GREATER-THAN SIGN
	0x3F: 0x003F,         // QUESTION MARK
	0x40: 0x0040,         // COMMERCIAL AT
	0x41: 0x0041,         // LATIN CAPITAL LETTER A
	0x42: 0x0042,         // LATIN CAPITAL LETTER B
	0x43: 0x0043,         // LATIN CAPITAL LETTER C
	0x44: 0x0044,         // LATIN CAPITAL LETTER D
	0x45: 0x0045,         // LATIN CAPITAL LETTER E
	0x46: 0x0046,         // LATIN CAPITAL LETTER F
	0x47: 0x0047,         // LATIN CAPITAL LETTER G
	0x48: 0x0048,         // LATIN CAPITAL LETTER H
	0x49: 0x0049,         // LATIN CAPITAL LETTER I
	0x4A: 0x004A,         // LATIN CAPITAL LETTER J
	0x4B: 0x004B,         // LATIN CAPITAL LETTER K
	0x4C: 0x004C,         // LATIN CAPITAL LETTER L
	0x4D: 0x004D,         // LATIN CAPITAL LETTER M
	0x4E: 0x004E,         // LATIN CAPITAL LETTER N
	0x4F: 0x004F,         // LATIN CAPITAL LETTER O
	0x50: 0x0050,         // LATIN CAPITAL LETTER P
	0x51: 0x0051,         // LATIN CAPITAL LETTER Q
	0x52: 0x0052,         // LATIN CAPITAL LETTER R
	0x53: 0x0053,         // LATIN CAPITAL LETTER S
	0x54: 0x0054,         // LATIN CAPITAL LETTER T
	0x55: 0x0055,         // LATIN CAPITAL LETTER U
	0x56: 0x0056,         // LATIN CAPITAL LETTER V
	0x57: 0x0057,         // LATIN CAPITAL LETTER W
	0x58: 0x0058,         // LATIN CAPITAL LETTER X
	0x59: 0x0059,         // LATIN CAPITAL LETTER Y
	0x5A: 0x005A,         // LATIN CAPITAL LETTER Z
	0x5B: 0x005B,         // LEFT SQUARE BRACKET
	0x5C: 0x005C,         // REVERSE SOLIDUS
	0x5D: 0x005D,         // RIGHT SQUARE BRACKET
	0x5E: 0x005E,         // CIRCUMFLEX ACCENT
	0x5F: 0x005F,         // LOW LINE
	0x60: 0x0060,         // GRAVE ACCENT
	0x61: 0x0061,         // LATIN SMALL LETTER A
	0x62: 0x0062,         // LATIN SMALL LETTER B
	0x63: 0x0063,         // LATIN SMALL LETTER C
	0x64: 0x0064,         // LATIN SMALL LETTER D
	0x65: 0x0065,         // LATIN SMALL LETTER E
	0x66: 0x0066,         // LATIN SMALL LETTER F
	0x67: 0x0067,         // LATIN SMALL LETTER G
	0x68: 0x0068,         // LATIN SMALL LETTER H
	0x69: 0x0069,         // LATIN SMALL LETTER I
	0x6A: 0x006A,         // LATIN SMALL LETTER J
	0x6B: 0x006B,         // LATIN SMALL LETTER K
	0x6C: 0x006C,         // LATIN SMALL LETTER L
	0x6D: 0x006D,         // LATIN SMALL LETTER M
	0x6E: 0x006E,         // LATIN SMALL LETTER N
	0x6F: 0x006F,         // LATIN SMALL LETTER O
	0x70: 0x0070,         // LATIN SMALL LETTER P
	0x71: 0x0071,         // LATIN SMALL LETTER Q
	0x72: 0x0072,         // LATIN SMALL LETTER R
	0x73: 0x0073,         // LATIN SMALL LETTER S
	0x74: 0x0074,         // LATIN SMALL LETTER T
	0x75: 0x0075,         // LATIN SMALL LETTER U
	0x76: 0x0076,         // LATIN SMALL LETTER V
	0x77: 0x0077,         // LATIN SMALL LETTER W
	0x78: 0x0078,         // LATIN SMALL LETTER X
	0x79: 0x0079,         // LATIN SMALL LETTER Y
	0x7A: 0x007A,         // LATIN SMALL LETTER Z
	0x7B: 0x007B,         // LEFT CURLY BRACKET
	0x7C: 0x007C,         // VERTICAL LINE
	0x7D: 0x007D,         // RIGHT CURLY BRACKET
	0x7E: 0x007E,         // TILDE
	0x7F: 0x007F,         // DELETE
	0x80: utf8.RuneError, // UNDEFINED
	0x81: utf8.RuneError, // UNDEFINED
	0x82: utf8.RuneError, // UNDEFINED
	0x83: utf8.RuneError, // UNDEFINED
	0x84: utf8.RuneError, // UNDEFINED
	0x85: utf8.RuneError, // UNDEFINED
	0x86: utf8.RuneError, // UNDEFINED
	0x87: utf8.RuneError, // UNDEFINED
	0x88: utf8.RuneError, // UNDEFINED
	0x89: utf8.RuneError, // UNDEFINED
	0x8A: utf8.RuneError, // UNDEFINED
	0x8B: utf8.RuneError, // UNDEFINED
	0x8C: utf8.RuneError, // UNDEFINED
	0x8D: utf8.RuneError, // UNDEFINED
	0x8E: utf8.RuneError, // UNDEFINED
	0x8F: utf8.RuneError, // UNDEFINED
	0x90: utf8.RuneError, // UNDEFINED
	0x91: utf8.RuneError, // UNDEFINED
	0x92: utf8.RuneError, // UNDEFINED
	0x93: utf8.RuneError, // UNDEFINED
	0x94: utf8.RuneError, // UNDEFINED
	0x95: utf8.RuneError, // UNDEFINED
	0x96: utf8.RuneError, // UNDEFINED
	0x97: utf8.RuneError, // UNDEFINED
	0x98: utf8.RuneError, // UNDEFINED
	0x99: utf8.RuneError, // UNDEFINED
	0x9A: utf8.RuneError, // UNDEFINED
	0x9B: utf8.RuneError, // UNDEFINED
	0x9C: utf8.RuneError, // UNDEFINED
	0x9D: utf8.RuneError, // UNDEFINED
	0x9E: utf8.RuneError, // UNDEFINED
	0x9F: utf8.RuneError, // UNDEFINED
	0xA0: utf8.RuneError, // UNDEFINED
	0xA1: utf8.RuneError, // UNDEFINED
	0xA2: utf8.RuneError, // UNDEFINED
	0xA3: utf8.RuneError, // UNDEFINED
	0xA4: utf8.RuneError, // UNDEFINED
	0xA5: utf8.RuneError, // UNDEFINED
	0xA6: utf8.RuneError, // UNDEFINED
	0xA7: utf8.RuneError, // UNDEFINED
	0xA8: utf8.RuneError, // UNDEFINED
	0xA9: utf8.RuneError, // UNDEFINED
	0xAA: utf8.RuneError, // UNDEFINED
	0xAB: utf8.RuneError, // UNDEFINED
	0xAC: utf8.RuneError, // UNDEFINED
	0xAD: utf8.RuneError, // UNDEFINED
	0xAE: utf8.RuneError, // UNDEFINED
	0xAF: utf8.RuneError, // UNDEFINED
	0xB0: utf8.RuneError, // UNDEFINED
	0xB1: utf8.RuneError, // UNDEFINED
	0xB2: utf8.RuneError, // UNDEFINED
	0xB3: utf8.RuneError, // UNDEFINED
	0xB4: utf8.RuneError, // UNDEFINED
	0xB5: utf8.RuneError, // UNDEFINED
	0xB6: utf8.RuneError, // UNDEFINED
	0xB7: utf8.RuneError, // UNDEFINED
	0xB8: utf8.RuneError, // UNDEFINED
	0xB9: utf8.RuneError, // UNDEFINED
	0xBA: utf8.RuneError, // UNDEFINED
	0xBB: utf8.RuneError, // UNDEFINED
	0xBC: utf8.RuneError, // UNDEFINED
	0xBD: utf8.RuneError, // UNDEFINED
	0xBE: utf8.RuneError, // UNDEFINED
	0xBF: utf8.RuneError, // UNDEFINED
	0xC0: utf8.RuneError, // UNDEFINED
	0xC1: utf8.RuneError, // UNDEFINED
	0xC2: utf8.RuneError, // UNDEFINED
	0xC3: utf8.RuneError, // UNDEFINED
	0xC4: utf8.RuneError, // UNDEFINED
	0xC5: utf8.RuneError, // UNDEFINED
	0xC6: utf8.RuneError, // UNDEFINED
	0xC7: utf8.RuneError, // UNDEFINED
	0xC8: utf8.RuneError, // UNDEFINED
	0xC9: utf8.RuneError, // UNDEFINED
	0xCA: utf8.RuneError, // UNDEFINED
	0xCB: utf8.RuneError, // UNDEFINED
	0xCC: utf8.RuneError, // UNDEFINED
	0xCD: utf8.RuneError, // UNDEFINED
	0xCE: utf8.RuneError, // UNDEFINED
	0xCF: utf8.RuneError, // UNDEFINED
	0xD0: utf8.RuneError, // UNDEFINED
	0xD1: utf8.RuneError, // UNDEFINED
	0xD2: utf8.RuneError, // UNDEFINED
	0xD3: utf8.RuneError, // UNDEFINED
	0xD4: utf8.RuneError, // UNDEFINED
	0xD5: utf8.RuneError, // UNDEFINED
	0xD6: utf8.RuneError, // UNDEFINED
	0xD7: utf8.RuneError, // UNDEFINED
	0xD8: utf8.RuneError, // UNDEFINED
	0xD9: utf8.RuneError, // UNDEFINED
	0xDA: utf8.RuneError, // UNDEFINED
	0xDB: utf8.RuneError, // UNDEFINED
	0xDC: utf8.RuneError, // UNDEFINED
	0xDD: utf8.RuneError, // UNDEFINED
	0xDE: utf8.RuneError, // UNDEFINED
	0xDF: utf8.RuneError, // UNDEFINED
	0xE0: utf8.RuneError, // UNDEFINED
	0xE1: utf8.RuneError, // UNDEFINED
	0xE2: utf8.RuneError, // UNDEFINED
	0xE3: utf8.RuneError, // UNDEFINED
	0xE4: utf8.RuneError, // UNDEFINED
	0xE5: utf8.RuneError, // UNDEFINED
	0xE6: utf8.RuneError, // UNDEFINED
	0xE7: utf8.RuneError, // UNDEFINED
	0xE8: utf8.RuneError, // UNDEFINED
	0xE9: utf8.RuneError, // UNDEFINED
	0xEA: utf8.RuneError, // UNDEFINED
	0xEB: utf8.RuneError, // UNDEFINED
	0xEC: utf8.RuneError, // UNDEFINED
	0xED: utf8.RuneError, // UNDEFINED
	0xEE: utf8.RuneError, // UNDEFINED
	0xEF: utf8.RuneError, // UNDEFINED
	0xF0: utf8.RuneError, // UNDEFINED
	0xF1: utf8.RuneError, // UNDEFINED
	0xF2: utf8.RuneError, // UNDEFINED
	0xF3: utf8.RuneError, // UNDEFINED
	0xF4: utf8.RuneError, // UNDEFINED
	0xF5: utf8.RuneError, // UNDEFINED
	0xF6: utf8.RuneError, // UNDEFINED
	0xF7: utf8.RuneError, // UNDEFINED
	0xF8: utf8.RuneError, // UNDEFINED
	0xF9: utf8.RuneError, // UNDEFINED
	0xFA: utf8.RuneError, // UNDEFINED
	0xFB: utf8.RuneError, // UNDEFINED
	0xFC: utf8.RuneError, // UNDEFINED
	0xFD: utf8.RuneError, // UNDEFINED
	0xFE: utf8.RuneError, // UNDEFINED
	0xFF: utf8.RuneError, // UNDEFINED
}