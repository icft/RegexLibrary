package dfa

import "reflect"

func SameStates(s1 *State, s2 *State) bool {
	var tr = false
	for k, v := range s1.Dtran {
		for k1,v1 := range s2.Dtran {
			if k == k1 {
				tr = true
				if v != v1 {
					return false
				}
			}
		}
		if !tr {
			return false
		}
	}
	return true
}


func MakeTransition(state *State, tr string) *State {
	for k, v := range state.Dtran {
		if k == tr {
			return v
		}
	}
	return nil
}

func Equal(mas1 []*State, mas2 []*State) bool {
	if len(mas1) != len(mas2) {
		return false
	}
	for i:=0; i < len(mas1); i++ {
		if mas1[i].Name == mas2[i].Name && reflect.DeepEqual(mas1[i].Dtran, mas2[i].Dtran) {
			continue
		} else {
			return false
		}
	}
	return true
}

func Minimization(dfa *DFA) {
	var states = ListOfStates(dfa)
	var exits []*State
	for _, v := range FindExit(dfa.InitialState) {
		exits = append(exits, v)
	}
	for i, v := range states {
		for _, val := range exits {
			if v == val {
				copy(states[i:], states[i+1:])
				states[len(states)-1] = nil
				states = states[:len(states)-1]
			}
		}
	}
	var split, new_split [][]*State
	split = append(split, states)
	split = append(split, exits)
	var table = make(map[*State][]*State)
	for _, i := range split {
		for _, j := range i {
			table[j] = i
		}
	}
	for len(new_split) != len(split) {
		var new_table = make(map[*State][]*State)
		if len(new_split) != 0 {
			split = new_split
			new_split = nil
		}
		for _, v := range split {
			if len(v) == 1 {
				new_split = append(new_split, v)
				if _, err := new_table[v[0]]; !err {
					new_table[v[0]] = v
				}
				continue
			}
			for j, a := range v {
				if _, err := new_table[a]; !err {
					var tmp []*State
					tmp = append(tmp, a)
					new_table[a] = tmp
					new_split = append(new_split, tmp)
				}
				for k := j + 1; k < len(v); k++ {
					var s = v[k]
					var add = true
					if _, err := new_table[s]; !err {
						for _, m := range dfa.Alphabet {
							var r1 = MakeTransition(a, m)
							var r2 = MakeTransition(s, m)
							if r1 != nil && r2 != nil {
								if Equal(new_table[r1], new_table[r2]) {
									add = false
									break
								}
							} else if !(r1 == nil && r2 == nil) {
								add = false
								break
							}
						}
					}
					if add {
						if _, err := new_table[s]; !err {
							new_table[a] = append(new_table[a], s)
							new_table[s] = new_table[a]
						}
					}
				}
			}
		}
		table = new_table
		new_table = nil
	}
	split = new_split
	for _, v := range split {
		if len(v) > 1 {
			var m =  v[0]
			for i:=1; i < len(v); i++ {
				var predecessors = GetPredecessors(dfa, v[i].Name)
				var successors = GetSuccessors(dfa, v[i].Name)
				for _, state := range predecessors {
					var k string
					for key, val := range state.Dtran {
						if val == v[i] {
							delete(state.Dtran, key)
							k = key
							break
						}
					}
					state.Dtran[k] = m
				}
				for _, state := range successors {
					var k string
					for key, val := range m.Dtran {
						if val == state {
							delete(m.Dtran, k)
							k = key
							break
						}
					}
					m.Dtran[k] = state
				}
			}
		}
	}
}