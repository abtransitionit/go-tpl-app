# tpl-repo-go

This repository serves as a standardized template for all future GitHub GO projects within the organization.  
----
[![Main CI](https://github.com/abtransitionit/tpl-repo-go/actions/workflows/ci-main.yaml/badge.svg?branch=main)](https://github.com/abtransitionit/tpl-repo-go/actions/workflows/ci-main.yaml)
[![LICENSE](https://img.shields.io/badge/license-Apache_2.0-blue.svg)](https://choosealicense.com/licenses/apache-2.0/)

----


# Repository Contribution & Governance

This repository includes the following standard governance and documentation components:

| Component | Description |
| - | - |
| Licensing | Predefined open-source license. |
| [Code of Conduct](.github/CODE_OF_CONDUCT.md) | A community standards for all participants. |
| [Contributing Guide](.github/CONTRIBUTING.md) | Explains how to contribute, including reporting issues, submitting pull requests, and development workflow. |
| Continuous Integration | Automated build, vet, and test executed on each push and pull request. |
| README | This document. Provides project overview, purpose, structure, and onboarding information for users and contributors. |


# Getting Started  

## Create a repository from this template (e.g. `mygoprj`)
**Step 1. create and clone**

- on `github.com` create `mygoprj` : an empty git repo without `README` and `.gitcore`
- git clone the tpl repo: `tpl-repo-go` into `mygoprj`
```shell
git clone https://github.com/abtransitionit/tpl-repo-go.git mygoprj
```
- reset history and init repo
```shell
cd mygoprj
rm -rf .git
git init -b main  
```
- update `GO` path and version in the file `go.mod`
```shell
# do update
go mod init github.com/abtransitionit/mygoprj
go 1.24.2
# check update
cat go.mod
```

- update the following files if necessary:
  - `.git/config`
  - `.gitignore`

- commit the code "initial setup from template"
```shell
git remote add origin https://github.com/abtransitionit/mygoprj.git
```
- push the code
```shell
git push -u origin main
```

**Step 2. Update the README**

- update `tpl-repo-go` to `mygoprj`
- review each sections and update/add content when needed

# Continuous Integration

This repository uses GitHub Actions (Workflow file: `.github/workflows/ci.yaml`) for automated tasks and quality control.

The CI pipeline performs:

- module validation (`go mod tidy`)
- static analysis (`go vet`)
- tests execution (`go test`)



