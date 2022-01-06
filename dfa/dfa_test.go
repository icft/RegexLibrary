package dfa

import (
"testing"
)


func Check(tmp *DFA, str string) bool {
	var state = tmp.InitialState
	for _, s := range str {
		var n *State
		for k, v := range state.Dtran {
			if k == string(s) {
				n = v
				break
			}
		}
		if n == nil {
			return false
		}
		state = n
	}
	if state.Receive {
		return true
	} else {
		return false
	}
}

func Equals(d1, d2 *DFA) bool {
	if len(d1.Alphabet) != len(d2.Alphabet) {
		return false
	}
	for i:=0; i<len(d1.Alphabet);i++ {
		if d1.Alphabet[i] != d2.Alphabet[i] {
			return false
		}
	}
	return eq(d1.InitialState, d2.InitialState)
}

func eq(state1, state2 *State) bool {
	if state1.Name != state2.Name {
		return false
	}
	if state1.Name == state2.Name && state1.Receive == state2.Receive {
		if len(state1.Dtran) != len(state2.Dtran) {
			return false
		}
		var k1, v1 = GetKeysValues(state1.Dtran)
		var k2, v2 = GetKeysValues(state2.Dtran)
		for i:=0; i<len(state1.Dtran);i++ {
			if k1[i] != k2[i] || v1[i].Name != v2[i].Name {
				return false
			} else if state1 == v1[i] && state2 == v2[i] && len(state1.Dtran) == 1 {
				return true
			} else {
				eq(v1[i], v2[i])
			}
		}
	}
	return true
}

func TestCompile(t *testing.T) {
	var d = Compile("a+")
	if !Check(d, "aaaa") {
		t.Error("Fail")
	}
	d = Compile("ab+")
	if !Check(d, "ab") {
		t.Error("Fail")
	}
	if !Check(d, "abbbbbb") {
		t.Error("Fail")
	}
	if Check(d, "a") {
		t.Error("Fail")
	}
	d = Compile("a|b")
	if !Check(d, "a") {
		t.Error("Fail")
	}
	if !Check(d, "b") {
		t.Error("Fail")
	}
	d = Compile("v{2,3}d")
	if !Check(d, "vvd") {
		t.Error("Fail")
	}
	if !Check(d, "vvvd") {
		t.Error("Fail")
	}
	if Check(d, "vvvvvvd") {
		t.Error("Fail")
	}
	d = Compile("a#+r#|")
	if !Check(d, "a+r|") {
		t.Error("Fail")
	}
	d = Compile("v{2,}d")
	if Check(d, "vd") {
		t.Error("Fail")
	}
	if !Check(d, "vvvvvvvvvvvvvvvvvvvvd") {
		t.Error("Fail")
	}
	d = Compile("a(5:as+)d")
	if !Check(d, "aassssssd") {
		t.Error("Fail")
	}
	if Check(d, "aaasd") {
		t.Error("Fail")
	}
}

func TestSearch(t *testing.T) {
	if Search("((mephi|mfti)$)", "aaamephiaaa") != "mephi" {
		t.Error("Fail")
	}
	if Search("((mep{3,}hi)$)", "mepphi") == "mepphi" {
		t.Error("Fail")
	}
	if Search("((mep{3,}hi)$)", "meppphi") != "meppphi" {
		t.Error("Fail")
	}
	if Search("((a+)$)","aaaa") != "aaaa" {
		t.Error("Fail")
	}
	if Search("((a#+)$)", "da+1") != "a+" {
		t.Error("Fail")
	}
}

func TestCreateRE(t *testing.T) {
	var d = Compile("a+cbs")
	if CreateRE(d) != "aa*cbs" {
		t.Error("Fail")
	}
	d = Compile("a((a|b)|h)c")
	if  CreateRE(d) != "a((ac)|(bc)|(hc))" && CreateRE(d) != "a((ac)|(hc)|(bc))" &&
		CreateRE(d) != "a((bc)|(ac)|(hc))" && CreateRE(d) != "a((bc)|(hc)|(ac))" &&
		CreateRE(d) != "a((hc)|(ac)|(bc))" && CreateRE(d) != "a((hc)|(bc)|(ac))" {
		t.Error("Fail")
	}
	d = Compile("ad{3,}v")
	if CreateRE(d) != "adddd*v" {
		t.Error("Fail")
	}
	d = Compile("ad{1,2}v")
	if CreateRE(d) != "ad((dv)|(v))" && CreateRE(d) == "ad((v)|(dv))" {
		t.Error("Fail")
	}
	d = Compile("a#|b")
	if CreateRE(d) != "a#|b" {
		t.Error("Fail")
	}
	d = Compile("(mephi|mf)((ti)###+#+#+ar+my)")
	if CreateRE(d) != "m((ft)|(ephit))i###+#+#+arr*my" &&
		CreateRE(d) != "m((ephi)|(f))ti###+#+#+arr*my" {
		t.Error("Fail")
	}
}

