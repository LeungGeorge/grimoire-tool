package snowflake

import "testing"

func TestNewGUID(t *testing.T) {
	t.Log("start test NewGUID:")

	g, err := NewGUID(1)
	if err != nil {
		t.Fatal(err)
		return
	}

	tmp := uint64(1)
	for i := 0; i < 10; i++ {
		n, _ := g.NextID()
		g.GetIncreaseID(&tmp)
		t.Log(n)
	}
}

func BenchmarkNewGUIDV2(b *testing.B) {

	g, err := NewGUID(1)
	if err != nil {
		b.Fatal(err)
		return
	}

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			g.NextID()
		}
	})
}

func BenchmarkNewGUID(b *testing.B) {

	g, err := NewGUID(1)
	if err != nil {
		b.Fatal(err)
		return
	}

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			g.NextID()
		}
	})
}
