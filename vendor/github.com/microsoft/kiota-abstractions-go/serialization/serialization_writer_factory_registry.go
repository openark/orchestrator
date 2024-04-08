package serialization

import (
	"errors"
	"strings"
)

// SerializationWriterFactoryRegistry is a factory holds a list of all the registered factories for the various types of nodes.
type SerializationWriterFactoryRegistry struct {
	// ContentTypeAssociatedFactories list of factories that are registered by content type.
	ContentTypeAssociatedFactories map[string]SerializationWriterFactory
}

// DefaultSerializationWriterFactoryInstance is the default singleton instance of the registry to be used when registering new factories that should be available by default.
var DefaultSerializationWriterFactoryInstance = &SerializationWriterFactoryRegistry{
	ContentTypeAssociatedFactories: make(map[string]SerializationWriterFactory),
}

// GetValidContentType returns the valid content type for the SerializationWriterFactoryRegistry
func (m *SerializationWriterFactoryRegistry) GetValidContentType() (string, error) {
	return "", errors.New("the registry supports multiple content types. Get the registered factory instead")
}

// GetSerializationWriter returns the relevant SerializationWriter instance for the given content type
func (m *SerializationWriterFactoryRegistry) GetSerializationWriter(contentType string) (SerializationWriter, error) {
	if contentType == "" {
		return nil, errors.New("the content type is empty")
	}
	vendorSpecificContentType := strings.Split(contentType, ";")[0]
	factory, ok := m.ContentTypeAssociatedFactories[vendorSpecificContentType]
	if ok {
		return factory.GetSerializationWriter(contentType)
	}
	cleanedContentType := contentTypeVendorCleanupPattern.ReplaceAllString(vendorSpecificContentType, "")
	factory, ok = m.ContentTypeAssociatedFactories[cleanedContentType]
	if ok {
		return factory.GetSerializationWriter(cleanedContentType)
	}
	return nil, errors.New("Content type " + cleanedContentType + " does not have a factory registered to be parsed")
}
