package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TopicModelingSettings 
type TopicModelingSettings struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Indicates whether the themes model should dynamically optimize the number of generated topics. To learn more, see Adjust maximum number of themes dynamically.
    dynamicallyAdjustTopicCount *bool
    // Indicates whether the themes model should exclude numbers while parsing document texts. To learn more, see Include numbers in themes.
    ignoreNumbers *bool
    // Indicates whether themes model is enabled for the case.
    isEnabled *bool
    // The OdataType property
    odataType *string
    // The total number of topics that the themes model will generate for a review set. To learn more, see Maximum number of themes.
    topicCount *int32
}
// NewTopicModelingSettings instantiates a new topicModelingSettings and sets the default values.
func NewTopicModelingSettings()(*TopicModelingSettings) {
    m := &TopicModelingSettings{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateTopicModelingSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTopicModelingSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTopicModelingSettings(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TopicModelingSettings) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDynamicallyAdjustTopicCount gets the dynamicallyAdjustTopicCount property value. Indicates whether the themes model should dynamically optimize the number of generated topics. To learn more, see Adjust maximum number of themes dynamically.
func (m *TopicModelingSettings) GetDynamicallyAdjustTopicCount()(*bool) {
    return m.dynamicallyAdjustTopicCount
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TopicModelingSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["dynamicallyAdjustTopicCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDynamicallyAdjustTopicCount(val)
        }
        return nil
    }
    res["ignoreNumbers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIgnoreNumbers(val)
        }
        return nil
    }
    res["isEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsEnabled(val)
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["topicCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTopicCount(val)
        }
        return nil
    }
    return res
}
// GetIgnoreNumbers gets the ignoreNumbers property value. Indicates whether the themes model should exclude numbers while parsing document texts. To learn more, see Include numbers in themes.
func (m *TopicModelingSettings) GetIgnoreNumbers()(*bool) {
    return m.ignoreNumbers
}
// GetIsEnabled gets the isEnabled property value. Indicates whether themes model is enabled for the case.
func (m *TopicModelingSettings) GetIsEnabled()(*bool) {
    return m.isEnabled
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *TopicModelingSettings) GetOdataType()(*string) {
    return m.odataType
}
// GetTopicCount gets the topicCount property value. The total number of topics that the themes model will generate for a review set. To learn more, see Maximum number of themes.
func (m *TopicModelingSettings) GetTopicCount()(*int32) {
    return m.topicCount
}
// Serialize serializes information the current object
func (m *TopicModelingSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("dynamicallyAdjustTopicCount", m.GetDynamicallyAdjustTopicCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("ignoreNumbers", m.GetIgnoreNumbers())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isEnabled", m.GetIsEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("topicCount", m.GetTopicCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TopicModelingSettings) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDynamicallyAdjustTopicCount sets the dynamicallyAdjustTopicCount property value. Indicates whether the themes model should dynamically optimize the number of generated topics. To learn more, see Adjust maximum number of themes dynamically.
func (m *TopicModelingSettings) SetDynamicallyAdjustTopicCount(value *bool)() {
    m.dynamicallyAdjustTopicCount = value
}
// SetIgnoreNumbers sets the ignoreNumbers property value. Indicates whether the themes model should exclude numbers while parsing document texts. To learn more, see Include numbers in themes.
func (m *TopicModelingSettings) SetIgnoreNumbers(value *bool)() {
    m.ignoreNumbers = value
}
// SetIsEnabled sets the isEnabled property value. Indicates whether themes model is enabled for the case.
func (m *TopicModelingSettings) SetIsEnabled(value *bool)() {
    m.isEnabled = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *TopicModelingSettings) SetOdataType(value *string)() {
    m.odataType = value
}
// SetTopicCount sets the topicCount property value. The total number of topics that the themes model will generate for a review set. To learn more, see Maximum number of themes.
func (m *TopicModelingSettings) SetTopicCount(value *int32)() {
    m.topicCount = value
}
