package main

import (
	"errors"
	"fmt"
)

// SetOnSlice - множество на базе слайса
type SetOnSlice[T comparable] struct {
	items []T
}

// NewSetOnSlice - создает множество на базе слайса
func NewSetOnSlice[T comparable]() *SetOnSlice[T] {
	return &SetOnSlice[T]{
		items: []T{},
	}
}

// Print - возвращает строковое представление множества
func (s *SetOnSlice[T]) Print() string {
	result := ""
	for _, item := range s.items {
		result += fmt.Sprintf("%d ", item)
	}
	return result
}

// Empty - проверяет, пустое ли множество
func (s *SetOnSlice[T]) Empty() bool {
	return len(s.items) == 0
}

// Add - добавляет значение в множества
func (s *SetOnSlice[T]) Add(item T) bool {
	if s.contains(item) {
		return false
	}
	s.items = append(s.items, item)
	return true
}

func (s *SetOnSlice[T]) contains(item T) bool {
	for _, val := range s.items {
		if val == item {
			return true
		}
	}
	return false
}

// Remove - удаляет значение из множества
func (s *SetOnSlice[T]) Remove(item T) bool {
	for i, val := range s.items {
		if val == item {
			s.items = append(s.items[:i], s.items[i+1:]...)
			return true
		}
	}
	return false
}

// Union - объединяет множества
func (s *SetOnSlice[T]) Union(otherSet *SetOnSlice[T]) error {
	if otherSet == nil {
		return errors.New("множество не должно быть пустым")
	}
	for _, item := range otherSet.items {
		s.Add(item)
	}
	return nil
}

// Intersect - пересечение множеств
func (s *SetOnSlice[T]) Intersect(otherSet *SetOnSlice[T]) error {
	if otherSet == nil {
		return errors.New("множество не должно быть пустым")
	}
	for _, item := range s.items {
		if !otherSet.contains(item) {
			s.Remove(item)
		}
	}
	return nil
}

// Difference - разность множеств
func (s *SetOnSlice[T]) Difference(otherSet *SetOnSlice[T]) error {
	if otherSet == nil {
		return errors.New("множество не должно быть пустым")
	}
	for _, item := range otherSet.items {
		s.Remove(item)
	}
	return nil
}

// IsSubset - проверяет, является ли текущее множество подмножеством otherSet
func (s *SetOnSlice[T]) IsSubset(otherSet *SetOnSlice[T]) bool {
	for _, item := range s.items {
		if !otherSet.contains(item) {
			return false
		}
	}
	return true
}

func main() {
	set := NewSetOnSlice[float64]()

	isEmpty := set.Empty() // множество пустое?
	fmt.Println(isEmpty)   // true
}
