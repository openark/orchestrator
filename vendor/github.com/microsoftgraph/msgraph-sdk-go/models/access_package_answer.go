package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AccessPackageAnswer 
type AccessPackageAnswer struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The answeredQuestion property
    answeredQuestion AccessPackageQuestionable
    // The localized display value shown to the requestor and approvers.
    displayValue *string
    // The OdataType property
    odataType *string
}
// NewAccessPackageAnswer instantiates a new accessPackageAnswer and sets the default values.
func NewAccessPackageAnswer()(*AccessPackageAnswer) {
    m := &AccessPackageAnswer{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAccessPackageAnswerFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAccessPackageAnswerFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.accessPackageAnswerString":
                        return NewAccessPackageAnswerString(), nil
                }
            }
        }
    }
    return NewAccessPackageAnswer(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AccessPackageAnswer) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAnsweredQuestion gets the answeredQuestion property value. The answeredQuestion property
func (m *AccessPackageAnswer) GetAnsweredQuestion()(AccessPackageQuestionable) {
    return m.answeredQuestion
}
// GetDisplayValue gets the displayValue property value. The localized display value shown to the requestor and approvers.
func (m *AccessPackageAnswer) GetDisplayValue()(*string) {
    return m.displayValue
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AccessPackageAnswer) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["answeredQuestion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAccessPackageQuestionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAnsweredQuestion(val.(AccessPackageQuestionable))
        }
        return nil
    }
    res["displayValue"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayValue(val)
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
func (m *AccessPackageAnswer) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *AccessPackageAnswer) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("answeredQuestion", m.GetAnsweredQuestion())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("displayValue", m.GetDisplayValue())
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
func (m *AccessPackageAnswer) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAnsweredQuestion sets the answeredQuestion property value. The answeredQuestion property
func (m *AccessPackageAnswer) SetAnsweredQuestion(value AccessPackageQuestionable)() {
    m.answeredQuestion = value
}
// SetDisplayValue sets the displayValue property value. The localized display value shown to the requestor and approvers.
func (m *AccessPackageAnswer) SetDisplayValue(value *string)() {
    m.displayValue = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AccessPackageAnswer) SetOdataType(value *string)() {
    m.odataType = value
}
