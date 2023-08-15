# NOTE: This file is for HashiCorp specific licensing automation and can be deleted after creating a new repo with this template.
schema_version = 1

project {
  license        = "MPL-2.0"
  copyright_year = 2021

  header_ignore = [
    # documentation
    "docs/**/*.yaml",

    # examples used within documentation (prose)
    "examples/**",
    "tests/**",

    # skip the header for generated files
    "internal/provider/schema.go",
    "internal/resource/schema.go",
    "internal/datasource/schema.go",
    "internal/models/models.go",
    "internal/models/encoders.go",
    "internal/models/decoders.go",

    # GitHub issue template configuration
    ".github/ISSUE_TEMPLATE/*.yml",

    # golangci-lint tooling configuration
    ".golangci.yml",

    # GoReleaser tooling configuration
    ".goreleaser.yml",
  ]
}
