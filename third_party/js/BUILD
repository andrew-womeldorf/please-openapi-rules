subinclude("///js//build_defs:js")
package(default_visibility = ["PUBLIC"])

yarn_library(
    name = "@asyncapi--specs",
    package_name = "specs",
    scope = "@asyncapi",
    version = "2.13.1",
)

yarn_library(
    name = "@babel--code-frame",
    package_name = "code-frame",
    scope = "@babel",
    version = "7.16.7",
    deps = [
        ":@babel--highlight",
    ],
)

yarn_library(
    name = "@babel--generator",
    package_name = "generator",
    scope = "@babel",
    version = "7.17.7",
    deps = [
        ":@babel--types",
        ":jsesc",
        ":source-map",
    ],
)

yarn_library(
    name = "@babel--helper-annotate-as-pure",
    package_name = "helper-annotate-as-pure",
    scope = "@babel",
    version = "7.16.7",
    deps = [
        ":@babel--types",
    ],
)

yarn_library(
    name = "@babel--helper-environment-visitor",
    package_name = "helper-environment-visitor",
    scope = "@babel",
    version = "7.16.7",
    deps = [
        ":@babel--types",
    ],
)

yarn_library(
    name = "@babel--helper-function-name",
    package_name = "helper-function-name",
    scope = "@babel",
    version = "7.16.7",
    deps = [
        ":@babel--helper-get-function-arity",
        ":@babel--template",
        ":@babel--types",
    ],
)

yarn_library(
    name = "@babel--helper-get-function-arity",
    package_name = "helper-get-function-arity",
    scope = "@babel",
    version = "7.16.7",
    deps = [
        ":@babel--types",
    ],
)

yarn_library(
    name = "@babel--helper-hoist-variables",
    package_name = "helper-hoist-variables",
    scope = "@babel",
    version = "7.16.7",
    deps = [
        ":@babel--types",
    ],
)

yarn_library(
    name = "@babel--helper-module-imports",
    package_name = "helper-module-imports",
    scope = "@babel",
    version = "7.16.7",
    deps = [
        ":@babel--types",
    ],
)

yarn_library(
    name = "@babel--helper-split-export-declaration",
    package_name = "helper-split-export-declaration",
    scope = "@babel",
    version = "7.16.7",
    deps = [
        ":@babel--types",
    ],
)

yarn_library(
    name = "@babel--helper-validator-identifier",
    package_name = "helper-validator-identifier",
    scope = "@babel",
    version = "7.16.7",
)

yarn_library(
    name = "@babel--highlight",
    package_name = "highlight",
    scope = "@babel",
    version = "7.16.10",
    deps = [
        ":@babel--helper-validator-identifier",
        ":chalk",
        ":js-tokens",
    ],
)

yarn_library(
    name = "@babel--parser",
    package_name = "parser",
    scope = "@babel",
    version = "7.17.8",
)

yarn_library(
    name = "@babel--runtime",
    package_name = "runtime",
    scope = "@babel",
    version = "7.17.8",
    deps = [
        ":regenerator-runtime",
    ],
)

yarn_library(
    name = "@babel--template",
    package_name = "template",
    scope = "@babel",
    version = "7.16.7",
    deps = [
        ":@babel--code-frame",
        ":@babel--parser",
        ":@babel--types",
    ],
)

yarn_library(
    name = "@babel--traverse",
    package_name = "traverse",
    scope = "@babel",
    version = "7.17.3",
    deps = [
        ":@babel--code-frame",
        ":@babel--generator",
        ":@babel--helper-environment-visitor",
        ":@babel--helper-function-name",
        ":@babel--helper-hoist-variables",
        ":@babel--helper-split-export-declaration",
        ":@babel--parser",
        ":@babel--types",
        ":debug",
        ":globals",
    ],
)

yarn_library(
    name = "@babel--types",
    package_name = "types",
    scope = "@babel",
    version = "7.17.0",
    deps = [
        ":@babel--helper-validator-identifier",
        ":to-fast-properties",
    ],
)

yarn_library(
    name = "@emotion--is-prop-valid",
    package_name = "is-prop-valid",
    scope = "@emotion",
    version = "1.1.2",
    deps = [
        ":@emotion--memoize",
    ],
)

yarn_library(
    name = "@emotion--memoize",
    package_name = "memoize",
    scope = "@emotion",
    version = "0.7.5",
)

yarn_library(
    name = "@emotion--stylis",
    package_name = "stylis",
    scope = "@emotion",
    version = "0.8.5",
)

yarn_library(
    name = "@emotion--unitless",
    package_name = "unitless",
    scope = "@emotion",
    version = "0.7.5",
)

yarn_library(
    name = "@exodus--schemasafe",
    package_name = "schemasafe",
    scope = "@exodus",
    version = "1.0.0-rc.6",
)

yarn_library(
    name = "@jsep-plugin--regex",
    package_name = "regex",
    scope = "@jsep-plugin",
    version = "1.0.2",
)

yarn_library(
    name = "@jsep-plugin--ternary",
    package_name = "ternary",
    scope = "@jsep-plugin",
    version = "1.1.2",
)

yarn_library(
    name = "@nodelib--fs.scandir",
    package_name = "fs.scandir",
    scope = "@nodelib",
    version = "2.1.5",
    deps = [
        ":@nodelib--fs.stat",
        ":run-parallel",
    ],
)

yarn_library(
    name = "@nodelib--fs.stat",
    package_name = "fs.stat",
    scope = "@nodelib",
    version = "2.0.5",
)

yarn_library(
    name = "@nodelib--fs.walk",
    package_name = "fs.walk",
    scope = "@nodelib",
    version = "1.2.8",
    deps = [
        ":@nodelib--fs.scandir",
        ":fastq",
    ],
)

yarn_library(
    name = "@redocly--ajv",
    package_name = "ajv",
    scope = "@redocly",
    version = "8.6.4",
    deps = [
        ":fast-deep-equal",
        ":json-schema-traverse",
        ":require-from-string",
        ":uri-js",
    ],
)

yarn_library(
    name = "@redocly--openapi-cli",
    package_name = "openapi-cli",
    scope = "@redocly",
    version = "1.0.0-beta.92",
    deps = [
        ":@redocly--openapi-core",
        ":@types--node",
        ":assert-node-version",
        ":chokidar",
        ":colorette",
        ":glob",
        ":glob-promise",
        ":handlebars",
        ":portfinder",
        ":simple-websocket",
        ":yargs",
    ],
)

yarn_library(
    name = "@redocly--openapi-core",
    package_name = "openapi-core",
    scope = "@redocly",
    version = "1.0.0-beta.92",
    deps = [
        ":@redocly--ajv",
        ":@types--node",
        ":colorette",
        ":js-levenshtein",
        ":js-yaml",
        ":lodash.isequal",
        ":minimatch",
        ":node-fetch",
        ":pluralize",
        ":yaml-ast-parser",
    ],
)

yarn_library(
    name = "@redocly--react-dropdown-aria",
    package_name = "react-dropdown-aria",
    scope = "@redocly",
    version = "2.0.12",
)

yarn_library(
    name = "@rollup--plugin-commonjs",
    package_name = "plugin-commonjs",
    scope = "@rollup",
    version = "21.0.3",
    deps = [
        ":@rollup--pluginutils",
        ":commondir",
        ":estree-walker",
        ":glob",
        ":is-reference",
        ":magic-string",
        ":resolve",
    ],
)

yarn_library(
    name = "@rollup--pluginutils",
    package_name = "pluginutils",
    scope = "@rollup",
    version = "3.1.0",
    deps = [
        ":@types--estree",
        ":estree-walker",
        ":picomatch",
    ],
)

yarn_library(
    name = "@stoplight--better-ajv-errors",
    package_name = "better-ajv-errors",
    scope = "@stoplight",
    version = "1.0.1",
    deps = [
        ":jsonpointer",
        ":leven",
    ],
)

yarn_library(
    name = "@stoplight--json",
    package_name = "json",
    scope = "@stoplight",
    version = "3.18.0",
    deps = [
        ":@stoplight--ordered-object-literal",
        ":@stoplight--types",
        ":jsonc-parser",
        ":lodash",
        ":safe-stable-stringify",
    ],
)

