# 📒 **QuickNote**

> Create and share notes quickly and easily.

[![GitHub License](https://img.shields.io/github/license/Sn0wo2/QuickNote)](LICENSE)
![GitHub Downloads (all assets, all releases)](https://img.shields.io/github/downloads/Sn0wo2/QuickNote/total)

[![Build](https://github.com/Sn0wo2/QuickNote/actions/workflows/build.yml/badge.svg)](https://github.com/Sn0wo2/QuickNote/actions/workflows/Build.yml)
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

## 🎉 **Demo**

| Status | `Developing` |
|--------|--------------|

---

## 📚 **Docs**

| Status | `Developing` |
|--------|--------------|

---

## 📥 **Download**

[![GitHub release](https://img.shields.io/github/v/release/Sn0wo2/QuickNote?logo=github)](https://github.com/Sn0wo2/QuickNote/releases)

---

## ⚙️ **Build Instructions**

### ✅ **Using GitHub Actions**

Check:

* [`build.yml`](https://github.com/Sn0wo2/QuickNote/blob/main/.github/workflows/build.yml)
* [`.goreleaser.yml`](https://github.com/Sn0wo2/QuickNote/blob/main/LICENSE)

---

### 🔧 **Manual Build**

```bash
go build -mod=readonly -trimpath \
  -tags="mysql postgres sqlite sqlserver" \
  -o="QuickNote" \
  -ldflags="-s -w -buildid= -extldflags=-static" \
  -gcflags="all=-d=ssa/check_bce/debug=0" \
  -asmflags="-trimpath" main.go

cd Frontend

bun install
bun run build

mv dist/* ../

cd ../ && ./QuickNote(.exe)
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