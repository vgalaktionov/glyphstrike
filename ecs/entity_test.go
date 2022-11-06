package ecs_test

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vgalaktionov/roguelike-go/ecs"
)

func TestAddEntity(t *testing.T) {
	w := ecs.NewWorld()

	e1 := ecs.AddEntity(w)

	assert.Equal(t, e1, ecs.Entity(0), "entity IDs start from 0")

	e2 := ecs.AddEntity(w)

	assert.Equal(t, e2, ecs.Entity(1), "entity IDs are sequential")
}

func BenchmarkAddEntity(b *testing.B) {
	w := ecs.NewWorld()
	for i := 0; i < 1_000_000; i++ {
		ecs.AddEntity(w, testComp1{}, testComp2{})
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ecs.AddEntity(w)
	}
}

func TestHasEntity(t *testing.T) {
	w := ecs.NewWorld()

	result := ecs.HasEntity(w, ecs.Entity(0))

	assert.False(t, result, "should correctly report entity missing")

	ecs.AddEntity(w)

	result = ecs.HasEntity(w, ecs.Entity(0))

	assert.True(t, result, "should correctly report entity exists")
}

func BenchmarkHasEntity(b *testing.B) {
	w := ecs.NewWorld()
	for i := 0; i < 1_000_000; i++ {
		if i%2 == 1 {
			ecs.AddEntity(w)
		}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {

		ecs.HasEntity(w, ecs.Entity(i%1_000_000))
	}
}

type testComp1 struct{}

func (testComp1) CID() ecs.CID {
	return ecs.CID(0)
}

type testComp2 struct {
	Value int
}

func (testComp2) CID() ecs.CID {
	return ecs.CID(1)
}

func TestAddEntityComponent(t *testing.T) {
	w := ecs.NewWorld()

	e1 := ecs.AddEntity(w, testComp1{}, testComp2{})

	assert.Equal(t, e1, ecs.Entity(0), "entity IDs start from 0")
	assert.True(t, ecs.HasEntityComponent[testComp1](w, e1), "should add component correctly")
	assert.True(t, ecs.HasEntityComponent[testComp2](w, e1), "should add another component correctly")

}

func BenchmarkAddEntityComponent(b *testing.B) {
	w := ecs.NewWorld()
	for i := 0; i < 1_000_000; i++ {
		ecs.AddEntity(w, testComp1{}, testComp2{})
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ecs.AddEntity(w, testComp1{}, testComp2{})
	}
}

func BenchmarkHasEntityComponent(b *testing.B) {
	w := ecs.NewWorld()
	for i := 0; i < 1_000_000; i++ {
		if i%2 == 1 {
			ecs.AddEntity(w)
		} else {
			ecs.AddEntity(w, testComp1{})
		}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ecs.HasEntityComponent[testComp1](w, ecs.Entity(i%1_000_000))
	}
}

func TestRemoveEntity(t *testing.T) {
	w := ecs.NewWorld()

	e1 := ecs.AddEntity(w, testComp1{}, testComp2{})

	ecs.RemoveEntity(w, e1)

	assert.False(t, ecs.HasEntity(w, e1), "should remove entity correctly")
	assert.False(t, ecs.HasEntityComponent[testComp1](w, e1), "should remove component correctly")
	assert.False(t, ecs.HasEntityComponent[testComp2](w, e1), "should remove another component correctly")
}

func BenchmarkRemoveEntity(b *testing.B) {
	w := ecs.NewWorld()
	for i := 0; i < 1_000_000; i++ {
		ecs.AddEntity(w, testComp1{}, testComp2{})
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ecs.RemoveEntity(w, ecs.Entity(i%1_000_000))
	}
}

func TestQueryEntitiesIter(t *testing.T) {
	w := ecs.NewWorld()

	ecs.AddEntity(w, testComp1{}, testComp2{})
	ecs.AddEntity(w)
	ecs.AddEntity(w, testComp1{})
	ecs.AddEntity(w, testComp2{})
	ecs.AddEntity(w, testComp1{}, testComp2{})

	results := []int{}
	for _, r := range ecs.QueryEntitiesIter(w, testComp1{}, testComp2{}) {
		results = append(results, int(r))
	}
	sort.Ints(results)

	assert.Equal(t, results, []int{0, 4}, "should query multiple entities correctly")
}

func BenchmarkQueryEntitiesIter(b *testing.B) {
	w := ecs.NewWorld()
	for i := 0; i < 1_000_000; i++ {
		if i%2 == 1 {
			ecs.AddEntity(w, testComp1{}, testComp2{})
		} else {
			ecs.AddEntity(w, testComp1{})
		}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for range ecs.QueryEntitiesIter(w, testComp1{}, testComp2{}) {
			continue
		}
	}
}

func TestQueryEntitiesSingle(t *testing.T) {
	w := ecs.NewWorld()

	ecs.AddEntity(w)
	e2 := ecs.AddEntity(w, testComp1{})

	result, err := ecs.QueryEntitiesSingle(w, testComp1{})
	assert.Nil(t, err, "should not error on existing component entity")
	assert.Equal(t, result, e2, "should find existing component entity")

	_, err = ecs.QueryEntitiesSingle(w, testComp1{}, testComp2{})
	assert.Error(t, err, "should error on missing component entity")
}

func BenchmarkQueryEntitiesSingle(b *testing.B) {
	w := ecs.NewWorld()
	for i := 0; i < 1_000_000; i++ {
		ecs.AddEntity(w, testComp1{})
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if i%2 == 1 {
			ecs.QueryEntitiesSingle(w, testComp1{})
		} else {
			ecs.QueryEntitiesSingle(w, testComp1{}, testComp2{})
		}
	}
}
