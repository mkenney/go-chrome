All notable changes to this project are documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/), and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

# v1.0.0-rc8 - 2019-06-21
#### Changed
* Enable graceful shutdown of socket listeners
* Update `godep` dependency versions
* Refactor process handling to be `context`-based


# v1.0.0-rc7 - 2019-04-12
#### Changed
* Fix an issue with taking the address of a pointer reference, see #132


# v1.0.0-rc6 - 2018-09-24
#### Added
* Enable handling socket errors
* Add initial error codes
* Example documentation
* Add useful docker-compose

#### Changed
* Cleanup logging code

#### Removed
- n/a


# v1.0.0-rc5 - 2018-09-15
#### Changed
* Fixes an issue with transitive dependencies


# v1.0.0-rc4 - 2018-08-22
#### Changed
* Replace logging package
* Various log cleanup


# v1.0.0-rc3 - 2018-07-23
#### Changed
* Updates related to upstream dependencies


# v1.0.0-rc2 - 2018-06-01
#### Added
* Add test coverage for socket timeouts

#### Changed
* Fix an issue with zombie listen process
* Fix an issue with zombie stop process


# v1.0.0-rc1 - 2018-05-30
#### Changed
* Initial refactoring
* Refactor test scripts
* Gocyclo
