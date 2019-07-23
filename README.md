Signaller
=========

*Simple Matrix server written in Go.*

Currently only memory backend are supported.

[![Build Status](https://travis-ci.com/nxshock/signaller.svg?branch=master)](https://travis-ci.com/nxshock/signaller)
[![Coverage Status](https://coveralls.io/repos/github/nxshock/signaller/badge.svg)](https://coveralls.io/github/nxshock/signaller)

Implemented from [Client-Server API](https://matrix.org/docs/spec/client_server/latest):

- [x] [5.4.1 GET /_matrix/client/r0/login](https://matrix.org/docs/spec/client_server/latest#get-matrix-client-r0-login)
- [x] [5.4.2 POST /_matrix/client/r0/login](https://matrix.org/docs/spec/client_server/latest#post-matrix-client-r0-login)
- [x] [5.4.3 POST /_matrix/client/r0/logout](https://matrix.org/docs/spec/client_server/latest#post-matrix-client-r0-logout)
- [x] [5.7.1 GET /_matrix/client/r0/account/whoami](https://matrix.org/docs/spec/client_server/latest#get-matrix-client-r0-account-whoami)