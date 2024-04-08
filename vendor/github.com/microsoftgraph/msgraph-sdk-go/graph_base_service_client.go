package msgraphsdkgo

import (
    i25911dc319edd61cbac496af7eab5ef20b6069a42515e22ec6a9bc97bf598488 "github.com/microsoft/kiota-serialization-json-go"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i4bcdc892e61ac17e2afc10b5e2b536b29f4fd6c1ad30f4a5a68df47495db3347 "github.com/microsoft/kiota-serialization-form-go"
    i7294a22093d408fdca300f11b81a887d89c47b764af06c8b803e2323973fdb83 "github.com/microsoft/kiota-serialization-text-go"
    i009f47bbce65ccdb7303730eed71e6bab3ae2f8e4e918bc9e94341d28624af97 "github.com/microsoftgraph/msgraph-sdk-go/serviceprincipals"
    i07d47a144340607d6d6dbd93575e531530e4f1cc6091c947ea0766f7951ffd34 "github.com/microsoftgraph/msgraph-sdk-go/shares"
    i0906e75d8a44bf92212e084e1d2f62d03887dcec6a5c8535e92ccc04c1e5fdec "github.com/microsoftgraph/msgraph-sdk-go/solutions"
    i185698f71f6301975f0627ee999e6e91920d8fa9c00bdef3487b9f349e2df04e "github.com/microsoftgraph/msgraph-sdk-go/directoryobjects"
    i1a1369b1521a8ac4885166fd68eae4247248a891006fea464d2eea2a271b2cdb "github.com/microsoftgraph/msgraph-sdk-go/permissiongrants"
    i1b75be7b5675627960b4672ab148be21ff379d5cbc0e62f6bc5b97d54464f8b5 "github.com/microsoftgraph/msgraph-sdk-go/teamstemplates"
    i1be0f1b1da466bc62355d411ef490acbd8dc0ec5ca4d3448c7eb73e5caffafc3 "github.com/microsoftgraph/msgraph-sdk-go/education"
    i1d6652ecc686b20c37a9a3448b26db8187e284e1a4017cab8876b02b97557436 "github.com/microsoftgraph/msgraph-sdk-go/grouplifecyclepolicies"
    i1dc06c4b7f499cb445a6c55e466abd6d7466bb35a2683c675909db23c57898e7 "github.com/microsoftgraph/msgraph-sdk-go/authenticationmethodconfigurations"
    i20b08d3949f1191430a14a315e0758a1f131dc59bbdc93e654f1dd447a6af14c "github.com/microsoftgraph/msgraph-sdk-go/auditlogs"
    i286f3babd79fe9ec3b0f52b6ed5910842c0adaeff02be1843d0e01c56d9ba6d9 "github.com/microsoftgraph/msgraph-sdk-go/search"
    i2a252d42835bdab6d88bf938595da6cf029001f9ca970d6f599cecf0ca27f8e5 "github.com/microsoftgraph/msgraph-sdk-go/directoryroletemplates"
    i32d45c1243c349600fbe53b2f9641bb59857a3326037587cbe4e347b46ad207e "github.com/microsoftgraph/msgraph-sdk-go/identitygovernance"
    i35d7bbcc8f7e8b8e9525ea0ee5b3c51c3a1a58f9ed512b727d181bfcd08eb032 "github.com/microsoftgraph/msgraph-sdk-go/security"
    i3e9b5129e2bb8b32b0374f7afe2536be6674d73df6c41d7c529f5a5432c4e0aa "github.com/microsoftgraph/msgraph-sdk-go/agreementacceptances"
    i4794c103c0d044c27a3ca3af0a0e498e93a9863420c1a4e7a29ef37590053c7b "github.com/microsoftgraph/msgraph-sdk-go/groupsettings"
    i4a624e38d68c2a9fc4db1ea915bcaffde116f967f58ec2c99e2ea8bbff3690e1 "github.com/microsoftgraph/msgraph-sdk-go/schemaextensions"
    i4ac7f0a844871066493521918f268cafe2a25c71c28a98221ea3f22d5153090f "github.com/microsoftgraph/msgraph-sdk-go/policies"
    i4c91eeb51f03f9d59a342065f7c6ee027ad1fe84ada6b1946b8162c5ae146cfb "github.com/microsoftgraph/msgraph-sdk-go/devices"
    i51b9802eedc1a25686534d117657be902df58c07e90ac6ea84501100998084d9 "github.com/microsoftgraph/msgraph-sdk-go/communications"
    i5310ba7d4cfddbf5de4c1be94a30f9ca8c747c30a87e76587ce88d1cbfff01b4 "github.com/microsoftgraph/msgraph-sdk-go/applicationtemplates"
    i535d6c02ba98f73ff3a8c1c12a035ba5de51606f93aa2c0babdfed56fe505550 "github.com/microsoftgraph/msgraph-sdk-go/certificatebasedauthconfiguration"
    i58857a108d6e260e56ef0dd7e783668388f113eb436006780703ac59f0abb3b1 "github.com/microsoftgraph/msgraph-sdk-go/privacy"
    i61686672307beee899fe5a14188df42982da47730f55a14800b102cd10ab2d72 "github.com/microsoftgraph/msgraph-sdk-go/localizations"
    i62c2771f3f3a1e5e085aedcde54473e9f043cc57b9ce4dd88980a77aca7a5a10 "github.com/microsoftgraph/msgraph-sdk-go/identityproviders"
    i638650494f9db477daff56d31ff923f5c100f72df0257ed7fa5c222cb1a77a94 "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement"
    i663c30678b300c2c4b619c4964b4326e471e4da61a44d7c39f752349da7a468e "github.com/microsoftgraph/msgraph-sdk-go/identityprotection"
    i6bf2d83eea06710580ad0d54b886ac4e14cbab0d1d84937f340f02b99f8f5738 "github.com/microsoftgraph/msgraph-sdk-go/reports"
    i71117da372286e863c042a526ec1361696ab14b838a5b77db5bc54386d436543 "github.com/microsoftgraph/msgraph-sdk-go/me"
    i738daeb889f22c1e163aee5a37a094b55b1d815dc76d4802d64e4e1b2e44206c "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement"
    i79ca23a9ac0659e1330dd29e049fe157787d5af6695ead2ff8263396db68d027 "github.com/microsoftgraph/msgraph-sdk-go/identity"
    i7c9d1b36ac198368c1d8bed014b43e2a518b170ee45bf02c8bbe64544a50539a "github.com/microsoftgraph/msgraph-sdk-go/admin"
    i7d140130aac6882792a019b5ebe51fe8d69dfd63ec213c2e3cd98282ce2d0428 "github.com/microsoftgraph/msgraph-sdk-go/appcatalogs"
    i80d5f91f6f8d9dc3428331303d1837675adde9653ceda73f120faa5f0545ac4b "github.com/microsoftgraph/msgraph-sdk-go/tenantrelationships"
    i86cada4d4a5f2f8a9d1e7a85eacd70a661ea7b20d2737008c0719e95b5be3e16 "github.com/microsoftgraph/msgraph-sdk-go/oauth2permissiongrants"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    i93194122344a685a2f9264205dc6d89a5ba39afdcea57fd0ade8f54b6f137c02 "github.com/microsoftgraph/msgraph-sdk-go/applications"
    i9429d7aae2f5c1dabbecc9411e8ad2b733d29338bc0c0436eeccc94605c461b7 "github.com/microsoftgraph/msgraph-sdk-go/print"
    i957076b10ba162b23efec7b94dd26b84c6475d285449c1cbc9c5b85910d36a12 "github.com/microsoftgraph/msgraph-sdk-go/domains"
    ia3e0f7c2d21d5c73ecb8a7552177d0fe444ae0522290dd1c4b5559e449b118af "github.com/microsoftgraph/msgraph-sdk-go/places"
    ia4b736f581ceef30e9ef8cebd9a6c2b932628e087982ff3dd2c9a0f1a920a918 "github.com/microsoftgraph/msgraph-sdk-go/compliance"
    ia6e876e3ed2d92c29c13dbc8c37513bc38d0d5f05ab9321e43a25ff336912a2d "github.com/microsoftgraph/msgraph-sdk-go/groups"
    iaca6694a878291d0e4021155b406c19d3080cdfc382b456e43c71264d4d9e519 "github.com/microsoftgraph/msgraph-sdk-go/domaindnsrecords"
    ib14d748b564c787931c10f1c7ba6856eeddea29a5b9e5c5c27eb1224ff65e5c4 "github.com/microsoftgraph/msgraph-sdk-go/directory"
    ib3217193884e00033cb8182cac52178dfa3b20ce9c4eb48e37a6217882d956ae "github.com/microsoftgraph/msgraph-sdk-go/external"
    ib33fc5e9889e020c0c572578957f59819123a589c61fd7f3eb37eb7958b525ee "github.com/microsoftgraph/msgraph-sdk-go/datapolicyoperations"
    ib68fa8e66bda853b3a33c491e8a66ca665897dab129192b2c97289266c4a1415 "github.com/microsoftgraph/msgraph-sdk-go/informationprotection"
    ibaef614e7692eebc6aaa8080b8ac29169fdf539f24925bc1de4465a3fcdac177 "github.com/microsoftgraph/msgraph-sdk-go/chats"
    ic5e701d75e87f15ce153687b00984a314f7eeea8cfdc77cd9ad648e5ccbc7fbd "github.com/microsoftgraph/msgraph-sdk-go/invitations"
    ic949a0bb5066d68760e8502a7f9db83f571d9e01e38fad4aadf7268188e52df0 "github.com/microsoftgraph/msgraph-sdk-go/organization"
    icabdee72951e77325f237b36d388a199c87e65f67652b6bb85723aba847d7e83 "github.com/microsoftgraph/msgraph-sdk-go/connections"
    ice10f31b9db59ba91184d2b882172edb754f885050cf0830aa2b7c8ff880556b "github.com/microsoftgraph/msgraph-sdk-go/scopedrolememberships"
    id007bc768abbff1131aab64890cdcd0411159a946e9df27140c5f7cf8f249647 "github.com/microsoftgraph/msgraph-sdk-go/subscribedskus"
    id2ac823944414906187dbe4e6ca3b5e46886b9db738d2c1c27de6df8b1bebd61 "github.com/microsoftgraph/msgraph-sdk-go/groupsettingtemplates"
    id4615a956cb1e7edabf8f5a4bc131d1ceca9a13d0f79ae0e122997452a9a0a4e "github.com/microsoftgraph/msgraph-sdk-go/directoryroles"
    id81f15a01b3ceaefa8b1b55f4ee944912f2179aafc4d873f0a2eaf0853eeccd0 "github.com/microsoftgraph/msgraph-sdk-go/authenticationmethodspolicy"
    idb8230b65f4a369c23b4d9b41ebe568c657c92f8f77fe36d16d64528b3a317a3 "github.com/microsoftgraph/msgraph-sdk-go/subscriptions"
    ie05ac24b652f7d895cca374316c093c4ca40dd2df0f1518c465233d6432b1ef9 "github.com/microsoftgraph/msgraph-sdk-go/teamwork"
    ie3631868038c44f490dbc03525ac7249d0523c29cc45cbb25b2aebcf470d6c0c "github.com/microsoftgraph/msgraph-sdk-go/contracts"
    ie66b913c1bc1c536bc8db5d185910e9318f621374e016f95e36e9d59b7127f63 "github.com/microsoftgraph/msgraph-sdk-go/planner"
    ieaa2790c8b7fa361674e69e4a385e279c8c641adf79d86e5b0ca566591a507e8 "github.com/microsoftgraph/msgraph-sdk-go/agreements"
    iefc72d8a17962d4db125c50866617eaa15d662c6e3fb13735d477380dcc0dbe3 "github.com/microsoftgraph/msgraph-sdk-go/drives"
    if398f5c2f1cb53106e045240edd469d82f1854899fd95cfdf8f559b19375750c "github.com/microsoftgraph/msgraph-sdk-go/branding"
    if39bc788926a05e976b265ecfc616408ca12af399df9ce3a2bb348fe89708057 "github.com/microsoftgraph/msgraph-sdk-go/teams"
    if51cca2652371587dbc02e65260e291435a6a8f7f2ffb419f26c3b9d2a033f57 "github.com/microsoftgraph/msgraph-sdk-go/contacts"
    if5372351befdb652f617b1ee71fbf092fa8dd2a161ba9c021bc265628b6ea82b "github.com/microsoftgraph/msgraph-sdk-go/sites"
    if5555fa41b6637688bcf8c25c62a041258f4dc6eacb38ad42d91c66f222ee182 "github.com/microsoftgraph/msgraph-sdk-go/rolemanagement"
    if6ffd1464db2d9c22e351b03e4c00ebd24a5353cd70ffb7f56cfad1c3ceec329 "github.com/microsoftgraph/msgraph-sdk-go/users"
)