yarn_library(
    name = "@stoplight--json-ref-readers",
    package_name = "json-ref-readers",
    scope = "@stoplight",
    version = "1.2.2",
    deps = [
        ":node-fetch",
        ":tslib",
    ],
)

yarn_library(
    name = "@stoplight--json-ref-resolver",
    package_name = "json-ref-resolver",
    scope = "@stoplight",
    version = "3.1.3",
    deps = [
        ":@stoplight--json",
        ":@stoplight--path",
        ":@stoplight--types",
        ":@types--urijs",
        ":dependency-graph",
        ":fast-memoize",
        ":immer",
        ":lodash.get",
        ":lodash.set",
        ":tslib",
        ":urijs",
    ],
)

yarn_library(
    name = "@stoplight--lifecycle",
    package_name = "lifecycle",
    scope = "@stoplight",
    version = "2.3.2",
    deps = [
        ":wolfy87-eventemitter",
    ],
)

yarn_library(
    name = "@stoplight--ordered-object-literal",
    package_name = "ordered-object-literal",
    scope = "@stoplight",
    version = "1.0.2",
)

yarn_library(
    name = "@stoplight--path",
    package_name = "path",
    scope = "@stoplight",
    version = "1.3.2",
)

yarn_library(
    name = "@stoplight--spectral-cli",
    package_name = "spectral-cli",
    scope = "@stoplight",
    version = "6.3.0",
    deps = [
        ":@rollup--plugin-commonjs",
        ":@stoplight--json",
        ":@stoplight--path",
        ":@stoplight--spectral-core",
        ":@stoplight--spectral-parsers",
        ":@stoplight--spectral-ref-resolver",
        ":@stoplight--spectral-ruleset-bundler",
        ":@stoplight--spectral-ruleset-migrator",
        ":@stoplight--spectral-rulesets",
        ":@stoplight--spectral-runtime",
        ":@stoplight--types",
        ":chalk",
        ":cliui",
        ":eol",
        ":fast-glob",
        ":lodash",
        ":pony-cause",
        ":proxy-agent",
        ":stacktracey",
        ":strip-ansi",
        ":text-table",
        ":tslib",
        ":yargs",
    ],
)

yarn_library(
    name = "@stoplight--spectral-core",
    package_name = "spectral-core",
    scope = "@stoplight",
    version = "1.11.1",
    deps = [
        ":@stoplight--better-ajv-errors",
        ":@stoplight--json",
        ":@stoplight--lifecycle",
        ":@stoplight--path",
        ":@stoplight--spectral-parsers",
        ":@stoplight--spectral-ref-resolver",
        ":@stoplight--spectral-runtime",
        ":@stoplight--types",
        ":ajv",
        ":ajv-errors",
        ":ajv-formats",
        ":blueimp-md5",
        ":json-schema",
        ":jsonpath-plus",
        ":lodash",
        ":lodash.topath",
        ":minimatch",
        ":nimma",
        ":pony-cause",
        ":simple-eval",
        ":tslib",
    ],
)

yarn_library(
    name = "@stoplight--spectral-formats",
    package_name = "spectral-formats",
    scope = "@stoplight",
    version = "1.1.0",
    deps = [
        ":@stoplight--json",
        ":@stoplight--spectral-core",
        ":@types--json-schema",
        ":tslib",
    ],
)

yarn_library(
    name = "@stoplight--spectral-functions",
    package_name = "spectral-functions",
    scope = "@stoplight",
    version = "1.6.1",
    deps = [
        ":@stoplight--better-ajv-errors",
        ":@stoplight--json",
        ":@stoplight--spectral-core",
        ":@stoplight--spectral-formats",
        ":@stoplight--spectral-runtime",
        ":@stoplight--types",
        ":ajv",
        ":ajv-draft-04",
        ":ajv-errors",
        ":ajv-formats",
        ":lodash",
        ":tslib",
    ],
)

yarn_library(
    name = "@stoplight--spectral-parsers",
    package_name = "spectral-parsers",
    scope = "@stoplight",
    version = "1.0.1",
    deps = [
        ":@stoplight--json",
        ":@stoplight--types",
        ":@stoplight--yaml",
        ":tslib",
    ],
)

yarn_library(
    name = "@stoplight--spectral-ref-resolver",
    package_name = "spectral-ref-resolver",
    scope = "@stoplight",
    version = "1.0.1",
    deps = [
        ":@stoplight--json-ref-readers",
        ":@stoplight--json-ref-resolver",
        ":@stoplight--spectral-runtime",
        ":dependency-graph",
        ":tslib",
    ],
)

yarn_library(
    name = "@stoplight--spectral-ruleset-bundler",
    package_name = "spectral-ruleset-bundler",
    scope = "@stoplight",
    version = "1.2.1",
    deps = [
        ":@rollup--plugin-commonjs",
        ":@stoplight--path",
        ":@stoplight--spectral-core",
        ":@stoplight--spectral-formats",
        ":@stoplight--spectral-functions",
        ":@stoplight--spectral-parsers",
        ":@stoplight--spectral-ref-resolver",
        ":@stoplight--spectral-ruleset-migrator",
        ":@stoplight--spectral-rulesets",
        ":@stoplight--spectral-runtime",
        ":@stoplight--types",
        ":@types--node",
        ":pony-cause",
        ":rollup",
        ":tslib",
        ":validate-npm-package-name",
    ],
)

yarn_library(
    name = "@stoplight--spectral-ruleset-migrator",
    package_name = "spectral-ruleset-migrator",
    scope = "@stoplight",
    version = "1.7.3",
    deps = [
        ":@stoplight--json",
        ":@stoplight--ordered-object-literal",
        ":@stoplight--path",
        ":@stoplight--spectral-functions",
        ":@stoplight--spectral-runtime",
        ":@stoplight--types",
        ":@stoplight--yaml",
        ":@types--node",
        ":ajv",
        ":ast-types",
        ":astring",
        ":reserved",
        ":tslib",
        ":validate-npm-package-name",
    ],
)

yarn_library(
    name = "@stoplight--spectral-rulesets",
    package_name = "spectral-rulesets",
    scope = "@stoplight",
    version = "1.7.0",
    deps = [
        ":@asyncapi--specs",
        ":@stoplight--better-ajv-errors",
        ":@stoplight--json",
        ":@stoplight--spectral-core",
        ":@stoplight--spectral-formats",
        ":@stoplight--spectral-functions",
        ":@stoplight--spectral-runtime",
        ":@stoplight--types",
        ":@types--json-schema",
        ":ajv",
        ":ajv-formats",
        ":json-schema-traverse",
        ":lodash",
        ":tslib",
    ],
)

yarn_library(
    name = "@stoplight--spectral-runtime",
    package_name = "spectral-runtime",
    scope = "@stoplight",
    version = "1.1.2",
    deps = [
        ":@stoplight--json",
        ":@stoplight--path",
        ":@stoplight--types",
        ":abort-controller",
        ":lodash",
        ":node-fetch",
        ":tslib",
    ],
)

yarn_library(
    name = "@stoplight--types",
    package_name = "types",
    scope = "@stoplight",
    version = "12.5.0",
    deps = [
        ":@types--json-schema",
        ":utility-types",
    ],
)

yarn_library(
    name = "@stoplight--yaml",
    package_name = "yaml",
    scope = "@stoplight",
    version = "4.2.2",
    deps = [
        ":@stoplight--ordered-object-literal",
        ":@stoplight--types",
        ":@stoplight--yaml-ast-parser",
        ":tslib",
    ],
)

yarn_library(
    name = "@stoplight--yaml-ast-parser",
    package_name = "yaml-ast-parser",
    scope = "@stoplight",
    version = "0.0.48",
)

yarn_library(
    name = "@tootallnate--once",
    package_name = "once",
    scope = "@tootallnate",
    version = "1.1.2",
)

yarn_library(
    name = "@types--estree",
    package_name = "estree",
    scope = "@types",
    version = "0.0.51",
)

yarn_library(
    name = "@types--glob",
    package_name = "glob",
    scope = "@types",
    version = "7.2.0",
    deps = [
        ":@types--minimatch",
        ":@types--node",
    ],
)

yarn_library(
    name = "@types--json-schema",
    package_name = "json-schema",
    scope = "@types",
    version = "7.0.11",
)

