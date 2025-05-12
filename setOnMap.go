package main

import (
	"fmt"
)

// SetOnMap - множество на базе хэш таблицы
type SetOnMap[T comparable] map[T]struct{}

// NewSetOnMap - создает множество на базе слайса
func NewSetOnMap[T comparable]() SetOnMap[T] {
	return SetOnMap[T]{}
}

// Print - возвращает строковое представление множества
func (s SetOnMap[T]) Print() string {
	result := ""
	for key := range s {
		result += fmt.Sprintf("%v ", key)
	}
	return result
}

// Empty - проверяет, пустое ли множество
func (s SetOnMap[T]) Empty() bool {
	return len(s) == 0
}

// Add - добавляет значение в множество
func (s SetOnMap[T]) Add(item T) bool {
	// проверяем есть ли уже такое значение во множестве, если есть то ничего не добавляем и возвращаем false
	if _, ok := s[item]; ok {
		return false
	}
	s[item] = struct{}{} // добавляем значение во множество
	return true
}

// Remove - удаляет значение из множества
func (s SetOnMap[T]) Remove(item T) bool {
	if _, ok := s[item]; ok {
		delete(s, item)
		return true
	}
	return false
}

// Union - объединяет множества
func (s SetOnMap[T]) Union(otherSet SetOnMap[T]) {
	for key := range otherSet {
		s.Add(key)
	}
}

// Intersect - пересечение множеств
func (s SetOnMap[T]) Intersect(otherSet SetOnMap[T]) {
	for key := range s {
		if _, ok := otherSet[key]; !ok {
			delete(s, key)
		}
	}
}

// Difference - разность множеств
func (s SetOnMap[T]) Difference(otherSet SetOnMap[T]) {
	for key := range s {
		if _, ok := otherSet[key]; ok {
			delete(s, key)
		}
	}
}

// IsSubset - проверяет, является ли otherSet подмножеством текущего множества
func (s SetOnMap[T]) IsSubset(otherSet SetOnMap[T]) bool {
	for key := range otherSet {
		if _, ok := s[key]; !ok {
			return false
		}
	}
	return true
}
