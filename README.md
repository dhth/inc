# inc

✨ Overview
---

`inc` gives you the next semver for your repo, based on git tags.

```bash
$ inc
# 2.37.0 👉 2.37.1
```

💾 Installation
---

**homebrew**:

```sh
brew install dhth/tap/inc
```

**go**:

```sh
go install github.com/dhth/inc@latest
```


⚡️ Usage
---

`inc` is dead simple to use.

```bash
inc

# will print out the next version as follows, and will also copy it to your
# system clipboard
# 2.37.0 👉 2.37.1

# flags
inc -type='minor'
inc -copy-version=false
inc -print-version=false
inc -print-fancy=false
```
