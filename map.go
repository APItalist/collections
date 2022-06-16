package collections

type Map[K, V comparable] interface {
    Get(K) (V, error)
    GetOrDefault(K, defaultValue V) V
    ContainsKey(K) bool
    ContainsValue(V) bool
    IsEmpty() bool
    Stream() Stream[MapEntry[K, V]]
}

type MapEntry[K, V comparable] struct {
    K
    V
}

type MutableMap[K, V comparable] interface {
    Map[K, V]

    Put(K, V)
    PutAll(Map[K, V])
    PutIfAbsent(K, V) V
    RemoveKey(K) V
    Remove(K, V)
    Replace(K, V)
    Size() uint
    Values() Collection[V]
}

type ImmutableMap[K, V comparable] interface {
    Map[K, V]

    WithPut(K, V) ImmutableMap[K, V]
    WithPutAll(Map[K, V]) ImmutableMap[K, V]
    WithPutIfAbsent(K, V) ImmutableMap[K, V]
    WithRemovedKey(K) ImmutableMap[K, V]
    WithRemoved(K, V) ImmutableMap[K, V]
    WithReplaced(K, V) ImmutableMap[K, V]
}
