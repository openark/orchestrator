package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Bundle 
type Bundle struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // If the bundle is an [album][], then the album property is included
    album Albumable
    // Number of children contained immediately within this container.
    childCount *int32
    // The OdataType property
    odataType *string
}
// NewBundle instantiates a new bundle and sets the default values.
func NewBundle()(*Bundle) {
    m := &Bundle{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateBundleFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateBundleFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewBundle(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Bundle) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAlbum gets the album property value. If the bundle is an [album][], then the album property is included
func (m *Bundle) GetAlbum()(Albumable) {
    return m.album
}
// GetChildCount gets the childCount property value. Number of children contained immediately within this container.
func (m *Bundle) GetChildCount()(*int32) {
    return m.childCount
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Bundle) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["album"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAlbumFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAlbum(val.(Albumable))
        }
        return nil
    }
    res["childCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetChildCount(val)
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
func (m *Bundle) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *Bundle) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("album", m.GetAlbum())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("childCount", m.GetChildCount())
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
func (m *Bundle) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAlbum sets the album property value. If the bundle is an [album][], then the album property is included
func (m *Bundle) SetAlbum(value Albumable)() {
    m.album = value
}
// SetChildCount sets the childCount property value. Number of children contained immediately within this container.
func (m *Bundle) SetChildCount(value *int32)() {
    m.childCount = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *Bundle) SetOdataType(value *string)() {
    m.odataType = value
}