yarn_library(
    name = "@types--minimatch",
    package_name = "minimatch",
    scope = "@types",
    version = "3.0.5",
)

yarn_library(
    name = "@types--node",
    package_name = "node",
    scope = "@types",
    version = "17.0.23",
)

yarn_library(
    name = "@types--urijs",
    package_name = "urijs",
    scope = "@types",
    version = "1.19.19",
)

yarn_library(
    name = "abort-controller",
    package_name = "abort-controller",
    version = "3.0.0",
    deps = [
        ":event-target-shim",
    ],
)

yarn_library(
    name = "acorn",
    package_name = "acorn",
    version = "8.7.0",
)

yarn_library(
    name = "acorn-walk",
    package_name = "acorn-walk",
    version = "8.2.0",
)

yarn_library(
    name = "agent-base",
    package_name = "agent-base",
    version = "6.0.2",
    deps = [
        ":debug",
    ],
)

yarn_library(
    name = "ajv",
    package_name = "ajv",
    version = "8.11.0",
    deps = [
        ":fast-deep-equal",
        ":json-schema-traverse",
        ":require-from-string",
        ":uri-js",
    ],
)

yarn_library(
    name = "ajv-draft-04",
    package_name = "ajv-draft-04",
    version = "1.0.0",
)

yarn_library(
    name = "ajv-errors",
    package_name = "ajv-errors",
    version = "3.0.0",
)

yarn_library(
    name = "ajv-formats",
    package_name = "ajv-formats",
    version = "2.1.1",
    deps = [
        ":ajv",
    ],
)

yarn_library(
    name = "ansi-regex",
    package_name = "ansi-regex",
    version = "5.0.1",
)

yarn_library(
    name = "ansi-styles",
    package_name = "ansi-styles",
    version = "4.3.0",
    deps = [
        ":color-convert",
    ],
)

yarn_library(
    name = "anymatch",
    package_name = "anymatch",
    version = "3.1.2",
    deps = [
        ":normalize-path",
        ":picomatch",
    ],
)

yarn_library(
    name = "argparse",
    package_name = "argparse",
    version = "2.0.1",
)

yarn_library(
    name = "as-table",
    package_name = "as-table",
    version = "1.0.55",
    deps = [
        ":printable-characters",
    ],
)

yarn_library(
    name = "asn1.js",
    package_name = "asn1.js",
    version = "5.4.1",
    deps = [
        ":bn.js",
        ":inherits",
        ":minimalistic-assert",
        ":safer-buffer",
    ],
)

yarn_library(
    name = "assert",
    package_name = "assert",
    version = "1.5.0",
    deps = [
        ":object-assign",
        ":util",
    ],
)

yarn_library(
    name = "assert-node-version",
    package_name = "assert-node-version",
    version = "1.0.3",
    deps = [
        ":expected-node-version",
        ":semver",
    ],
)

yarn_library(
    name = "ast-types",
    package_name = "ast-types",
    version = "0.14.2",
    deps = [
        ":tslib",
    ],
)

yarn_library(
    name = "astring",
    package_name = "astring",
    version = "1.8.1",
)

yarn_library(
    name = "async",
    package_name = "async",
    version = "2.6.3",
    deps = [
        ":lodash",
    ],
)

yarn_library(
    name = "babel-plugin-styled-components",
    package_name = "babel-plugin-styled-components",
    version = "2.0.6",
    deps = [
        ":@babel--helper-annotate-as-pure",
        ":@babel--helper-module-imports",
        ":babel-plugin-syntax-jsx",
        ":lodash",
        ":picomatch",
    ],
)

yarn_library(
    name = "babel-plugin-syntax-jsx",
    package_name = "babel-plugin-syntax-jsx",
    version = "6.18.0",
)

yarn_library(
    name = "balanced-match",
    package_name = "balanced-match",
    version = "1.0.2",
)

yarn_library(
    name = "base64-js",
    package_name = "base64-js",
    version = "1.5.1",
)

yarn_library(
    name = "binary-extensions",
    package_name = "binary-extensions",
    version = "2.2.0",
)

yarn_library(
    name = "blueimp-md5",
    package_name = "blueimp-md5",
    version = "2.18.0",
)

yarn_library(
    name = "bn.js",
    package_name = "bn.js",
    version = "5.2.0",
)

yarn_library(
    name = "brace-expansion",
    package_name = "brace-expansion",
    version = "1.1.11",
    deps = [
        ":balanced-match",
        ":concat-map",
    ],
)

yarn_library(
    name = "braces",
    package_name = "braces",
    version = "3.0.2",
    deps = [
        ":fill-range",
    ],
)

yarn_library(
    name = "brorand",
    package_name = "brorand",
    version = "1.1.0",
)

yarn_library(
    name = "browserify-aes",
    package_name = "browserify-aes",
    version = "1.2.0",
    deps = [
        ":buffer-xor",
        ":cipher-base",
        ":create-hash",
        ":evp_bytestokey",
        ":inherits",
        ":safe-buffer",
    ],
)

yarn_library(
    name = "browserify-cipher",
    package_name = "browserify-cipher",
    version = "1.0.1",
    deps = [
        ":browserify-aes",
        ":browserify-des",
        ":evp_bytestokey",
    ],
)

yarn_library(
    name = "browserify-des",
    package_name = "browserify-des",
    version = "1.0.2",
    deps = [
        ":cipher-base",
        ":des.js",
        ":inherits",
        ":safe-buffer",
    ],
)

yarn_library(
    name = "browserify-rsa",
    package_name = "browserify-rsa",
    version = "4.1.0",
    deps = [
        ":bn.js",
        ":randombytes",
    ],
)

yarn_library(
    name = "browserify-sign",
    package_name = "browserify-sign",
    version = "4.2.1",
    deps = [
        ":bn.js",
        ":browserify-rsa",
        ":create-hash",
        ":create-hmac",
        ":elliptic",
        ":inherits",
        ":parse-asn1",
        ":readable-stream",
        ":safe-buffer",
    ],
)

yarn_library(
    name = "browserify-zlib",
    package_name = "browserify-zlib",
    version = "0.2.0",
    deps = [
        ":pako",
    ],
)

yarn_library(
    name = "buffer",
    package_name = "buffer",
    version = "4.9.2",
    deps = [
        ":base64-js",
        ":ieee754",
        ":isarray",
    ],
)

yarn_library(
    name = "buffer-xor",
    package_name = "buffer-xor",
    version = "1.0.3",
)

yarn_library(
    name = "builtin-status-codes",
    package_name = "builtin-status-codes",
    version = "3.0.0",
)

yarn_library(
    name = "builtins",
    package_name = "builtins",
    version = "1.0.3",
)

yarn_library(
    name = "bytes",
    package_name = "bytes",
    version = "3.1.2",
)

yarn_library(
    name = "call-me-maybe",
    package_name = "call-me-maybe",
    version = "1.0.1",
)

yarn_library(
    name = "camelize",
    package_name = "camelize",
    version = "1.0.0",
)

yarn_library(
    name = "chalk",
    package_name = "chalk",
    version = "4.1.2",
    deps = [
        ":ansi-styles",
        ":supports-color",
    ],
)

yarn_library(
    name = "chokidar",
    package_name = "chokidar",
    version = "3.5.3",
    deps = [
        ":anymatch",
        ":braces",
        ":glob-parent",
        ":is-binary-path",
        ":is-glob",
        ":normalize-path",
        ":readdirp",
        ":fsevents",
    ],
)

yarn_library(
    name = "cipher-base",
    package_name = "cipher-base",
    version = "1.0.4",
    deps = [
        ":inherits",
        ":safe-buffer",
    ],
)

yarn_library(
    name = "classnames",
    package_name = "classnames",
    version = "2.3.1",
)

yarn_library(
    name = "cliui",
    package_name = "cliui",
    version = "7.0.4",
    deps = [
        ":string-width",
        ":strip-ansi",
        ":wrap-ansi",
    ],
)

yarn_library(
    name = "clsx",
    package_name = "clsx",
    version = "1.1.1",
)

yarn_library(
    name = "color-convert",
    package_name = "color-convert",
    version = "2.0.1",
    deps = [
        ":color-name",
    ],
)

