package _97_serialize_and_deserialize_binary_tree

import (
	"math"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
	null int
}

func Constructor() Codec {
	return Codec{null: math.MinInt64}
}

func (c *Codec) Serialize(root *TreeNode) string {
	if root == nil {
		return ``
	}
	s := make([]int, 0)
	q := []*TreeNode{root}
	for len(q) > 0 {
		v := q[0]
		q = q[1:]

		if v == nil {
			s = append(s, c.null)
		} else {
			s = append(s, v.Val)
			q = append(q, v.Left, v.Right)
		}
	}
	return c.serializeSlice(s)
}

func (c *Codec) Deserialize(data string) *TreeNode {
	if len(data) == 0 {
		return nil
	}

	vs := c.deserializeSlice(data)
	n := len(vs)
	root := &TreeNode{Val: vs[0]}
	i := 1
	q := []*TreeNode{root}
	for len(q) > 0 {
		v := q[0]
		q = q[1:]
		for j := 0; j < 2; j++ {
			var newNode *TreeNode
			if i < n && vs[i] != c.null {
				newNode = &TreeNode{Val: vs[i]}
				q = append(q, newNode)
			}
			if j&1 == 0 {
				v.Left = newNode
			} else {
				v.Right = newNode
			}
			i++
		}
	}
	return root
}

func (c *Codec) serializeSlice(s []int) string {
	n := len(s)
	var sb strings.Builder
	for i := 0; i < n; i++ {
		if i > 0 {
			sb.WriteByte(',')
		}
		if s[i] == c.null {
			sb.WriteString("null")
		} else {
			sb.WriteString(strconv.Itoa(s[i]))
		}
	}
	res := sb.String()
	for strings.HasSuffix(res, ",null") {
		res = res[:len(res)-5]
	}
	return res
}

func (c *Codec) deserializeSlice(data string) []int {
	splits := strings.Split(data, `,`)
	s := make([]int, 0, len(splits))
	for _, split := range splits {
		if split == "null" {
			s = append(s, c.null)
		} else {
			v, _ := strconv.Atoi(split)
			s = append(s, v)
		}

	}
	return s
}
