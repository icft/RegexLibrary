package dfa


func GetPredecessors(dfa *DFA, name string) (pred []*State) {
	var leftStates = []*State{dfa.InitialState}
	var seenStates []*State
	for len(leftStates) != 0 {
		var state = leftStates[len(leftStates)-1]
		leftStates = leftStates[:len(leftStates)-1]
		if !FindSeen(state, seenStates) {
			seenStates = append(seenStates, state)
			for _, v := range state.Dtran {
				if v.Name == name {
					pred = append(pred, state)
				}
				leftStates = append(leftStates, v)
			}
		}
	}
	return
}

func GetSuccessors(dfa *DFA, name string) (s []*State) {
	var leftStates = []*State{dfa.InitialState}
	var seenStates []*State
	for len(leftStates) != 0 {
		var state = leftStates[len(leftStates)-1]
		leftStates = leftStates[:len(leftStates)-1]
		if !FindSeen(state, seenStates) {
			seenStates = append(seenStates, state)
			for _, v := range state.Dtran {
				if state.Name == name && v.Name != name {
					s = append(s, v)
				}
				leftStates = append(leftStates, v)
			}
		}
	}
	return
}

func GetTransition(state *State, next string) string {
	for k, v := range state.Dtran {
		if v.Name == next {
			return k
		}
	}
	return ""
}

func GetState(dfa *DFA, name string) *State {
	var leftStates = []*State{dfa.InitialState}
	var seenStates []*State
	for len(leftStates) != 0 {
		var state = leftStates[len(leftStates)-1]
		leftStates = leftStates[:len(leftStates)-1]
		if !FindSeen(state, seenStates) {
			seenStates = append(seenStates, state)
			if state.Name == name {
				return state
			}
			for _, v := range state.Dtran {
				leftStates = append(leftStates, v)
			}
		}
	}
	return nil
}

func CheckSelfLoop(dfa *DFA, name string) bool {
	var s = GetState(dfa, name)
	for _, v := range s.Dtran {
		if v.Name == name {
			return true
		}
	}
	return false
}

type Pair struct {
	str string
	state *State
}

