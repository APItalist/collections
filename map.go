package collections

// Map represents a data structure that is stored in a key-value fashion. Values can be looked up by keys.
//
// The K and V type parameters specify they ey and value types, respectively.
// TKeys specifies the type of the key set returned from the Keys() call, while TValues is the type of the Collection
// of values returned from the Values() call.
type Map[K, V comparable, TKeys Set[K], TValues Collection[V]] interface {
	// Keys returns a set of keys within the map.
	Keys() TKeys

	// Values returns a collection of all values.
	Values() TValues

	// Get returns a value of the specified key, or an ErrKeyNotFound if the specified key is not found in the map.
	Get(K) (V, error)
	// GetOrDefault returns a value of the specified key, or the defaultValue if the specified key is not found in the
	// map.
	GetOrDefault(K, defaultValue V) V
	// ContainsKey returns true if the specified key is present in the map.
	ContainsKey(K) bool
	// ContainsValue returns true if the specified map contains te specified value.
	ContainsValue(V) bool
	// IsEmpty returns true if there are no items in the current map.
	IsEmpty() bool
	// Size returns the number of map entries.
	Size() uint
	// Stream creates a processable stream of map entries for all elements in the current map.
	Stream() Stream[MapEntry[K, V]]
}

// MapEntry is a simple data structure that holds one map key and one map value.
type MapEntry[K, V comparable] struct {
	Key   K
	Value V
}

// MutableMap is a map that can be changed.
//
// The K and V type parameters specify they ey and value types, respectively.
// TKeys specifies the type of the key set returned from the Keys() call, while TValues is the type of the Collection
// of values returned from the Values() call.
type MutableMap[K, V comparable, TKeys Set[K], TValues Collection[V]] interface {
	Map[K, V, TKeys, TValues]

	// Put sets the specified key to contain the specified value.
	Put(K, V)

	// PutAll puts all values from the passed map into the current map.
	PutAll(Map[K, V, Set[K], Collection[V]])

	// PutIfAbsent sets the specified value to the specified key if, and only if, the key does not yet exist. If the
	// current value is set, that value is returned.
	PutIfAbsent(K, V) *V

	// RemoveKey removes the specified key from the map if present. If the specified key was present, the old value is
	// returned.
	RemoveKey(K) *V

	// Remove removes a specified key/value combination if present.
	Remove(K, V)

	// Replace replaces the given value associated with the specified key only if it is already present. In this case
	// the old value is returned.
	Replace(K, V) *V
}

// ImmutableMap is a map that cannot be changed, but contains helper functions to create a modified copy of the current
// map.
//
// The K and V type parameters specify they ey and value types, respecitvely.
// TKeys specifies the type of the key set returned from the Keys() call, while TValues is the type of the Collection
// of values returned from the Values() call.
type ImmutableMap[K, V comparable, TKeys Set[K], TValues Collection[V]] interface {
	Map[K, V, TKeys, TValues]

	// WithPut creates a copy of the current map with the specified value set on the specified key.
	WithPut(K, V) ImmutableMap[K, V, TKeys, TValues]

	// WithPutAll creates a copied map with all keys and values set from the passed map.
	WithPutAll(Map[K, V, Set[K], Set[V]]) ImmutableMap[K, V, TKeys, TValues]

	// WithPutIfAbsent creates a copied map with the specified value set on the key only if the key was previously not
	// set.
	WithPutIfAbsent(K, V) ImmutableMap[K, V, TKeys, TValues]

	// WithRemovedKey creates a copied map with the specified key removed.
	WithRemovedKey(K) ImmutableMap[K, V, TKeys, TValues]

	// WithRemoved creates a copied map with the specified key removed only if it was set to the specified value.
	WithRemoved(K, V) ImmutableMap[K, V, TKeys, TValues]

	// WithReplaced creates a copied map with the specified key set to the specified value only if it was previously
	// set.
	WithReplaced(K, V) ImmutableMap[K, V, TKeys, TValues]
}
