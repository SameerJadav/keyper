# Keyper

Keyper is a CLI tool for effortlessly managing your environment variables.
Save environment variables locally and retrieve them with just one command.
Built using only the standard libraries of Go. Keyper is simple, useful, and blazingly fast.

# Table of Contents

- [Installation](#installation)
- [Usage](#usage)
  - [Save environment variables](#save-environment-variables)
  - [Retrieve environment variables](#retrieve-environment-variables)
  - [Remove environment variables](#remove-environment-variables)
  - [Purge project or environment](#purge-project-or-environment)
- [Contributing](#contributing)
- [License](#license)

## Installation

To install Keyper, ensure you have Go installed on your machine.
If you haven't installed Go yet, you can follow the [official Go installation instructions](https://go.dev/doc/install).

```
go install github.com/SameerJadav/keyper@latest
```

This will install a Go binary that automatically binds to your `$GOPATH`.

## Usage

Keyper is designed for simplicity. It's a [grug brain tool](https://grugbrain.dev/#grug-on-tools).

### Save environment variables

```
keyper set todo-list -e prod -f .env
```

This command sets environment variables for the "todo-list" project.
It uses the production environment (-e prod) and loads variables from .env file (-f .env).
You can also add `key=value` pairs directly in the command instead of using a file.
Run `keyper set --help` for more information.

### Retrieve environment variables

```
keyper get todo-list -e dev -o .env
```

This command retrieves environment variables for the "todo-list" project.
It fetches variables from the development environment (-e dev) and saves them to a file named .env (-o .env).
If -e is not specified, it retrieves variables for all environments.
If -o is not specified, variables are printed to the console.
Run `keyper get --help` for more information.

### Remove environment variables

```
keyper remove todo-list -force -e prod API_KEY SECRET_TOKEN
```

This command removes specified environment variables (API_KEY and SECRET_TOKEN) from the "todo-list" project in the production environment (-e prod).
The -force flag skips the confirmation prompt.
If -e is omitted, it removes the variables from all environments.
Run `keyper remove --help` for more information.

**Note: Use with caution as this action cannot be undone.**

### Purge project or environment

```
keyper purge todo-list -force -e prod
```

This command purges the "todo-list" project's production environment (-e prod).
The -force flag skips the confirmation prompt.
If -e is omitted, it purges the entire project and all its environments.
Run `keyper purge --help` for more information.

**Note: Use with caution as this action cannot be undone.**

## Contributing

If you'd like to contribute to Keyper, please feel free to open an issue or submit a pull request. I welcome your suggestions and improvements!

## License

Keyper is open-source and available under the [MIT License](./LICENSE).
