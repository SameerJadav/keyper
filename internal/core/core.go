package core

const SET_CMD_USAGE = `Usage: keyper set <project> [flags] [key=value...]

Set environment variables for a project.

Arguments:
  <project>    Name of the project (required)

Flags:
  -e, --environment string    Specify the environment (e.g., dev, staging, prod) (default "dev")
  -f, --file string           Path to a .env file to load variables from
  --overwrite                 Overwrite existing variables without warning

Environment Variables:
  key=value    Specify environment variables directly (can be multiple)

Examples:
  keyper set myapp DB_HOST=localhost DB_PORT=5432
  keyper set myapp -e prod --file prod.env`

const GET_CMD_USAGE = `Usage: keyper get <project> [flags]

Retrieve environment variables for a project.

Arguments:
  <project>    Name of the project (required)

Flags:
  -e, --environment string    Specify the environment (e.g., dev, staging, prod)
                              If not specified, variables for all environments will be retrieved
  -o, --out string            Specify the output file path
                              If not specified, variables will be printed to stdout

Examples:
  keyper get myapp
  keyper get myapp -e prod
  keyper get myapp -o .env
  keyper get myapp -e dev -o .env

Notes:
  - Only one project's environment variables can be retrieved at a time.
  - If no environment is specified, variables for all environments will be retrieved.
  - If no output file is specified, variables will be printed to the console.
  - The output format is compatible with .env files.`

const REMOVE_CMD_USAGE = `Usage: keyper remove <project> [flags] <key...>

Remove environment variables from a project.

Arguments:
  <project>    Name of the project (required)
  <key...>     One or more keys to remove (required)

Flags:
  -e, --environment string    Specify the environment (e.g., dev, staging, prod)
                              If not specified, keys will be removed from all environments
  -f, --force                 Force removal without confirmation

Examples:
  keyper remove myapp DB_HOST DB_PORT
  keyper remove myapp -e prod API_KEY SECRET_TOKEN
  keyper remove myapp --environment staging --force DEPRECATED_VAR

Notes:
  - At least one key must be specified for removal.
  - If no environment is specified, the keys will be removed from all environments.
  - By default, the command will ask for confirmation before removing keys.
  - Use the --force flag to skip the confirmation prompt.
  - This action cannot be undone, so use with caution.`

const PURGE_CMD_USAGE = `Usage: keyper purge <project> [flags]

Remove all environment variables for a project or a specific environment.

Arguments:
  <project>    Name of the project to purge (required)

Flags:
  -e, --environment string    Specify the environment to purge (e.g., dev, staging, prod)
                              If not specified, all environments for the project will be purged
  -f, --force                 Force purge without confirmation

Examples:
  keyper purge myapp
  keyper purge myapp -e prod
  keyper purge myapp --environment staging --force

Notes:
  - If no environment is specified, all environments for the project will be purged.
  - By default, the command will ask for confirmation before purging.
  - Use the --force flag to skip the confirmation prompt.
  - This action cannot be undone, so use with caution.`
