package glru

import (
	"testing"
)

func TestNEWGLRU(t *testing.T) {
	_, err := NEWGLRU(10)
	if err != nil {
		t.Fatal(err)
		return
	}
}

func TestSet(t *testing.T) {
	lru, err := NEWGLRU(1)
	if err != nil {
		t.Fatal(err)
		return
	}
	evicted := lru.Set("a", 1)
	if !evicted {
		t.Logf("return as expected")
	}
	evicted = lru.Set("b", 2)
	if evicted {
		t.Logf("return as expected")
	}
}

func TestGet(t *testing.T) {
	lru, err := NEWGLRU(1)
	if err != nil {
		t.Fatal(err)
		return
	}
	lru.Set("a", 1)
	v, ok := lru.Get("a")
	if ok {
		t.Logf("value: %v\n", v)
	}

	v, ok = lru.Get("b")
	if !ok {
		t.Logf("not set value of b\n")
	}
}

func TestContains(t *testing.T) {
	lru, err := NEWGLRU(1)
	if err != nil {
		t.Fatal(err)
		return
	}
	lru.Set("a", 1)
	exist := lru.Contains("a")
	if exist {
		t.Logf("exist")
	}

	exist = lru.Contains("b")
	if !exist {
		t.Logf("not exist")
	}
}

func TestPeek(t *testing.T) {
	lru, err := NEWGLRU(1)
	if err != nil {
		t.Fatal(err)
		return
	}
	lru.Set("a", 1)
	v, ok := lru.Peek("a")
	if ok {
		t.Logf("peek value of a is: %v", v)
	}
}

func TestPurge(t *testing.T) {
	lru, err := NEWGLRU(1)
	if err != nil {
		t.Fatal(err)
		return
	}
	lru.Set("a", 1)
	t.Log(lru.Keys())
	lru.Purge()
	t.Log(lru.Keys())
}

func TestKeys(t *testing.T) {
	lru, err := NEWGLRU(1)
	if err != nil {
		t.Fatal(err)
		return
	}
	lru.Set("a", 1)
	t.Log(lru.Keys())
}

func TestLen(t *testing.T) {
	lru, err := NEWGLRU(1)
	if err != nil {
		t.Fatal(err)
		return
	}
	lru.Set("a", 1)
	t.Log(lru.Len())
}
