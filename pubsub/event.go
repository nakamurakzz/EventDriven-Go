package pubsub

import (
	"context"
)

type EventType int

const (
	None EventType = iota
	LevelUp
	QuestClear
)

type Eventer interface {
	EventType() EventType
}

func Event[T Eventer](ctx context.Context) []T {
	m := GetManager(ctx)
	var obj T
	events := m.Events(obj.EventType())
	s := make([]T, len(events))
	for _, e := range events {
		s = append(s, e.(T))
	}
	return s
}

func (e LevelUpEvent) EventType() EventType {
	return LevelUp
}

func (e QuestClearEvent) EventType() EventType {
	return QuestClear
}

type LevelUpEvent struct {
	Level int
}

type QuestClearEvent struct {
	QuestId   int
	ClearRank string
}

type typeManagerKey struct{}

func SetManager(ctx context.Context, m *Manager) context.Context {
	return context.WithValue(ctx, typeManagerKey{}, m)
}

func GetManager(ctx context.Context) *Manager {
	m, ok := ctx.Value(typeManagerKey{}).(*Manager)
	if !ok {
		return nil
	}
	return m
}

type Manager struct {
	events map[EventType][]Eventer
}

func NewManager() *Manager {
	return &Manager{}
}

func (m *Manager) Push(event Eventer) {
	t := event.EventType()
	m.events[t] = append(m.events[t], event)
}

func (m *Manager) Events(t EventType) []Eventer {
	return m.events[t]
}

func Exec() {
	m := NewManager()
	ctx := SetManager(context.Background(), m)
	doBusiness(ctx)
}

func doBusiness(ctx context.Context) {
	m := GetManager(ctx)
	m.Push(QuestClearEvent{QuestId: 1, ClearRank: "S"})
	m.Push(&LevelUpEvent{Level: 10})
	m.Push(QuestClearEvent{QuestId: 2, ClearRank: "A"})

	for _, e := range m.Events(QuestClear) {
		switch e := e.(type) {
		case QuestClearEvent:
			println("quest clear", e.QuestId, e.ClearRank)
		}
	}

	for _, e := range m.Events(LevelUp) {
		switch e := e.(type) {
		case *LevelUpEvent:
			println("level up", e.Level)
		}
	}
}
