package leetcode567

import (
	"fmt"
)

func checkInclusion(s1 string, s2 string) bool {
	s1_map := make(map[string]int)
	copy_map := make(map[string]int)
	for _, v := range s1 {
		s1_map[string(v)] += 1
	}
	fmt.Println(s1_map)
	for id, v := range s1_map {
		copy_map[id] = v
	}

	ptr_in := 0
	ptr_fi := 0
	count := 0
	for ptr_fi < len(s2) {
		fmt.Println(s1)
		fmt.Println(s2)
		fmt.Printf("s2[%d] = %s\n", ptr_in, string(s2[ptr_fi]))
		if v, ok := copy_map[string(s2[ptr_fi])]; ok && v != 0 {
			copy_map[string(s2[ptr_fi])] -= 1
			count += 1
			if count == len(s1) {
				fmt.Printf("\n\n\n\n")
				return true
			}
			ptr_fi += 1
		} else {
			fmt.Printf("restoring map...\n")
			for id, v := range s1_map {
				copy_map[id] = v
			}
			count = 0
			ptr_in += 1
			ptr_fi = ptr_in
		}
		fmt.Printf("count = %d\t", count)
		fmt.Println(copy_map)
		fmt.Printf("\n")
	}
	fmt.Printf("\n\n\n\n")
	return false
}
