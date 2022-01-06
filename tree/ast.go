package tree

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

var ops = []string{"(", ")", "\\", "{", "}", "|", "+", "#", "."}
var id = 1

func Find(s string) bool {
	for i := 0; i < len(ops); i++ {
		if s == ops[i] {
			return true
		}
	}
	return false
}

func AddConcatenations(regex string) (newRegex string) {
	newRegex = ""
	var i = 0
	for i < len(regex)-1 {
		var Flag = Find(string(regex[i]))
		var Check = Find(string(regex[i+1]))
		if !Flag {
			if Check {
				if regex[i+1] == '(' {
					newRegex += string(regex[i]) + "."
					i++
				} else if regex[i+1] == '\\' {
					newRegex += string(regex[i]) + "."
					i++
					newRegex += string(regex[i])
					i++
					for unicode.IsDigit(rune(regex[i])) {
						newRegex += string(regex[i])
						i++
					}
					newRegex += "."
				} else if regex[i+1] == '{' {
					for regex[i] != '}' {
						newRegex += string(regex[i])
						i++
					}
					if regex[i+1] != ')' {
						newRegex += string(regex[i]) + "."
					} else {
						newRegex += string(regex[i])
					}
					i++
				} else if regex[i+1] == '#' {
					newRegex += string(regex[i]) + "."
					i++
				} else {
					newRegex += string(regex[i])
					i++
				}
			} else {
				newRegex += string(regex[i]) + "."
				i++
			}
		} else {
			if regex[i] == ')' {
				if regex[i+1] == '(' || !Find(string(regex[i+1])) {
					newRegex += string(regex[i]) + "."
					i++
				} else if regex[i+1] == '#' || regex[i+1] == '\\' {
					newRegex += string(regex[i]) + "."
					i++
				} else if regex[i+1] == '{' {
					newRegex += string(regex[i])
					i++
					for regex[i] != '}' {
						newRegex += string(regex[i])
						i++
					}
					if regex[i+1] != ')' {
						newRegex += string(regex[i]) + "."
					} else {
						newRegex += string(regex[i])
					}
					i++
				} else {
					newRegex += string(regex[i])
					i++
				}
			} else if regex[i] == '+' {
				if regex[i+1] == '.' {
					newRegex += string(regex[i]) + "."
					i+=2
				} else {
					newRegex += string(regex[i]) + "."
					i++
				}
			} else if regex[i] == '#' {
				newRegex += string(regex[i]) + string(regex[i+1]) + "."
				i += 2
			} else if regex[i] == '}' {
				if regex[i+1] != ')' {
					newRegex += string(regex[i]) + "."
				} else {
					newRegex += string(regex[i])
				}
			} else {
				newRegex += string(regex[i])
				i++
			}
		}
	}
	newRegex += string(regex[len(regex)-1])
	var copyNewRegex string
	var minus = 0
	for _, c := range newRegex {
		copyNewRegex += string(c)
	}
	var Brack = false
	var IndBracket = -1
	for ind, char := range copyNewRegex {
		if char == '(' {
			Brack = true
			IndBracket = ind
		}
		if char == ')' {
			Brack = false
			IndBracket = -1
		}
		if char == ':' && Brack && (ind-IndBracket) > 0 {
			for j := IndBracket + -minus; j < ind-1-minus; j++ {
				if newRegex[j] == '.' {
					newRegex = newRegex[:j] + newRegex[j+1:]
					minus++
				}
			}
			newRegex = newRegex[:ind-minus-1] + ":" + newRegex[ind-minus+2:]
			minus += 2
		}
	}
	return
}

