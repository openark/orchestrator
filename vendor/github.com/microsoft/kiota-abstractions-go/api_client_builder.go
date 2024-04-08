package abstractions

import (
	"github.com/microsoft/kiota-abstractions-go/store"
	sync "sync"

	s "github.com/microsoft/kiota-abstractions-go/serialization"
)

var serializerMutex sync.Mutex
var deserializerMutex sync.Mutex

// RegisterDefaultSerializer registers the default serializer to the registry singleton to be used by the request adapter.
func RegisterDefaultSerializer(metaFactory func() s.SerializationWriterFactory) {
	factory := metaFactory()
	contentType, err := factory.GetValidContentType()
	if err == nil && contentType != "" {
		serializerMutex.Lock()
		s.DefaultSerializationWriterFactoryInstance.ContentTypeAssociatedFactories[contentType] = factory
		serializerMutex.Unlock()
	}
}

// RegisterDefaultDeserializer registers the default deserializer to the registry singleton to be used by the request adapter.
func RegisterDefaultDeserializer(metaFactory func() s.ParseNodeFactory) {
	factory := metaFactory()
	contentType, err := factory.GetValidContentType()
	if err == nil && contentType != "" {
		deserializerMutex.Lock()
		s.DefaultParseNodeFactoryInstance.ContentTypeAssociatedFactories[contentType] = factory
		deserializerMutex.Unlock()
	}
}

// EnableBackingStoreForSerializationWriterFactory Enables the backing store on default serialization writers and the given serialization writer.
func EnableBackingStoreForSerializationWriterFactory(factory s.SerializationWriterFactory) s.SerializationWriterFactory {
	switch v := factory.(type) {
	case *s.SerializationWriterFactoryRegistry:
		enableBackingStoreForSerializationRegistry(v)
	default:
		factory = store.NewBackingStoreSerializationWriterProxyFactory(factory)
		enableBackingStoreForSerializationRegistry(s.DefaultSerializationWriterFactoryInstance)
	}
	return factory
}

func enableBackingStoreForSerializationRegistry(registry *s.SerializationWriterFactoryRegistry) {
	for key, value := range registry.ContentTypeAssociatedFactories {
		if _, ok := value.(*store.BackingStoreSerializationWriterProxyFactory); !ok {
			registry.ContentTypeAssociatedFactories[key] = store.NewBackingStoreSerializationWriterProxyFactory(value)
		}
	}
}

// EnableBackingStoreForParseNodeFactory Enables the backing store on default parse nodes factories and the given parse node factory.
func EnableBackingStoreForParseNodeFactory(factory s.ParseNodeFactory) s.ParseNodeFactory {
	switch v := factory.(type) {
	case *s.ParseNodeFactoryRegistry:
		enableBackingStoreForParseNodeRegistry(v)
	default:
		factory = store.NewBackingStoreParseNodeFactory(factory)
		enableBackingStoreForParseNodeRegistry(s.DefaultParseNodeFactoryInstance)
	}
	return factory
}

func enableBackingStoreForParseNodeRegistry(registry *s.ParseNodeFactoryRegistry) {
	for key, value := range registry.ContentTypeAssociatedFactories {
		if _, ok := value.(*store.BackingStoreParseNodeFactory); !ok {
			registry.ContentTypeAssociatedFactories[key] = store.NewBackingStoreParseNodeFactory(value)
		}
	}
}
