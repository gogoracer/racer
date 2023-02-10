package turnsignal

type Accessor[T any] func() T
type Setter[T any] func(T)

type effect struct {
	id           int
	execute      func()
	dependencies []*effectSet
}

type effectSet struct {
	arr  []*effect
	keys map[int]struct{}
}

func newEffectSet() *effectSet {
	return &effectSet{
		arr:  []*effect{},
		keys: map[int]struct{}{},
	}
}

func (s *effectSet) add(e *effect) {
	if _, ok := s.keys[e.id]; !ok {
		s.keys[e.id] = struct{}{}
		s.arr = append(s.arr, e)
	}
}

func (s *effectSet) remove(e *effect) {
	if _, ok := s.keys[e.id]; ok {
		delete(s.keys, e.id)
		for i, v := range s.arr {
			if v == e {
				s.arr = append(s.arr[:i], s.arr[i+1:]...)
				break
			}
		}
	}
}

func (s *effectSet) execute() {
	for _, e := range s.arr {
		e.execute()
	}
}

var (
	effectContextStack []*effect
	nextEffectID       int
)

func CreateSignal[T any](value T) (Accessor[T], Setter[T]) {
	subscriptions := newEffectSet()
	val := value

	read := func() T {
		headIndex := len(effectContextStack) - 1
		if headIndex >= 0 {
			running := effectContextStack[headIndex]

			// subscribe
			subscriptions.add(running)
			running.dependencies = append(running.dependencies, subscriptions)
		}
		return val
	}

	write := func(nextValue T) {
		val = nextValue
		subscriptions.execute()
	}

	return read, write
}

func CreateEffect(fn func()) {
	var e *effect
	cleanup := func() {
		for _, s := range e.dependencies {
			s.remove(e)
		}
		e.dependencies = e.dependencies[:0]
	}
	e = &effect{
		id: nextEffectID,
		execute: func() {
			cleanup()
			effectContextStack = append(effectContextStack, e)
			fn()
			effectContextStack = effectContextStack[:len(effectContextStack)-1]
		},
		dependencies: nil,
	}
	nextEffectID++

	e.execute()
}

func CreateMemo[T any](fn func() T, initial T) Accessor[T] {
	getMemo, setMemo := CreateSignal(initial)
	CreateEffect(func() {
		setMemo(fn())
	})
	return getMemo
}
