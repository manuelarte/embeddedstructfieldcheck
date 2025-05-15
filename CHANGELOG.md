# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [v0.3.0] 2025-05-15

### Added

- Adding `forbid-mutex` setting.

## [v0.2.0] 2025-05-12

### Added

- Renaming project to `embeddedstructfieldcheck`.
- Supporting `SuggestedFix` for missing space between embedded fields and regular fields.

## [v0.1.0] 2025-04-24

### Added

- Added check that embedded types should be at the top of the field list of a struct.
- Added check that there must be an empty line separating embedded fields from regular fields.