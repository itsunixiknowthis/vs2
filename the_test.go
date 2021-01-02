package vs2

import "testing"

func TestA(t *testing.T) {
    emb := Emb{}
    if "T1.F" != emb.F() { t.Errorf("!") }
    if "T2.G" != emb.G() { t.Errorf("?") }
}