func ReplaceRepeat(regex string) string {
	var i = 0
	var repeat string = ""
	var add string = ""
	var start, end int = -1, -1
	var endRepeated = -1
	var repeatBracket = false
	for i < len(regex) {
		if i > 1 && regex[i-1] == '(' {
			repeat = ""
			for regex[i] != ')' {
				if regex[i] == '{' {
					break
				}
				repeat += string(regex[i])
				i++
			}
		}
		if regex[i] == '{' {
			endRepeated = i - 1
			if regex[i-1] == ')' {
				repeatBracket = true
			}
			i = i + 1
			var x, y string = "", ""
			for regex[i] != ',' {
				x += string(regex[i])
				i++
			}
			i++
			for regex[i] != '}' {
				y += string(regex[i])
				i++
			}
			end = i
			add = ""
			if !repeatBracket {
				if y == "" {
					var IntX, _ = strconv.Atoi(x)
					add += strings.Repeat(string(regex[endRepeated])+".", IntX)
					add = add[:len(add)-1]
					add += "+"
				} else {
					var IntX, _ = strconv.Atoi(x)
					var IntY, _ = strconv.Atoi(y)
					for j := 0; j <= IntY-IntX; j++ {
						add += strings.Repeat(string(regex[endRepeated])+".", IntX+j)
						add = add[:len(add)-1]
						if j != IntY-IntX {
							add += "|"
						}
					}
				}
				regex = regex[:endRepeated] + "(" + add + ")" + regex[end+1:]
			} else {
				if y == "" {
					var IntX, _ = strconv.Atoi(x)
					add += strings.Repeat(repeat+".", IntX-2)
					add = add[:len(add)-1]
					add += ".(" + repeat + ")+"
				} else {
					var IntX, _ = strconv.Atoi(x)
					var IntY, _ = strconv.Atoi(y)
					for j := 0; j <= IntY-IntX; j++ {
						add += strings.Repeat(repeat+".", IntX+j)
						add = add[:len(add)-1]
						if j != IntY-IntX {
							add += "|"
						}
					}
				}
				regex = regex[:start] + "(" + add + ")" + regex[end+1:]
			}
		}
		i++
	}
	return regex
}

func CreateTokens(Regex string) (tokens []string) {
	var c = ""
	var FigBr = false
	var shield = false
	for _, r := range Regex {
		if (r == '.' || r == '|' || r == '+') && !FigBr && !shield {
			tokens = append(tokens, string(r))
		} else if r == '{' {
			FigBr = true
			c += string(r)
		} else if r == '}' {
			FigBr = false
			c += string(r)
			tokens = append(tokens, c)
			c = ""
		} else if FigBr {
			c += string(r)
		} else if r == '\\' {
			c += string(r)
			shield = true
		} else if shield {
			if r == ')' {
				tokens = append(tokens, c)
				tokens = append(tokens, string(r))
				c = ""
				shield = false
			} else if r == '.' {
				tokens = append(tokens, c)
				tokens = append(tokens, string(r))
				c = ""
				shield = false
			} else {
				c += string(r)
			}
		} else if !FigBr {
			tokens = append(tokens, string(r))
		} else if r == '(' || r == ')' && !FigBr {
			if c != "" {
				tokens = append(tokens, c)
			}
			tokens = append(tokens, string(r))
		}
	}
	var minus = 0
	var Brack = false
	var indBracket = -1
	var copyTokens []string
	for _, ind := range tokens {
		copyTokens = append(copyTokens, ind)
	}
	for ind, char := range copyTokens {
		if char == "(" {
			Brack = true
			indBracket = ind
		}
		if char == ")" {
			Brack = false
			indBracket = -1
		}
		if char == ":" && Brack && (ind-indBracket) > 0 {
			var s = ""
			var Ind = ind - minus
			var IndBracket = indBracket - minus
			var shift = 0 - minus
			for j := IndBracket + 1; j < Ind; j++ {
				s += tokens[j]
				minus++
			}
			minus--
			copy(tokens[IndBracket+2:], tokens[Ind:])
			tokens[IndBracket+1] = s
			tokens = tokens[:len(tokens)-(shift+minus)]
		}
	}
	return
}

func ReplaceGroup(tokens []string) {
	var i = 0
	for i < len(tokens) {
		if tokens[i] == "(" && tokens[i+2] == ":" {

		}
	}
}

func isDigit(str string) bool {
	for i := range str {
		if str[i] < '0' || str[i] > '9' {
			return false
		}
	}
	return true
}