yarn_library(
    name = "color-name",
    package_name = "color-name",
    version = "1.1.4",
)

yarn_library(
    name = "colorette",
    package_name = "colorette",
    version = "1.4.0",
)

yarn_library(
    name = "commondir",
    package_name = "commondir",
    version = "1.0.1",
)

yarn_library(
    name = "concat-map",
    package_name = "concat-map",
    version = "0.0.1",
)

yarn_library(
    name = "console-browserify",
    package_name = "console-browserify",
    version = "1.2.0",
)

yarn_library(
    name = "constants-browserify",
    package_name = "constants-browserify",
    version = "1.0.0",
)

yarn_library(
    name = "create-ecdh",
    package_name = "create-ecdh",
    version = "4.0.4",
    deps = [
        ":bn.js",
        ":elliptic",
    ],
)

yarn_library(
    name = "create-hash",
    package_name = "create-hash",
    version = "1.2.0",
    deps = [
        ":cipher-base",
        ":inherits",
        ":md5.js",
        ":ripemd160",
        ":sha.js",
    ],
)

yarn_library(
    name = "create-hmac",
    package_name = "create-hmac",
    version = "1.1.7",
    deps = [
        ":cipher-base",
        ":create-hash",
        ":inherits",
        ":ripemd160",
        ":safe-buffer",
        ":sha.js",
    ],
)

yarn_library(
    name = "crypto-browserify",
    package_name = "crypto-browserify",
    version = "3.12.0",
    deps = [
        ":browserify-cipher",
        ":browserify-sign",
        ":create-ecdh",
        ":create-hash",
        ":create-hmac",
        ":diffie-hellman",
        ":inherits",
        ":pbkdf2",
        ":public-encrypt",
        ":randombytes",
        ":randomfill",
    ],
)

yarn_library(
    name = "css-color-keywords",
    package_name = "css-color-keywords",
    version = "1.0.0",
)

yarn_library(
    name = "css-to-react-native",
    package_name = "css-to-react-native",
    version = "3.0.0",
    deps = [
        ":camelize",
        ":css-color-keywords",
        ":postcss-value-parser",
    ],
)

yarn_library(
    name = "data-uri-to-buffer",
    package_name = "data-uri-to-buffer",
    version = "3.0.1",
)

yarn_library(
    name = "debug",
    package_name = "debug",
    version = "4.3.4",
    deps = [
        ":ms",
    ],
)

yarn_library(
    name = "decko",
    package_name = "decko",
    version = "1.2.0",
)

yarn_library(
    name = "deep-is",
    package_name = "deep-is",
    version = "0.1.4",
)

yarn_library(
    name = "degenerator",
    package_name = "degenerator",
    version = "3.0.2",
    deps = [
        ":ast-types",
        ":escodegen",
        ":esprima",
        ":vm2",
    ],
)

yarn_library(
    name = "depd",
    package_name = "depd",
    version = "2.0.0",
)

yarn_library(
    name = "dependency-graph",
    package_name = "dependency-graph",
    version = "0.11.0",
)

yarn_library(
    name = "des.js",
    package_name = "des.js",
    version = "1.0.1",
    deps = [
        ":inherits",
        ":minimalistic-assert",
    ],
)

yarn_library(
    name = "diffie-hellman",
    package_name = "diffie-hellman",
    version = "5.0.3",
    deps = [
        ":bn.js",
        ":miller-rabin",
        ":randombytes",
    ],
)

yarn_library(
    name = "domain-browser",
    package_name = "domain-browser",
    version = "1.2.0",
)

yarn_library(
    name = "dompurify",
    package_name = "dompurify",
    version = "2.3.6",
)

yarn_library(
    name = "elliptic",
    package_name = "elliptic",
    version = "6.5.4",
    deps = [
        ":bn.js",
        ":brorand",
        ":hash.js",
        ":hmac-drbg",
        ":inherits",
        ":minimalistic-assert",
        ":minimalistic-crypto-utils",
    ],
)

yarn_library(
    name = "emoji-regex",
    package_name = "emoji-regex",
    version = "8.0.0",
)

yarn_library(
    name = "eol",
    package_name = "eol",
    version = "0.9.1",
)

yarn_library(
    name = "es6-promise",
    package_name = "es6-promise",
    version = "3.3.1",
)

yarn_library(
    name = "escalade",
    package_name = "escalade",
    version = "3.1.1",
)

yarn_library(
    name = "escodegen",
    package_name = "escodegen",
    version = "1.14.3",
    deps = [
        ":esprima",
        ":estraverse",
        ":esutils",
        ":optionator",
        ":source-map",
    ],
)

yarn_library(
    name = "esprima",
    package_name = "esprima",
    version = "4.0.1",
)

yarn_library(
    name = "estraverse",
    package_name = "estraverse",
    version = "4.3.0",
)

yarn_library(
    name = "estree-walker",
    package_name = "estree-walker",
    version = "2.0.2",
)

yarn_library(
    name = "esutils",
    package_name = "esutils",
    version = "2.0.3",
)

yarn_library(
    name = "event-target-shim",
    package_name = "event-target-shim",
    version = "5.0.1",
)

yarn_library(
    name = "eventemitter3",
    package_name = "eventemitter3",
    version = "4.0.7",
)

yarn_library(
    name = "events",
    package_name = "events",
    version = "3.3.0",
)

yarn_library(
    name = "evp_bytestokey",
    package_name = "evp_bytestokey",
    version = "1.0.3",
    deps = [
        ":md5.js",
        ":safe-buffer",
    ],
)

yarn_library(
    name = "expected-node-version",
    package_name = "expected-node-version",
    version = "1.0.2",
)

yarn_library(
    name = "fast-deep-equal",
    package_name = "fast-deep-equal",
    version = "3.1.3",
)

yarn_library(
    name = "fast-glob",
    package_name = "fast-glob",
    version = "3.2.7",
    deps = [
        ":@nodelib--fs.stat",
        ":@nodelib--fs.walk",
        ":glob-parent",
        ":merge2",
        ":micromatch",
    ],
)

yarn_library(
    name = "fast-levenshtein",
    package_name = "fast-levenshtein",
    version = "2.0.6",
)

yarn_library(
    name = "fast-memoize",
    package_name = "fast-memoize",
    version = "2.5.2",
)

yarn_library(
    name = "fast-safe-stringify",
    package_name = "fast-safe-stringify",
    version = "2.1.1",
)

yarn_library(
    name = "fastq",
    package_name = "fastq",
    version = "1.13.0",
    deps = [
        ":reusify",
    ],
)

yarn_library(
    name = "file-uri-to-path",
    package_name = "file-uri-to-path",
    version = "2.0.0",
)

yarn_library(
    name = "fill-range",
    package_name = "fill-range",
    version = "7.0.1",
    deps = [
        ":to-regex-range",
    ],
)

yarn_library(
    name = "foreach",
    package_name = "foreach",
    version = "2.0.5",
)

yarn_library(
    name = "fs-extra",
    package_name = "fs-extra",
    version = "8.1.0",
    deps = [
        ":graceful-fs",
        ":jsonfile",
        ":universalify",
    ],
)

yarn_library(
    name = "fs.realpath",
    package_name = "fs.realpath",
    version = "1.0.0",
)

yarn_library(
    name = "fsevents",
    package_name = "fsevents",
    version = "2.3.2",
)

yarn_library(
    name = "ftp",
    package_name = "ftp",
    version = "0.3.10",
    deps = [
        ":readable-stream",
        ":xregexp",
    ],
)

yarn_library(
    name = "function-bind",
    package_name = "function-bind",
    version = "1.1.1",
)

yarn_library(
    name = "get-caller-file",
    package_name = "get-caller-file",
    version = "2.0.5",
)

yarn_library(
    name = "get-source",
    package_name = "get-source",
    version = "2.0.12",
    deps = [
        ":data-uri-to-buffer",
        ":source-map",
    ],
)

yarn_library(
    name = "get-uri",
    package_name = "get-uri",
    version = "3.0.2",
    deps = [
        ":@tootallnate--once",
        ":data-uri-to-buffer",
        ":debug",
        ":file-uri-to-path",
        ":fs-extra",
        ":ftp",
    ],
)

