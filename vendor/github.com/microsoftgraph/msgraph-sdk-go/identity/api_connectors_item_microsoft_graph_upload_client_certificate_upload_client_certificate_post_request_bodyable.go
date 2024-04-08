package identity

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ApiConnectorsItemMicrosoftGraphUploadClientCertificateUploadClientCertificatePostRequestBodyable 
type ApiConnectorsItemMicrosoftGraphUploadClientCertificateUploadClientCertificatePostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetPassword()(*string)
    GetPkcs12Value()(*string)
    SetPassword(value *string)()
    SetPkcs12Value(value *string)()
}
