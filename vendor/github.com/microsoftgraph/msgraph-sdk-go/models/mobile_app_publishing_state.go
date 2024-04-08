package models
import (
    "errors"
)
// Indicates the publishing state of an app.
type MobileAppPublishingState int

const (
    // The app is not yet published.
    NOTPUBLISHED_MOBILEAPPPUBLISHINGSTATE MobileAppPublishingState = iota
    // The app is pending service-side processing.
    PROCESSING_MOBILEAPPPUBLISHINGSTATE
    // The app is published.
    PUBLISHED_MOBILEAPPPUBLISHINGSTATE
)

func (i MobileAppPublishingState) String() string {
    return []string{"notPublished", "processing", "published"}[i]
}
func ParseMobileAppPublishingState(v string) (any, error) {
    result := NOTPUBLISHED_MOBILEAPPPUBLISHINGSTATE
    switch v {
        case "notPublished":
            result = NOTPUBLISHED_MOBILEAPPPUBLISHINGSTATE
        case "processing":
            result = PROCESSING_MOBILEAPPPUBLISHINGSTATE
        case "published":
            result = PUBLISHED_MOBILEAPPPUBLISHINGSTATE
        default:
            return 0, errors.New("Unknown MobileAppPublishingState value: " + v)
    }
    return &result, nil
}
func SerializeMobileAppPublishingState(values []MobileAppPublishingState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
