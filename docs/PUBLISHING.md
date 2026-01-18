# Publishing Tardigrade-Mod to GitHub

## Step-by-Step Guide to Make Your Module Available

### 1. Initialize Git Repository (if not already done)

```bash
cd /home/ricardo/Programing/tardigrade-mod
git init
```

### 2. Create .gitignore File

```bash
cat > .gitignore << 'EOF'
# Binaries
*.exe
*.exe~
*.dll
*.so
*.dylib
*.test
*.out

# Database files (optional - remove if you want to include test dbs)
*.db

# IDE
.vscode/
.idea/
*.swp
*.swo
*~

# OS
.DS_Store
Thumbs.db

# Test coverage
*.coverprofile
coverage.html
EOF
```

### 3. Add All Files to Git

```bash
git add .
git status  # Review what will be committed
```

### 4. Commit Your Changes

```bash
git commit -m "Release v0.3.0 - Added flexible field support"
```

### 5. Create GitHub Repository

1. Go to https://github.com/new
2. Repository name: `tardigrade-mod`
3. Description: "A lightweight, file-based NoSQL database library for Go applications"
4. Make it **Public** (required for go get to work)
5. **DO NOT** initialize with README (you already have one)
6. Click "Create repository"

### 6. Link Local Repository to GitHub

Replace `YOUR_GITHUB_USERNAME` with your actual GitHub username:

```bash
git remote add origin https://github.com/YOUR_GITHUB_USERNAME/tardigrade-mod.git
git branch -M main
git push -u origin main
```

### 7. Create a Git Tag for Version 0.3.0

```bash
git tag v0.3.0
git push origin v0.3.0
```

### 8. Verify Module Path in go.mod

Your `go.mod` should have:

```
module github.com/YOUR_GITHUB_USERNAME/tardigrade-mod

go 1.20
```

If it doesn't match, update it:

```bash
# Update go.mod with correct path
go mod edit -module github.com/YOUR_GITHUB_USERNAME/tardigrade-mod
git add go.mod
git commit -m "Update module path"
git push
```

### 9. Test Installation from Another Project

```bash
# In a different directory
mkdir test-tardigrade
cd test-tardigrade
go mod init test-app

# Install your module
go get github.com/YOUR_GITHUB_USERNAME/tardigrade-mod@v0.3.0
```

### 10. Using Your Module in Other Projects

In any Go project:

```go
// go.mod
module myapp

go 1.20

require github.com/YOUR_GITHUB_USERNAME/tardigrade-mod v0.3.0
```

```go
// main.go
package main

import (
    "fmt"
    "github.com/YOUR_GITHUB_USERNAME/tardigrade-mod"
)

func main() {
    tar := tardigrade.Tardigrade{}
    tar.CreateDB("myapp.db")
    
    tar.AddFlexFieldVariadic("user:1", "myapp.db",
        "name", "John Doe",
        "email", "john@example.com")
    
    result := tar.SelectFlexByID(1, "json", "myapp.db")
    fmt.Println(result)
}
```

Then run:
```bash
go mod tidy
go run main.go
```

## Future Updates

When you make changes and want to release a new version:

```bash
# Make your changes
git add .
git commit -m "Description of changes"
git push

# Create new version tag
git tag v0.3.1
git push origin v0.3.1
```

Users can then update to the new version:
```bash
go get github.com/YOUR_GITHUB_USERNAME/tardigrade-mod@v0.3.1
```

## Important Notes

1. **Repository must be public** for `go get` to work without authentication
2. **Use semantic versioning**: v0.3.0, v0.3.1, v1.0.0, etc.
3. **Tag format**: Must start with 'v' (e.g., v0.3.0, not 0.3.0)
4. **Module path**: Must match your GitHub username and repo name
5. **Go proxy**: Changes may take a few minutes to appear on pkg.go.dev

## Verify Your Module

After publishing, check:
- https://pkg.go.dev/github.com/YOUR_GITHUB_USERNAME/tardigrade-mod
- It may take 10-15 minutes to appear on pkg.go.dev

## Quick Command Summary

```bash
# Initial setup
git init
git add .
git commit -m "Release v0.3.0 - Added flexible field support"
git remote add origin https://github.com/YOUR_GITHUB_USERNAME/tardigrade-mod.git
git branch -M main
git push -u origin main

# Tag the release
git tag v0.3.0
git push origin v0.3.0

# Future updates
git add .
git commit -m "Your changes"
git push
git tag v0.3.1
git push origin v0.3.1
```
