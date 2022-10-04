# go-algorithms
This is a collection of algorithms and data-structure implemenetations in Go made by myself!

## How this works
I work with a TDD methodology. Each folder has its [name].go file and its [name]_test.go file. In order to run the go code just go to the desired path and run the following command:
```bash
go test [-v] // The -v flag is optional
```

## Data Structures
- [Linked List](https://github.com/ggcr/go-algorithms/tree/master/Data-Structures/linked-list)

## Neet-Code 150
### Arrays & Hashing
- [Contains Duplicate](https://github.com/ggcr/go-algorithms/tree/master/Array-And-Hashing/contains-duplicate) ([Leetcode 217](https://leetcode.com/problems/contains-duplicate/)): Brute force O(N^2), Merge Sort method O(N logN) and Map solution O(N).
- [Valid Anagram](https://github.com/ggcr/go-algorithms/tree/master/Array-And-Hashing/valid-anagram) ([Leetcode 242](https://leetcode.com/problems/valid-anagram/)): An efficient Hashmap solution O(N) and another implementation for detecting sub-anagrams of t contained in s!
- [Two Sum](https://github.com/ggcr/go-algorithms/tree/master/Array-And-Hashing/two-sum) ([Leetcode 1](https://leetcode.com/problems/two-sum/)): Brute force O(N^2) and efficient O(N) using a hashmap.
- [Group Anagrams](https://github.com/ggcr/go-algorithms/tree/master/Array-And-Hashing/group-anagrams) ([Leetcode 49](https://leetcode.com/problems/group-anagrams/)): Brute force O(N^3) and efficient O(M*N) using a hashmap.
- [Top K Frequent Elements](https://github.com/ggcr/go-algorithms/tree/master/Array-And-Hashing/top-k-frequent-elements) ([Leetcode 347](https://leetcode.com/problems/top-k-frequent-elements)): Not efficient solution in O(n log n) time and an efficient Bucket Sort solution in O(n).
- [Product of Array expect Self](https://github.com/ggcr/go-algorithms/tree/master/Array-And-Hashing/product-of-array-except-self/) ([Leetcode 238](https://leetcode.com/problems/product-of-array-except-self/)): Efficient solution in O(N), there is a simpler formula but does the trick.
- [Valid Sudoku](https://github.com/ggcr/go-algorithms/tree/master/Array-And-Hashing/valid-sudoku) ([Leetcode 36](https://leetcode.com/problems/valid-sudoku)): Efficient solution in O(N*M) **faster than 100%** ⭐.
- [Encode and Decode Strings](https://github.com/ggcr/go-algorithms/tree/master/Array-And-Hashing/encode-decode-strings) ([Leetcode 271](https://leetcode.com/problems/encode-and-decode-strings/)): Two O(N) solutions (with and without the standard library).
- [Longest Consecutive Sequence](https://github.com/ggcr/go-algorithms/tree/master/Array-And-Hashing/longest-consecutive-sequence) ([Leetcode 128](https://leetcode.com/problems/longest-consecutive-sequence/)): O(N) solution: first we build up the hashset, then we count steaks above and below (value+1 and value-1).
### Two Pointers
- [Valid Palindrome](https://github.com/ggcr/go-algorithms/tree/master/Two-Pointers/valid-palyndrome) ([Leetcode 125](https://leetcode.com/problems/valid-palyndrome)): O(N) efficient solution, working with Go strings, runes or bytes is always a pleasure for its standard lib! **Faster than 100%** ⭐.
- [Two Sum II](https://github.com/ggcr/go-algorithms/tree/master/Two-Pointers/two-sum-2) ([Leetcode 167](https://leetcode.com/problems/two-sum-ii-input-array-is-sorted)): O(N^2) brute force solution and O(N) efficient solution, because the array is sorted we don't need the hashmap approach.
- 🆘 [3Sum](https://github.com/ggcr/go-algorithms/tree/master/Two-Pointers/3Sum) ([Leetcode 15](https://leetcode.com/problems/3sum/)): O(N^3) brute force solution and O(N^2) Better approach in which we avoid duplicates by sorting the array and not using as first value (i) twice the same value. Then we reduce the rest of the two nums to a two sum problem.
- [Container With Most Water](https://github.com/ggcr/go-algorithms/tree/master/Two-Pointers/Container-with-most-water) ([Leetcode 11](https://leetcode.com/problems/container-with-most-water/)): O(N) time. It's not a hard problem it just requires problem specification.
### Sliding Window
- [Best Time to Buy and Sell Stock](https://github.com/ggcr/go-algorithms/tree/master/Sliding-Window/best-time-to-buy-and-sell-stock) ([Leetcode 121](https://leetcode.com/problems/best-time-to-buy-and-sell-stock/)): Brute force O(N^2) and Sliding Window with 2 pointers in O(N) time. 
- [Longests Substring without repeating characters](https://github.com/ggcr/go-algorithms/tree/master/Sliding-Window/longest-substring-without-repeating-characters) ([Leetcode 3](https://leetcode.com/problems/longest-substring-without-repeating-characters/)): O(N) time. It was tricky because our reference pointer is the tail, and typically we are so used to have the reference pointer the leading one. **Faster than 100%** ⭐.
- 🆘 [Longest Repeating Character Replacement](https://github.com/ggcr/go-algorithms/tree/master/Sliding-Window/longest-repeating-character-replacement) ([Leetcode 424](https://leetcode.com/problems/longest-repeating-character-replacement/)): In O(N) time. This one was hard, did not know how to approach in a clear way. 

## Other
- [skiena-1.30](https://github.com/ggcr/go-algorithms/tree/master/Other/skiena-1.30): Nearest Neighbour vs Closest Pair heuristic in the TSP problem.  
- [leetcode-739](https://github.com/ggcr/go-algorithms/tree/master/Other/leetcode-739) ([Problem in Leetcode](https://leetcode.com/problems/daily-temperatures/)): Brute force implementation O(N^2)
- [leetcode-61](https://github.com/ggcr/go-algorithms/tree/master/Other/leetcode-61) ([Problem in Leetcode](https://leetcode.com/problems/rotate-list/)): Linked List implementation from scratch and Rotation algorithm in O(N^2) and an improved version in O(N) time **faster than 100%** ⭐.
- [leetcode-324](https://github.com/ggcr/go-algorithms/tree/master/Other/leetcode-324) ([Problem in Leetcode](https://leetcode.com/problems/wiggle-sort-ii/)): Wiggle Sort. Clear approach, but not optimal.

