package main

import "sync"

type TheSyncMap[K, V any] struct {
	m *sync.Map
}
type Key any
type Value any

func (s *TheSyncMap[K, V]) Init() *TheSyncMap[K, V] { //初始化
	s.m = new(sync.Map)
	return s
}
func (s *TheSyncMap[K, V]) Set(k Key, v Value) *TheSyncMap[K, V] { //设置
	s.m.Store(k, v)
	return s
}
func (s *TheSyncMap[K, V]) Get(k Key) (Value, bool) { //取值
	v, ok := s.m.Load(k)
	return v, ok
}
func (s *TheSyncMap[K, V]) Update(k Key, v Value) *TheSyncMap[K, V] { //更新值
	ok := s.Contain(k)
	if !ok {
		return s
	}
	s.Set(k, v)
	return s
}
func (s *TheSyncMap[K, V]) Delete(k Key) *TheSyncMap[K, V] { //删除元素
	s.m.Delete(k)
	return s
}
func (s *TheSyncMap[K, V]) Contain(k Key) bool { //是否蕴含
	_, ok := s.m.Load(k)
	return ok
}
func (s *TheSyncMap[K, V]) SetOrGet(k Key, v Value) (Value, bool) { //设置或取值 如果之前已经有这个键的，那么返回值和1，没有则赋值，bool返回0
	actual, ok := s.m.LoadOrStore(k, v)
	return actual, ok
}
func (s *TheSyncMap[K, V]) GetAndDelete(k Key) (Value, bool) { //取值或删除 有则返回以前的值 bool报告键是否存在
	actual, ok := s.m.LoadAndDelete(k)
	return actual, ok
}
