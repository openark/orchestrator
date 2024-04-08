package models

import (
    i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IosVppEBook 
type IosVppEBook struct {
    ManagedEBook
    // The Apple ID associated with Vpp token.
    appleId *string
    // Genres.
    genres []string
    // Language.
    language *string
    // Seller.
    seller *string
    // Total license count.
    totalLicenseCount *int32
    // Used license count.
    usedLicenseCount *int32
    // The Vpp token's organization name.
    vppOrganizationName *string
    // The Vpp token ID.
    vppTokenId *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID
}
// NewIosVppEBook instantiates a new IosVppEBook and sets the default values.
func NewIosVppEBook()(*IosVppEBook) {
    m := &IosVppEBook{
        ManagedEBook: *NewManagedEBook(),
    }
    odataTypeValue := "#microsoft.graph.iosVppEBook"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateIosVppEBookFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateIosVppEBookFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIosVppEBook(), nil
}
// GetAppleId gets the appleId property value. The Apple ID associated with Vpp token.
func (m *IosVppEBook) GetAppleId()(*string) {
    return m.appleId
}
// GetFieldDeserializers the deserialization information for the current model
func (m *IosVppEBook) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ManagedEBook.GetFieldDeserializers()
    res["appleId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppleId(val)
        }
        return nil
    }
    res["genres"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetGenres(res)
        }
        return nil
    }
    res["language"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLanguage(val)
        }
        return nil
    }
    res["seller"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSeller(val)
        }
        return nil
    }
    res["totalLicenseCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalLicenseCount(val)
        }
        return nil
    }
    res["usedLicenseCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUsedLicenseCount(val)
        }
        return nil
    }
    res["vppOrganizationName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVppOrganizationName(val)
        }
        return nil
    }
    res["vppTokenId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetUUIDValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVppTokenId(val)
        }
        return nil
    }
    return res
}
// GetGenres gets the genres property value. Genres.
func (m *IosVppEBook) GetGenres()([]string) {
    return m.genres
}
// GetLanguage gets the language property value. Language.
func (m *IosVppEBook) GetLanguage()(*string) {
    return m.language
}
// GetSeller gets the seller property value. Seller.
func (m *IosVppEBook) GetSeller()(*string) {
    return m.seller
}
// GetTotalLicenseCount gets the totalLicenseCount property value. Total license count.
func (m *IosVppEBook) GetTotalLicenseCount()(*int32) {
    return m.totalLicenseCount
}
// GetUsedLicenseCount gets the usedLicenseCount property value. Used license count.
func (m *IosVppEBook) GetUsedLicenseCount()(*int32) {
    return m.usedLicenseCount
}
// GetVppOrganizationName gets the vppOrganizationName property value. The Vpp token's organization name.
func (m *IosVppEBook) GetVppOrganizationName()(*string) {
    return m.vppOrganizationName
}
// GetVppTokenId gets the vppTokenId property value. The Vpp token ID.
func (m *IosVppEBook) GetVppTokenId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID) {
    return m.vppTokenId
}
// Serialize serializes information the current object
func (m *IosVppEBook) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ManagedEBook.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("appleId", m.GetAppleId())
        if err != nil {
            return err
        }
    }
    if m.GetGenres() != nil {
        err = writer.WriteCollectionOfStringValues("genres", m.GetGenres())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("language", m.GetLanguage())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("seller", m.GetSeller())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("totalLicenseCount", m.GetTotalLicenseCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("usedLicenseCount", m.GetUsedLicenseCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("vppOrganizationName", m.GetVppOrganizationName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteUUIDValue("vppTokenId", m.GetVppTokenId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAppleId sets the appleId property value. The Apple ID associated with Vpp token.
func (m *IosVppEBook) SetAppleId(value *string)() {
    m.appleId = value
}
// SetGenres sets the genres property value. Genres.
func (m *IosVppEBook) SetGenres(value []string)() {
    m.genres = value
}
// SetLanguage sets the language property value. Language.
func (m *IosVppEBook) SetLanguage(value *string)() {
    m.language = value
}
// SetSeller sets the seller property value. Seller.
func (m *IosVppEBook) SetSeller(value *string)() {
    m.seller = value
}
// SetTotalLicenseCount sets the totalLicenseCount property value. Total license count.
func (m *IosVppEBook) SetTotalLicenseCount(value *int32)() {
    m.totalLicenseCount = value
}
// SetUsedLicenseCount sets the usedLicenseCount property value. Used license count.
func (m *IosVppEBook) SetUsedLicenseCount(value *int32)() {
    m.usedLicenseCount = value
}
// SetVppOrganizationName sets the vppOrganizationName property value. The Vpp token's organization name.
func (m *IosVppEBook) SetVppOrganizationName(value *string)() {
    m.vppOrganizationName = value
}
// SetVppTokenId sets the vppTokenId property value. The Vpp token ID.
func (m *IosVppEBook) SetVppTokenId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)() {
    m.vppTokenId = value
}