yarn_library(
    name = "glob",
    package_name = "glob",
    version = "7.2.0",
    deps = [
        ":fs.realpath",
        ":inflight",
        ":inherits",
        ":minimatch",
        ":once",
        ":path-is-absolute",
    ],
)

yarn_library(
    name = "glob-parent",
    package_name = "glob-parent",
    version = "5.1.2",
    deps = [
        ":is-glob",
    ],
)

yarn_library(
    name = "glob-promise",
    package_name = "glob-promise",
    version = "3.4.0",
    deps = [
        ":@types--glob",
    ],
)

yarn_library(
    name = "globals",
    package_name = "globals",
    version = "11.12.0",
)

yarn_library(
    name = "graceful-fs",
    package_name = "graceful-fs",
    version = "4.2.10",
)

yarn_library(
    name = "handlebars",
    package_name = "handlebars",
    version = "4.7.7",
    deps = [
        ":minimist",
        ":neo-async",
        ":source-map",
        ":wordwrap",
        ":uglify-js",
    ],
)

yarn_library(
    name = "has",
    package_name = "has",
    version = "1.0.3",
    deps = [
        ":function-bind",
    ],
)

yarn_library(
    name = "has-flag",
    package_name = "has-flag",
    version = "4.0.0",
)

yarn_library(
    name = "hash-base",
    package_name = "hash-base",
    version = "3.1.0",
    deps = [
        ":inherits",
        ":readable-stream",
        ":safe-buffer",
    ],
)

yarn_library(
    name = "hash.js",
    package_name = "hash.js",
    version = "1.1.7",
    deps = [
        ":inherits",
        ":minimalistic-assert",
    ],
)

yarn_library(
    name = "hmac-drbg",
    package_name = "hmac-drbg",
    version = "1.0.1",
    deps = [
        ":hash.js",
        ":minimalistic-assert",
        ":minimalistic-crypto-utils",
    ],
)

yarn_library(
    name = "hoist-non-react-statics",
    package_name = "hoist-non-react-statics",
    version = "3.3.2",
    deps = [
        ":react-is",
    ],
)

yarn_library(
    name = "http-errors",
    package_name = "http-errors",
    version = "2.0.0",
    deps = [
        ":depd",
        ":inherits",
        ":setprototypeof",
        ":statuses",
        ":toidentifier",
    ],
)

yarn_library(
    name = "http-proxy-agent",
    package_name = "http-proxy-agent",
    version = "4.0.1",
    deps = [
        ":@tootallnate--once",
        ":agent-base",
        ":debug",
    ],
)

yarn_library(
    name = "http2-client",
    package_name = "http2-client",
    version = "1.3.5",
)

yarn_library(
    name = "https-browserify",
    package_name = "https-browserify",
    version = "1.0.0",
)

yarn_library(
    name = "https-proxy-agent",
    package_name = "https-proxy-agent",
    version = "5.0.0",
    deps = [
        ":agent-base",
        ":debug",
    ],
)

yarn_library(
    name = "iconv-lite",
    package_name = "iconv-lite",
    version = "0.4.24",
    deps = [
        ":safer-buffer",
    ],
)

yarn_library(
    name = "ieee754",
    package_name = "ieee754",
    version = "1.2.1",
)

yarn_library(
    name = "immer",
    package_name = "immer",
    version = "9.0.12",
)

yarn_library(
    name = "inflight",
    package_name = "inflight",
    version = "1.0.6",
    deps = [
        ":once",
        ":wrappy",
    ],
)

yarn_library(
    name = "inherits",
    package_name = "inherits",
    version = "2.0.4",
)

yarn_library(
    name = "ip",
    package_name = "ip",
    version = "1.1.5",
)

yarn_library(
    name = "is-binary-path",
    package_name = "is-binary-path",
    version = "2.1.0",
    deps = [
        ":binary-extensions",
    ],
)

yarn_library(
    name = "is-core-module",
    package_name = "is-core-module",
    version = "2.8.1",
    deps = [
        ":has",
    ],
)

yarn_library(
    name = "is-extglob",
    package_name = "is-extglob",
    version = "2.1.1",
)

yarn_library(
    name = "is-fullwidth-code-point",
    package_name = "is-fullwidth-code-point",
    version = "3.0.0",
)

yarn_library(
    name = "is-glob",
    package_name = "is-glob",
    version = "4.0.3",
    deps = [
        ":is-extglob",
    ],
)

yarn_library(
    name = "is-number",
    package_name = "is-number",
    version = "7.0.0",
)

yarn_library(
    name = "is-reference",
    package_name = "is-reference",
    version = "1.2.1",
    deps = [
        ":@types--estree",
    ],
)

yarn_library(
    name = "isarray",
    package_name = "isarray",
    version = "2.0.5",
)

yarn_library(
    name = "js-levenshtein",
    package_name = "js-levenshtein",
    version = "1.1.6",
)

yarn_library(
    name = "js-tokens",
    package_name = "js-tokens",
    version = "4.0.0",
)

yarn_library(
    name = "js-yaml",
    package_name = "js-yaml",
    version = "4.1.0",
    deps = [
        ":argparse",
    ],
)

yarn_library(
    name = "jsep",
    package_name = "jsep",
    version = "1.3.4",
)

yarn_library(
    name = "jsesc",
    package_name = "jsesc",
    version = "2.5.2",
)

yarn_library(
    name = "json-pointer",
    package_name = "json-pointer",
    version = "0.6.2",
    deps = [
        ":foreach",
    ],
)

yarn_library(
    name = "json-schema",
    package_name = "json-schema",
    version = "0.4.0",
)

yarn_library(
    name = "json-schema-traverse",
    package_name = "json-schema-traverse",
    version = "1.0.0",
)

yarn_library(
    name = "jsonc-parser",
    package_name = "jsonc-parser",
    version = "2.2.1",
)

yarn_library(
    name = "jsonfile",
    package_name = "jsonfile",
    version = "4.0.0",
    deps = [
        ":graceful-fs",
    ],
)

yarn_library(
    name = "jsonpath-plus",
    package_name = "jsonpath-plus",
    version = "6.0.1",
)

yarn_library(
    name = "jsonpointer",
    package_name = "jsonpointer",
    version = "5.0.0",
)

yarn_library(
    name = "leven",
    package_name = "leven",
    version = "3.1.0",
)

yarn_library(
    name = "levn",
    package_name = "levn",
    version = "0.3.0",
    deps = [
        ":prelude-ls",
        ":type-check",
    ],
)

yarn_library(
    name = "lodash",
    package_name = "lodash",
    version = "4.17.21",
)

yarn_library(
    name = "lodash.get",
    package_name = "lodash.get",
    version = "4.4.2",
)

yarn_library(
    name = "lodash.isequal",
    package_name = "lodash.isequal",
    version = "4.5.0",
)

yarn_library(
    name = "lodash.set",
    package_name = "lodash.set",
    version = "4.3.2",
)

yarn_library(
    name = "lodash.topath",
    package_name = "lodash.topath",
    version = "4.5.2",
)

yarn_library(
    name = "loose-envify",
    package_name = "loose-envify",
    version = "1.4.0",
    deps = [
        ":js-tokens",
    ],
)

yarn_library(
    name = "lru-cache",
    package_name = "lru-cache",
    version = "5.1.1",
    deps = [
        ":yallist",
    ],
)

yarn_library(
    name = "lunr",
    package_name = "lunr",
    version = "2.3.9",
)

yarn_library(
    name = "magic-string",
    package_name = "magic-string",
    version = "0.25.9",
    deps = [
        ":sourcemap-codec",
    ],
)

yarn_library(
    name = "mark.js",
    package_name = "mark.js",
    version = "8.11.1",
)

yarn_library(
    name = "marked",
    package_name = "marked",
    version = "4.0.12",
)

yarn_library(
    name = "md5.js",
    package_name = "md5.js",
    version = "1.3.5",
    deps = [
        ":hash-base",
        ":inherits",
        ":safe-buffer",
    ],
)

yarn_library(
    name = "merge2",
    package_name = "merge2",
    version = "1.4.1",
)

yarn_library(
    name = "micromatch",
    package_name = "micromatch",
    version = "4.0.5",
    deps = [
        ":braces",
        ":picomatch",
    ],
)