// GraphBaseServiceClient the main entry point of the SDK, exposes the configuration and the fluent API.
type GraphBaseServiceClient struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// Admin provides operations to manage the admin singleton.
func (m *GraphBaseServiceClient) Admin()(*i7c9d1b36ac198368c1d8bed014b43e2a518b170ee45bf02c8bbe64544a50539a.AdminRequestBuilder) {
    return i7c9d1b36ac198368c1d8bed014b43e2a518b170ee45bf02c8bbe64544a50539a.NewAdminRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// AgreementAcceptances provides operations to manage the collection of agreementAcceptance entities.
func (m *GraphBaseServiceClient) AgreementAcceptances()(*i3e9b5129e2bb8b32b0374f7afe2536be6674d73df6c41d7c529f5a5432c4e0aa.AgreementAcceptancesRequestBuilder) {
    return i3e9b5129e2bb8b32b0374f7afe2536be6674d73df6c41d7c529f5a5432c4e0aa.NewAgreementAcceptancesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// AgreementAcceptancesById provides operations to manage the collection of agreementAcceptance entities.
func (m *GraphBaseServiceClient) AgreementAcceptancesById(id string)(*i3e9b5129e2bb8b32b0374f7afe2536be6674d73df6c41d7c529f5a5432c4e0aa.AgreementAcceptanceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["agreementAcceptance%2Did"] = id
    }
    return i3e9b5129e2bb8b32b0374f7afe2536be6674d73df6c41d7c529f5a5432c4e0aa.NewAgreementAcceptanceItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Agreements provides operations to manage the collection of agreement entities.
func (m *GraphBaseServiceClient) Agreements()(*ieaa2790c8b7fa361674e69e4a385e279c8c641adf79d86e5b0ca566591a507e8.AgreementsRequestBuilder) {
    return ieaa2790c8b7fa361674e69e4a385e279c8c641adf79d86e5b0ca566591a507e8.NewAgreementsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// AgreementsById provides operations to manage the collection of agreement entities.
func (m *GraphBaseServiceClient) AgreementsById(id string)(*ieaa2790c8b7fa361674e69e4a385e279c8c641adf79d86e5b0ca566591a507e8.AgreementItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["agreement%2Did"] = id
    }
    return ieaa2790c8b7fa361674e69e4a385e279c8c641adf79d86e5b0ca566591a507e8.NewAgreementItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// AppCatalogs provides operations to manage the appCatalogs singleton.
func (m *GraphBaseServiceClient) AppCatalogs()(*i7d140130aac6882792a019b5ebe51fe8d69dfd63ec213c2e3cd98282ce2d0428.AppCatalogsRequestBuilder) {
    return i7d140130aac6882792a019b5ebe51fe8d69dfd63ec213c2e3cd98282ce2d0428.NewAppCatalogsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Applications provides operations to manage the collection of application entities.
func (m *GraphBaseServiceClient) Applications()(*i93194122344a685a2f9264205dc6d89a5ba39afdcea57fd0ade8f54b6f137c02.ApplicationsRequestBuilder) {
    return i93194122344a685a2f9264205dc6d89a5ba39afdcea57fd0ade8f54b6f137c02.NewApplicationsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ApplicationsById provides operations to manage the collection of application entities.
func (m *GraphBaseServiceClient) ApplicationsById(id string)(*i93194122344a685a2f9264205dc6d89a5ba39afdcea57fd0ade8f54b6f137c02.ApplicationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["application%2Did"] = id
    }
    return i93194122344a685a2f9264205dc6d89a5ba39afdcea57fd0ade8f54b6f137c02.NewApplicationItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// ApplicationTemplates provides operations to manage the collection of applicationTemplate entities.
func (m *GraphBaseServiceClient) ApplicationTemplates()(*i5310ba7d4cfddbf5de4c1be94a30f9ca8c747c30a87e76587ce88d1cbfff01b4.ApplicationTemplatesRequestBuilder) {
    return i5310ba7d4cfddbf5de4c1be94a30f9ca8c747c30a87e76587ce88d1cbfff01b4.NewApplicationTemplatesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ApplicationTemplatesById provides operations to manage the collection of applicationTemplate entities.
func (m *GraphBaseServiceClient) ApplicationTemplatesById(id string)(*i5310ba7d4cfddbf5de4c1be94a30f9ca8c747c30a87e76587ce88d1cbfff01b4.ApplicationTemplateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["applicationTemplate%2Did"] = id
    }
    return i5310ba7d4cfddbf5de4c1be94a30f9ca8c747c30a87e76587ce88d1cbfff01b4.NewApplicationTemplateItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// AuditLogs provides operations to manage the auditLogRoot singleton.
func (m *GraphBaseServiceClient) AuditLogs()(*i20b08d3949f1191430a14a315e0758a1f131dc59bbdc93e654f1dd447a6af14c.AuditLogsRequestBuilder) {
    return i20b08d3949f1191430a14a315e0758a1f131dc59bbdc93e654f1dd447a6af14c.NewAuditLogsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// AuthenticationMethodConfigurations provides operations to manage the collection of authenticationMethodConfiguration entities.
func (m *GraphBaseServiceClient) AuthenticationMethodConfigurations()(*i1dc06c4b7f499cb445a6c55e466abd6d7466bb35a2683c675909db23c57898e7.AuthenticationMethodConfigurationsRequestBuilder) {
    return i1dc06c4b7f499cb445a6c55e466abd6d7466bb35a2683c675909db23c57898e7.NewAuthenticationMethodConfigurationsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// AuthenticationMethodConfigurationsById provides operations to manage the collection of authenticationMethodConfiguration entities.
func (m *GraphBaseServiceClient) AuthenticationMethodConfigurationsById(id string)(*i1dc06c4b7f499cb445a6c55e466abd6d7466bb35a2683c675909db23c57898e7.AuthenticationMethodConfigurationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["authenticationMethodConfiguration%2Did"] = id
    }
    return i1dc06c4b7f499cb445a6c55e466abd6d7466bb35a2683c675909db23c57898e7.NewAuthenticationMethodConfigurationItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// AuthenticationMethodsPolicy provides operations to manage the authenticationMethodsPolicy singleton.
func (m *GraphBaseServiceClient) AuthenticationMethodsPolicy()(*id81f15a01b3ceaefa8b1b55f4ee944912f2179aafc4d873f0a2eaf0853eeccd0.AuthenticationMethodsPolicyRequestBuilder) {
    return id81f15a01b3ceaefa8b1b55f4ee944912f2179aafc4d873f0a2eaf0853eeccd0.NewAuthenticationMethodsPolicyRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Branding provides operations to manage the organizationalBranding singleton.
func (m *GraphBaseServiceClient) Branding()(*if398f5c2f1cb53106e045240edd469d82f1854899fd95cfdf8f559b19375750c.BrandingRequestBuilder) {
    return if398f5c2f1cb53106e045240edd469d82f1854899fd95cfdf8f559b19375750c.NewBrandingRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// CertificateBasedAuthConfiguration provides operations to manage the collection of certificateBasedAuthConfiguration entities.
func (m *GraphBaseServiceClient) CertificateBasedAuthConfiguration()(*i535d6c02ba98f73ff3a8c1c12a035ba5de51606f93aa2c0babdfed56fe505550.CertificateBasedAuthConfigurationRequestBuilder) {
    return i535d6c02ba98f73ff3a8c1c12a035ba5de51606f93aa2c0babdfed56fe505550.NewCertificateBasedAuthConfigurationRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// CertificateBasedAuthConfigurationById provides operations to manage the collection of certificateBasedAuthConfiguration entities.
func (m *GraphBaseServiceClient) CertificateBasedAuthConfigurationById(id string)(*i535d6c02ba98f73ff3a8c1c12a035ba5de51606f93aa2c0babdfed56fe505550.CertificateBasedAuthConfigurationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["certificateBasedAuthConfiguration%2Did"] = id
    }
    return i535d6c02ba98f73ff3a8c1c12a035ba5de51606f93aa2c0babdfed56fe505550.NewCertificateBasedAuthConfigurationItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Chats provides operations to manage the collection of chat entities.
func (m *GraphBaseServiceClient) Chats()(*ibaef614e7692eebc6aaa8080b8ac29169fdf539f24925bc1de4465a3fcdac177.ChatsRequestBuilder) {
    return ibaef614e7692eebc6aaa8080b8ac29169fdf539f24925bc1de4465a3fcdac177.NewChatsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ChatsById provides operations to manage the collection of chat entities.
func (m *GraphBaseServiceClient) ChatsById(id string)(*ibaef614e7692eebc6aaa8080b8ac29169fdf539f24925bc1de4465a3fcdac177.ChatItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["chat%2Did"] = id
    }
    return ibaef614e7692eebc6aaa8080b8ac29169fdf539f24925bc1de4465a3fcdac177.NewChatItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Communications provides operations to manage the cloudCommunications singleton.
func (m *GraphBaseServiceClient) Communications()(*i51b9802eedc1a25686534d117657be902df58c07e90ac6ea84501100998084d9.CommunicationsRequestBuilder) {
    return i51b9802eedc1a25686534d117657be902df58c07e90ac6ea84501100998084d9.NewCommunicationsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Compliance provides operations to manage the compliance singleton.
func (m *GraphBaseServiceClient) Compliance()(*ia4b736f581ceef30e9ef8cebd9a6c2b932628e087982ff3dd2c9a0f1a920a918.ComplianceRequestBuilder) {
    return ia4b736f581ceef30e9ef8cebd9a6c2b932628e087982ff3dd2c9a0f1a920a918.NewComplianceRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Connections provides operations to manage the collection of externalConnection entities.
func (m *GraphBaseServiceClient) Connections()(*icabdee72951e77325f237b36d388a199c87e65f67652b6bb85723aba847d7e83.ConnectionsRequestBuilder) {
    return icabdee72951e77325f237b36d388a199c87e65f67652b6bb85723aba847d7e83.NewConnectionsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ConnectionsById provides operations to manage the collection of externalConnection entities.
func (m *GraphBaseServiceClient) ConnectionsById(id string)(*icabdee72951e77325f237b36d388a199c87e65f67652b6bb85723aba847d7e83.ExternalConnectionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["externalConnection%2Did"] = id
    }
    return icabdee72951e77325f237b36d388a199c87e65f67652b6bb85723aba847d7e83.NewExternalConnectionItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// NewGraphBaseServiceClient instantiates a new GraphBaseServiceClient and sets the default values.
func NewGraphBaseServiceClient(requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GraphBaseServiceClient) {
    m := &GraphBaseServiceClient{
    }
    m.pathParameters = make(map[string]string);
    m.urlTemplate = "{+baseurl}";
    m.requestAdapter = requestAdapter
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultSerializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriterFactory { return i25911dc319edd61cbac496af7eab5ef20b6069a42515e22ec6a9bc97bf598488.NewJsonSerializationWriterFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultSerializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriterFactory { return i7294a22093d408fdca300f11b81a887d89c47b764af06c8b803e2323973fdb83.NewTextSerializationWriterFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultSerializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriterFactory { return i4bcdc892e61ac17e2afc10b5e2b536b29f4fd6c1ad30f4a5a68df47495db3347.NewFormSerializationWriterFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultDeserializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNodeFactory { return i25911dc319edd61cbac496af7eab5ef20b6069a42515e22ec6a9bc97bf598488.NewJsonParseNodeFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultDeserializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNodeFactory { return i7294a22093d408fdca300f11b81a887d89c47b764af06c8b803e2323973fdb83.NewTextParseNodeFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultDeserializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNodeFactory { return i4bcdc892e61ac17e2afc10b5e2b536b29f4fd6c1ad30f4a5a68df47495db3347.NewFormParseNodeFactory() })
    if m.requestAdapter.GetBaseUrl() == "" {
        m.requestAdapter.SetBaseUrl("https://graph.microsoft.com/v1.0")
    }
    m.pathParameters["baseurl"] = m.requestAdapter.GetBaseUrl()
    return m
}
// Contacts provides operations to manage the collection of orgContact entities.
func (m *GraphBaseServiceClient) Contacts()(*if51cca2652371587dbc02e65260e291435a6a8f7f2ffb419f26c3b9d2a033f57.ContactsRequestBuilder) {
    return if51cca2652371587dbc02e65260e291435a6a8f7f2ffb419f26c3b9d2a033f57.NewContactsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ContactsById provides operations to manage the collection of orgContact entities.
func (m *GraphBaseServiceClient) ContactsById(id string)(*if51cca2652371587dbc02e65260e291435a6a8f7f2ffb419f26c3b9d2a033f57.OrgContactItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["orgContact%2Did"] = id
    }
    return if51cca2652371587dbc02e65260e291435a6a8f7f2ffb419f26c3b9d2a033f57.NewOrgContactItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Contracts provides operations to manage the collection of contract entities.
func (m *GraphBaseServiceClient) Contracts()(*ie3631868038c44f490dbc03525ac7249d0523c29cc45cbb25b2aebcf470d6c0c.ContractsRequestBuilder) {
    return ie3631868038c44f490dbc03525ac7249d0523c29cc45cbb25b2aebcf470d6c0c.NewContractsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ContractsById provides operations to manage the collection of contract entities.
func (m *GraphBaseServiceClient) ContractsById(id string)(*ie3631868038c44f490dbc03525ac7249d0523c29cc45cbb25b2aebcf470d6c0c.ContractItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["contract%2Did"] = id
    }
    return ie3631868038c44f490dbc03525ac7249d0523c29cc45cbb25b2aebcf470d6c0c.NewContractItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// DataPolicyOperations provides operations to manage the collection of dataPolicyOperation entities.
func (m *GraphBaseServiceClient) DataPolicyOperations()(*ib33fc5e9889e020c0c572578957f59819123a589c61fd7f3eb37eb7958b525ee.DataPolicyOperationsRequestBuilder) {
    return ib33fc5e9889e020c0c572578957f59819123a589c61fd7f3eb37eb7958b525ee.NewDataPolicyOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// DataPolicyOperationsById provides operations to manage the collection of dataPolicyOperation entities.
func (m *GraphBaseServiceClient) DataPolicyOperationsById(id string)(*ib33fc5e9889e020c0c572578957f59819123a589c61fd7f3eb37eb7958b525ee.DataPolicyOperationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["dataPolicyOperation%2Did"] = id
    }
    return ib33fc5e9889e020c0c572578957f59819123a589c61fd7f3eb37eb7958b525ee.NewDataPolicyOperationItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// DeviceAppManagement provides operations to manage the deviceAppManagement singleton.
func (m *GraphBaseServiceClient) DeviceAppManagement()(*i638650494f9db477daff56d31ff923f5c100f72df0257ed7fa5c222cb1a77a94.DeviceAppManagementRequestBuilder) {
    return i638650494f9db477daff56d31ff923f5c100f72df0257ed7fa5c222cb1a77a94.NewDeviceAppManagementRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// DeviceManagement provides operations to manage the deviceManagement singleton.
func (m *GraphBaseServiceClient) DeviceManagement()(*i738daeb889f22c1e163aee5a37a094b55b1d815dc76d4802d64e4e1b2e44206c.DeviceManagementRequestBuilder) {
    return i738daeb889f22c1e163aee5a37a094b55b1d815dc76d4802d64e4e1b2e44206c.NewDeviceManagementRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Devices provides operations to manage the collection of device entities.
func (m *GraphBaseServiceClient) Devices()(*i4c91eeb51f03f9d59a342065f7c6ee027ad1fe84ada6b1946b8162c5ae146cfb.DevicesRequestBuilder) {
    return i4c91eeb51f03f9d59a342065f7c6ee027ad1fe84ada6b1946b8162c5ae146cfb.NewDevicesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// DevicesById provides operations to manage the collection of device entities.
func (m *GraphBaseServiceClient) DevicesById(id string)(*i4c91eeb51f03f9d59a342065f7c6ee027ad1fe84ada6b1946b8162c5ae146cfb.DeviceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["device%2Did"] = id
    }
    return i4c91eeb51f03f9d59a342065f7c6ee027ad1fe84ada6b1946b8162c5ae146cfb.NewDeviceItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Directory provides operations to manage the directory singleton.
func (m *GraphBaseServiceClient) Directory()(*ib14d748b564c787931c10f1c7ba6856eeddea29a5b9e5c5c27eb1224ff65e5c4.DirectoryRequestBuilder) {
    return ib14d748b564c787931c10f1c7ba6856eeddea29a5b9e5c5c27eb1224ff65e5c4.NewDirectoryRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// DirectoryObjects provides operations to manage the collection of directoryObject entities.
func (m *GraphBaseServiceClient) DirectoryObjects()(*i185698f71f6301975f0627ee999e6e91920d8fa9c00bdef3487b9f349e2df04e.DirectoryObjectsRequestBuilder) {
    return i185698f71f6301975f0627ee999e6e91920d8fa9c00bdef3487b9f349e2df04e.NewDirectoryObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// DirectoryObjectsById provides operations to manage the collection of directoryObject entities.
func (m *GraphBaseServiceClient) DirectoryObjectsById(id string)(*i185698f71f6301975f0627ee999e6e91920d8fa9c00bdef3487b9f349e2df04e.DirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return i185698f71f6301975f0627ee999e6e91920d8fa9c00bdef3487b9f349e2df04e.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// DirectoryRoles provides operations to manage the collection of directoryRole entities.
func (m *GraphBaseServiceClient) DirectoryRoles()(*id4615a956cb1e7edabf8f5a4bc131d1ceca9a13d0f79ae0e122997452a9a0a4e.DirectoryRolesRequestBuilder) {
    return id4615a956cb1e7edabf8f5a4bc131d1ceca9a13d0f79ae0e122997452a9a0a4e.NewDirectoryRolesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// DirectoryRolesById provides operations to manage the collection of directoryRole entities.
func (m *GraphBaseServiceClient) DirectoryRolesById(id string)(*id4615a956cb1e7edabf8f5a4bc131d1ceca9a13d0f79ae0e122997452a9a0a4e.DirectoryRoleItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryRole%2Did"] = id
    }
    return id4615a956cb1e7edabf8f5a4bc131d1ceca9a13d0f79ae0e122997452a9a0a4e.NewDirectoryRoleItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// DirectoryRoleTemplates provides operations to manage the collection of directoryRoleTemplate entities.
func (m *GraphBaseServiceClient) DirectoryRoleTemplates()(*i2a252d42835bdab6d88bf938595da6cf029001f9ca970d6f599cecf0ca27f8e5.DirectoryRoleTemplatesRequestBuilder) {
    return i2a252d42835bdab6d88bf938595da6cf029001f9ca970d6f599cecf0ca27f8e5.NewDirectoryRoleTemplatesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// DirectoryRoleTemplatesById provides operations to manage the collection of directoryRoleTemplate entities.
func (m *GraphBaseServiceClient) DirectoryRoleTemplatesById(id string)(*i2a252d42835bdab6d88bf938595da6cf029001f9ca970d6f599cecf0ca27f8e5.DirectoryRoleTemplateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryRoleTemplate%2Did"] = id
    }
    return i2a252d42835bdab6d88bf938595da6cf029001f9ca970d6f599cecf0ca27f8e5.NewDirectoryRoleTemplateItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// DomainDnsRecords provides operations to manage the collection of domainDnsRecord entities.
func (m *GraphBaseServiceClient) DomainDnsRecords()(*iaca6694a878291d0e4021155b406c19d3080cdfc382b456e43c71264d4d9e519.DomainDnsRecordsRequestBuilder) {
    return iaca6694a878291d0e4021155b406c19d3080cdfc382b456e43c71264d4d9e519.NewDomainDnsRecordsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// DomainDnsRecordsById provides operations to manage the collection of domainDnsRecord entities.
func (m *GraphBaseServiceClient) DomainDnsRecordsById(id string)(*iaca6694a878291d0e4021155b406c19d3080cdfc382b456e43c71264d4d9e519.DomainDnsRecordItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["domainDnsRecord%2Did"] = id
    }
    return iaca6694a878291d0e4021155b406c19d3080cdfc382b456e43c71264d4d9e519.NewDomainDnsRecordItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Domains provides operations to manage the collection of domain entities.
func (m *GraphBaseServiceClient) Domains()(*i957076b10ba162b23efec7b94dd26b84c6475d285449c1cbc9c5b85910d36a12.DomainsRequestBuilder) {
    return i957076b10ba162b23efec7b94dd26b84c6475d285449c1cbc9c5b85910d36a12.NewDomainsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// DomainsById provides operations to manage the collection of domain entities.
func (m *GraphBaseServiceClient) DomainsById(id string)(*i957076b10ba162b23efec7b94dd26b84c6475d285449c1cbc9c5b85910d36a12.DomainItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["domain%2Did"] = id
    }
    return i957076b10ba162b23efec7b94dd26b84c6475d285449c1cbc9c5b85910d36a12.NewDomainItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Drives provides operations to manage the collection of drive entities.
func (m *GraphBaseServiceClient) Drives()(*iefc72d8a17962d4db125c50866617eaa15d662c6e3fb13735d477380dcc0dbe3.DrivesRequestBuilder) {
    return iefc72d8a17962d4db125c50866617eaa15d662c6e3fb13735d477380dcc0dbe3.NewDrivesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// DrivesById provides operations to manage the collection of drive entities.
func (m *GraphBaseServiceClient) DrivesById(id string)(*iefc72d8a17962d4db125c50866617eaa15d662c6e3fb13735d477380dcc0dbe3.DriveItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["drive%2Did"] = id
    }
    return iefc72d8a17962d4db125c50866617eaa15d662c6e3fb13735d477380dcc0dbe3.NewDriveItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Education provides operations to manage the educationRoot singleton.
func (m *GraphBaseServiceClient) Education()(*i1be0f1b1da466bc62355d411ef490acbd8dc0ec5ca4d3448c7eb73e5caffafc3.EducationRequestBuilder) {
    return i1be0f1b1da466bc62355d411ef490acbd8dc0ec5ca4d3448c7eb73e5caffafc3.NewEducationRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// External provides operations to manage the external singleton.
func (m *GraphBaseServiceClient) External()(*ib3217193884e00033cb8182cac52178dfa3b20ce9c4eb48e37a6217882d956ae.ExternalRequestBuilder) {
    return ib3217193884e00033cb8182cac52178dfa3b20ce9c4eb48e37a6217882d956ae.NewExternalRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// GroupLifecyclePolicies provides operations to manage the collection of groupLifecyclePolicy entities.
func (m *GraphBaseServiceClient) GroupLifecyclePolicies()(*i1d6652ecc686b20c37a9a3448b26db8187e284e1a4017cab8876b02b97557436.GroupLifecyclePoliciesRequestBuilder) {
    return i1d6652ecc686b20c37a9a3448b26db8187e284e1a4017cab8876b02b97557436.NewGroupLifecyclePoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// GroupLifecyclePoliciesById provides operations to manage the collection of groupLifecyclePolicy entities.
func (m *GraphBaseServiceClient) GroupLifecyclePoliciesById(id string)(*i1d6652ecc686b20c37a9a3448b26db8187e284e1a4017cab8876b02b97557436.GroupLifecyclePolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["groupLifecyclePolicy%2Did"] = id
    }
    return i1d6652ecc686b20c37a9a3448b26db8187e284e1a4017cab8876b02b97557436.NewGroupLifecyclePolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Groups provides operations to manage the collection of group entities.
func (m *GraphBaseServiceClient) Groups()(*ia6e876e3ed2d92c29c13dbc8c37513bc38d0d5f05ab9321e43a25ff336912a2d.GroupsRequestBuilder) {
    return ia6e876e3ed2d92c29c13dbc8c37513bc38d0d5f05ab9321e43a25ff336912a2d.NewGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// GroupsById provides operations to manage the collection of group entities.
func (m *GraphBaseServiceClient) GroupsById(id string)(*ia6e876e3ed2d92c29c13dbc8c37513bc38d0d5f05ab9321e43a25ff336912a2d.GroupItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["group%2Did"] = id
    }
    return ia6e876e3ed2d92c29c13dbc8c37513bc38d0d5f05ab9321e43a25ff336912a2d.NewGroupItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// GroupSettings provides operations to manage the collection of groupSetting entities.
func (m *GraphBaseServiceClient) GroupSettings()(*i4794c103c0d044c27a3ca3af0a0e498e93a9863420c1a4e7a29ef37590053c7b.GroupSettingsRequestBuilder) {
    return i4794c103c0d044c27a3ca3af0a0e498e93a9863420c1a4e7a29ef37590053c7b.NewGroupSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// GroupSettingsById provides operations to manage the collection of groupSetting entities.
func (m *GraphBaseServiceClient) GroupSettingsById(id string)(*i4794c103c0d044c27a3ca3af0a0e498e93a9863420c1a4e7a29ef37590053c7b.GroupSettingItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["groupSetting%2Did"] = id
    }
    return i4794c103c0d044c27a3ca3af0a0e498e93a9863420c1a4e7a29ef37590053c7b.NewGroupSettingItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// GroupSettingTemplates provides operations to manage the collection of groupSettingTemplate entities.
func (m *GraphBaseServiceClient) GroupSettingTemplates()(*id2ac823944414906187dbe4e6ca3b5e46886b9db738d2c1c27de6df8b1bebd61.GroupSettingTemplatesRequestBuilder) {
    return id2ac823944414906187dbe4e6ca3b5e46886b9db738d2c1c27de6df8b1bebd61.NewGroupSettingTemplatesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// GroupSettingTemplatesById provides operations to manage the collection of groupSettingTemplate entities.
func (m *GraphBaseServiceClient) GroupSettingTemplatesById(id string)(*id2ac823944414906187dbe4e6ca3b5e46886b9db738d2c1c27de6df8b1bebd61.GroupSettingTemplateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["groupSettingTemplate%2Did"] = id
    }
    return id2ac823944414906187dbe4e6ca3b5e46886b9db738d2c1c27de6df8b1bebd61.NewGroupSettingTemplateItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Identity provides operations to manage the identityContainer singleton.
func (m *GraphBaseServiceClient) Identity()(*i79ca23a9ac0659e1330dd29e049fe157787d5af6695ead2ff8263396db68d027.IdentityRequestBuilder) {
    return i79ca23a9ac0659e1330dd29e049fe157787d5af6695ead2ff8263396db68d027.NewIdentityRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// IdentityGovernance provides operations to manage the identityGovernance singleton.
func (m *GraphBaseServiceClient) IdentityGovernance()(*i32d45c1243c349600fbe53b2f9641bb59857a3326037587cbe4e347b46ad207e.IdentityGovernanceRequestBuilder) {
    return i32d45c1243c349600fbe53b2f9641bb59857a3326037587cbe4e347b46ad207e.NewIdentityGovernanceRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// IdentityProtection provides operations to manage the identityProtectionRoot singleton.
func (m *GraphBaseServiceClient) IdentityProtection()(*i663c30678b300c2c4b619c4964b4326e471e4da61a44d7c39f752349da7a468e.IdentityProtectionRequestBuilder) {
    return i663c30678b300c2c4b619c4964b4326e471e4da61a44d7c39f752349da7a468e.NewIdentityProtectionRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// IdentityProviders provides operations to manage the collection of identityProvider entities.
func (m *GraphBaseServiceClient) IdentityProviders()(*i62c2771f3f3a1e5e085aedcde54473e9f043cc57b9ce4dd88980a77aca7a5a10.IdentityProvidersRequestBuilder) {
    return i62c2771f3f3a1e5e085aedcde54473e9f043cc57b9ce4dd88980a77aca7a5a10.NewIdentityProvidersRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// IdentityProvidersById provides operations to manage the collection of identityProvider entities.
func (m *GraphBaseServiceClient) IdentityProvidersById(id string)(*i62c2771f3f3a1e5e085aedcde54473e9f043cc57b9ce4dd88980a77aca7a5a10.IdentityProviderItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["identityProvider%2Did"] = id
    }
    return i62c2771f3f3a1e5e085aedcde54473e9f043cc57b9ce4dd88980a77aca7a5a10.NewIdentityProviderItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// InformationProtection provides operations to manage the informationProtection singleton.
func (m *GraphBaseServiceClient) InformationProtection()(*ib68fa8e66bda853b3a33c491e8a66ca665897dab129192b2c97289266c4a1415.InformationProtectionRequestBuilder) {
    return ib68fa8e66bda853b3a33c491e8a66ca665897dab129192b2c97289266c4a1415.NewInformationProtectionRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Invitations provides operations to manage the collection of invitation entities.
func (m *GraphBaseServiceClient) Invitations()(*ic5e701d75e87f15ce153687b00984a314f7eeea8cfdc77cd9ad648e5ccbc7fbd.InvitationsRequestBuilder) {
    return ic5e701d75e87f15ce153687b00984a314f7eeea8cfdc77cd9ad648e5ccbc7fbd.NewInvitationsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// InvitationsById provides operations to manage the collection of invitation entities.
func (m *GraphBaseServiceClient) InvitationsById(id string)(*ic5e701d75e87f15ce153687b00984a314f7eeea8cfdc77cd9ad648e5ccbc7fbd.InvitationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["invitation%2Did"] = id
    }
    return ic5e701d75e87f15ce153687b00984a314f7eeea8cfdc77cd9ad648e5ccbc7fbd.NewInvitationItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Localizations provides operations to manage the collection of organizationalBrandingLocalization entities.
func (m *GraphBaseServiceClient) Localizations()(*i61686672307beee899fe5a14188df42982da47730f55a14800b102cd10ab2d72.LocalizationsRequestBuilder) {
    return i61686672307beee899fe5a14188df42982da47730f55a14800b102cd10ab2d72.NewLocalizationsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// LocalizationsById provides operations to manage the collection of organizationalBrandingLocalization entities.
func (m *GraphBaseServiceClient) LocalizationsById(id string)(*i61686672307beee899fe5a14188df42982da47730f55a14800b102cd10ab2d72.OrganizationalBrandingLocalizationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["organizationalBrandingLocalization%2Did"] = id
    }
    return i61686672307beee899fe5a14188df42982da47730f55a14800b102cd10ab2d72.NewOrganizationalBrandingLocalizationItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Me provides operations to manage the user singleton.
func (m *GraphBaseServiceClient) Me()(*i71117da372286e863c042a526ec1361696ab14b838a5b77db5bc54386d436543.MeRequestBuilder) {
    return i71117da372286e863c042a526ec1361696ab14b838a5b77db5bc54386d436543.NewMeRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Oauth2PermissionGrants provides operations to manage the collection of oAuth2PermissionGrant entities.
func (m *GraphBaseServiceClient) Oauth2PermissionGrants()(*i86cada4d4a5f2f8a9d1e7a85eacd70a661ea7b20d2737008c0719e95b5be3e16.Oauth2PermissionGrantsRequestBuilder) {
    return i86cada4d4a5f2f8a9d1e7a85eacd70a661ea7b20d2737008c0719e95b5be3e16.NewOauth2PermissionGrantsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Oauth2PermissionGrantsById provides operations to manage the collection of oAuth2PermissionGrant entities.
func (m *GraphBaseServiceClient) Oauth2PermissionGrantsById(id string)(*i86cada4d4a5f2f8a9d1e7a85eacd70a661ea7b20d2737008c0719e95b5be3e16.OAuth2PermissionGrantItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["oAuth2PermissionGrant%2Did"] = id
    }
    return i86cada4d4a5f2f8a9d1e7a85eacd70a661ea7b20d2737008c0719e95b5be3e16.NewOAuth2PermissionGrantItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Organization provides operations to manage the collection of organization entities.
func (m *GraphBaseServiceClient) Organization()(*ic949a0bb5066d68760e8502a7f9db83f571d9e01e38fad4aadf7268188e52df0.OrganizationRequestBuilder) {
    return ic949a0bb5066d68760e8502a7f9db83f571d9e01e38fad4aadf7268188e52df0.NewOrganizationRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// OrganizationById provides operations to manage the collection of organization entities.
func (m *GraphBaseServiceClient) OrganizationById(id string)(*ic949a0bb5066d68760e8502a7f9db83f571d9e01e38fad4aadf7268188e52df0.OrganizationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["organization%2Did"] = id
    }
    return ic949a0bb5066d68760e8502a7f9db83f571d9e01e38fad4aadf7268188e52df0.NewOrganizationItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// PermissionGrants provides operations to manage the collection of resourceSpecificPermissionGrant entities.
func (m *GraphBaseServiceClient) PermissionGrants()(*i1a1369b1521a8ac4885166fd68eae4247248a891006fea464d2eea2a271b2cdb.PermissionGrantsRequestBuilder) {
    return i1a1369b1521a8ac4885166fd68eae4247248a891006fea464d2eea2a271b2cdb.NewPermissionGrantsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// PermissionGrantsById provides operations to manage the collection of resourceSpecificPermissionGrant entities.
func (m *GraphBaseServiceClient) PermissionGrantsById(id string)(*i1a1369b1521a8ac4885166fd68eae4247248a891006fea464d2eea2a271b2cdb.ResourceSpecificPermissionGrantItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["resourceSpecificPermissionGrant%2Did"] = id
    }
    return i1a1369b1521a8ac4885166fd68eae4247248a891006fea464d2eea2a271b2cdb.NewResourceSpecificPermissionGrantItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Places provides operations to manage the collection of place entities.
func (m *GraphBaseServiceClient) Places()(*ia3e0f7c2d21d5c73ecb8a7552177d0fe444ae0522290dd1c4b5559e449b118af.PlacesRequestBuilder) {
    return ia3e0f7c2d21d5c73ecb8a7552177d0fe444ae0522290dd1c4b5559e449b118af.NewPlacesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// PlacesById provides operations to manage the collection of place entities.
func (m *GraphBaseServiceClient) PlacesById(id string)(*ia3e0f7c2d21d5c73ecb8a7552177d0fe444ae0522290dd1c4b5559e449b118af.PlaceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["place%2Did"] = id
    }
    return ia3e0f7c2d21d5c73ecb8a7552177d0fe444ae0522290dd1c4b5559e449b118af.NewPlaceItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Planner provides operations to manage the planner singleton.
func (m *GraphBaseServiceClient) Planner()(*ie66b913c1bc1c536bc8db5d185910e9318f621374e016f95e36e9d59b7127f63.PlannerRequestBuilder) {
    return ie66b913c1bc1c536bc8db5d185910e9318f621374e016f95e36e9d59b7127f63.NewPlannerRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Policies provides operations to manage the policyRoot singleton.
func (m *GraphBaseServiceClient) Policies()(*i4ac7f0a844871066493521918f268cafe2a25c71c28a98221ea3f22d5153090f.PoliciesRequestBuilder) {
    return i4ac7f0a844871066493521918f268cafe2a25c71c28a98221ea3f22d5153090f.NewPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Print provides operations to manage the print singleton.
func (m *GraphBaseServiceClient) Print()(*i9429d7aae2f5c1dabbecc9411e8ad2b733d29338bc0c0436eeccc94605c461b7.PrintRequestBuilder) {
    return i9429d7aae2f5c1dabbecc9411e8ad2b733d29338bc0c0436eeccc94605c461b7.NewPrintRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Privacy provides operations to manage the privacy singleton.
func (m *GraphBaseServiceClient) Privacy()(*i58857a108d6e260e56ef0dd7e783668388f113eb436006780703ac59f0abb3b1.PrivacyRequestBuilder) {
    return i58857a108d6e260e56ef0dd7e783668388f113eb436006780703ac59f0abb3b1.NewPrivacyRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Reports provides operations to manage the reportRoot singleton.
func (m *GraphBaseServiceClient) Reports()(*i6bf2d83eea06710580ad0d54b886ac4e14cbab0d1d84937f340f02b99f8f5738.ReportsRequestBuilder) {
    return i6bf2d83eea06710580ad0d54b886ac4e14cbab0d1d84937f340f02b99f8f5738.NewReportsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// RoleManagement provides operations to manage the roleManagement singleton.
func (m *GraphBaseServiceClient) RoleManagement()(*if5555fa41b6637688bcf8c25c62a041258f4dc6eacb38ad42d91c66f222ee182.RoleManagementRequestBuilder) {
    return if5555fa41b6637688bcf8c25c62a041258f4dc6eacb38ad42d91c66f222ee182.NewRoleManagementRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// SchemaExtensions provides operations to manage the collection of schemaExtension entities.
func (m *GraphBaseServiceClient) SchemaExtensions()(*i4a624e38d68c2a9fc4db1ea915bcaffde116f967f58ec2c99e2ea8bbff3690e1.SchemaExtensionsRequestBuilder) {
    return i4a624e38d68c2a9fc4db1ea915bcaffde116f967f58ec2c99e2ea8bbff3690e1.NewSchemaExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// SchemaExtensionsById provides operations to manage the collection of schemaExtension entities.
func (m *GraphBaseServiceClient) SchemaExtensionsById(id string)(*i4a624e38d68c2a9fc4db1ea915bcaffde116f967f58ec2c99e2ea8bbff3690e1.SchemaExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["schemaExtension%2Did"] = id
    }
    return i4a624e38d68c2a9fc4db1ea915bcaffde116f967f58ec2c99e2ea8bbff3690e1.NewSchemaExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// ScopedRoleMemberships provides operations to manage the collection of scopedRoleMembership entities.
func (m *GraphBaseServiceClient) ScopedRoleMemberships()(*ice10f31b9db59ba91184d2b882172edb754f885050cf0830aa2b7c8ff880556b.ScopedRoleMembershipsRequestBuilder) {
    return ice10f31b9db59ba91184d2b882172edb754f885050cf0830aa2b7c8ff880556b.NewScopedRoleMembershipsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ScopedRoleMembershipsById provides operations to manage the collection of scopedRoleMembership entities.
func (m *GraphBaseServiceClient) ScopedRoleMembershipsById(id string)(*ice10f31b9db59ba91184d2b882172edb754f885050cf0830aa2b7c8ff880556b.ScopedRoleMembershipItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["scopedRoleMembership%2Did"] = id
    }
    return ice10f31b9db59ba91184d2b882172edb754f885050cf0830aa2b7c8ff880556b.NewScopedRoleMembershipItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Search provides operations to manage the searchEntity singleton.
func (m *GraphBaseServiceClient) Search()(*i286f3babd79fe9ec3b0f52b6ed5910842c0adaeff02be1843d0e01c56d9ba6d9.SearchRequestBuilder) {
    return i286f3babd79fe9ec3b0f52b6ed5910842c0adaeff02be1843d0e01c56d9ba6d9.NewSearchRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Security provides operations to manage the security singleton.
func (m *GraphBaseServiceClient) Security()(*i35d7bbcc8f7e8b8e9525ea0ee5b3c51c3a1a58f9ed512b727d181bfcd08eb032.SecurityRequestBuilder) {
    return i35d7bbcc8f7e8b8e9525ea0ee5b3c51c3a1a58f9ed512b727d181bfcd08eb032.NewSecurityRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ServicePrincipals provides operations to manage the collection of servicePrincipal entities.
func (m *GraphBaseServiceClient) ServicePrincipals()(*i009f47bbce65ccdb7303730eed71e6bab3ae2f8e4e918bc9e94341d28624af97.ServicePrincipalsRequestBuilder) {
    return i009f47bbce65ccdb7303730eed71e6bab3ae2f8e4e918bc9e94341d28624af97.NewServicePrincipalsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ServicePrincipalsById provides operations to manage the collection of servicePrincipal entities.
func (m *GraphBaseServiceClient) ServicePrincipalsById(id string)(*i009f47bbce65ccdb7303730eed71e6bab3ae2f8e4e918bc9e94341d28624af97.ServicePrincipalItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["servicePrincipal%2Did"] = id
    }
    return i009f47bbce65ccdb7303730eed71e6bab3ae2f8e4e918bc9e94341d28624af97.NewServicePrincipalItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Shares provides operations to manage the collection of sharedDriveItem entities.
func (m *GraphBaseServiceClient) Shares()(*i07d47a144340607d6d6dbd93575e531530e4f1cc6091c947ea0766f7951ffd34.SharesRequestBuilder) {
    return i07d47a144340607d6d6dbd93575e531530e4f1cc6091c947ea0766f7951ffd34.NewSharesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// SharesById provides operations to manage the collection of sharedDriveItem entities.
func (m *GraphBaseServiceClient) SharesById(id string)(*i07d47a144340607d6d6dbd93575e531530e4f1cc6091c947ea0766f7951ffd34.SharedDriveItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["sharedDriveItem%2Did"] = id
    }
    return i07d47a144340607d6d6dbd93575e531530e4f1cc6091c947ea0766f7951ffd34.NewSharedDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Sites provides operations to manage the collection of site entities.
func (m *GraphBaseServiceClient) Sites()(*if5372351befdb652f617b1ee71fbf092fa8dd2a161ba9c021bc265628b6ea82b.SitesRequestBuilder) {
    return if5372351befdb652f617b1ee71fbf092fa8dd2a161ba9c021bc265628b6ea82b.NewSitesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// SitesById provides operations to manage the collection of site entities.
func (m *GraphBaseServiceClient) SitesById(id string)(*if5372351befdb652f617b1ee71fbf092fa8dd2a161ba9c021bc265628b6ea82b.SiteItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["site%2Did"] = id
    }
    return if5372351befdb652f617b1ee71fbf092fa8dd2a161ba9c021bc265628b6ea82b.NewSiteItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Solutions provides operations to manage the solutionsRoot singleton.
func (m *GraphBaseServiceClient) Solutions()(*i0906e75d8a44bf92212e084e1d2f62d03887dcec6a5c8535e92ccc04c1e5fdec.SolutionsRequestBuilder) {
    return i0906e75d8a44bf92212e084e1d2f62d03887dcec6a5c8535e92ccc04c1e5fdec.NewSolutionsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// SubscribedSkus provides operations to manage the collection of subscribedSku entities.
func (m *GraphBaseServiceClient) SubscribedSkus()(*id007bc768abbff1131aab64890cdcd0411159a946e9df27140c5f7cf8f249647.SubscribedSkusRequestBuilder) {
    return id007bc768abbff1131aab64890cdcd0411159a946e9df27140c5f7cf8f249647.NewSubscribedSkusRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// SubscribedSkusById provides operations to manage the collection of subscribedSku entities.
func (m *GraphBaseServiceClient) SubscribedSkusById(id string)(*id007bc768abbff1131aab64890cdcd0411159a946e9df27140c5f7cf8f249647.SubscribedSkuItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["subscribedSku%2Did"] = id
    }
    return id007bc768abbff1131aab64890cdcd0411159a946e9df27140c5f7cf8f249647.NewSubscribedSkuItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Subscriptions provides operations to manage the collection of subscription entities.
func (m *GraphBaseServiceClient) Subscriptions()(*idb8230b65f4a369c23b4d9b41ebe568c657c92f8f77fe36d16d64528b3a317a3.SubscriptionsRequestBuilder) {
    return idb8230b65f4a369c23b4d9b41ebe568c657c92f8f77fe36d16d64528b3a317a3.NewSubscriptionsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// SubscriptionsById provides operations to manage the collection of subscription entities.
func (m *GraphBaseServiceClient) SubscriptionsById(id string)(*idb8230b65f4a369c23b4d9b41ebe568c657c92f8f77fe36d16d64528b3a317a3.SubscriptionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["subscription%2Did"] = id
    }
    return idb8230b65f4a369c23b4d9b41ebe568c657c92f8f77fe36d16d64528b3a317a3.NewSubscriptionItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Teams provides operations to manage the collection of team entities.
func (m *GraphBaseServiceClient) Teams()(*if39bc788926a05e976b265ecfc616408ca12af399df9ce3a2bb348fe89708057.TeamsRequestBuilder) {
    return if39bc788926a05e976b265ecfc616408ca12af399df9ce3a2bb348fe89708057.NewTeamsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// TeamsById provides operations to manage the collection of team entities.
func (m *GraphBaseServiceClient) TeamsById(id string)(*if39bc788926a05e976b265ecfc616408ca12af399df9ce3a2bb348fe89708057.TeamItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["team%2Did"] = id
    }
    return if39bc788926a05e976b265ecfc616408ca12af399df9ce3a2bb348fe89708057.NewTeamItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// TeamsTemplates provides operations to manage the collection of teamsTemplate entities.
func (m *GraphBaseServiceClient) TeamsTemplates()(*i1b75be7b5675627960b4672ab148be21ff379d5cbc0e62f6bc5b97d54464f8b5.TeamsTemplatesRequestBuilder) {
    return i1b75be7b5675627960b4672ab148be21ff379d5cbc0e62f6bc5b97d54464f8b5.NewTeamsTemplatesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// TeamsTemplatesById provides operations to manage the collection of teamsTemplate entities.
func (m *GraphBaseServiceClient) TeamsTemplatesById(id string)(*i1b75be7b5675627960b4672ab148be21ff379d5cbc0e62f6bc5b97d54464f8b5.TeamsTemplateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamsTemplate%2Did"] = id
    }
    return i1b75be7b5675627960b4672ab148be21ff379d5cbc0e62f6bc5b97d54464f8b5.NewTeamsTemplateItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Teamwork provides operations to manage the teamwork singleton.
func (m *GraphBaseServiceClient) Teamwork()(*ie05ac24b652f7d895cca374316c093c4ca40dd2df0f1518c465233d6432b1ef9.TeamworkRequestBuilder) {
    return ie05ac24b652f7d895cca374316c093c4ca40dd2df0f1518c465233d6432b1ef9.NewTeamworkRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// TenantRelationships provides operations to manage the tenantRelationship singleton.
func (m *GraphBaseServiceClient) TenantRelationships()(*i80d5f91f6f8d9dc3428331303d1837675adde9653ceda73f120faa5f0545ac4b.TenantRelationshipsRequestBuilder) {
    return i80d5f91f6f8d9dc3428331303d1837675adde9653ceda73f120faa5f0545ac4b.NewTenantRelationshipsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Users provides operations to manage the collection of user entities.
func (m *GraphBaseServiceClient) Users()(*if6ffd1464db2d9c22e351b03e4c00ebd24a5353cd70ffb7f56cfad1c3ceec329.UsersRequestBuilder) {
    return if6ffd1464db2d9c22e351b03e4c00ebd24a5353cd70ffb7f56cfad1c3ceec329.NewUsersRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// UsersById provides operations to manage the collection of user entities.
func (m *GraphBaseServiceClient) UsersById(id string)(*if6ffd1464db2d9c22e351b03e4c00ebd24a5353cd70ffb7f56cfad1c3ceec329.UserItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["user%2Did"] = id
    }
    return if6ffd1464db2d9c22e351b03e4c00ebd24a5353cd70ffb7f56cfad1c3ceec329.NewUserItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
