package dfa

func DekartMul(mas1 []*State, mas2 []*State) (res []*State, m map[string]bool) {
	for _, i := range mas1 {
		for _, j := range mas2 {
			var str =  i.Name+","+j.Name
			res = append(res, &State{Name: i.Name+","+j.Name})
			if m == nil {
				m = make(map[string]bool)
			}
			m[str] = false
		}
	}
	return
}

func Mul(dfa1 *DFA, dfa2 *DFA) ([]*State, map[string]bool) {
	var names1 = ListOfStates(dfa1)
	var names2 = ListOfStates(dfa2)
	var stateList, reach = DekartMul(names1, names2)
	var trans = GetTransitions(dfa1.InitialState, dfa2.InitialState)
	reach[stateList[0].Name] = true
	for i:=0; i < len(stateList); i++ {
		var m = make(map[string]map[string]*State)
		var n []string
		var k string = ""
		for j, v := range stateList[i].Name {
			if v != ',' {
				k += string(v)
			} else {
				m[k] = trans[k]
				n = append(n, k)
				k = ""
			}
			if j == len(stateList[i].Name)-1 {
				m[k] = trans[k]
				n = append(n, k)
			}
		}
		var nextName string = ""
		for key, v := range m[n[0]] {
			nextName = v.Name+","
			for z:=1; z < len(n); z++ {
				for kz, vz := range m[n[z]] {
					if key == kz {
						nextName += vz.Name+","
						break
					}
				}
			}
			nextName = nextName[:len(nextName)-1]
			for a:=0; a < len(stateList); a++ {
				if stateList[a].Name == nextName {
					if stateList[i].Dtran == nil {
						stateList[i].Dtran = make(map[string]*State)
					}
					if reach[stateList[a].Name] == false {
						reach[stateList[a].Name] = reach[stateList[i].Name]
					}
					stateList[i].Dtran[key] = stateList[a]
					break
				}
			}
		}
	}
	return stateList, reach
}
