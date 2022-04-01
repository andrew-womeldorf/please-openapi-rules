package(default_visibility = ["PUBLIC"])

# Version 6 is not published to Maven yet, and is considered unstable.
# By default, this plugin uses version 6, but version 5.4.0 is kept here for reference.
# If the implementing project would prefer to use v5.4.0, they can change the tool used when declaring the plugin.
# See /.plzconfig
maven_jar(
    name = "openapi-generator-cli",
    id = "org.openapitools:openapi-generator-cli:5.4.0",
    hash = "f3ed312310e390324b33ba2ffff290ce812935207a1493ec5c098d0a441be51c",
    sources = False,
)

remote_file(
    name = "openapi-generator-cli-6",
    url = "https://oss.sonatype.org/content/repositories/snapshots/org/openapitools/openapi-generator-cli/6.0.0-SNAPSHOT/openapi-generator-cli-6.0.0-20220330.230127-109.jar",
    hashes = ["e54795d354d3ed3f64797630a850e06ba89f56f2295bf9dc7733517f251e124a"],
    out = "openapi-generator-cli-6.jar"
)
