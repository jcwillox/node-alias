# Node Alias

[![GitHub Release](https://img.shields.io/github/v/release/jcwillox/node-alias?style=flat-square)](https://github.com/jcwillox/node-alias/releases/latest)
[![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/jcwillox/node-alias?style=flat-square&label=go)](https://github.com/jcwillox/node-alias/blob/main/go.mod)
[![License](https://img.shields.io/github/license/jcwillox/node-alias?style=flat-square)](https://github.com/jcwillox/node-alias/blob/main/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/jcwillox/node-alias?style=flat-square)](https://goreportcard.com/report/github.com/jcwillox/node-alias)

A fast, lightweight proxy for Node.js package managers that works across `npm`, `yarn`, `pnpm`, and `bun`. It provides smart shortcuts for common commands and seamless integration with your shell.

## Features

- **Universal Package Manager Support**: Works with `npm`, `yarn`, `pnpm`, and `bun` without any configuration.
- **Smart Shortcuts**: Quickly run common commands like install, add, remove, and run scripts with minimal typing.
- **Automatic Tool Detection**: Automatically uses tools like `taze` for outdated checks and `npkill` for cleaning `node_modules` if they are installed.
- **File Execution**: Directly execute JavaScript, TypeScript, and Python files with appropriate runtimes.

## Installation

You can download the latest release from the [releases page](https://github.com/jcwillox/node-alias/releases/latest). Or use one of the following methods:

### Script

```bash
sh -c "$(curl -fsSL jcwillox.com/l/node-alias)"
```

Windows

```powershell
iwr -useb jcwillox.com/l/node-alias-ps1 | iex
```

## Usage

### Set up an Alias

The power of `node-alias` comes from setting up shell aliases. At minimum, set up the `n` alias:

**Bash/Zsh/Fish:**
```bash
alias n="node-alias"
```

You can also add additional shorthands to save even more characters:

```bash
alias ni="node-alias install"
alias nr="node-alias run"
# alias just for bun pkg manager
alias b="NODE_ALIAS_MANAGER=bun node-alias"
```

**PowerShell:**
```powershell
New-Alias n node-alias.exe
```

Once set up, you can use `n` instead of typing out your full package manager command!

### Common Commands

Using the `n` alias, you can run package manager commands quickly:

```bash
n install                # Install dependencies
n i                      # Same as install
n add lodash             # Add a dependency
n a lodash               # Same as add
n a -d lodash            # Add a dev dependency and save exact version
n remove lodash          # Remove a dependency
n rm lodash              # Same as remove
n list                   # List installed packages
n ls                     # Same as list
n run dev                # Run a npm script
n r dev                  # Same as run
```

### Smart Shortcuts

`node-alias` includes intelligent shortcuts that automatically run the right tool:

- **`n o`** – Check for outdated packages
  - Uses [taze](https://github.com/antfu/taze) if available, otherwise falls back to the package manager's outdated command

- **`n k`** – Clean up node_modules
  - Uses [npkill](https://github.com/voidcosmos/npkill) if installed, otherwise forwards the command line to the package manager

### Proxy and Execute Files

`node-alias` can also directly execute Node.js files, and figure out the right runtime to use.

```bash
n script.js              # Run any .js file
n script.ts              # Run TypeScript files (using tsx or bun if available, or node if supported)
```

And Python files:

```bash
n script.py              # Run Python (using uv or python)
```

### Use a Specific Package Manager

`node-alias` automatically detects your project's package manager (npm, yarn, pnpm, or bun) and uses the correct one. You can also manually override it:

```bash
# Use bun explicitly
b install                # Uses bun (requires `alias b="NODE_ALIAS_MANAGER=bun node-alias"`)
```

### Shell Completions

`node-alias` provides intelligent completions for multiple shells, for linux packages (deb, rpm, apk) these are built-in. Otherwise, up-to-date copies are located in the [completions/](./completions) directory.

You can also generate them yourself with the following command:

```bash
node-alias completion [bash|zsh|fish|pwsh]
```

You can write this to a file then source it in your shell configuration.

