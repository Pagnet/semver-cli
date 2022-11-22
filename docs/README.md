# Semver-Cli Manual

Example command using cobra and cobradoc.

```
semver-cli [command] [global flags] [command flags]
```

### Global Flags

```
```

### Main Commands

* [semver-cli semver](#semver-cli-semver)

### Documentation Commands

* [semver-cli markdown](#semver-cli-markdown)

### Additional Commands

* [semver-cli completion](#semver-cli-completion)
* [semver-cli completion bash](#semver-cli-completion-bash)
* [semver-cli completion fish](#semver-cli-completion-fish)
* [semver-cli completion help](#semver-cli-completion-help)
* [semver-cli completion powershell](#semver-cli-completion-powershell)
* [semver-cli completion zsh](#semver-cli-completion-zsh)
* [semver-cli help](#semver-cli-help)
* [semver-cli version](#semver-cli-version)

# Main Commands

## `semver-cli semver`

Validation of semver

```
semver-cli semver [flags]
```

### Command Flags

```
      --e int[=1]    Equal (default 1)
      --eg int[=2]   Equal or greater (default 2)
      --el int[=3]   Equal or less (default 3)
      --g int[=4]    Greater (default 4)
  -h, --help         help for semver
      --l int[=5]    Less (default 5)
      --v1 string    Version 1
      --v2 string    Version 2
```

# Documentation Commands

## `semver-cli markdown`

Genereate markdown documentation

```
semver-cli markdown [flags]
```

### Command Flags

```
  -h, --help   help for markdown
```

# Additional Commands

## `semver-cli completion`

Generate the autocompletion script for semver-cli for the specified shell.
See each sub-command's help for details on how to use the generated script.


```
semver-cli completion [flags]
```

### Command Flags

```
  -h, --help   help for completion
```

## `semver-cli completion bash`

Generate the autocompletion script for the bash shell.

This script depends on the 'bash-completion' package.
If it is not installed already, you can install it via your OS's package manager.

To load completions in your current shell session:

	source <(semver-cli completion bash)

To load completions for every new session, execute once:

#### Linux:

	semver-cli completion bash > /etc/bash_completion.d/semver-cli

#### macOS:

	semver-cli completion bash > $(brew --prefix)/etc/bash_completion.d/semver-cli

You will need to start a new shell for this setup to take effect.


```
semver-cli completion bash
```

### Command Flags

```
  -h, --help              help for bash
      --no-descriptions   disable completion descriptions
```

## `semver-cli completion fish`

Generate the autocompletion script for the fish shell.

To load completions in your current shell session:

	semver-cli completion fish | source

To load completions for every new session, execute once:

	semver-cli completion fish > ~/.config/fish/completions/semver-cli.fish

You will need to start a new shell for this setup to take effect.


```
semver-cli completion fish [flags]
```

### Command Flags

```
  -h, --help              help for fish
      --no-descriptions   disable completion descriptions
```

## `semver-cli completion help`

Help about any command

```
semver-cli completion help [command] [flags]
```

### Command Flags

```
  -h, --help   help for help
```

## `semver-cli completion powershell`

Generate the autocompletion script for powershell.

To load completions in your current shell session:

	semver-cli completion powershell | Out-String | Invoke-Expression

To load completions for every new session, add the output of the above command
to your powershell profile.


```
semver-cli completion powershell [flags]
```

### Command Flags

```
  -h, --help              help for powershell
      --no-descriptions   disable completion descriptions
```

## `semver-cli completion zsh`

Generate the autocompletion script for the zsh shell.

If shell completion is not already enabled in your environment you will need
to enable it.  You can execute the following once:

	echo "autoload -U compinit; compinit" >> ~/.zshrc

To load completions in your current shell session:

	source <(semver-cli completion zsh); compdef _semver-cli semver-cli

To load completions for every new session, execute once:

#### Linux:

	semver-cli completion zsh > "${fpath[1]}/_semver-cli"

#### macOS:

	semver-cli completion zsh > $(brew --prefix)/share/zsh/site-functions/_semver-cli

You will need to start a new shell for this setup to take effect.


```
semver-cli completion zsh [flags]
```

### Command Flags

```
  -h, --help              help for zsh
      --no-descriptions   disable completion descriptions
```

## `semver-cli help`

Help about any command

```
semver-cli help [command] [flags]
```

### Command Flags

```
  -h, --help   help for help
```

## `semver-cli version`

A version of software

```
semver-cli version [flags]
```

### Command Flags

```
  -h, --help   help for version
```