yarn_library(
    name = "miller-rabin",
    package_name = "miller-rabin",
    version = "4.0.1",
    deps = [
        ":bn.js",
        ":brorand",
    ],
)

yarn_library(
    name = "minimalistic-assert",
    package_name = "minimalistic-assert",
    version = "1.0.1",
)

yarn_library(
    name = "minimalistic-crypto-utils",
    package_name = "minimalistic-crypto-utils",
    version = "1.0.1",
)

yarn_library(
    name = "minimatch",
    package_name = "minimatch",
    version = "3.1.2",
    deps = [
        ":brace-expansion",
    ],
)

yarn_library(
    name = "minimist",
    package_name = "minimist",
    version = "1.2.6",
)

yarn_library(
    name = "mkdirp",
    package_name = "mkdirp",
    version = "1.0.4",
)

yarn_library(
    name = "mobx",
    package_name = "mobx",
    version = "6.5.0",
)

yarn_library(
    name = "mobx-react",
    package_name = "mobx-react",
    version = "7.3.0",
    deps = [
        ":mobx-react-lite",
    ],
)

yarn_library(
    name = "mobx-react-lite",
    package_name = "mobx-react-lite",
    version = "3.3.0",
)

yarn_library(
    name = "ms",
    package_name = "ms",
    version = "2.1.3",
)

yarn_library(
    name = "neo-async",
    package_name = "neo-async",
    version = "2.6.2",
)

yarn_library(
    name = "netmask",
    package_name = "netmask",
    version = "2.0.2",
)

yarn_library(
    name = "nimma",
    package_name = "nimma",
    version = "0.2.0",
    deps = [
        ":@jsep-plugin--regex",
        ":@jsep-plugin--ternary",
        ":astring",
        ":jsep",
        ":jsonpath-plus",
        ":lodash.topath",
    ],
)

yarn_library(
    name = "node-fetch",
    package_name = "node-fetch",
    version = "2.6.7",
    deps = [
        ":whatwg-url",
    ],
)

yarn_library(
    name = "node-fetch-h2",
    package_name = "node-fetch-h2",
    version = "2.3.0",
    deps = [
        ":http2-client",
    ],
)

yarn_library(
    name = "node-libs-browser",
    package_name = "node-libs-browser",
    version = "2.2.1",
    deps = [
        ":assert",
        ":browserify-zlib",
        ":buffer",
        ":console-browserify",
        ":constants-browserify",
        ":crypto-browserify",
        ":domain-browser",
        ":events",
        ":https-browserify",
        ":os-browserify",
        ":path-browserify",
        ":process",
        ":punycode",
        ":querystring-es3",
        ":readable-stream",
        ":stream-browserify",
        ":stream-http",
        ":string_decoder",
        ":timers-browserify",
        ":tty-browserify",
        ":url",
        ":util",
        ":vm-browserify",
    ],
)

yarn_library(
    name = "node-readfiles",
    package_name = "node-readfiles",
    version = "0.2.0",
    deps = [
        ":es6-promise",
    ],
)

yarn_library(
    name = "normalize-path",
    package_name = "normalize-path",
    version = "3.0.0",
)

yarn_library(
    name = "oas-kit-common",
    package_name = "oas-kit-common",
    version = "1.0.8",
    deps = [
        ":fast-safe-stringify",
    ],
)

yarn_library(
    name = "oas-linter",
    package_name = "oas-linter",
    version = "3.2.2",
    deps = [
        ":@exodus--schemasafe",
        ":should",
        ":yaml",
    ],
)

yarn_library(
    name = "oas-resolver",
    package_name = "oas-resolver",
    version = "2.5.6",
    deps = [
        ":node-fetch-h2",
        ":oas-kit-common",
        ":reftools",
        ":yaml",
        ":yargs",
    ],
)

yarn_library(
    name = "oas-schema-walker",
    package_name = "oas-schema-walker",
    version = "1.1.5",
)

yarn_library(
    name = "oas-validator",
    package_name = "oas-validator",
    version = "5.0.8",
    deps = [
        ":call-me-maybe",
        ":oas-kit-common",
        ":oas-linter",
        ":oas-resolver",
        ":oas-schema-walker",
        ":reftools",
        ":should",
        ":yaml",
    ],
)

yarn_library(
    name = "object-assign",
    package_name = "object-assign",
    version = "4.1.1",
)

yarn_library(
    name = "once",
    package_name = "once",
    version = "1.4.0",
    deps = [
        ":wrappy",
    ],
)

yarn_library(
    name = "openapi-sampler",
    package_name = "openapi-sampler",
    version = "1.2.1",
    deps = [
        ":@types--json-schema",
        ":json-pointer",
    ],
)

yarn_library(
    name = "optionator",
    package_name = "optionator",
    version = "0.8.3",
    deps = [
        ":deep-is",
        ":fast-levenshtein",
        ":levn",
        ":prelude-ls",
        ":type-check",
        ":word-wrap",
    ],
)

yarn_library(
    name = "os-browserify",
    package_name = "os-browserify",
    version = "0.3.0",
)

yarn_library(
    name = "pac-proxy-agent",
    package_name = "pac-proxy-agent",
    version = "5.0.0",
    deps = [
        ":@tootallnate--once",
        ":agent-base",
        ":debug",
        ":get-uri",
        ":http-proxy-agent",
        ":https-proxy-agent",
        ":pac-resolver",
        ":raw-body",
        ":socks-proxy-agent",
    ],
)

yarn_library(
    name = "pac-resolver",
    package_name = "pac-resolver",
    version = "5.0.0",
    deps = [
        ":degenerator",
        ":ip",
        ":netmask",
    ],
)

yarn_library(
    name = "pako",
    package_name = "pako",
    version = "1.0.11",
)

yarn_library(
    name = "parse-asn1",
    package_name = "parse-asn1",
    version = "5.1.6",
    deps = [
        ":asn1.js",
        ":browserify-aes",
        ":evp_bytestokey",
        ":pbkdf2",
        ":safe-buffer",
    ],
)

yarn_library(
    name = "path-browserify",
    package_name = "path-browserify",
    version = "1.0.1",
)

yarn_library(
    name = "path-is-absolute",
    package_name = "path-is-absolute",
    version = "1.0.1",
)

yarn_library(
    name = "path-parse",
    package_name = "path-parse",
    version = "1.0.7",
)

yarn_library(
    name = "pbkdf2",
    package_name = "pbkdf2",
    version = "3.1.2",
    deps = [
        ":create-hash",
        ":create-hmac",
        ":ripemd160",
        ":safe-buffer",
        ":sha.js",
    ],
)

yarn_library(
    name = "perfect-scrollbar",
    package_name = "perfect-scrollbar",
    version = "1.5.5",
)

yarn_library(
    name = "picomatch",
    package_name = "picomatch",
    version = "2.3.1",
)

yarn_library(
    name = "pluralize",
    package_name = "pluralize",
    version = "8.0.0",
)

yarn_library(
    name = "polished",
    package_name = "polished",
    version = "4.2.1",
    deps = [
        ":@babel--runtime",
    ],
)

yarn_library(
    name = "pony-cause",
    package_name = "pony-cause",
    version = "1.1.1",
)

yarn_library(
    name = "portfinder",
    package_name = "portfinder",
    version = "1.0.28",
    deps = [
        ":async",
        ":debug",
        ":mkdirp",
    ],
)

yarn_library(
    name = "postcss-value-parser",
    package_name = "postcss-value-parser",
    version = "4.2.0",
)

yarn_library(
    name = "prelude-ls",
    package_name = "prelude-ls",
    version = "1.1.2",
)

yarn_library(
    name = "printable-characters",
    package_name = "printable-characters",
    version = "1.0.42",
)

yarn_library(
    name = "prismjs",
    package_name = "prismjs",
    version = "1.27.0",
)

yarn_library(
    name = "process",
    package_name = "process",
    version = "0.11.10",
)

yarn_library(
    name = "prop-types",
    package_name = "prop-types",
    version = "15.8.1",
    deps = [
        ":loose-envify",
        ":object-assign",
        ":react-is",
    ],
)

