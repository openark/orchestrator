package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// LookupColumn 
type LookupColumn struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Indicates whether multiple values can be selected from the source.
    allowMultipleValues *bool
    // Indicates whether values in the column should be able to exceed the standard limit of 255 characters.
    allowUnlimitedLength *bool
    // The name of the lookup source column.
    columnName *string
    // The unique identifier of the lookup source list.
    listId *string
    // The OdataType property
    odataType *string
    // If specified, this column is a secondary lookup, pulling an additional field from the list item looked up by the primary lookup. Use the list item looked up by the primary as the source for the column named here.
    primaryLookupColumnId *string
}
// NewLookupColumn instantiates a new lookupColumn and sets the default values.
func NewLookupColumn()(*LookupColumn) {
    m := &LookupColumn{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateLookupColumnFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateLookupColumnFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewLookupColumn(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *LookupColumn) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAllowMultipleValues gets the allowMultipleValues property value. Indicates whether multiple values can be selected from the source.
func (m *LookupColumn) GetAllowMultipleValues()(*bool) {
    return m.allowMultipleValues
}
// GetAllowUnlimitedLength gets the allowUnlimitedLength property value. Indicates whether values in the column should be able to exceed the standard limit of 255 characters.
func (m *LookupColumn) GetAllowUnlimitedLength()(*bool) {
    return m.allowUnlimitedLength
}
// GetColumnName gets the columnName property value. The name of the lookup source column.
func (m *LookupColumn) GetColumnName()(*string) {
    return m.columnName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *LookupColumn) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["allowMultipleValues"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowMultipleValues(val)
        }
        return nil
    }
    res["allowUnlimitedLength"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowUnlimitedLength(val)
        }
        return nil
    }
    res["columnName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetColumnName(val)
        }
        return nil
    }
    res["listId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetListId(val)
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
    res["primaryLookupColumnId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrimaryLookupColumnId(val)
        }
        return nil
    }
    return res
}
// GetListId gets the listId property value. The unique identifier of the lookup source list.
func (m *LookupColumn) GetListId()(*string) {
    return m.listId
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *LookupColumn) GetOdataType()(*string) {
    return m.odataType
}
// GetPrimaryLookupColumnId gets the primaryLookupColumnId property value. If specified, this column is a secondary lookup, pulling an additional field from the list item looked up by the primary lookup. Use the list item looked up by the primary as the source for the column named here.
func (m *LookupColumn) GetPrimaryLookupColumnId()(*string) {
    return m.primaryLookupColumnId
}
// Serialize serializes information the current object
func (m *LookupColumn) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("allowMultipleValues", m.GetAllowMultipleValues())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("allowUnlimitedLength", m.GetAllowUnlimitedLength())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("columnName", m.GetColumnName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("listId", m.GetListId())
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
        err := writer.WriteStringValue("primaryLookupColumnId", m.GetPrimaryLookupColumnId())
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
func (m *LookupColumn) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAllowMultipleValues sets the allowMultipleValues property value. Indicates whether multiple values can be selected from the source.
func (m *LookupColumn) SetAllowMultipleValues(value *bool)() {
    m.allowMultipleValues = value
}
// SetAllowUnlimitedLength sets the allowUnlimitedLength property value. Indicates whether values in the column should be able to exceed the standard limit of 255 characters.
func (m *LookupColumn) SetAllowUnlimitedLength(value *bool)() {
    m.allowUnlimitedLength = value
}
// SetColumnName sets the columnName property value. The name of the lookup source column.
func (m *LookupColumn) SetColumnName(value *string)() {
    m.columnName = value
}
// SetListId sets the listId property value. The unique identifier of the lookup source list.
func (m *LookupColumn) SetListId(value *string)() {
    m.listId = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *LookupColumn) SetOdataType(value *string)() {
    m.odataType = value
}
// SetPrimaryLookupColumnId sets the primaryLookupColumnId property value. If specified, this column is a secondary lookup, pulling an additional field from the list item looked up by the primary lookup. Use the list item looked up by the primary as the source for the column named here.
func (m *LookupColumn) SetPrimaryLookupColumnId(value *string)() {
    m.primaryLookupColumnId = value
}
