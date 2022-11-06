package ecs_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vgalaktionov/roguelike-go/ecs"
)

func TestMustGetEntityComponent(t *testing.T) {
	w := ecs.NewWorld()

	e1 := ecs.AddEntity(w, testComp1{})
	result := ecs.MustGetEntityComponent[testComp1](w, e1)

	assert.Equal(t, result, testComp1{}, "should correctly retrieve component")

	assert.Panics(t, assert.PanicTestFunc(func() {
		ecs.MustGetEntityComponent[testComp2](w, e1)
	}), "missing component will be nil")
}

func BenchmarkMustGetEntityComponent(b *testing.B) {
	b.StopTimer()
	w := ecs.NewWorld()
	for i := 0; i < 1000; i++ {
		ecs.AddEntity(w, testComp1{})
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		ecs.MustGetEntityComponent[testComp1](w, ecs.Entity(i%1000))
	}
}

func TestGetEntityComponent(t *testing.T) {
	w := ecs.NewWorld()

	e1 := ecs.AddEntity(w, testComp1{})
	result, err := ecs.GetEntityComponent[testComp1](w, e1)

	assert.Nil(t, err, "should not error")
	assert.Equal(t, result, testComp1{}, "should correctly retrieve component")

	_, err = ecs.GetEntityComponent[testComp2](w, e1)
	assert.Error(t, err, "should error")
}

func BenchmarkGetEntityComponent(b *testing.B) {
	b.StopTimer()
	w := ecs.NewWorld()
	for i := 0; i < 1000; i++ {
		ecs.AddEntity(w, testComp1{})
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		ecs.GetEntityComponent[testComp1](w, ecs.Entity(i%1000))
	}
}

func TestSetEntityComponent(t *testing.T) {
	w := ecs.NewWorld()

	e1 := ecs.AddEntity(w, testComp2{Value: 1})
	ecs.SetEntityComponent(w, testComp2{Value: 2}, e1)

	assert.Equal(t, ecs.MustGetEntityComponent[testComp2](w, e1).Value, 2, "should correctly set component")
}

func BenchmarkSetEntityComponent(b *testing.B) {
	b.StopTimer()
	w := ecs.NewWorld()
	for i := 0; i < 1000; i++ {
		ecs.AddEntity(w, testComp1{})
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		ecs.SetEntityComponent(w, testComp1{}, ecs.Entity(i%1000))
	}
}

func TestRemoveEntityComponent(t *testing.T) {
	w := ecs.NewWorld()

	e1 := ecs.AddEntity(w, testComp1{}, testComp2{})
	ecs.RemoveEntityComponent[testComp1](w, e1)

	assert.False(t, ecs.HasEntityComponent[testComp1](w, e1), "should correctly remove component")
	assert.True(t, ecs.HasEntityComponent[testComp2](w, e1), "should not remove other component")
}

func BenchmarkRemoveEntityComponent(b *testing.B) {
	b.StopTimer()
	w := ecs.NewWorld()
	for i := 0; i < 1000; i++ {
		ecs.AddEntity(w, testComp1{}, testComp2{})
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		ecs.RemoveEntityComponent[testComp1](w, ecs.Entity(i&1000))
	}
}
