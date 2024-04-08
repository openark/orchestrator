package serialization

import (
	"errors"
	re "regexp"
	"strings"
)

// ParseNodeFactoryRegistry holds a list of all the registered factories for the various types of nodes.
type ParseNodeFactoryRegistry struct {
	ContentTypeAssociatedFactories map[string]ParseNodeFactory
}

// DefaultParseNodeFactoryInstance is the default singleton instance of the registry to be used when registering new factories that should be available by default.
var DefaultParseNodeFactoryInstance = &ParseNodeFactoryRegistry{
	ContentTypeAssociatedFactories: make(map[string]ParseNodeFactory),
}

// GetValidContentType returns the valid content type for the ParseNodeFactoryRegistry
func (m *ParseNodeFactoryRegistry) GetValidContentType() (string, error) {
	return "", errors.New("the registry supports multiple content types. Get the registered factory instead")
}

var contentTypeVendorCleanupPattern = re.MustCompile("[^/]+\\+")

// GetRootParseNode returns a new ParseNode instance that is the root of the content
func (m *ParseNodeFactoryRegistry) GetRootParseNode(contentType string, content []byte) (ParseNode, error) {
	if contentType == "" {
		return nil, errors.New("contentType is required")
	}
	if content == nil {
		return nil, errors.New("content is required")
	}
	vendorSpecificContentType := strings.Split(contentType, ";")[0]
	factory, ok := m.ContentTypeAssociatedFactories[vendorSpecificContentType]
	if ok {
		return factory.GetRootParseNode(vendorSpecificContentType, content)
	}
	cleanedContentType := contentTypeVendorCleanupPattern.ReplaceAllString(vendorSpecificContentType, "")
	factory, ok = m.ContentTypeAssociatedFactories[cleanedContentType]
	if ok {
		return factory.GetRootParseNode(cleanedContentType, content)
	}
	return nil, errors.New("content type " + cleanedContentType + " does not have a factory registered to be parsed")
}
