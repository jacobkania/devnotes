# WIP: Devnotes

```
A tool for efficient note-taking, built for those who are often working in the terminal.
```

Devnotes saves all notes within an SQLite database, and will allow creating and querying of notes using the `dn` command-line utility.

# Usage

Devnotes is currently a work-in-progress. Therefore, usage is intended towards development and debugging, for now.

### Prerequisites: 

1. Pull this repository
2. [Build the binary](#building), it will be saved in `bin/dn` (dn for DevNotes)

### Flags

* `-t` "today": will print all notes from today to the command line

### Writing a new note

You can plainly use the `dn` command, along with your note, to write something new.

Example: `dn Followed-up with John, decided to continue working on the project`

# Building

The commands for this project are contained within a Makefile.

The primary command to know for development is `make build`. Running that command in the root directory of the project will create the binary `dn` within the `bin` folder.

# Development

### Project structure

The `dn` command line utility is handled within `cmd/term`. All logic related to the command line utility lives there.

The `/config` folder contains the setup of the application config, which is persisted for the application within `~/.config/devnotes/config.yml`.

The `/db` folder contains everything the application needs for querying and storing records

The `/migrate` folder contains all database migrations. See [migrations](#migrations)

### Migrations

For now, the migrations are bundled within the binary when it is built. That way, when a user updates to a new version, all of their migrations will keep their database up to date.

To create one, run `make migrate` to begin a prompt which will generate the migration files.

Then, write the sql for the migration within the newly generated files.
