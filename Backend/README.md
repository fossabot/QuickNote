# QuickNote

Create and share notes quickly and easily.

## Frontend

> [Sn0wo2/QuickNote-WEB](https://github.com/Sn0wo2/QuickNote-WEB)

---

[![License](https://img.shields.io/badge/license-GPL3.0-green.svg)](LICENSE)
[![ci](https://github.com/Sn0wo2/QuickNote/actions/workflows/ci.yml/badge.svg)](https://github.com/Sn0wo2/QuickNote/actions/workflows/ci.yml)
[![lint](https://github.com/Sn0wo2/QuickNote/actions/workflows/lint.yml/badge.svg)](https://github.com/Sn0wo2/QuickNote/actions/workflows/lint.yml)
[![go-release](https://github.com/Sn0wo2/QuickNote/actions/workflows/release.yml/badge.svg)](https://github.com/Sn0wo2/QuickNote/actions/workflows/release.yml)
[![Dependabot Updates](https://github.com/Sn0wo2/QuickNote/actions/workflows/dependabot/dependabot-updates/badge.svg)](https://github.com/Sn0wo2/QuickNote/actions/workflows/dependabot/dependabot-updates)


---

* Progress: `Developing`

--- 

## Features

- [X] No login
- [X] High performance
- [ ] Encryption
- [ ] Compression
- [X] Markdown preview support
- [ ] Dark mode
- [X] Note Sharing
- [ ] Note history

### Support Database:

- MySQL, MariaDB, TIDB, Aurora
- PostgreSQL, CockroachDB, AlloyDB
- SQLite3
- Microsoft SQL Server

---

## Docs

- `Developing`
---

## Build Instructions

```bash
# Install dependencies (if not already installed)
npm install

# Build the production-ready frontend files
npm run build
````

After building the frontend, move the generated static files to the backend's `./static` directory. This ensures the backend can serve the frontend correctly.

```bash
# Create static directory if it doesn't exist
mkdir -p ./static

# Copy the build output to the static directory
cp -r ./frontend/dist/* ./static/
```

```bash
# Compile the Go backend
go build -trimpath -o="QuickNote.exe" -ldflags="-s -w -buildid=" main.go
```

> **Note:**
>
> * Make sure the `./static` directory is in the same directory as your backend executable.
> * Adjust the frontend build output path (`./frontend/dist`) to match your actual project structure.

After these steps, run your backend executable (`./QuickNote`), and it will serve the frontend from the `./static` directory.

---

## Contributors

![Contributors](https://contrib.rocks/image?repo=Sn0wo2/QuickNote)

---

## Star History

<a href="https://www.star-history.com/#Sn0wo2/QuickNote&Date">
 <picture>
   <source media="(prefers-color-scheme: dark)" srcset="https://api.star-history.com/svg?repos=Sn0wo2/QuickNote&type=Date&theme=dark" />
   <source media="(prefers-color-scheme: light)" srcset="https://api.star-history.com/svg?repos=Sn0wo2/QuickNote&type=Date" />
   <img alt="Star History Chart" src="https://api.star-history.com/svg?repos=Sn0wo2/QuickNote&type=Date" />
 </picture>
</a>