func CreateTransitions(tmp *DFA, k *State, flag bool) *State {
	if k.Name == tmp.InitialState.Name {
		return k
	}
	var p = GetPredecessors(tmp, k.Name)[0]
	var s = GetSuccessors(tmp, k.Name)[0]
	var loop = CheckSelfLoop(tmp, k.Name)
	/*if k.Name == "T" {
		fmt.Println(p, s, loop)
	}*/
	var pTs, sTp, pLoop, sLoop string = "", "", "", ""
	if !flag {
		if GetTransition(p, k.Name) == "+" || GetTransition(p, k.Name) == "|" ||
			GetTransition(p, k.Name) == "{" || GetTransition(p, k.Name) == "}" ||
			GetTransition(p, k.Name) == "\\" || GetTransition(p, k.Name) == "." ||
			GetTransition(p, k.Name) == "#" {
			pTs = "#"+GetTransition(p, k.Name)
		} else  {
			pTs = GetTransition(p, k.Name)
		}
	} else {
		if GetTransition(p, s.Name) == "+" || GetTransition(p, s.Name) == "|" ||
			GetTransition(p, s.Name) == "{" || GetTransition(p, s.Name) == "}" ||
			GetTransition(p, s.Name) == "\\" || GetTransition(p, s.Name) == "." ||
			GetTransition(p, s.Name) == "#" {
			pTs += "#"+GetTransition(p, s.Name)
		} else  {
			pTs += GetTransition(p, s.Name)
		}
		if GetTransition(p, k.Name) == "+" || GetTransition(p, k.Name) == "|" ||
			GetTransition(p, k.Name) == "{" || GetTransition(p, k.Name) == "}" ||
			GetTransition(p, k.Name) == "\\" || GetTransition(p, k.Name) == "." ||
			GetTransition(p, k.Name) == "#" {
			pTs += "#"+GetTransition(p, k.Name)
		} else  {
			pTs += GetTransition(p, k.Name)
		}
	}
	if loop {
		if GetTransition(k, k.Name) == "+" || GetTransition(k, k.Name) == "|" ||
			GetTransition(k, k.Name) == "{" || GetTransition(k, k.Name) == "}" ||
			GetTransition(k, k.Name) == "\\" || GetTransition(k, k.Name) == "." ||
			GetTransition(k, k.Name) == "#" {
			pTs += "#"+GetTransition(k, k.Name) + "*"
		} else  {
			pTs += GetTransition(k, k.Name) + "*"
		}
	}
	if GetTransition(k, s.Name) == "+" || GetTransition(k, s.Name) == "|" ||
		GetTransition(k, s.Name) == "{" || GetTransition(k, s.Name) == "}" ||
		GetTransition(k, s.Name) == "\\" || GetTransition(k, s.Name) == "." ||
		GetTransition(k, s.Name) == "#" {
		pTs += "#"+GetTransition(k, s.Name)
	} else  {
		pTs += GetTransition(k, s.Name)
	}
	if GetTransition(s, p.Name) == "+" || GetTransition(s, p.Name) == "|" ||
		GetTransition(s, p.Name) == "{" || GetTransition(s, p.Name) == "}" ||
		GetTransition(s, p.Name) == "\\" || GetTransition(s, p.Name) == "." ||
		GetTransition(s, p.Name) == "#" {
		sTp = "#" + GetTransition(s, p.Name)
	} else  {
		sTp = GetTransition(s, p.Name)
	}
	if GetTransition(s, k.Name) == "+" || GetTransition(s, k.Name) == "|" ||
		GetTransition(s, k.Name) == "{" || GetTransition(s, k.Name) == "}" ||
		GetTransition(s, k.Name) == "\\" || GetTransition(s, k.Name) == "." ||
		GetTransition(s, k.Name) == "#" {
		sTp += "#" + GetTransition(s, k.Name)
	} else  {
		sTp += GetTransition(s, k.Name)
	}
	if sTp != "" {
		if loop {
			if GetTransition(k, k.Name) == "+" || GetTransition(k, k.Name) == "|" ||
				GetTransition(k, k.Name) == "{" || GetTransition(k, k.Name) == "}" ||
				GetTransition(k, k.Name) == "\\" || GetTransition(k, k.Name) == "." ||
				GetTransition(k, k.Name) == "#" {
				sTp += "#" + GetTransition(k, k.Name) + "*"
			} else  {
				sTp += GetTransition(k, k.Name) + "*"
			}
		}
		if GetTransition(k, p.Name) == "+" || GetTransition(k, p.Name) == "|" ||
			GetTransition(k, p.Name) == "{" || GetTransition(k, p.Name) == "}" ||
			GetTransition(k, p.Name) == "\\" || GetTransition(k, p.Name) == "." ||
			GetTransition(k, p.Name) == "#" {
			sTp += "#" + GetTransition(k, p.Name)
		} else  {
			sTp += GetTransition(k, p.Name)
		}
	}
	var tr1_1, tr1_2 string
	if GetTransition(k, p.Name) == "+" || GetTransition(k, p.Name) == "|" ||
		GetTransition(k, p.Name) == "{" || GetTransition(k, p.Name) == "}" ||
		GetTransition(k, p.Name) == "\\" || GetTransition(k, p.Name) == "." ||
		GetTransition(k, p.Name) == "#" {
		tr1_1 = "#"+GetTransition(k, p.Name)
	} else  {
		tr1_1 = GetTransition(k, p.Name)
	}
	if GetTransition(p, k.Name) == "+" || GetTransition(p, k.Name) == "|" ||
		GetTransition(p, k.Name) == "{" || GetTransition(p, k.Name) == "}" ||
		GetTransition(p, k.Name) == "\\" || GetTransition(p, k.Name) == "." ||
		GetTransition(p, k.Name) == "#" {
		tr1_2 = "#"+GetTransition(p, k.Name)
	} else  {
		tr1_2 = GetTransition(p, k.Name)
	}
	if tr1_1 != "" && tr1_2 != "" {
		pLoop += tr1_2
		if loop {
			if GetTransition(k, k.Name) == "+" || GetTransition(k, k.Name) == "|" ||
				GetTransition(k, k.Name) == "{" || GetTransition(k, k.Name) == "}" ||
				GetTransition(k, k.Name) == "\\" || GetTransition(k, k.Name) == "." ||
				GetTransition(k, k.Name) == "#" {
				pLoop += "#" + GetTransition(k, k.Name) + "*"
			} else  {
				pLoop += GetTransition(k, k.Name) + "*"
			}
		}
		pLoop += tr1_1
	}
	var tr2_1, tr2_2 string
	if GetTransition(k, s.Name) == "+" || GetTransition(k, s.Name) == "|" ||
		GetTransition(k, s.Name) == "{" || GetTransition(k, s.Name) == "}" ||
		GetTransition(k, s.Name) == "\\" || GetTransition(k, s.Name) == "." ||
		GetTransition(k, s.Name) == "#" {
		tr2_1 = "#"+GetTransition(k, p.Name)
	} else  {
		tr2_1 = GetTransition(k, p.Name)
	}
	if GetTransition(s, k.Name) == "+" || GetTransition(s, k.Name) == "|" ||
		GetTransition(s, k.Name) == "{" || GetTransition(s, k.Name) == "}" ||
		GetTransition(s, k.Name) == "\\" || GetTransition(s, k.Name) == "." ||
		GetTransition(s, k.Name) == "#" {
		tr2_2 = "#"+GetTransition(p, k.Name)
	} else  {
		tr1_2 = GetTransition(p, k.Name)
	}
	if tr2_1 != "" && tr2_2 != "" {
		sLoop += tr2_2
		if loop {
			if GetTransition(k, k.Name) == "+" || GetTransition(k, k.Name) == "|" ||
				GetTransition(k, k.Name) == "{" || GetTransition(k, k.Name) == "}" ||
				GetTransition(k, k.Name) == "\\" || GetTransition(k, k.Name) == "." ||
				GetTransition(k, k.Name) == "#" {
				sLoop += "#" + GetTransition(k, k.Name) + "*"
			} else  {
				sLoop += GetTransition(k, k.Name) + "*"
			}
		}
		sLoop += tr2_1
	}
	delete(p.Dtran, GetTransition(p, k.Name))
	//fmt.Printf("%s %s %s %s\n", pTs, sTp, pLoop, sLoop)
	if pTs != "" {
		p.Dtran[pTs] = s
	}
	if sTp != "" {
		s.Dtran[sTp] = p
	}
	if pLoop != "" {
		p.Dtran[pLoop] = p
	}
	if sLoop != "" {
		s.Dtran[sLoop] = s
	}
	return s
}

