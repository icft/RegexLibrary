package dfa

import (
	"reflect"
	"regex/tree"
	"strings"
)

type State struct {
	Name        string
	StateNumber []int
	Dtran       map[string]*State
	num 		int
	Receive		bool
}

var num = 1

var DictNames = map[int]string{
	1:  "A",
	2:  "B",
	3:  "C",
	4:  "D",
	5:  "E",
	6:  "F",
	7:  "G",
	8:  "H",
	9:  "I",
	10: "J",
	11: "K",
	12: "L",
	13: "M",
	14: "N",
	15: "O",
	16: "P",
	17: "Q",
	18: "R",
	19: "S",
	20: "T",
	21: "U",
	22: "V",
	23: "W",
	24: "X",
	25: "Y",
	26: "Z",
}

type DFA struct {
	tree               tree.Ast
	FollowPos          [][]int
	InitialStateNumber []int
	InitialState       *State
	NumberStates	   int
	Alphabet 		   []string
}

func InitDFA(root tree.Ast, FollowPos [][]int, FirstPos []int) *DFA {
	return &DFA{tree: root, FollowPos: FollowPos, InitialStateNumber: FirstPos, NumberStates: 0}
}

func Convert(dfa *DFA, node *tree.Node, LeafNodes map[string][]int) *State {
	dfa.InitialState = &State{Name: DictNames[num], StateNumber: dfa.InitialStateNumber}
	dfa.NumberStates++
	num++
	var leftStates = []*State{dfa.InitialState}
	var seenStates []*State
	for len(leftStates) != 0 {
		var state = leftStates[len(leftStates)-1]
		leftStates = leftStates[:len(leftStates)-1]
		if !FindSeen(state, seenStates) {
			seenStates = append(seenStates, state)
			for k, v := range LeafNodes {
				var i []int
				for _, c := range state.StateNumber {
					if tree.FindMerge(v, c) {
						i = append(i, c)
					}
				}
				//fmt.Println(i)
				if len(i) != 0 {
					var nextStateNumber []int
					for _, c := range i {
						nextStateNumber = tree.Merge(nextStateNumber, dfa.FollowPos[c-1])
					}
					var broken = false
					for _, seen := range seenStates {
						if reflect.DeepEqual(nextStateNumber, seen.StateNumber) {
							if state.Dtran == nil {
								state.Dtran = make(map[string]*State)
							}
							state.Dtran[k] = seen
							broken = true
							break
						}
					}
					if !broken {
						var a, b = num, 1
						if num > 26 {
							a = num % 26
							b += num / 26
						}
						var nextState = &State{Name: strings.Repeat(DictNames[a], b),
							StateNumber: nextStateNumber}
						dfa.NumberStates++
						num++
						if state.Dtran == nil {
							state.Dtran = make(map[string]*State)
						}
						if !Find(dfa, k) {
							dfa.Alphabet = append(dfa.Alphabet, k)
						}
						state.Dtran[k] = nextState
						leftStates = append(leftStates, nextState)
					}
				}
			}
		}
	}
	return dfa.InitialState
}

func SetReceive(start *State) {
	var leftStates = []*State{start}
	var seenStates []*State
	for len(leftStates) != 0 {
		var state = leftStates[len(leftStates)-1]
		leftStates = leftStates[:len(leftStates)-1]
		if !FindSeen(state, seenStates) {
			seenStates = append(seenStates, state)
			k, _ := GetKeysValues(state.Dtran)
			if len(state.Dtran) == 0 || (len(state.Dtran)==1 && state.Dtran[k[0]] == state) {
				state.Receive = true
			} else {
				state.Receive = false
			}
			for _, v := range state.Dtran {
				leftStates = append(leftStates, v)
			}
		}
	}
}

func Compile(regex string) (d *DFA) {
	var SyntaxTree tree.Ast
	regex = "(("+regex+")$)"
	SyntaxTree.ID = make(map[int]string)
	SyntaxTree.LeafNodes = make(map[string][]int)
	SyntaxTree.Root, SyntaxTree.FollowPos = tree.CreateTree(regex, SyntaxTree.ID, SyntaxTree.LeafNodes, SyntaxTree.FollowPos)
	d = InitDFA(SyntaxTree, SyntaxTree.FollowPos, SyntaxTree.Root.FirstPos)
	Convert(d, SyntaxTree.Root, SyntaxTree.LeafNodes)
	SetReceive(d.InitialState)
	return
}