func CreateNodes(tokens []string, idMap map[int]string, leafNodes map[string][]int) (nodes []*Node) {
	var i = 0
	for i < len(tokens) {
		var tok = tokens[i]
		if tok == "(" || tok == ")" {
			nodes = append(nodes, &Node{Type: Bracket, Val: tok})
			i++
		} else if len(tok) == 1 && !Find(tok) && tokens[i+1] != ":" {
			var n *Node
			if tok != "^" {
				n = &Node{
					Id:       id,
					Parent:   nil,
					Left:     nil,
					Right:    nil,
					Type:     LeafNode,
					Val:      tok,
					FirstPos: []int{id},
					LastPos:  []int{id},
					Nullable: false,
				}
				idMap[id] = tok
			} else {
				n = &Node{
					Id:       id,
					Parent:   nil,
					Left:     nil,
					Right:    nil,
					Type:     LeafNode,
					Val:      tok,
					FirstPos: []int{id},
					LastPos:  []int{id},
					Nullable: true,
				}
				idMap[id] = tok
			}
			id++
			if n.Val != "^" && n.Val != "$" {
				_, ok := leafNodes[n.Val]
				if ok {
					leafNodes[n.Val] = append(leafNodes[n.Val], n.Id)
				} else {
					leafNodes[n.Val] = []int{n.Id}
				}
			}
			nodes = append(nodes, n)
			i++
		} else if len(tok) == 1 && Find(tok) && tokens[i+1] != ":" {
			if tok == "." {
				nodes = append(nodes, &Node{
					Id:       -1,
					Parent:   nil,
					Left:     nil,
					Right:    nil,
					Type:     Concat,
					Val:      tok,
					FirstPos: nil,
					LastPos:  nil,
					Nullable: false,
				})
				i++
			} else if tok == "|" {
				nodes = append(nodes, &Node{
					Id:       -1,
					Parent:   nil,
					Left:     nil,
					Right:    nil,
					Type:     Or,
					Val:      tok,
					FirstPos: nil,
					LastPos:  nil,
					Nullable: false,
				})
				i++
			} else if tok == "+" {
				nodes = append(nodes, &Node{
					Id:       -1,
					Parent:   nil,
					Left:     nil,
					Right:    nil,
					Type:     Concat,
					Val:      ".",
					FirstPos: nil,
					LastPos:  nil,
					Nullable: false,
				})
				var n = &Node{
					Id:       id,
					Parent:   nil,
					Left:     nil,
					Right:    nil,
					Type:     LeafNode,
					Val:      tokens[i-1],
					FirstPos: []int{id},
					LastPos:  []int{id},
					Nullable: false,
				}
				nodes = append(nodes, n)
				idMap[id] = tok
				id++
				_, ok := leafNodes[n.Val]
				if ok {
					leafNodes[n.Val] = append(leafNodes[n.Val], n.Id)
				} else {
					leafNodes[n.Val] = []int{n.Id}
				}
				nodes = append(nodes, &Node{
					Id:       -1,
					Parent:   nil,
					Left:     nil,
					Right:    nil,
					Type:     Star,
					Val:      "*",
					FirstPos: nil,
					LastPos:  nil,
					Nullable: false,
				})
				i++
			} else if tok == "#" {
				nodes = append(nodes, &Node{
					Id:       -1,
					Parent:   nil,
					Left:     nil,
					Right:    nil,
					Type:     Sharp,
					Val:      tok,
					FirstPos: nil,
					LastPos:  nil,
					Nullable: false,
				})
				nodes = append(nodes, &Node{
					Id:       id,
					Parent:   nil,
					Left:     nil,
					Right:    nil,
					Type:     LeafNode,
					Val:      tokens[i+1],
					FirstPos: []int{id},
					LastPos:  []int{id},
					Nullable: false,
				})
				idMap[id] = tokens[i+1]
				leafNodes[tokens[i+1]] = append(leafNodes[tokens[i+1]], id)
				id++
				i += 2
			}
		} else {
			if tok[0] == '\\' {
				nodes = append(nodes, &Node{
					Id:       -1,
					Parent:   nil,
					Left:     nil,
					Right:    nil,
					Type:     Reference,
					Val:      tok[1:],
					FirstPos: nil,
					LastPos:  nil,
					Nullable: false,
				})
				i++
			} else if isDigit(tok) {
				nodes = append(nodes, &Node{
					Id:       -1,
					Parent:   nil,
					Left:     nil,
					Right:    nil,
					Type:     Group,
					Val:      tok,
					FirstPos: nil,
					LastPos:  nil,
					Nullable: false,
				})
				i += 2
			}
		}
	}
	return nodes
}

func ClosestBrackets(tokens []*Node) (first int, second int) {
	first = 0
	second = len(tokens) - 1
	var currF = 0
	var currS = 0
	var priority = true
	for i, tok := range tokens {
		if tok.Val == "(" && tokens[i+2].Val != ":" {
			currF = i
			priority = true
		} else if tok.Val == ")" {
			if priority {
				currS = i
				if currS-currF < second-first {
					second = currS
					first = currF
				}
				priority = false
			}
		}
	}
	return
}

func FindMerge(a []int, x int) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

