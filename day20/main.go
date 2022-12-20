package main

import (
	"fmt"
	"github.com/mrbarge/aoc2022/helper"
	"math"
	"os"
)

const DECRYPT_KEY = 811589153

type Node struct {
	val  int64
	len  int
	prev *Node
	next *Node
}

func makeNodes(lines []int, partTwo bool) []*Node {
	nodes := make([]*Node, 0)

	llen := len(lines)
	node := Node{
		val:  int64(lines[0]),
		next: nil,
		prev: nil,
		len:  llen,
	}
	if partTwo {
		node.val *= DECRYPT_KEY
	}
	nodes = append(nodes, &node)

	for k := 1; k < llen; k++ {
		n := Node{
			val:  int64(lines[k]),
			len:  len(lines),
			prev: nodes[k-1],
			next: nil,
		}
		if partTwo {
			n.val *= DECRYPT_KEY
		}
		nodes[k-1].next = &n
		nodes = append(nodes, &n)
	}
	nodes[0].prev = nodes[llen-1]
	nodes[llen-1].next = nodes[0]

	return nodes
}

func problem(lines []int, partTwo bool) (int64, error) {
	count := int64(0)

	nodes := makeNodes(lines, partTwo)

	var zeroNode *Node

	loops := 1
	if partTwo {
		loops = 10
	}
	for i := 0; i < loops; i++ {
		for _, v := range nodes {
			if v.val == 0 {
				zeroNode = v
			}
			shuffle(v)
		}
	}

	ni := zeroNode
	for i := 0; i < 3001; i++ {
		if i%1000 == 0 {
			count += ni.val
		}
		ni = ni.next
	}

	return count, nil
}

func shuffle(n *Node) {

	n.prev.next = n.next
	n.next.prev = n.prev

	numMoves := int(math.Abs(float64(n.val)))
	dest := n

	for i := 0; i < (numMoves % (n.len - 1)); i++ {
		if n.val < 0 {
			dest.next = dest.prev
			dest.prev = dest.prev.prev
		} else {
			dest.next = dest.next.next
			dest.prev = dest.prev.next
		}
	}

	dest.prev.next = dest
	dest.next.prev = dest
}

func main() {
	fh, _ := os.Open("input.txt")
	lines, err := helper.ReadLinesAsInt(fh)
	if err != nil {
		fmt.Printf("Unable to read input: %v\n", err)
		return
	}

	ans, err := problem(lines, false)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Part one: %v\n", ans)

	ans, err = problem(lines, true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Part two: %v\n", ans)

}
