package(
    default_visibility = ["PUBLIC"],
    python_wheel_name_scheme = [
        "https://files.pythonhosted.org/packages/py3/{initial}/{package_name}/{package_name}-{version}-py3-none-any.whl",
        "https://files.pythonhosted.org/packages/py2.py3/{initial}/{package_name}/{package_name}-{version}-py2.py3-none-any.whl",
    ],
    python_wheel_repo = "use/your/own",
)

pip_library(
    name = "anyio",
    version = "3.5.0",
    deps = [
        ":idna",
        ":sniffio",
    ],
)

pip_library(
    name = "attrs",
    version = "21.4.0",
)

pip_library(
    name = "certifi",
    version = "2021.10.8",
    deps = [
    ],
)

pip_library(
    name = "charset_normalizer",
    version = "2.0.12",
    deps = [
    ],
)

pip_library(
    name = "click",
    version = "8.1.2",
    deps = [
    ],
)

pip_library(
    name = "colorama",
    version = "0.4.4",
)

pip_library(
    name = "curlify",
    version = "2.2.1",
    deps = [
        ":requests",
    ],
)

pip_library(
    name = "graphql-core",
    version = "3.2.1",
    deps = [
    ],
)

pip_library(
    name = "hypothesis",
    version = "6.43.3",
    zip_safe = False,
    deps = [
        ":attrs",
        ":sortedcontainers",
    ],
)

pip_library(
    name = "hypothesis_graphql",
    version = "0.5.1",
    deps = [
        ":attrs",
        ":graphql-core",
        ":hypothesis",
    ],
)

pip_library(
    name = "hypothesis_jsonschema",
    version = "0.22.0",
    deps = [
        ":hypothesis",
        ":jsonschema",
    ],
)

pip_library(
    name = "idna",
    version = "3.3",
)

pip_library(
    name = "iniconfig",
    version = "1.1.1",
)

pip_library(
    name = "jsonschema",
    version = "4.4.0",
    deps = [
        ":attrs",
        ":pyrsistent",
        ":typing_extensions",
    ],
)

pip_library(
    name = "junit-xml",
    version = "1.9",
    deps = [
        ":six",
    ],
)

pip_library(
    name = "multidict",
    version = "6.0.2",
    deps = [
    ],
)

pip_library(
    name = "packaging",
    version = "21.3",
    deps = [
        ":pyparsing",
    ],
)

pip_library(
    name = "pluggy",
    version = "1.0.0",
    deps = [
    ],
)

pip_library(
    name = "py",
    licences = ["..."],
    version = "1.11.0",
)

pip_library(
    name = "pyparsing",
    version = "3.0.8",
    deps = [
    ],
)

pip_library(
    name = "pyrsistent",
    version = "0.18.1",
)

pip_library(
    name = "pytest",
    version = "7.1.1",
    deps = [
        ":attrs",
        ":iniconfig",
        ":packaging",
        ":pluggy",
        ":py",
        ":tomli",
    ],
)

pip_library(
    name = "pytest_subtests",
    version = "0.7.0",
    deps = [
        ":pytest",
    ],
)

pip_library(
    name = "pyyaml",
    version = "6.0",
)

pip_library(
    name = "requests",
    version = "2.27.1",
    deps = [
        ":certifi",
        ":charset_normalizer",
        ":idna",
        ":urllib3",
    ],
)

pip_library(
    name = "schemathesis",
    version = "3.14.1",
    deps = [
        ":attrs",
        ":click",
        ":colorama",
        ":curlify",
        ":hypothesis",
        ":hypothesis_graphql",
        ":hypothesis_jsonschema",
        ":jsonschema",
        ":junit-xml",
        ":pytest",
        ":pytest_subtests",
        ":pyyaml",
        ":requests",
        ":starlette",
        ":tomli",
        ":tomli-w",
        ":typing_extensions",
        ":werkzeug",
        ":yarl",
    ],
)

pip_library(
    name = "six",
    version = "1.16.0",
)

pip_library(
    name = "sniffio",
    version = "1.2.0",
)

pip_library(
    name = "sortedcontainers",
    version = "2.4.0",
)

pip_library(
    name = "starlette",
    version = "0.19.0",
    deps = [
        ":anyio",
        ":typing_extensions",
    ],
)

pip_library(
    name = "tomli",
    version = "2.0.1",
)

pip_library(
    name = "tomli-w",
    version = "1.0.0",
)

pip_library(
    name = "typing_extensions",
    version = "4.2.0",
)

pip_library(
    name = "urllib3",
    version = "1.26.9",
)

pip_library(
    name = "werkzeug",
    version = "2.1.1",
)

pip_library(
    name = "yarl",
    version = "1.7.2",
    deps = [
        ":idna",
        ":multidict",
        ":typing_extensions",
    ],
)
