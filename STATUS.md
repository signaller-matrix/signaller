# API implementation status

Implemented from [Client-Server API](https://matrix.org/docs/spec/client_server/latest):

## [2 API Standards](https://matrix.org/docs/spec/client_server/latest#api-standards)

- [x] [2.1 GET /_matrix/client/versions](https://matrix.org/docs/spec/client_server/latest#get-matrix-client-versions)

## [4 Server Discovery](https://matrix.org/docs/spec/client_server/latest#server-discovery)

### [4.1 Well-known URI](https://matrix.org/docs/spec/client_server/latest#well-known-uri)

- [x] [4.1.1 GET /.well-known/matrix/client](https://matrix.org/docs/spec/client_server/latest#get-well-known-matrix-client)

## [5 Client Authentication](https://matrix.org/docs/spec/client_server/latest#client-authentication)

### [5.4 Login](https://matrix.org/docs/spec/client_server/latest#login)

- [x] [5.4.1 GET /_matrix/client/r0/login](https://matrix.org/docs/spec/client_server/latest#get-matrix-client-r0-login)
- [x] [5.4.2 POST /_matrix/client/r0/login](https://matrix.org/docs/spec/client_server/latest#post-matrix-client-r0-login)
- [x] [5.4.3 POST /_matrix/client/r0/logout](https://matrix.org/docs/spec/client_server/latest#post-matrix-client-r0-logout)
- [x] [5.4.4 POST /_matrix/client/r0/logout/all](https://matrix.org/docs/spec/client_server/latest#post-matrix-client-r0-logout-all)

### [5.5 Account registration and management](https://matrix.org/docs/spec/client_server/latest#account-registration-and-management)

- [x] [5.5.1 POST /_matrix/client/r0/register](https://matrix.org/docs/spec/client_server/r0.5.0#post-matrix-client-r0-register)
- [ ] [5.5.2 POST /_matrix/client/r0/register/email/requestToken](https://matrix.org/docs/spec/client_server/latest#post-matrix-client-r0-register-email-requesttoken)
- [ ] [5.5.3 POST /_matrix/client/r0/register/msisdn/requestToken](https://matrix.org/docs/spec/client_server/latest#post-matrix-client-r0-register-msisdn-requesttoken)
- [x] [5.5.4 POST /_matrix/client/r0/account/password](https://matrix.org/docs/spec/client_server/latest#post-matrix-client-r0-account-password)
- [ ] [5.5.5 POST /_matrix/client/r0/account/password/email/requestToken](https://matrix.org/docs/spec/client_server/latest#post-matrix-client-r0-account-password-email-requesttoken)
- [ ] [5.5.6 POST /_matrix/client/r0/account/password/msisdn/requestToken](https://matrix.org/docs/spec/client_server/latest#post-matrix-client-r0-account-password-msisdn-requesttoken)
- [ ] [5.5.7 POST /_matrix/client/r0/account/deactivate](https://matrix.org/docs/spec/client_server/latest#post-matrix-client-r0-account-deactivate)
- [x] [5.5.8 GET /_matrix/client/r0/register/available](https://matrix.org/docs/spec/client_server/latest#get-matrix-client-r0-register-available)

### [5.6 Adding Account Administrative Contact Information](https://matrix.org/docs/spec/client_server/latest#adding-account-administrative-contact-information)

- [ ] [5.6.1 GET /_matrix/client/r0/account/3pid](https://matrix.org/docs/spec/client_server/latest#get-matrix-client-r0-account-3pid)
- [ ] [5.6.2 POST /_matrix/client/r0/account/3pid](https://matrix.org/docs/spec/client_server/latest#post-matrix-client-r0-account-3pid)
- [ ] [5.6.3 POST /_matrix/client/r0/account/3pid/delete](https://matrix.org/docs/spec/client_server/latest#post-matrix-client-r0-account-3pid-delete)
- [ ] [5.6.4 POST /_matrix/client/r0/account/3pid/email/requestToken](https://matrix.org/docs/spec/client_server/latest#post-matrix-client-r0-account-3pid-email-requesttoken)
- [ ] [5.6.5 POST /_matrix/client/r0/account/3pid/msisdn/requestToken](https://matrix.org/docs/spec/client_server/latest#post-matrix-client-r0-account-3pid-msisdn-requesttoken)

### [5.7 Current account information](https://matrix.org/docs/spec/client_server/latest#current-account-information)

- [x] [5.7.1 GET /_matrix/client/r0/account/whoami](https://matrix.org/docs/spec/client_server/latest#get-matrix-client-r0-account-whoami)

## [6 Capabilities negotiation](https://matrix.org/docs/spec/client_server/latest#capabilities-negotiation)

- [x] [6.1 GET /_matrix/client/r0/capabilities](https://matrix.org/docs/spec/client_server/latest#get-matrix-client-r0-capabilities)

## [8 Filtering](https://matrix.org/docs/spec/client_server/latest#filtering)

### [8.2 API endpoints](https://matrix.org/docs/spec/client_server/latest#api-endpoints)

- [ ] [8.2.1 POST /_matrix/client/r0/user/{userId}/filter](https://matrix.org/docs/spec/client_server/latest#post-matrix-client-r0-user-userid-filter)
- [ ] [8.2.2 GET /_matrix/client/r0/user/{userId}/filter/{filterId}](https://matrix.org/docs/spec/client_server/latest#get-matrix-client-r0-user-userid-filter-filterid)

### [9.4 Syncing](https://matrix.org/docs/spec/client_server/latest#syncing)

- [ ] [9.4.1 GET /_matrix/client/r0/sync](https://matrix.org/docs/spec/client_server/latest#get-matrix-client-r0-sync)
- [ ] ~~[9.4.2 GET /_matrix/client/r0/events DEPRECATED](https://matrix.org/docs/spec/client_server/latest#get-matrix-client-r0-events)~~
- [ ] ~~[9.4.3 GET /_matrix/client/r0/initialSync DEPRECATED](https://matrix.org/docs/spec/client_server/latest#get-matrix-client-r0-initialsync)~~
- [ ] ~~[9.4.4 GET /_matrix/client/r0/events/{eventId} DEPRECATED](https://matrix.org/docs/spec/client_server/latest#get-matrix-client-r0-events-eventid)~~

### [9.5 Getting events for a room](https://matrix.org/docs/spec/client_server/latest#getting-events-for-a-room)

- [ ] [9.5.1 GET /_matrix/client/r0/rooms/{roomId}/event/{eventId}](https://matrix.org/docs/spec/client_server/latest#get-matrix-client-r0-rooms-roomid-event-eventid)
- [ ] [9.5.2 GET /_matrix/client/r0/rooms/{roomId}/state/{eventType}/{stateKey}](https://matrix.org/docs/spec/client_server/latest#get-matrix-client-r0-rooms-roomid-state-eventtype-statekey)
- [ ] [9.5.3 GET /_matrix/client/r0/rooms/{roomId}/state](https://matrix.org/docs/spec/client_server/latest#get-matrix-client-r0-rooms-roomid-state)
- [ ] [9.5.4 GET /_matrix/client/r0/rooms/{roomId}/members](https://matrix.org/docs/spec/client_server/latest#get-matrix-client-r0-rooms-roomid-members)
- [ ] [9.5.5 GET /_matrix/client/r0/rooms/{roomId}/joined_members](https://matrix.org/docs/spec/client_server/latest#get-matrix-client-r0-rooms-roomid-joined-members)
- [ ] [9.5.6 GET /_matrix/client/r0/rooms/{roomId}/messages](https://matrix.org/docs/spec/client_server/latest#get-matrix-client-r0-rooms-roomid-messages)
- [ ] ~~[9.5.7 GET /_matrix/client/r0/rooms/{roomId}/initialSync DEPRECATED](https://matrix.org/docs/spec/client_server/latest#get-matrix-client-r0-rooms-roomid-initialsync)~~

### [9.6 Sending events to a room](https://matrix.org/docs/spec/client_server/latest#sending-events-to-a-room)

- [ ] [9.6.1 PUT /_matrix/client/r0/rooms/{roomId}/state/{eventType}/{stateKey}](https://matrix.org/docs/spec/client_server/latest#put-matrix-client-r0-rooms-roomid-state-eventtype-statekey)
- [ ] [9.6.2 PUT /_matrix/client/r0/rooms/{roomId}/send/{eventType}/{txnId}](https://matrix.org/docs/spec/client_server/latest#put-matrix-client-r0-rooms-roomid-send-eventtype-txnid)

## [9.7 Redactions](https://matrix.org/docs/spec/client_server/latest#redactions)

### [9.7.2 Client behaviour](https://matrix.org/docs/spec/client_server/latest#client-behaviour)

- [ ] [9.7.2.1 PUT /_matrix/client/r0/rooms/{roomId}/redact/{eventId}/{txnId}](https://matrix.org/docs/spec/client_server/latest#put-matrix-client-r0-rooms-roomid-redact-eventid-txnid)

### [10.1 Creation](https://matrix.org/docs/spec/client_server/latest#creation)

- [x] [10.1.1 POST /_matrix/client/r0/createRoom](https://matrix.org/docs/spec/client_server/latest#post-matrix-client-r0-createroom)

### [10.2 Room aliases](https://matrix.org/docs/spec/client_server/latest#room-aliases)

- [ ] [10.2.1 PUT /_matrix/client/r0/directory/room/{roomAlias}](https://matrix.org/docs/spec/client_server/latest#put-matrix-client-r0-directory-room-roomalias)
- [ ] [10.2.2 GET /_matrix/client/r0/directory/room/{roomAlias}](https://matrix.org/docs/spec/client_server/latest#get-matrix-client-r0-directory-room-roomalias)
- [ ] [10.2.3 DELETE /_matrix/client/r0/directory/room/{roomAlias}](https://matrix.org/docs/spec/client_server/latest#delete-matrix-client-r0-directory-room-roomalias)

### [10.4 Room membership](https://matrix.org/docs/spec/client_server/latest#room-membership)

- [x] [10.4.1 GET /_matrix/client/r0/joined_rooms](https://matrix.org/docs/spec/client_server/latest#get-matrix-client-r0-joined-rooms)

- [ ] **[10.4.2.1 POST /_matrix/client/r0/rooms/{roomId}/invite](https://matrix.org/docs/spec/client_server/latest#post-matrix-client-r0-rooms-roomid-invite)**
- [ ] [10.4.2.2 POST /_matrix/client/r0/rooms/{roomId}/join](https://matrix.org/docs/spec/client_server/latest#post-matrix-client-r0-rooms-roomid-join)
- [ ] [10.4.2.3 POST /_matrix/client/r0/join/{roomIdOrAlias}](https://matrix.org/docs/spec/client_server/latest#post-matrix-client-r0-join-roomidoralias)

- [x] [10.4.3.1 POST /_matrix/client/r0/rooms/{roomId}/leave](https://matrix.org/docs/spec/client_server/latest#post-matrix-client-r0-rooms-roomid-leave)
- [ ] [10.4.3.2 POST /_matrix/client/r0/rooms/{roomId}/forget](https://matrix.org/docs/spec/client_server/latest#post-matrix-client-r0-rooms-roomid-forget)
- [ ] [10.4.3.3 POST /_matrix/client/r0/rooms/{roomId}/kick](https://matrix.org/docs/spec/client_server/latest#post-matrix-client-r0-rooms-roomid-kick)

- [ ] [10.4.4.1 POST /_matrix/client/r0/rooms/{roomId}/ban](https://matrix.org/docs/spec/client_server/latest#post-matrix-client-r0-rooms-roomid-ban)
- [ ] [10.4.4.2 POST /_matrix/client/r0/rooms/{roomId}/unban](https://matrix.org/docs/spec/client_server/latest#post-matrix-client-r0-rooms-roomid-unban)

### [10.5 Listing rooms](https://matrix.org/docs/spec/client_server/latest#listing-rooms)

- [x] [10.5.1 GET /_matrix/client/r0/directory/list/room/{roomId}](https://matrix.org/docs/spec/client_server/latest#get-matrix-client-r0-directory-list-room-roomid)
- [x] [10.5.2 PUT /_matrix/client/r0/directory/list/room/{roomId}](https://matrix.org/docs/spec/client_server/latest#put-matrix-client-r0-directory-list-room-roomid)

- [x] [10.5.3 GET /_matrix/client/r0/publicRooms](https://matrix.org/docs/spec/client_server/latest#get-matrix-client-r0-publicrooms)
- [x] [10.5.4 POST /_matrix/client/r0/publicRooms](https://matrix.org/docs/spec/client_server/latest#post-matrix-client-r0-publicrooms)

## [11 User Data](https://matrix.org/docs/spec/client_server/latest#user-data)

### [11.1 User Directory](https://matrix.org/docs/spec/client_server/latest#user-directory)

- [ ] [11.1.1 POST /_matrix/client/r0/user_directory/search](https://matrix.org/docs/spec/client_server/latest#post-matrix-client-r0-user-directory-search)

### [11.2 Profiles](https://matrix.org/docs/spec/client_server/latest#profiles)

- [ ] [11.2.1 PUT /_matrix/client/r0/profile/{userId}/displayname](https://matrix.org/docs/spec/client_server/latest#put-matrix-client-r0-profile-userid-displayname)
- [ ] [11.2.2 GET /_matrix/client/r0/profile/{userId}/displayname](https://matrix.org/docs/spec/client_server/latest#get-matrix-client-r0-profile-userid-displayname)
- [ ] [11.2.3 PUT /_matrix/client/r0/profile/{userId}/avatar_url](https://matrix.org/docs/spec/client_server/latest#put-matrix-client-r0-profile-userid-avatar-url)
- [ ] [11.2.4 GET /_matrix/client/r0/profile/{userId}/avatar_url](https://matrix.org/docs/spec/client_server/latest#get-matrix-client-r0-profile-userid-avatar-url)
- [ ] [11.2.5 GET /_matrix/client/r0/profile/{userId}](https://matrix.org/docs/spec/client_server/latest#get-matrix-client-r0-profile-userid)

## [13.3 Voice over IP](https://matrix.org/docs/spec/client_server/latest#voice-over-ip)

## [13.4 Typing Notifications](https://matrix.org/docs/spec/client_server/latest#id93)

## [13.5 Receipts](https://matrix.org/docs/spec/client_server/latest#id97)

## [13.6 Fully read markers](https://matrix.org/docs/spec/client_server/latest#id102)

## [13.7 Presence](https://matrix.org/docs/spec/client_server/latest#id106)

## [13.8 Content repository](https://matrix.org/docs/spec/client_server/latest#id110)

## [13.9 Send-to-Device messaging](https://matrix.org/docs/spec/client_server/latest#id114)

- [x] [13.10.1.1 GET /_matrix/client/r0/devices](https://matrix.org/docs/spec/client_server/latest#get-matrix-client-r0-devices)
- [ ] [13.10.1.2 GET /_matrix/client/r0/devices/{deviceId}](https://matrix.org/docs/spec/client_server/latest#get-matrix-client-r0-devices-deviceid)
- [ ] [13.10.1.3 PUT /_matrix/client/r0/devices/{deviceId}](https://matrix.org/docs/spec/client_server/latest#put-matrix-client-r0-devices-deviceid)
- [ ] [13.10.1.4 DELETE /_matrix/client/r0/devices/{deviceId}](https://matrix.org/docs/spec/client_server/latest#delete-matrix-client-r0-devices-deviceid)
- [ ] [13.10.1.5 POST /_matrix/client/r0/delete_devices](https://matrix.org/docs/spec/client_server/latest#post-matrix-client-r0-delete-devices)

## [13.11 End-to-End Encryption](https://matrix.org/docs/spec/client_server/latest#id120)

## [13.12 Room History Visibility](https://matrix.org/docs/spec/client_server/latest#room-history-visibility)

## [13.13 Push Notifications](https://matrix.org/docs/spec/client_server/latest#id134)

## [13.14 Third party invites](https://matrix.org/docs/spec/client_server/latest#third-party-invites)

## [13.15 Server Side Search](https://matrix.org/docs/spec/client_server/latest#id149)

## [13.16 Guest Access](https://matrix.org/docs/spec/client_server/latest#guest-access)

## [13.17 Room Previews](https://matrix.org/docs/spec/client_server/latest#id161)

## [13.18 Room Tagging](https://matrix.org/docs/spec/client_server/latest#room-tagging)

## [13.19 Client Config](https://matrix.org/docs/spec/client_server/latest#id171)

## [13.20 Server Administration](https://matrix.org/docs/spec/client_server/latest#id175)

## [13.21 Event Context](https://matrix.org/docs/spec/client_server/latest#id177)

## [13.22 SSO client login](https://matrix.org/docs/spec/client_server/latest#sso-client-login)

## [13.26 Reporting Content](https://matrix.org/docs/spec/client_server/latest#id195)

## [13.27 Third Party Networks](https://matrix.org/docs/spec/client_server/latest#id198)

## [13.28 OpenID](https://matrix.org/docs/spec/client_server/latest#id199)

## [13.31 Room Upgrades](https://matrix.org/docs/spec/client_server/latest#id205)
