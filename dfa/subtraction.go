package dfa


func Difference(dfa1 *DFA, dfa2 *DFA) *DFA {
	var stateList, reach = Mul(dfa1, dfa2)
	var NewStateList []*State
	for _, v := range stateList {
		if reach[v.Name] {
			NewStateList = append(NewStateList, v)
		}
	}
	var res = &DFA{InitialState: NewStateList[0]}
	var Exits []string
	var exits1, exits2, list2 = FindExit(dfa1.InitialState), FindExit(dfa2.InitialState), ListOfStates(dfa2)
	var ex []*State
	for _, v := range list2 {
		if !FindInExits(exits2, v.Name) {
			ex = append(ex, v)
		}
	}
	for _, v := range exits1 {
		for _, f := range ex {
			Exits = append(Exits, v.Name+","+f.Name)
		}
	}
	for _, v := range NewStateList {
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