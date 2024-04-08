package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OrganizationalBrandingProperties 
type OrganizationalBrandingProperties struct {
    Entity
    // Color that will appear in place of the background image in low-bandwidth connections. We recommend that you use the primary color of your banner logo or your organization color. Specify this in hexadecimal format, for example, white is #FFFFFF.
    backgroundColor *string
    // Image that appears as the background of the sign-in page. The allowed types are PNG or JPEG not smaller than 300 KB and not larger than 1920 × 1080 pixels. A smaller image will reduce bandwidth requirements and make the page load faster.
    backgroundImage []byte
    // A relative URL for the backgroundImage property that is combined with a CDN base URL from the cdnList to provide the version served by a CDN. Read-only.
    backgroundImageRelativeUrl *string
    // A banner version of your company logo that appears on the sign-in page. The allowed types are PNG or JPEG no larger than 36 × 245 pixels. We recommend using a transparent image with no padding around the logo.
    bannerLogo []byte
    // A relative url for the bannerLogo property that is combined with a CDN base URL from the cdnList to provide the read-only version served by a CDN. Read-only.
    bannerLogoRelativeUrl *string
    // A list of base URLs for all available CDN providers that are serving the assets of the current resource. Several CDN providers are used at the same time for high availability of read requests. Read-only.
    cdnList []string
    // Text that appears at the bottom of the sign-in box. You can use this to communicate additional information, such as the phone number to your help desk or a legal statement. This text must be Unicode and not exceed 1024 characters.
    signInPageText *string
    // A square version of your company logo that appears in Windows 10 out-of-box experiences (OOBE) and when Windows Autopilot is enabled for deployment. Allowed types are PNG or JPEG no larger than 240 x 240 pixels and no more than 10 KB in size. We recommend using a transparent image with no padding around the logo.
    squareLogo []byte
    // A relative url for the squareLogo property that is combined with a CDN base URL from the cdnList to provide the version served by a CDN. Read-only.
    squareLogoRelativeUrl *string
    // String that shows as the hint in the username textbox on the sign-in screen. This text must be a Unicode, without links or code, and can't exceed 64 characters.
    usernameHintText *string
}
// NewOrganizationalBrandingProperties instantiates a new organizationalBrandingProperties and sets the default values.
func NewOrganizationalBrandingProperties()(*OrganizationalBrandingProperties) {
    m := &OrganizationalBrandingProperties{
        Entity: *NewEntity(),
    }
    return m
}
// CreateOrganizationalBrandingPropertiesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOrganizationalBrandingPropertiesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.organizationalBranding":
                        return NewOrganizationalBranding(), nil
                    case "#microsoft.graph.organizationalBrandingLocalization":
                        return NewOrganizationalBrandingLocalization(), nil
                }
            }
        }
    }
    return NewOrganizationalBrandingProperties(), nil
}
// GetBackgroundColor gets the backgroundColor property value. Color that will appear in place of the background image in low-bandwidth connections. We recommend that you use the primary color of your banner logo or your organization color. Specify this in hexadecimal format, for example, white is #FFFFFF.
func (m *OrganizationalBrandingProperties) GetBackgroundColor()(*string) {
    return m.backgroundColor
}
// GetBackgroundImage gets the backgroundImage property value. Image that appears as the background of the sign-in page. The allowed types are PNG or JPEG not smaller than 300 KB and not larger than 1920 × 1080 pixels. A smaller image will reduce bandwidth requirements and make the page load faster.
func (m *OrganizationalBrandingProperties) GetBackgroundImage()([]byte) {
    return m.backgroundImage
}
// GetBackgroundImageRelativeUrl gets the backgroundImageRelativeUrl property value. A relative URL for the backgroundImage property that is combined with a CDN base URL from the cdnList to provide the version served by a CDN. Read-only.
func (m *OrganizationalBrandingProperties) GetBackgroundImageRelativeUrl()(*string) {
    return m.backgroundImageRelativeUrl
}
// GetBannerLogo gets the bannerLogo property value. A banner version of your company logo that appears on the sign-in page. The allowed types are PNG or JPEG no larger than 36 × 245 pixels. We recommend using a transparent image with no padding around the logo.
func (m *OrganizationalBrandingProperties) GetBannerLogo()([]byte) {
    return m.bannerLogo
}
// GetBannerLogoRelativeUrl gets the bannerLogoRelativeUrl property value. A relative url for the bannerLogo property that is combined with a CDN base URL from the cdnList to provide the read-only version served by a CDN. Read-only.
func (m *OrganizationalBrandingProperties) GetBannerLogoRelativeUrl()(*string) {
    return m.bannerLogoRelativeUrl
}
// GetCdnList gets the cdnList property value. A list of base URLs for all available CDN providers that are serving the assets of the current resource. Several CDN providers are used at the same time for high availability of read requests. Read-only.
func (m *OrganizationalBrandingProperties) GetCdnList()([]string) {
    return m.cdnList
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OrganizationalBrandingProperties) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["backgroundColor"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBackgroundColor(val)
        }
        return nil
    }
    res["backgroundImage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBackgroundImage(val)
        }
        return nil
    }
    res["backgroundImageRelativeUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBackgroundImageRelativeUrl(val)
        }
        return nil
    }
    res["bannerLogo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBannerLogo(val)
        }
        return nil
    }
    res["bannerLogoRelativeUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBannerLogoRelativeUrl(val)
        }
        return nil
    }
    res["cdnList"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetCdnList(res)
        }
        return nil
    }
    res["signInPageText"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSignInPageText(val)
        }
        return nil
    }
    res["squareLogo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSquareLogo(val)
        }
        return nil
    }
    res["squareLogoRelativeUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSquareLogoRelativeUrl(val)
        }
        return nil
    }
    res["usernameHintText"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUsernameHintText(val)
        }
        return nil
    }
    return res
}
// GetSignInPageText gets the signInPageText property value. Text that appears at the bottom of the sign-in box. You can use this to communicate additional information, such as the phone number to your help desk or a legal statement. This text must be Unicode and not exceed 1024 characters.
func (m *OrganizationalBrandingProperties) GetSignInPageText()(*string) {
    return m.signInPageText
}
// GetSquareLogo gets the squareLogo property value. A square version of your company logo that appears in Windows 10 out-of-box experiences (OOBE) and when Windows Autopilot is enabled for deployment. Allowed types are PNG or JPEG no larger than 240 x 240 pixels and no more than 10 KB in size. We recommend using a transparent image with no padding around the logo.
func (m *OrganizationalBrandingProperties) GetSquareLogo()([]byte) {
    return m.squareLogo
}
// GetSquareLogoRelativeUrl gets the squareLogoRelativeUrl property value. A relative url for the squareLogo property that is combined with a CDN base URL from the cdnList to provide the version served by a CDN. Read-only.
func (m *OrganizationalBrandingProperties) GetSquareLogoRelativeUrl()(*string) {
    return m.squareLogoRelativeUrl
}
// GetUsernameHintText gets the usernameHintText property value. String that shows as the hint in the username textbox on the sign-in screen. This text must be a Unicode, without links or code, and can't exceed 64 characters.
func (m *OrganizationalBrandingProperties) GetUsernameHintText()(*string) {
    return m.usernameHintText
}
// Serialize serializes information the current object
func (m *OrganizationalBrandingProperties) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("backgroundColor", m.GetBackgroundColor())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteByteArrayValue("backgroundImage", m.GetBackgroundImage())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("backgroundImageRelativeUrl", m.GetBackgroundImageRelativeUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteByteArrayValue("bannerLogo", m.GetBannerLogo())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("bannerLogoRelativeUrl", m.GetBannerLogoRelativeUrl())
        if err != nil {
            return err
        }
    }
    if m.GetCdnList() != nil {
        err = writer.WriteCollectionOfStringValues("cdnList", m.GetCdnList())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("signInPageText", m.GetSignInPageText())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteByteArrayValue("squareLogo", m.GetSquareLogo())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("squareLogoRelativeUrl", m.GetSquareLogoRelativeUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("usernameHintText", m.GetUsernameHintText())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetBackgroundColor sets the backgroundColor property value. Color that will appear in place of the background image in low-bandwidth connections. We recommend that you use the primary color of your banner logo or your organization color. Specify this in hexadecimal format, for example, white is #FFFFFF.
func (m *OrganizationalBrandingProperties) SetBackgroundColor(value *string)() {
    m.backgroundColor = value
}
// SetBackgroundImage sets the backgroundImage property value. Image that appears as the background of the sign-in page. The allowed types are PNG or JPEG not smaller than 300 KB and not larger than 1920 × 1080 pixels. A smaller image will reduce bandwidth requirements and make the page load faster.
func (m *OrganizationalBrandingProperties) SetBackgroundImage(value []byte)() {
    m.backgroundImage = value
}
// SetBackgroundImageRelativeUrl sets the backgroundImageRelativeUrl property value. A relative URL for the backgroundImage property that is combined with a CDN base URL from the cdnList to provide the version served by a CDN. Read-only.
func (m *OrganizationalBrandingProperties) SetBackgroundImageRelativeUrl(value *string)() {
    m.backgroundImageRelativeUrl = value
}
// SetBannerLogo sets the bannerLogo property value. A banner version of your company logo that appears on the sign-in page. The allowed types are PNG or JPEG no larger than 36 × 245 pixels. We recommend using a transparent image with no padding around the logo.
func (m *OrganizationalBrandingProperties) SetBannerLogo(value []byte)() {
    m.bannerLogo = value
}
// SetBannerLogoRelativeUrl sets the bannerLogoRelativeUrl property value. A relative url for the bannerLogo property that is combined with a CDN base URL from the cdnList to provide the read-only version served by a CDN. Read-only.
func (m *OrganizationalBrandingProperties) SetBannerLogoRelativeUrl(value *string)() {
    m.bannerLogoRelativeUrl = value
}
// SetCdnList sets the cdnList property value. A list of base URLs for all available CDN providers that are serving the assets of the current resource. Several CDN providers are used at the same time for high availability of read requests. Read-only.
func (m *OrganizationalBrandingProperties) SetCdnList(value []string)() {
    m.cdnList = value
}
// SetSignInPageText sets the signInPageText property value. Text that appears at the bottom of the sign-in box. You can use this to communicate additional information, such as the phone number to your help desk or a legal statement. This text must be Unicode and not exceed 1024 characters.
func (m *OrganizationalBrandingProperties) SetSignInPageText(value *string)() {
    m.signInPageText = value
}
// SetSquareLogo sets the squareLogo property value. A square version of your company logo that appears in Windows 10 out-of-box experiences (OOBE) and when Windows Autopilot is enabled for deployment. Allowed types are PNG or JPEG no larger than 240 x 240 pixels and no more than 10 KB in size. We recommend using a transparent image with no padding around the logo.
func (m *OrganizationalBrandingProperties) SetSquareLogo(value []byte)() {
    m.squareLogo = value
}
// SetSquareLogoRelativeUrl sets the squareLogoRelativeUrl property value. A relative url for the squareLogo property that is combined with a CDN base URL from the cdnList to provide the version served by a CDN. Read-only.
func (m *OrganizationalBrandingProperties) SetSquareLogoRelativeUrl(value *string)() {
    m.squareLogoRelativeUrl = value
}
// SetUsernameHintText sets the usernameHintText property value. String that shows as the hint in the username textbox on the sign-in screen. This text must be a Unicode, without links or code, and can't exceed 64 characters.
func (m *OrganizationalBrandingProperties) SetUsernameHintText(value *string)() {
    m.usernameHintText = value
}