func Merge(first []int, second []int) []int {
	var res []int
	res = append(res, first...)
	for _, c := range second {
		if !FindMerge(first, c) {
			res = append(res, c)
		}
	}
	sort.Ints(res)
	return res
}

func CopyTree(node *Node, leafNodes map[string][]int) *Node {
	if node == nil {
		return nil
	}
	var newnode *Node
	if node.Type == LeafNode {
		newnode = &Node{
			Id:       id,
			Parent:   nil,
			Left:     nil,
			Right:    nil,
			Type:     LeafNode,
			Val:      node.Val,
			FirstPos: node.FirstPos,
			LastPos:  node.LastPos,
			Nullable: node.Nullable,
		}
		leafNodes[node.Val] = append(leafNodes[node.Val], id)
		id++
	} else {
		newnode = &Node{
			Id:       -1,
			Parent:   nil,
			Left:     nil,
			Right:    nil,
			Type:     node.Type,
			Val:      node.Val,
			FirstPos: node.FirstPos,
			LastPos:  node.LastPos,
			Nullable: node.Nullable,
		}
	}
	newnode.Left = CopyTree(node.Left, leafNodes)
	newnode.Right = CopyTree(node.Right, leafNodes)
	return newnode
}

var m = make(map[string]*Node)

func CreateSubtree(nodes []*Node, first int, second int, followpos [][]int, leafNodes map[string][]int) ([]*Node, [][]int) {
	var currNode *Node
	var i int
	_, ok := m[nodes[first+1].Val]
	if nodes[first+1].Type == Group && !ok {
		var groupNum = nodes[first+1].Val
		i = first + 2
		second--
		for i <= second {
			if nodes[i].Type == Sharp {
				nodes[i].Left = nodes[i+1]
				nodes[i].Right = nil
				nodes[i].Nullable = nodes[i+1].Nullable
				nodes[i].FirstPos = nodes[i+1].FirstPos
				nodes[i].LastPos = nodes[i+1].LastPos
				nodes[i+1].Parent = nodes[i]
				currNode = nodes[i]
				copy(nodes[i+1:], nodes[i+2:])
				nodes = nodes[:len(nodes)-1]
				second--
				i++
			}
			i++
		}
		i = first + 2
		for i <= second {
			if nodes[i].Type == Star {
				nodes[i].Left = nodes[i-1]
				nodes[i].Right = nil
				nodes[i].Nullable = true
				nodes[i].FirstPos = nodes[i-1].FirstPos
				nodes[i].LastPos = nodes[i-1].LastPos
				nodes[i-1].Parent = nodes[i]
				for _, k := range nodes[i].LastPos {
					if followpos[k-1] != nil {
						followpos[k-1] = Merge(followpos[k-1], nodes[i].FirstPos)
					} else {
						followpos[k-1] = nodes[i].FirstPos
					}
				}
				currNode = nodes[i]
				copy(nodes[i-1:], nodes[i:])
				nodes = nodes[:len(nodes)-1]
				second--
				i--
			}
			i++
		}
		i = first + 3
		for i <= second-1 {
			if nodes[i].Type == Concat && nodes[i].Left == nil && nodes[i].Right == nil {
				nodes[i].Left = nodes[i-1]
				nodes[i].Right = nodes[i+1]
				nodes[i].Nullable = nodes[i-1].Nullable && nodes[i+1].Nullable
				if nodes[i-1].Nullable {
					nodes[i].FirstPos = Merge(nodes[i-1].FirstPos, nodes[i+1].FirstPos)
				} else {
					nodes[i].FirstPos = nodes[i-1].FirstPos
				}
				if nodes[i+1].Nullable {
					nodes[i].LastPos = Merge(nodes[i-1].LastPos, nodes[i+1].LastPos)
				} else {
					nodes[i].LastPos = nodes[i+1].LastPos
				}
				nodes[i-1].Parent = nodes[i]
				nodes[i+1].Parent = nodes[i]
				for _, k := range nodes[i].Left.LastPos {
					if followpos[k-1] != nil {
						followpos[k-1] = Merge(followpos[k-1], nodes[i].Right.FirstPos)
					} else {
						followpos[k-1] = nodes[i].Right.FirstPos
					}
				}
				currNode = nodes[i]
				copy(nodes[i:], nodes[i+2:])
				nodes[i-1] = currNode
				nodes = nodes[:len(nodes)-2]
				i -= 2
				second = second - 2
			}
			i++
		}
		i = first + 3
		for i <= second-1 {
			if i >= first+2 {
				if nodes[i].Type == Or && nodes[i].Left == nil && nodes[i].Right == nil {
					nodes[i].Left = nodes[i-1]
					nodes[i].Right = nodes[i+1]
					nodes[i].Nullable = nodes[i-1].Nullable || nodes[i+1].Nullable
					nodes[i].FirstPos = Merge(nodes[i-1].FirstPos, nodes[i+1].FirstPos)
					nodes[i].LastPos = Merge(nodes[i-1].LastPos, nodes[i+1].LastPos)
					nodes[i-1].Parent = nodes[i]
					nodes[i+1].Parent = nodes[i]
					currNode = nodes[i]
					copy(nodes[i:], nodes[i+2:])
					nodes[i-1] = currNode
					nodes = nodes[:len(nodes)-2]
					i -= 2
					second = second - 2
				}
				i++
			}
			i++
		}
		nodes[first+1].Left = currNode
		nodes[first+1].Right = nil
		nodes[first+1].Nullable = currNode.Nullable
		nodes[first+1].FirstPos = currNode.FirstPos
		nodes[first+1].LastPos = currNode.LastPos
		currNode.Parent = nodes[first+1]
		nodes[first] = nodes[first+1]
		copy(nodes[first+1:], nodes[second+2:])
		nodes = nodes[:len(nodes)-3]
		m[groupNum] = currNode
	} else if second-first != 2 {
		i = first + 1
		second -= 1
		for i <= second {
			if nodes[i].Type == Reference {
				var key = nodes[i].Val
				if _, ok = m[key]; !ok {
					panic("Group not previously mentioned\n")
				}
				nodes[i].Left = CopyTree(m[key], leafNodes)
				//fmt.Println(leafNodes)
				nodes[i].Nullable = m[key].Nullable
				nodes[i].FirstPos = m[key].FirstPos
				nodes[i].LastPos = m[key].LastPos
			}
			i++
		}
		i = first + 1
		for i <= second {
			if nodes[i].Type == Sharp {
				nodes[i].Left = nodes[i+1]
				nodes[i].Right = nil
				nodes[i].Nullable = nodes[i+1].Nullable
				nodes[i].FirstPos = nodes[i+1].FirstPos
				nodes[i].LastPos = nodes[i+1].LastPos
				nodes[i+1].Parent = nodes[i]
				currNode = nodes[i]
				copy(nodes[i+1:], nodes[i+2:])
				nodes = nodes[:len(nodes)-1]
				second--
				i++
			}
			i++
		}
		i = first + 1
		for i <= second {
			if nodes[i].Type == Star {
				nodes[i].Left = nodes[i-1]
				nodes[i].Right = nil
				nodes[i-1].Parent = nodes[i]
				nodes[i].Nullable = true
				nodes[i].FirstPos = nodes[i-1].FirstPos
				nodes[i].LastPos = nodes[i-1].LastPos
				for _, k := range nodes[i].LastPos {
					if followpos[k-1] != nil {
						followpos[k-1] = Merge(followpos[k-1], nodes[i].FirstPos)
					} else {
						followpos[k-1] = nodes[i].FirstPos
					}
				}
				currNode = nodes[i]
				copy(nodes[i-1:], nodes[i:])
				nodes = nodes[:len(nodes)-1]
				second--
				i--
			}
			i++
		}
		i = first + 1
		i = first + 1
		for i <= second-1 {
			if i >= first+2 {
				if nodes[i].Type == Concat && nodes[i].Left == nil && nodes[i].Right == nil {
					nodes[i].Left = nodes[i-1]
					nodes[i].Right = nodes[i+1]
					nodes[i].Nullable = nodes[i-1].Nullable && nodes[i+1].Nullable
					if nodes[i-1].Nullable {
						nodes[i].FirstPos = Merge(nodes[i-1].FirstPos, nodes[i+1].FirstPos)
					} else {
						nodes[i].FirstPos = nodes[i-1].FirstPos
					}
					if nodes[i+1].Nullable {
						nodes[i].LastPos = Merge(nodes[i-1].LastPos, nodes[i+1].LastPos)
					} else {
						nodes[i].LastPos = nodes[i+1].LastPos
					}
					nodes[i-1].Parent = nodes[i]
					nodes[i+1].Parent = nodes[i]
					for _, k := range nodes[i].Left.LastPos {
						if followpos[k-1] != nil {
							followpos[k-1] = Merge(followpos[k-1], nodes[i].Right.FirstPos)
						} else {
							followpos[k-1] = nodes[i].Right.FirstPos
						}
					}
					currNode = nodes[i]
					copy(nodes[i:], nodes[i+2:])
					nodes[i-1] = currNode
					nodes = nodes[:len(nodes)-2]
					i -= 2
					second = second - 2
				}
				i++
			}
			i++
		}
		i = first + 1
		for i <= second-1 {
			if i >= first+2 {
				if nodes[i].Type == Or && nodes[i].Left == nil && nodes[i].Right == nil {
					nodes[i].Left = nodes[i-1]
					nodes[i].Right = nodes[i+1]
					nodes[i].Nullable = nodes[i-1].Nullable || nodes[i+1].Nullable
					nodes[i].FirstPos = Merge(nodes[i-1].FirstPos, nodes[i+1].FirstPos)
					nodes[i].LastPos = Merge(nodes[i-1].LastPos, nodes[i+1].LastPos)
					nodes[i-1].Parent = nodes[i]
					nodes[i+1].Parent = nodes[i]
					currNode = nodes[i]
					copy(nodes[i:], nodes[i+2:])
					nodes[i-1] = currNode
					nodes = nodes[:len(nodes)-2]
					i -= 2
					second = second - 2
				}
				i++
			}
			i++
		}
		copy(nodes[first+1:], nodes[second+2:])
		nodes[first] = currNode
		nodes = nodes[:len(nodes)-2]
	} else {
		currNode = nodes[first+1]
		copy(nodes[first+1:], nodes[second+1:])
		nodes[first] = currNode
		nodes = nodes[:len(nodes)-2]
	}
	return nodes, followpos
}

