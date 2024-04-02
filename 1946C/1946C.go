package main

import (
	"bufio"
	"fmt"
	"os"
)

// treeCutting calculates the minimum number of tree cuts required to obtain at least k non-empty connected components of size at least x.
//
// Parameters:
// - edges: a 2D array representing the edges of the tree. Each element [i] contains the indices of the nodes connected to node i.
// - n: the number of nodes in the tree.
// - k: the minimum number of non-empty connected components required.
//
// Return:
// - int: the maximum number of x
func treeCutting(edges [][]int, n, k int) int {
	// will check the tree if size x is possible or not for cuts greater than or equal to k
	check := func(x int) bool {
		// total no. of cuts
		res := 0
		// dfs on tree to get size of a node
		var dfs func(int, int) int
		dfs = func(v, f int) int {
			sz := 1

			for _, u := range edges[v] {
				if u == f {
					continue
				}
				sz += dfs(u, v)
			}
			// here we are making sure if size of a node is greater than or equal to x and it is not root node
			// if it's true then we will cut the tree from here and then increment the no. of cuts (res) by 1
			if sz >= x && f != v {
				res++
				sz = 0
			}
			// finally return the size of tree after all the cuts made in the way
			return sz
		}
		// get the size of tree starting from root node after all the cuts
		t := dfs(0, 0)
		// check if total cuts are greater than k or not
		// if not then check it is equal to k and the size of root node is also greater than or equal to x
		return res > k || (res == k && t >= x)
	}

	low, high := 1, n/(k+1)+1

	for high-low > 1 {
		mid := (low + high) / 2
		if check(mid) {
			low = mid
		} else {
			high = mid
		}
	}

	return low
}

func main() {
	// check for test file
	r := os.Stdin
	if f, err := os.Open("test.txt"); err == nil {
		r = f
		defer f.Close()
	}
	// get testcases t
	scanner := bufio.NewScanner(r)
	scanner.Scan()
	var t int
	fmt.Sscanf(scanner.Text(), "%d", &t)
	// check each testcases t

	for ; t > 0; t-- {
		var n, k int
		scanner.Scan()
		fmt.Sscanf(scanner.Text(), "%d %d", &n, &k)

		adj := make([][]int, n)

		for i := 0; i < n-1; i++ {
			var v, u int
			scanner.Scan()
			fmt.Sscanf(scanner.Text(), "%d %d", &v, &u)
			v--
			u--

			adj[v] = append(adj[v], u)
			adj[u] = append(adj[u], v)
		}
		// print ans
		fmt.Println(treeCutting(adj, n, k))
	}
}