var stack []*State

func CreateRE(tmp *DFA) (regex string) {
	var start = tmp.InitialState
	for true {
		if start.Receive {
			break
		}
		if len(start.Dtran) == 2 && CheckSelfLoop(tmp, start.Name) {
			start = CreateTransitions(tmp, start, true)
		} else if len(start.Dtran) > 1  {
			for _, v := range start.Dtran {
				stack = append(stack, start)
				CreateSubOr(tmp, v)
			}
			for _, v := range start.Dtran {
				stack = append(stack, start)
				CreateSubOr(tmp, v)
			}
			var newMap = make(map[string]*State)
			var newKey = "("
			var _, end = GetKeysValues(start.Dtran)
			for k, _ := range start.Dtran {
				if k == "+" || k == "|" || k== "{" || k == "}" ||
					k == "\\" || k == "." || k == "#"{
					newKey += "(#" + k + ")|"
				} else {
					newKey += "(" + k + ")|"
				}
			}
			newKey = newKey[:len(newKey)-1]
			newKey += ")"
			newMap[newKey] = end[0]
			start.Dtran = newMap
			start = CreateTransitions(tmp, start, true)
		} else if len(start.Dtran) == 1 && start != tmp.InitialState {
			start = CreateTransitions(tmp, start, true)
		} else if len(start.Dtran) == 1 && start == tmp.InitialState {
			_, v := GetKeysValues(start.Dtran)
			start = v[0]
		}
	}
	k, _ := GetKeysValues(tmp.InitialState.Dtran)
	return k[0]
}

func CreateSubOr(tmp *DFA, state *State) {
	if len(GetPredecessors(tmp, state.Name)) > 1 {
		stack = stack[:len(stack)-1]
		return
	}
	if state.Receive {
		return
	}
	for len(stack) != 0 {
		//fmt.Println(state)
		if len(GetPredecessors(tmp, state.Name)) > 1 {
			stack = stack[:len(stack)-1]
			return
		}
		if len(state.Dtran) > 1 {
			for _, v := range state.Dtran {
				stack = append(stack, state)
				CreateSubOr(tmp, v)
			}
			var newMap = make(map[string]*State)
			var newKey = "("
			var _, end = GetKeysValues(state.Dtran)
			for k, _ := range state.Dtran {
				if k == "+" || k == "|" || k== "{" || k == "}" ||
					k == "\\" || k == "." || k == "#"{
					newKey += "(#" + k + ")|"
				} else {
					newKey += "(" + k + ")|"
				}
			}
			newKey = newKey[:len(newKey)-1]
			newKey += ")"
			newMap[newKey] = end[0]
			state.Dtran = newMap
			state = CreateTransitions(tmp, state, true)
		} else if len(state.Dtran) == 1 && state != tmp.InitialState {
			state = CreateTransitions(tmp, state, false)
		}
	}
	return
}