func CreateTree(regex string, idMap map[int]string, leafNodes map[string][]int, followpos [][]int) (*Node, [][]int) {
	var nodes = CreateNodes(CreateTokens(ReplaceRepeat(AddConcatenations(regex))), idMap, leafNodes)
	var first, second = ClosestBrackets(nodes)
	for i := 0; i < len(idMap); i++ {
		followpos = append(followpos, nil)
	}
	for second-first > 1 {
		nodes, followpos = CreateSubtree(nodes, first, second, followpos, leafNodes)
		first, second = ClosestBrackets(nodes)
	}
	id=1
	return nodes[0], followpos
}

func Print(nodes []*Node) {
	for i := 0; i < len(nodes); i++ {
		if nodes[i].Type == 1 {
			fmt.Printf("concat ")
		}
		if nodes[i].Type == 2 {
			fmt.Printf("or ")
		}
		if nodes[i].Type == 3 {
			fmt.Printf("star ")
		}
		if nodes[i].Type == 4 {
			fmt.Printf("sharp ")
		}
		if nodes[i].Type == 5 {
			fmt.Printf("group ")
		}
		if nodes[i].Type == 6 {
			fmt.Printf("reference ")
		}
		if nodes[i].Type == 7 {
			fmt.Printf("repeat ")
		}
		if nodes[i].Type == 8 {
			fmt.Printf("leaf ")
		}
		if nodes[i].Type == 9 {
			fmt.Printf("bracket ")
		}
	}
	fmt.Println()
}

func PrintTree(n *Node) {
	if n != nil {
		PrintSubTree(n)
	}
}

func PrintSubTree(n *Node) {
	var parentState *NodePrint
	if rootState != nil {
		fmt.Printf(" ")
		var s = rootState
		for s.ChildState != nil {
			if s.PrintLastChild {
				fmt.Printf(" ")
			} else {
				fmt.Printf("│ ")
			}
			s = s.ChildState
		}
		parentState = s
		if parentState.PrintLastChild {
			fmt.Printf("└")
		} else {
			fmt.Printf("├")
		}
	} else {
		parentState = nil
	}
	fmt.Printf("▶%s\n", n.Val)
	if n.Left != nil || n.Right != nil {
		var s NodePrint
		if parentState != nil {
			parentState.ChildState = &s
		} else {
			rootState = &s
		}
		s.ChildState = nil

		if n.Left != nil {
			s.PrintLastChild = n.Right == nil
			PrintSubTree(n.Left)
		}
		if n.Right != nil {
			s.PrintLastChild = true
			PrintSubTree(n.Right)
		}
		if parentState != nil {
			parentState.ChildState = nil
		} else {
			rootState = nil
		}
	}
}