yarn_library(
    name = "proxy-agent",
    package_name = "proxy-agent",
    version = "5.0.0",
    deps = [
        ":agent-base",
        ":debug",
        ":http-proxy-agent",
        ":https-proxy-agent",
        ":lru-cache",
        ":pac-proxy-agent",
        ":proxy-from-env",
        ":socks-proxy-agent",
    ],
)

yarn_library(
    name = "proxy-from-env",
    package_name = "proxy-from-env",
    version = "1.1.0",
)

yarn_library(
    name = "public-encrypt",
    package_name = "public-encrypt",
    version = "4.0.3",
    deps = [
        ":bn.js",
        ":browserify-rsa",
        ":create-hash",
        ":parse-asn1",
        ":randombytes",
        ":safe-buffer",
    ],
)

yarn_library(
    name = "punycode",
    package_name = "punycode",
    version = "1.3.2",
)

yarn_library(
    name = "querystring",
    package_name = "querystring",
    version = "0.2.0",
)

yarn_library(
    name = "querystring-es3",
    package_name = "querystring-es3",
    version = "0.2.1",
)

yarn_library(
    name = "queue-microtask",
    package_name = "queue-microtask",
    version = "1.2.3",
)

yarn_library(
    name = "randombytes",
    package_name = "randombytes",
    version = "2.1.0",
    deps = [
        ":safe-buffer",
    ],
)

yarn_library(
    name = "randomfill",
    package_name = "randomfill",
    version = "1.0.4",
    deps = [
        ":randombytes",
        ":safe-buffer",
    ],
)

yarn_library(
    name = "raw-body",
    package_name = "raw-body",
    version = "2.5.1",
    deps = [
        ":bytes",
        ":http-errors",
        ":iconv-lite",
        ":unpipe",
    ],
)

yarn_library(
    name = "react",
    package_name = "react",
    version = "17.0.2",
    deps = [
        ":loose-envify",
        ":object-assign",
    ],
)

yarn_library(
    name = "react-dom",
    package_name = "react-dom",
    version = "17.0.2",
    deps = [
        ":loose-envify",
        ":object-assign",
        ":scheduler",
    ],
)

yarn_library(
    name = "react-is",
    package_name = "react-is",
    version = "16.13.1",
)

yarn_library(
    name = "react-tabs",
    package_name = "react-tabs",
    version = "3.2.3",
    deps = [
        ":clsx",
        ":prop-types",
    ],
)

yarn_library(
    name = "readable-stream",
    package_name = "readable-stream",
    version = "3.6.0",
    deps = [
        ":inherits",
        ":string_decoder",
        ":util-deprecate",
    ],
)

yarn_library(
    name = "readdirp",
    package_name = "readdirp",
    version = "3.6.0",
    deps = [
        ":picomatch",
    ],
)

yarn_library(
    name = "redoc",
    package_name = "redoc",
    version = "2.0.0-rc.66",
    deps = [
        ":@redocly--openapi-core",
        ":@redocly--react-dropdown-aria",
        ":classnames",
        ":decko",
        ":dompurify",
        ":eventemitter3",
        ":json-pointer",
        ":lunr",
        ":mark.js",
        ":marked",
        ":mobx-react",
        ":openapi-sampler",
        ":path-browserify",
        ":perfect-scrollbar",
        ":polished",
        ":prismjs",
        ":prop-types",
        ":react-tabs",
        ":slugify",
        ":stickyfill",
        ":style-loader",
        ":swagger2openapi",
        ":url-template",
    ],
)

yarn_library(
    name = "redoc-cli-lib",
    package_name = "redoc-cli",
    version = "0.13.10",
    deps = [
        ":chokidar",
        ":handlebars",
        ":isarray",
        ":mkdirp",
        ":mobx",
        ":node-libs-browser",
        ":react",
        ":react-dom",
        ":redoc",
        ":styled-components",
        ":yargs",
    ],
)

yarn_library(
    name = "reftools",
    package_name = "reftools",
    version = "1.1.9",
)

yarn_library(
    name = "regenerator-runtime",
    package_name = "regenerator-runtime",
    version = "0.13.9",
)

yarn_library(
    name = "require-directory",
    package_name = "require-directory",
    version = "2.1.1",
)

yarn_library(
    name = "require-from-string",
    package_name = "require-from-string",
    version = "2.0.2",
)

yarn_library(
    name = "reserved",
    package_name = "reserved",
    version = "0.1.2",
)

yarn_library(
    name = "resolve",
    package_name = "resolve",
    version = "1.22.0",
    deps = [
        ":is-core-module",
        ":path-parse",
        ":supports-preserve-symlinks-flag",
    ],
)

yarn_library(
    name = "reusify",
    package_name = "reusify",
    version = "1.0.4",
)

yarn_library(
    name = "ripemd160",
    package_name = "ripemd160",
    version = "2.0.2",
    deps = [
        ":hash-base",
        ":inherits",
    ],
)

yarn_library(
    name = "rollup",
    package_name = "rollup",
    version = "2.67.3",
    deps = [
        ":fsevents",
    ],
)

yarn_library(
    name = "run-parallel",
    package_name = "run-parallel",
    version = "1.2.0",
    deps = [
        ":queue-microtask",
    ],
)

yarn_library(
    name = "safe-buffer",
    package_name = "safe-buffer",
    version = "5.2.1",
)

yarn_library(
    name = "safe-stable-stringify",
    package_name = "safe-stable-stringify",
    version = "1.1.1",
)

yarn_library(
    name = "safer-buffer",
    package_name = "safer-buffer",
    version = "2.1.2",
)

yarn_library(
    name = "scheduler",
    package_name = "scheduler",
    version = "0.20.2",
    deps = [
        ":loose-envify",
        ":object-assign",
    ],
)

yarn_library(
    name = "semver",
    package_name = "semver",
    version = "5.7.1",
)

yarn_library(
    name = "setimmediate",
    package_name = "setimmediate",
    version = "1.0.5",
)

yarn_library(
    name = "setprototypeof",
    package_name = "setprototypeof",
    version = "1.2.0",
)

yarn_library(
    name = "sha.js",
    package_name = "sha.js",
    version = "2.4.11",
    deps = [
        ":inherits",
        ":safe-buffer",
    ],
)

yarn_library(
    name = "shallowequal",
    package_name = "shallowequal",
    version = "1.1.0",
)

yarn_library(
    name = "should",
    package_name = "should",
    version = "13.2.3",
    deps = [
        ":should-equal",
        ":should-format",
        ":should-type",
        ":should-type-adaptors",
        ":should-util",
    ],
)

yarn_library(
    name = "should-equal",
    package_name = "should-equal",
    version = "2.0.0",
    deps = [
        ":should-type",
    ],
)

yarn_library(
    name = "should-format",
    package_name = "should-format",
    version = "3.0.3",
    deps = [
        ":should-type",
        ":should-type-adaptors",
    ],
)

yarn_library(
    name = "should-type",
    package_name = "should-type",
    version = "1.4.0",
)

yarn_library(
    name = "should-type-adaptors",
    package_name = "should-type-adaptors",
    version = "1.1.0",
    deps = [
        ":should-type",
        ":should-util",
    ],
)

yarn_library(
    name = "should-util",
    package_name = "should-util",
    version = "1.0.1",
)

yarn_library(
    name = "simple-eval",
    package_name = "simple-eval",
    version = "1.0.0",
    deps = [
        ":jsep",
    ],
)

yarn_library(
    name = "simple-websocket",
    package_name = "simple-websocket",
    version = "9.1.0",
    deps = [
        ":debug",
        ":queue-microtask",
        ":randombytes",
        ":readable-stream",
        ":ws",
    ],
)

yarn_library(
    name = "slugify",
    package_name = "slugify",
    version = "1.4.7",
)

yarn_library(
    name = "smart-buffer",
    package_name = "smart-buffer",
    version = "4.2.0",
)

yarn_library(
    name = "socks",
    package_name = "socks",
    version = "2.6.2",
    deps = [
        ":ip",
        ":smart-buffer",
    ],
)

yarn_library(
    name = "socks-proxy-agent",
    package_name = "socks-proxy-agent",
    version = "5.0.1",
    deps = [
        ":agent-base",
        ":debug",
        ":socks",
    ],
)

