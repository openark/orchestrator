package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ChangeNotificationEncryptedContentable 
type ChangeNotificationEncryptedContentable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetData()(*string)
    GetDataKey()(*string)
    GetDataSignature()(*string)
    GetEncryptionCertificateId()(*string)
    GetEncryptionCertificateThumbprint()(*string)
    GetOdataType()(*string)
    SetData(value *string)()
    SetDataKey(value *string)()
    SetDataSignature(value *string)()
    SetEncryptionCertificateId(value *string)()
    SetEncryptionCertificateThumbprint(value *string)()
    SetOdataType(value *string)()
}
