# Boot.Dev Gator in Go

Gator is a CLI-based RSS feed aggregator written in Go. It allows users to register, follow their favorite RSS feeds, and run a background scraper to keep track of new content.

## Prerequisites

Before installing and running Gator, ensure you have the following dependencies installed on your system:

- **Go**: Version `1.25.4` or higher.
- **PostgreSQL**: A running Postgres instance to store users, feeds, and posts.

---

## Installation

You can install the `gator` binary directly using Go's install command. Run the following in your terminal:

```bash
go install github.com/cmilliron/bootdev-gator-go@latest

```

> **Note:** Ensure your `$GOPATH/bin` (usually `~/go/bin`) is added to your system's `PATH` environment variable so you can run the `gator` command from anywhere.

---

## Configuration & Setup

Gator relies on a configuration file located in your user's root (home) directory to know how to connect to your database and keep track of the current user.

1. Locate the `gatorconfig.json.sample` file in the project repository.
2. Copy it to your root user directory named as `.gatorconfig.json`:

```bash
cp gatorconfig.json.sample ~/.gatorconfig.json

```

3. Open `~/.gatorconfig.json` and update it with your Postgres connection string and database credentials.

---

## Usage & Commands

Once configured, you can interact with Gator using the commands detailed below.

### Public Commands

These commands can be run by anyone and do not require an active login session.

| Command    | Usage                       | Description                                                                                                                                |
| ---------- | --------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------ |
| `register` | `gator register <username>` | Registers a new user in the system.                                                                                                        |
| `login`    | `gator login <username>`    | Switches the current active user in the config file.                                                                                       |
| `users`    | `gator users`               | Lists all registered users.                                                                                                                |
| `reset`    | `gator reset`               | **Warning:** Resets/clears all data from the database.                                                                                     |
| `feeds`    | `gator feeds`               | Lists all feeds available in the system.                                                                                                   |
| `agg`      | `gator agg <duration>`      | Starts the continuous feed scraper. The `<duration>` (e.g., `1m`, `1h`) defines the pause between pings to prevent DOSing the RSS servers. |

### Protected Commands

These commands require a user to be actively logged in via the `login` command.

| Command     | Usage                        | Description                                                                                   |
| ----------- | ---------------------------- | --------------------------------------------------------------------------------------------- |
| `addfeed`   | `gator addfeed <name> <url>` | Adds a new RSS feed to the database and automatically marks the current user as a follower.   |
| `follow`    | `gator follow <url>`         | Allows the current user to follow an existing feed via its URL.                               |
| `following` | `gator following`            | Lists all the RSS feeds the current user is actively following.                               |
| `unfollow`  | `gator unfollow <url>`       | Allows the current user to unfollow a specific feed.                                          |
| `browse`    | `gator browse [limit]`       | Browses posts from followed feeds in the database. `<limit>` is optional and defaults to `2`. |
