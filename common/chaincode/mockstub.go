package main

import (
	"fmt"
	"strconv"
)

type MockStub struct {
	State map[string]int
}

func NewMockStub() *MockStub {
	return &MockStub{State: make(map[string]int)}
}

func (s *MockStub) Put(key string, value int) {
	s.State[key] = value
}

func (s *MockStub) Get(key string) int {
	return s.State[key]
}

func (s *MockStub) PrintState() {
	for k, v := range s.State {
		fmt.Printf("%s = %d\n", k, v)
	}
}

func atoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
