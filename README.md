# go-algorithms
This is a collection of algorithms and data-structure implemenetations in Go made by myself!

## How this works
I work with a TDD methodology. Each folder has its [name].go file and its [name]_test.go file. In order to run the go code just go to the desired path and run the following command:
```bash
go test [-v] // The -v flag is optional
```

## Data Structures
- [Linked List](https://github.com/ggcr/go-algorithms/tree/master/linked-list)

## Neet-Code 150
### Arrays & Hashing
- [Contains Duplicate](https://github.com/ggcr/go-algorithms/tree/master/contains-duplicate) ([Leetcode 217](https://leetcode.com/problems/contains-duplicate/)): Brute force O(N^2), Merge Sort method O(N logN) and Map solution O(N).
- [Valid Anagram](https://github.com/ggcr/go-algorithms/tree/master/valid-anagram) ([Leetcode 242](https://leetcode.com/problems/valid-anagram/)): An efficient Hashmap solution O(N) and another implementation for detecting sub-anagrams of t contained in s!
- [Two Sum](https://github.com/ggcr/go-algorithms/tree/master/two-sum) ([Leetcode 1](https://leetcode.com/problems/two-sum/)): Brute force O(N^2) and efficient O(N) using a hashmap.
- [Group Anagrams](https://github.com/ggcr/go-algorithms/tree/master/group-anagrams) ([Leetcode 49](https://leetcode.com/problems/group-anagrams/)): Brute force O(N^3) and efficient O(M*N) using a hashmap.
- [Top K Frequent Elements](https://github.com/ggcr/go-algorithms/tree/master/top-k-frequent-elements) ([Leetcode 347](https://leetcode.com/problems/top-k-frequent-elements)): Not efficient solution in O(n log n) time and an efficient Bucket Sort solution in O(n).
- [Product of Array expect Self](https://github.com/ggcr/go-algorithms/tree/master/product-of-array-except-self/) ([Leetcode 238](https://leetcode.com/problems/product-of-array-except-self/)): Efficient solution in O(N), there is a simpler formula but does the trick.

## Other
- [skiena-1.30](https://github.com/ggcr/go-algorithms/tree/master/skiena-1.30): Nearest Neighbour vs Closest Pair heuristic in the TSP problem.  
- [leetcode-739](https://github.com/ggcr/go-algorithms/tree/master/leetcode-739) ([Problem in Leetcode](https://leetcode.com/problems/daily-temperatures/)): Brute force implementation O(N^2)
- [leetcode-61](https://github.com/ggcr/go-algorithms/tree/master/leetcode-61) ([Problem in Leetcode](https://leetcode.com/problems/rotate-list/)): Linked List implementation from scratch and Rotation algorithm in O(N^2) and an improved version in O(N) time **faster than 100%** ⭐️
- [leetcode-324](https://github.com/ggcr/go-algorithms/tree/master/leetcode-324) ([Problem in Leetcode](https://leetcode.com/problems/wiggle-sort-ii/)): Wiggle Sort. Clear approach, but not optimal.

