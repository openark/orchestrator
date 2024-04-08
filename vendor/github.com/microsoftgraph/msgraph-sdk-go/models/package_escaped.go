package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PackageEscaped 
type PackageEscaped struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The OdataType property
    odataType *string
    // A string indicating the type of package. While oneNote is the only currently defined value, you should expect other package types to be returned and handle them accordingly.
    typeEscaped *string
}
// NewPackageEscaped instantiates a new packageEscaped and sets the default values.
func NewPackageEscaped()(*PackageEscaped) {
    m := &PackageEscaped{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreatePackageEscapedFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePackageEscapedFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPackageEscaped(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PackageEscaped) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PackageEscaped) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetType(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *PackageEscaped) GetOdataType()(*string) {
    return m.odataType
}
// GetType gets the type property value. A string indicating the type of package. While oneNote is the only currently defined value, you should expect other package types to be returned and handle them accordingly.
func (m *PackageEscaped) GetType()(*string) {
    return m.typeEscaped
}
// Serialize serializes information the current object
func (m *PackageEscaped) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("type", m.GetType())
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
func (m *PackageEscaped) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *PackageEscaped) SetOdataType(value *string)() {
    m.odataType = value
}
// SetType sets the type property value. A string indicating the type of package. While oneNote is the only currently defined value, you should expect other package types to be returned and handle them accordingly.
func (m *PackageEscaped) SetType(value *string)() {
    m.typeEscaped = value
}
