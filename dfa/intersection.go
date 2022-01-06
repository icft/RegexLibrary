package dfa

import "fmt"

func Intersection(dfa1 *DFA, dfa2 *DFA) *DFA {
	var stateList, reach = Mul(dfa1, dfa2)
	var NewStateList []*State
	for _, v := range stateList {
		if reach[v.Name] {
			NewStateList = append(NewStateList, v)
		}
	}
	var res = &DFA{InitialState: NewStateList[0]}
	var Exits []string
	for _, v := range FindExit(dfa1.InitialState) {
		for _, f := range FindExit(dfa2.InitialState) {
			Exits = append(Exits, v.Name+","+f.Name)
		}
	}
	for _, v := range NewStateList {
		fmt.Println(FindInExitsString(Exits, v.Name))
		if FindInExitsString(Exits, v.Name) && v.Name != NewStateList[0].Name && reach[v.Name] == true {
			v.Receive = true
		}
	}
	var leftStates = []*State{res.InitialState}
	var seenStates []*State
	for len(leftStates) != 0 {
		var state = leftStates[len(leftStates)-1]
		leftStates = leftStates[:len(leftStates)-1]
		if !FindSeen(state, seenStates) {
			seenStates = append(seenStates, state)
			for k, v := range state.Dtran {
				leftStates = append(leftStates, v)
				if !Find(res, k) {
					res.Alphabet = append(res.Alphabet, k)
				}
			}
		}
	}
	return res
}