package dfa

import "reflect"

var startState *State

func Search(pattern interface{}, str string) string {
	var dfa *DFA
	if reflect.TypeOf(pattern) == reflect.TypeOf("11") {
		dfa = Compile(pattern.(string))
	} else {
		dfa = pattern.(*DFA)
	}
	startState = dfa.InitialState
	return search(dfa.InitialState, str, 0, "")
}

func search(state *State, str string, ind int, copystr string) string {
	//fmt.Println("String ", str)
	var finded bool
	var startInd = -1
	for ind < len(str) {
		finded = false
		var flag = false
		for k, v := range state.Dtran {
			//fmt.Println(v, state)
			if string(str[ind]) == k {
				if startInd == -1 {
					startInd = ind
				}
				copystr += k
				finded = true
				if v.Name == state.Name {
					flag = true
				}
				state = v
				break
			}
		}
		var selfLoop = false
		for _, v := range state.Dtran {
			if state == v {
				selfLoop = true
			}
		}
		ind++
		//fmt.Println(copystr)
		if flag {
			continue
		}
		if state.Receive && !selfLoop {
			return copystr
		}
		if !finded && state != startState {
			startInd = -1
			for _, v := range state.Dtran {
				return search(v, str, ind, copystr)
			}
		}

	}
	if copystr == str {
		return copystr
	}
	return ""
}