yarn_library(
    name = "source-map",
    package_name = "source-map",
    version = "0.6.1",
)

yarn_library(
    name = "sourcemap-codec",
    package_name = "sourcemap-codec",
    version = "1.4.8",
)

yarn_library(
    name = "stacktracey",
    package_name = "stacktracey",
    version = "2.1.8",
    deps = [
        ":as-table",
        ":get-source",
    ],
)

yarn_library(
    name = "statuses",
    package_name = "statuses",
    version = "2.0.1",
)

yarn_library(
    name = "stickyfill",
    package_name = "stickyfill",
    version = "1.1.1",
)

yarn_library(
    name = "stream-browserify",
    package_name = "stream-browserify",
    version = "2.0.2",
    deps = [
        ":inherits",
        ":readable-stream",
    ],
)

yarn_library(
    name = "stream-http",
    package_name = "stream-http",
    version = "2.8.3",
    deps = [
        ":builtin-status-codes",
        ":inherits",
        ":readable-stream",
        ":to-arraybuffer",
        ":xtend",
    ],
)

yarn_library(
    name = "string-width",
    package_name = "string-width",
    version = "4.2.3",
    deps = [
        ":emoji-regex",
        ":is-fullwidth-code-point",
        ":strip-ansi",
    ],
)

yarn_library(
    name = "string_decoder",
    package_name = "string_decoder",
    version = "1.1.1",
    deps = [
        ":safe-buffer",
    ],
)

yarn_library(
    name = "strip-ansi",
    package_name = "strip-ansi",
    version = "6.0.1",
    deps = [
        ":ansi-regex",
    ],
)

yarn_library(
    name = "style-loader",
    package_name = "style-loader",
    version = "3.3.1",
)

yarn_library(
    name = "styled-components",
    package_name = "styled-components",
    version = "5.3.5",
    deps = [
        ":@babel--helper-module-imports",
        ":@babel--traverse",
        ":@emotion--is-prop-valid",
        ":@emotion--stylis",
        ":@emotion--unitless",
        ":babel-plugin-styled-components",
        ":css-to-react-native",
        ":hoist-non-react-statics",
        ":shallowequal",
        ":supports-color",
    ],
)

yarn_library(
    name = "supports-color",
    package_name = "supports-color",
    version = "7.2.0",
    deps = [
        ":has-flag",
    ],
)

yarn_library(
    name = "supports-preserve-symlinks-flag",
    package_name = "supports-preserve-symlinks-flag",
    version = "1.0.0",
)

yarn_library(
    name = "swagger2openapi",
    package_name = "swagger2openapi",
    version = "7.0.8",
    deps = [
        ":call-me-maybe",
        ":node-fetch",
        ":node-fetch-h2",
        ":node-readfiles",
        ":oas-kit-common",
        ":oas-resolver",
        ":oas-schema-walker",
        ":oas-validator",
        ":reftools",
        ":yaml",
        ":yargs",
    ],
)

yarn_library(
    name = "text-table",
    package_name = "text-table",
    version = "0.2.0",
)

yarn_library(
    name = "timers-browserify",
    package_name = "timers-browserify",
    version = "2.0.12",
    deps = [
        ":setimmediate",
    ],
)

yarn_library(
    name = "to-arraybuffer",
    package_name = "to-arraybuffer",
    version = "1.0.1",
)

yarn_library(
    name = "to-fast-properties",
    package_name = "to-fast-properties",
    version = "2.0.0",
)

yarn_library(
    name = "to-regex-range",
    package_name = "to-regex-range",
    version = "5.0.1",
    deps = [
        ":is-number",
    ],
)

yarn_library(
    name = "toidentifier",
    package_name = "toidentifier",
    version = "1.0.1",
)

yarn_library(
    name = "tr46",
    package_name = "tr46",
    version = "0.0.3",
)

yarn_library(
    name = "tslib",
    package_name = "tslib",
    version = "2.3.1",
)

yarn_library(
    name = "tty-browserify",
    package_name = "tty-browserify",
    version = "0.0.0",
)

yarn_library(
    name = "type-check",
    package_name = "type-check",
    version = "0.3.2",
    deps = [
        ":prelude-ls",
    ],
)

yarn_library(
    name = "uglify-js",
    package_name = "uglify-js",
    version = "3.15.3",
)

yarn_library(
    name = "universalify",
    package_name = "universalify",
    version = "0.1.2",
)

yarn_library(
    name = "unpipe",
    package_name = "unpipe",
    version = "1.0.0",
)

yarn_library(
    name = "uri-js",
    package_name = "uri-js",
    version = "4.4.1",
    deps = [
        ":punycode",
    ],
)

yarn_library(
    name = "urijs",
    package_name = "urijs",
    version = "1.19.11",
)

yarn_library(
    name = "url",
    package_name = "url",
    version = "0.11.0",
    deps = [
        ":punycode",
        ":querystring",
    ],
)

yarn_library(
    name = "url-template",
    package_name = "url-template",
    version = "2.0.8",
)

yarn_library(
    name = "util",
    package_name = "util",
    version = "0.11.1",
    deps = [
        ":inherits",
    ],
)

yarn_library(
    name = "util-deprecate",
    package_name = "util-deprecate",
    version = "1.0.2",
)

yarn_library(
    name = "utility-types",
    package_name = "utility-types",
    version = "3.10.0",
)

yarn_library(
    name = "validate-npm-package-name",
    package_name = "validate-npm-package-name",
    version = "3.0.0",
    deps = [
        ":builtins",
    ],
)

yarn_library(
    name = "vm-browserify",
    package_name = "vm-browserify",
    version = "1.1.2",
)

yarn_library(
    name = "vm2",
    package_name = "vm2",
    version = "3.9.9",
    deps = [
        ":acorn",
        ":acorn-walk",
    ],
)

yarn_library(
    name = "webidl-conversions",
    package_name = "webidl-conversions",
    version = "3.0.1",
)

yarn_library(
    name = "whatwg-url",
    package_name = "whatwg-url",
    version = "5.0.0",
    deps = [
        ":tr46",
        ":webidl-conversions",
    ],
)

yarn_library(
    name = "wolfy87-eventemitter",
    package_name = "wolfy87-eventemitter",
    version = "5.2.9",
)

yarn_library(
    name = "word-wrap",
    package_name = "word-wrap",
    version = "1.2.3",
)

yarn_library(
    name = "wordwrap",
    package_name = "wordwrap",
    version = "1.0.0",
)

yarn_library(
    name = "wrap-ansi",
    package_name = "wrap-ansi",
    version = "7.0.0",
    deps = [
        ":ansi-styles",
        ":string-width",
        ":strip-ansi",
    ],
)

yarn_library(
    name = "wrappy",
    package_name = "wrappy",
    version = "1.0.2",
)

yarn_library(
    name = "ws",
    package_name = "ws",
    version = "7.5.7",
)

yarn_library(
    name = "xregexp",
    package_name = "xregexp",
    version = "2.0.0",
)

yarn_library(
    name = "xtend",
    package_name = "xtend",
    version = "4.0.2",
)

yarn_library(
    name = "y18n",
    package_name = "y18n",
    version = "5.0.8",
)

yarn_library(
    name = "yallist",
    package_name = "yallist",
    version = "3.1.1",
)

yarn_library(
    name = "yaml",
    package_name = "yaml",
    version = "1.10.2",
)

yarn_library(
    name = "yaml-ast-parser",
    package_name = "yaml-ast-parser",
    version = "0.0.43",
)

yarn_library(
    name = "yargs",
    package_name = "yargs",
    version = "17.4.0",
    deps = [
        ":cliui",
        ":escalade",
        ":get-caller-file",
        ":require-directory",
        ":string-width",
        ":y18n",
        ":yargs-parser",
    ],
)

yarn_library(
    name = "yargs-parser",
    package_name = "yargs-parser",
    version = "21.0.1",
)

yarn_binary(
    name = "openapi-cli",
    bin_name = "openapi",
    package = ":@redocly--openapi-cli",
)

yarn_binary(
    name = "redoc-cli",
    bin_name = "redoc-cli",
    package = ":redoc-cli-lib",
)

yarn_binary(
    name = "spectral-cli",
    bin_name = "spectral",
    package = ":@stoplight--spectral-cli",
)
