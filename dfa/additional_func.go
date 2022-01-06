package dfa

import (
	"fmt"
	"reflect"
)

func FindSeen(finded *State, mas []*State) bool {
	for _, a := range mas {
		if finded.Name == a.Name &&
			reflect.DeepEqual(finded.Name, a.Name) &&
			reflect.DeepEqual(finded.Dtran, a.Dtran) {
			return true
		}
	}
	return false
}

func Find(dfa *DFA, s string) bool {
	for _, v := range dfa.Alphabet {
		if v == s {
			return true
		}
	}
	return false
}

func Print(start *State) {
	var leftStates = []*State{start}
	var seenStates []*State
	for len(leftStates) != 0 {
		var state = leftStates[len(leftStates)-1]
		leftStates = leftStates[:len(leftStates)-1]
		if !FindSeen(state, seenStates) {
			seenStates = append(seenStates, state)
			fmt.Printf("<%s  %v %v>     ", state.Name, state.StateNumber, state.Receive)
			for k, v := range state.Dtran {
				fmt.Printf("Dtran<%s>=%s    ", k, v.Name)
				leftStates = append(leftStates, v)
			}
			fmt.Printf("\n")
		}
	}
}

func CopyDFA(dfa *DFA) *DFA {
	num=1
	var tmp = &DFA{tree: dfa.tree, FollowPos: dfa.FollowPos, InitialStateNumber: dfa.InitialStateNumber}
	Convert(tmp, dfa.tree.Root, dfa.tree.LeafNodes)
	return tmp
}

func FindExit(st *State) (s []*State) {
	var leftStates = []*State{st}
	var seenStates []*State
	for len(leftStates) != 0 {
		var state = leftStates[len(leftStates)-1]
		leftStates = leftStates[:len(leftStates)-1]
		if !FindSeen(state, seenStates) {
			seenStates = append(seenStates, state)
			for _, v := range state.Dtran {
				leftStates = append(leftStates, v)
			}
			if state.Receive {
				s = append(s, state)
			}
		}
	}
	return s
}

func GetKeysValues(m map[string]*State) (str []string, s []*State) {
	for k, v := range m {
		str = append(str, k)
		s = append(s, v)
	}
	return
}

func FindInExits(exits []*State, name string) bool {
	for _, v := range exits {
		if v.Name == name {
			return true
		}
	}
	return false
}

func FindInExitsString(exits []string, name string) bool {
	for _, v := range exits {
		if v == name {
			return true
		}
	}
	return false
}

func ListOfStates(dfa *DFA) (names []*State)  {
	var leftStates = []*State{dfa.InitialState}
	var seenStates []*State
	for len(leftStates) != 0 {
		var state = leftStates[len(leftStates)-1]
		leftStates = leftStates[:len(leftStates)-1]
		if !FindSeen(state, seenStates) {
			seenStates = append(seenStates, state)
			for _, v := range state.Dtran {
				leftStates = append(leftStates, v)
			}
			names = append(names, state)
		}
	}
	return
}

func GetTransitions(start1 *State, start2 *State) map[string]map[string]*State {
	var res = make(map[string]map[string]*State)
	var leftStates = []*State{start1}
	var seenStates []*State
	for len(leftStates) != 0 {
		var state = leftStates[len(leftStates)-1]
		leftStates = leftStates[:len(leftStates)-1]
		if !FindSeen(state, seenStates) {
			seenStates = append(seenStates, state)
			for _, v := range state.Dtran {
				leftStates = append(leftStates, v)
			}
			res[state.Name] = state.Dtran
		}
	}
	leftStates = []*State{start2}
	seenStates = nil
	if start2 == nil {
		return res
	} else {
		for len(leftStates) != 0 {
			var state = leftStates[len(leftStates)-1]
			leftStates = leftStates[:len(leftStates)-1]
			if !FindSeen(state, seenStates) {
				seenStates = append(seenStates, state)
				for _, v := range state.Dtran {
					leftStates = append(leftStates, v)
				}
				res[state.Name] = state.Dtran
			}
		}
	}
	return res
}


