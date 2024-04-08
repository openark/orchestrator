package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AuditResource a class containing the properties for Audit Resource.
type AuditResource struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Audit resource's type.
    auditResourceType *string
    // Display name.
    displayName *string
    // List of modified properties.
    modifiedProperties []AuditPropertyable
    // The OdataType property
    odataType *string
    // Audit resource's Id.
    resourceId *string
}
// NewAuditResource instantiates a new auditResource and sets the default values.
func NewAuditResource()(*AuditResource) {
    m := &AuditResource{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAuditResourceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAuditResourceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAuditResource(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AuditResource) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAuditResourceType gets the auditResourceType property value. Audit resource's type.
func (m *AuditResource) GetAuditResourceType()(*string) {
    return m.auditResourceType
}
// GetDisplayName gets the displayName property value. Display name.
func (m *AuditResource) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AuditResource) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["auditResourceType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuditResourceType(val)
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["modifiedProperties"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAuditPropertyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AuditPropertyable, len(val))
            for i, v := range val {
                res[i] = v.(AuditPropertyable)
            }
            m.SetModifiedProperties(res)
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
    res["resourceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResourceId(val)
        }
        return nil
    }
    return res
}
// GetModifiedProperties gets the modifiedProperties property value. List of modified properties.
func (m *AuditResource) GetModifiedProperties()([]AuditPropertyable) {
    return m.modifiedProperties
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *AuditResource) GetOdataType()(*string) {
    return m.odataType
}
// GetResourceId gets the resourceId property value. Audit resource's Id.
func (m *AuditResource) GetResourceId()(*string) {
    return m.resourceId
}
// Serialize serializes information the current object
func (m *AuditResource) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("auditResourceType", m.GetAuditResourceType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    if m.GetModifiedProperties() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetModifiedProperties()))
        for i, v := range m.GetModifiedProperties() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("modifiedProperties", cast)
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
        err := writer.WriteStringValue("resourceId", m.GetResourceId())
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
func (m *AuditResource) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAuditResourceType sets the auditResourceType property value. Audit resource's type.
func (m *AuditResource) SetAuditResourceType(value *string)() {
    m.auditResourceType = value
}
// SetDisplayName sets the displayName property value. Display name.
func (m *AuditResource) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetModifiedProperties sets the modifiedProperties property value. List of modified properties.
func (m *AuditResource) SetModifiedProperties(value []AuditPropertyable)() {
    m.modifiedProperties = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AuditResource) SetOdataType(value *string)() {
    m.odataType = value
}
// SetResourceId sets the resourceId property value. Audit resource's Id.
func (m *AuditResource) SetResourceId(value *string)() {
    m.resourceId = value
}
