package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EducationItemBody 
type EducationItemBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The content property
    content *string
    // The contentType property
    contentType *BodyType
    // The OdataType property
    odataType *string
}
// NewEducationItemBody instantiates a new educationItemBody and sets the default values.
func NewEducationItemBody()(*EducationItemBody) {
    m := &EducationItemBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateEducationItemBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEducationItemBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEducationItemBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EducationItemBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetContent gets the content property value. The content property
func (m *EducationItemBody) GetContent()(*string) {
    return m.content
}
// GetContentType gets the contentType property value. The contentType property
func (m *EducationItemBody) GetContentType()(*BodyType) {
    return m.contentType
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EducationItemBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["content"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContent(val)
        }
        return nil
    }
    res["contentType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseBodyType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentType(val.(*BodyType))
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
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *EducationItemBody) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *EducationItemBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("content", m.GetContent())
        if err != nil {
            return err
        }
    }
    if m.GetContentType() != nil {
        cast := (*m.GetContentType()).String()
        err := writer.WriteStringValue("contentType", &cast)
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EducationItemBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetContent sets the content property value. The content property
func (m *EducationItemBody) SetContent(value *string)() {
    m.content = value
}
// SetContentType sets the contentType property value. The contentType property
func (m *EducationItemBody) SetContentType(value *BodyType)() {
    m.contentType = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *EducationItemBody) SetOdataType(value *string)() {
    m.odataType = value
}
