# go-tpl-repo

This repository serves as a standardized template for all future GitHub GO projects within the organization.  

----
[![Main CI](https://github.com/abtransitionit/go-tpl-repo/actions/workflows/ci.yaml/badge.svg?branch=main)](https://github.com/abtransitionit/go-tpl-repo/actions/workflows/ci.yaml)
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
| CHANGELOG | Tracks project changes across versions. |
| README | This document. Provides project overview, purpose, structure, and onboarding information for users and contributors. |


# Getting Started  

## Create a repository from this template (e.g. `go-prjname`)
**Step 1. create and clone**

- on `github.com` create `go-prjname` : an empty git repo without `README` and `.gitcore`
- **locally**, git clone the template for `GO` projects:
```shell
git clone https://github.com/abtransitionit/go-tpl-repo.git go-prjname
```
- reset history and init repo
```shell
cd go-prjname
rm -rf .git
git init -b main  
```
- update `GO` path and version in the file `go.mod`
```shell
# do update
go mod init github.com/abtransitionit/go-prjname
go 1.26
# check update
cat go.mod
```

**Step 2. Update the config**

- update the following files if necessary:
  - `.git/config`
    ```shell
    git remote add origin https://github.com/abtransitionit/go-prjname.git
    ```
  - `.gitignore`
  - `.CGHANGELOG.md`


**Step 3. Update the README**

- update `tpl-repo-go` to `mygoprj`
- review each sections and update/add content to match the repo purpose

**Step 4. test the code**
```
go mod tidy
go vet ./...
go build ./...
```

**Step 5. Commit the code**

- 
```
git add .
git commit -m "initial setup from template go-tpl-repo"
```

**Step 6. push the code**

- 
```shell
git push -u origin main
```


**Step 7. Update the README**

- update `go-tpl-repo` to `go-prjname`
- review each sections and update/add content when needed

**Step 8. check the changes**

```sh
go vet ./...
go test ./...
go build ./...
```

**Step 9. push the code**

```shell
git push -u origin main
```

**Step 10. tag the foundation**

```shell
# create the tag
git tag -a v0.1.0 -m "first tag to freeze the foundation"

# push the tag
git push origin v0.1.0
```

# Nest practice
**managing errors**
```go
// generic
fmt.Errorf("<operation> <resource>: %w", err)

// example
fmt.Errorf("read file %q: %w", path, err)
fmt.Errorf("parse yaml %q: %w", path, err)
fmt.Errorf("open database %q: %w", dsn, err)
```

see the `CI` in action on github .

# Continuous Integration

This repository uses GitHub Actions (Workflow file: `.github/workflows/ci.yaml`) for automated tasks and quality control.

The CI pipeline performs:

- module validation (`go mod tidy`)
- static analysis (`go vet`)
- tests execution (`go test`)



