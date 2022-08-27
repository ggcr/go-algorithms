// Design an algorithm to encode a list of strings to a string. The encoded string is then sent over the network and is decoded back to the original list of strings.

// Implement the encode and decode methods.

// Input: dummy_input = ["Hello","World"]
// Output: ["Hello","World"]

package leetcode271

import (
	"strings"
)

const SECRET_KEY string = "(,//&6"

type Codec struct {
}

// Encodes a list of strings to a single string.
func (codec *Codec) Encode(strs []string) string {
	return strings.Join(strs[:], SECRET_KEY)
}

func (codec *Codec) EncodeManual(strs []string) string {
	res := ""
	for key, word := range strs {
		if key == len(strs)-1 {
			return res + word
		}
		res = res + word + SECRET_KEY
	}
	return ""
}

// Decodes a single string to a list of strings.
func (codec *Codec) Decode(strs string) []string {
	return strings.Split(strs, SECRET_KEY)
}

// We know the key, hence we know the length of the key, so we can check for the key at the string array and split accordingly.
func (codec *Codec) DecodeManual(strs string) []string {
	if len(strs) <= 1 {
		return []string{strs}
	}
	res := []string{}
	init_iterator := 0
	accumulative_str := ""
	for i := 0; i < len(strs); i += 1 {
		if i+6 >= len(strs) {
			if strs[:] == SECRET_KEY || strs[i:] == SECRET_KEY {
				res = append(res, accumulative_str)
				return append(res, "")
			}
			return append(res, strs[init_iterator:])
		}
		if strs[i:i+6] == SECRET_KEY {
			init_iterator = i + 6
			res = append(res, accumulative_str)
			accumulative_str = ""
		}
		if i >= init_iterator {
			accumulative_str = accumulative_str + string(strs[i])
		}
	}
	return []string{accumulative_str}
}

// Your Codec object will be instantiated and called as such:
// var codec Codec
// codec.Decode(codec.Encode(strs));
