# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [0.0.6] - 2022-05-14
### Changed
- Add args `timeout`, `flaky`, `size`, `exit` to `openapi_test`
- Explicitly kill background webserver in `openapi_test`

## [0.0.5] - 2022-04-20
### Added
- Added rule `openapi_test`
- Added `validate` tool

### Changed
- Renamed `CHANGELOG` to `CHANGELOG.md`

## [0.0.4] - 2022-04-18
### Added
- `openapi_generator` arg `deps`
- `openapi_generator` label `codegen`

### Changed
- `openapi_generator` args `spec` and `config` are now build targets
- Update example usage in README
- Update generator example to remove srcs

### Removed
- `openapi_generator` arg `srcs`

## [0.0.3] - 2022-04-10
### Changed
- `please-js` update to v0.0.3

## [0.0.2] - 2022-04-04
### Added
- `openapi_bundle` rule
- `openapi_docs` rule
- `openapi_preview` rule
- `openapi_lint` rule
- lint with spectral
- lint with redocly

## [0.0.1] - 2022-04-01
### Added
- `openapi_generate` build rule
- example using `openapi_generate`
- This CHANGELOG
- LICENSE
- README with usage examples

[Unreleased]: https://github.com/andrew-womeldorf/please-openapi-rules/compare/v0.0.5...HEAD
[0.0.5]: https://github.com/andrew-womeldorf/please-openapi-rules/compare/v0.0.4...v0.0.5
[0.0.4]: https://github.com/andrew-womeldorf/please-openapi-rules/compare/v0.0.3...v0.0.4
[0.0.3]: https://github.com/andrew-womeldorf/please-openapi-rules/compare/v0.0.2...v0.0.3
[0.0.2]: https://github.com/andrew-womeldorf/please-openapi-rules/compare/v0.0.1...v0.0.2
[0.0.1]: https://github.com/andrew-womeldorf/please-openapi-rules/releases/tag/v0.0.1
