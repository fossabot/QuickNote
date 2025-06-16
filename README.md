# 📒 **QuickNote**

> Create and share notes quickly and easily.

[![License](https://img.shields.io/badge/license-GPL3.0-green.svg)](LICENSE)
[![Build](https://github.com/Sn0wo2/QuickNote/actions/workflows/Build.yml/badge.svg)](https://github.com/Sn0wo2/QuickNote/actions/workflows/Build.yml)
[![Dependabot Updates](https://github.com/Sn0wo2/QuickNote/actions/workflows/dependabot/dependabot-updates/badge.svg)](https://github.com/Sn0wo2/QuickNote/actions/workflows/dependabot/dependabot-updates)

---

## 🚀 **Project Status**

| Status | `Developing` |
| ------ | ------------ |

---

## ✅ **Features**

* ✔️ No login required
* ✔️ High performance
* ✔️ Simple UI
* ✔️ Markdown preview support
* ✔️ Dark Mode
* ✔️ Note sharing

**Planned:**

* 🔒 Encryption
* 📦 Compression
* 🕑 Note history

---

## 🗃️ **Supported Databases**

* **Relational:**

    * MySQL, MariaDB, TiDB, Aurora
    * PostgreSQL, CockroachDB, AlloyDB
    * SQLite3
    * Microsoft SQL Server

---

## 📚 **Docs**

> Documentation: `Developing`

---

## ⚙️ **Build Instructions**

### ✅ **Using GitHub Actions**

Check:

* [`Build.yml`](https://github.com/Sn0wo2/QuickNote/blob/main/.github/workflows/Build.yml)
* [`.goreleaser.yml`](https://github.com/Sn0wo2/QuickNote/blob/main/LICENSE)

---

### 🔧 **Manual Build**

```bash
# 1️⃣ Build Frontend
cd Frontend

bun run install
bun run build

# Rename dist → static and move it to Backend directory

# 2️⃣ Build Backend
cd ../Backend

go build -mod=readonly -trimpath \
  -o="QuickNote(.exe)" \
  -ldflags="-s -w -buildid= -extldflags=-static" \
  -gcflags="all=-d=ssa/check_bce/debug=0" \
  -asmflags="-trimpath" main.go

# 3️⃣ Run
./QuickNote(.exe)
```

---

## 👥 **Contributors**

![Contributors](https://contrib.rocks/image?repo=Sn0wo2/QuickNote)

---

## ⭐ **Star History**

<a href="https://www.star-history.com/#Sn0wo2/QuickNote&Date">
 <picture>
   <source media="(prefers-color-scheme: dark)" srcset="https://api.star-history.com/svg?repos=Sn0wo2/QuickNote&type=Date&theme=dark" />
   <source media="(prefers-color-scheme: light)" srcset="https://api.star-history.com/svg?repos=Sn0wo2/QuickNote&type=Date" />
   <img alt="Star History Chart" src="https://api.star-history.com/svg?repos=Sn0wo2/QuickNote&type=Date" />
 </picture>
</a>

---

## 📄 **License**

Licensed under [GPL 3.0](LICENSE).