func TestMinimization(t *testing.T) {
	/*Нету*/
}

func TestIntersection(t *testing.T) {
	var d1, d2 = Compile("ab"), Compile("a")
	var dfa *DFA = &DFA{Alphabet: []string{"a"}}
	var s1, s2 = &State{Name: "A,D"}, &State{Name: "B,E"}
	dfa.InitialState = s1
	s1.Dtran = make(map[string]*State)
	s1.Dtran["a"] = s2
	s2.Receive = false
	/*	if !Equals(Intersection(d1, d2), dfa) {
		t.Error("Fail")
	}*/
	// ab+f+ abf+
	d1, d2 = Compile("ab+f"), Compile("abf")
	dfa = &DFA{Alphabet: []string{"a","b","f"}}
	s1, s2 = &State{Name: "F,J"}, &State{Name: "G,K"}
	var s3, s4 = &State{Name: "H,L"}, &State{Name: "I,M"}
	dfa.InitialState = s1
	s1.Dtran = make(map[string]*State)
	s2.Dtran = make(map[string]*State)
	s3.Dtran = make(map[string]*State)
	s1.Dtran["a"] = s2
	s2.Dtran["b"] = s3
	s3.Dtran["f"] = s4
	s2.Receive = false
	s3.Receive = false
	s4.Receive = false
	/*	if !Equals(Intersection(d1, d2), dfa) {
		t.Error("Fail")
	}*/
	d1, d2 = Compile("ab+f+"), Compile("abf+")
	s1, s2 = &State{Name: "N,R"}, &State{Name: "O,S"}
	s3, s4 = &State{Name: "P,T"}, &State{Name: "Q,U"}
	dfa.InitialState = s1
	s1.Dtran = make(map[string]*State)
	s2.Dtran = make(map[string]*State)
	s3.Dtran = make(map[string]*State)
	s4.Dtran = make(map[string]*State)
	s1.Dtran["a"] = s2
	s2.Dtran["b"] = s3
	s3.Dtran["f"] = s4
	s4.Dtran["f"] = s4
	s2.Receive = false
	s3.Receive = false
	s4.Receive = false
	Print(Intersection(d1, d2).InitialState)
	/*if !Equals(Intersection(d1, d2), dfa) {
		t.Error("Fail")
	}*/
}

func TestDifference(t *testing.T) {
	var d1, d2 = Compile("ab"), Compile("a")
	var dfa *DFA = &DFA{Alphabet: []string{"a"}}
	var s1, s2 = &State{Name: "A,D"}, &State{Name: "B,E"}
	dfa.InitialState = s1
	s1.Dtran = make(map[string]*State)
	s1.Dtran["a"] = s2
	s2.Receive = false
	if !Equals(Difference(d1, d2), dfa) {
		t.Error("Fail")
	}
	// ab+f+ abf+
	d1, d2 = Compile("ab+f"), Compile("abf")
	dfa = &DFA{Alphabet: []string{"a","b","f"}}
	s1, s2 = &State{Name: "F,J"}, &State{Name: "G,K"}
	var s3, s4 = &State{Name: "H,L"}, &State{Name: "I,M"}
	dfa.InitialState = s1
	s1.Dtran = make(map[string]*State)
	s2.Dtran = make(map[string]*State)
	s3.Dtran = make(map[string]*State)
	s1.Dtran["a"] = s2
	s2.Dtran["b"] = s3
	s3.Dtran["f"] = s4
	s2.Receive = false
	s3.Receive = false
	s4.Receive = true
	if !Equals(Difference(d1, d2), dfa) {
		t.Error("Fail")
	}
	d1, d2 = Compile("ab+f+"), Compile("abf+")
	s1, s2 = &State{Name: "N,R"}, &State{Name: "O,S"}
	s3, s4 = &State{Name: "P,T"}, &State{Name: "Q,U"}
	dfa.InitialState = s1
	s1.Dtran = make(map[string]*State)
	s2.Dtran = make(map[string]*State)
	s3.Dtran = make(map[string]*State)
	s4.Dtran = make(map[string]*State)
	s1.Dtran["a"] = s2
	s2.Dtran["b"] = s3
	s3.Dtran["f"] = s4
	s4.Dtran["f"] = s4
	s2.Receive = false
	s3.Receive = false
	s4.Receive = false
	if !Equals(Difference(d1, d2), dfa) {
		t.Error("Fail")
	}
}
