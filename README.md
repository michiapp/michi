# michi

A blazing-fast, local search multiplexer for your browser. Navigate the web with custom bangs, shortcuts, and session launchers, all powered by a tiny, self-hosted Go service.

## Features
- Shortcuts: Create, list, and delete shortcuts to quickly access frequently used URLs.
- Sessions: Create, list, and delete sessions to group related shortcuts together.
- Bangs: Create, list, and delete bangs to quickly access frequently used search queries.
- Local & Private: Your configurations are stored locally in a SQLite database, never leaving your machine.
- Blazing Fast: Runs as a tiny background service, providing instant redirects without any network latency or browser pop-up blockers.
- Cross-Platform: Built with Go, available for Linux, macOS (soon).

## Get Started

### 1. Installation

```bash
curl -fsSL https://michi.run/install.sh | sh
```

### 2. Start the Local Server

Run the `michi` HTTP server in the background:

```bash
michi serve
```

Or detach it to run it in the background:

```bash
michi serve --detach
```

The server will run on `http://localhost:5980` by default. And the default search page will be served on the root path `/`.
i.e. `http://localhost:5980/`

> Note that you can change this port through the configuration file, under ~/.michi/config.yaml

### 3. Configure Your Browser

Set `http://localhost:5980/search?q=%s` as your browser's default search engine.

Instructions for common browsers:
*   **Zen:** `Settings > Search > Search Shortcuts`
    *  Don't forget to set michi as your default search engine at the top of the page. 
*   **Chromium:** `Settings > Search engine > Manage search engines and site search > Add`
    *   **Search engine:** `michi`
    *   **URL with %s:** `http://localhost:5980/search?q=%s`
---

## Usage

Once configured, simply type into your browser's address bar:

- Bang Search:
```bash
!g my Go query
!yt epic jdm cars drifting
!gh michi
```

- Web Shortcut:
```bash 
#portal
#book
#repos
```

- Session Launcher:
```bash
@dev 
@learning
```

> All prefixes can be customized in the configuration file, under ~/.michi/config.yaml
---

## CLI Commands

`michi` offers commands to manage your bangs, shortcuts, sessions & it's lifecycle

```bash
michi serve
michi serve --detach
michi stop
michi doctor


michi shortcuts list 
michi shortcuts add --alias "michi" --url "https://github.com/OrbitalJin/michi"
...

michi sessions list 
michi sessions add --alias "sesh" \ 
                   --url "https://google.com" \ 
                   --url "https://github.com" 
...
```
> You can use `michi --help` to get a list of available commands

## Uninstall

if you want to uninstall the cli, you can run the following command

```bash
curl -fsSL https://michi.run/uninstall.sh | sh
```

> For data intergrity purposes, all user data stored under ~/.michi will be kept. Manually delete them if you want.

---

## Todo
- [ ] Fix macos installer
- [ ] michi commands i.e. :michi foo
- [ ] Analytics
- [ ] Firewall rules
- [ ] configurable redirects to avoid certain websites
- [ ] list down average reponse time in help command
- [x] Migrate to sqlc
- [x] Fix db changes not reflecting immediately
- [x] add shorthands for the cli commands
- [x] Add support for arm64 & darwin
- [x] serve default search page on / & forward to /search
- [x] CI/CD crossplatform build pipeline
- [x] curl & bash installer | uninstaller
- [x] Hydrate local user's copy of the database from embedded snapshot
- [x] Make sure to only store history in the local copy of the database
- [x] cli
- [x] Shortcuts e.g. repos => github.com/johndoe?tab=repositories
- [x] Bangs
- [x] History
- [x] Sessions
- [x] embedded templates
- [x] seperate router with templates & handlers
- [x] Setup database connection
- [x] Setup database migrations
- [x] scrape & dump duckduckgo's bang index into the relational db
- [x] Implement query & bang parsing 
- [x] Check bang matches against db and keep highest ranking one 
- [x] Implement service layer 
- [x] Implement url resolving
- [x] fix cors
- [x] Implement provider fallback
- [x] Speed it up
- [x] clean up api & router
- [x] implement caching using sync.Map
- [x] Implement features: shortcuts #, sessions @ and history $
- [x] Refactor config to use yaml 
- [x] build cli
- [x] Embed snapshot of the database & hydrate a local version on the user's machine

## cli 
- [x] Implement copy to clipboard
- [x] History, delete, copy, list 
- [x] Server lifecycle management
- [x] shortcuts list, add, delete
- [x] sessions list, add, delete
- [x] bangs list, add, Delete


## Features 
- [x] Schema (html/json or relational)
- [x] Repository
- [x] Service
- [x] Parser
- [x] Handler

## Shortcuts
- [x] Cache
- [x] Handler function
- [x] Service
- [x] Repository
- [x] Parse them seperately from bangs
- [x] Handler struct
- [x] Use interfaces for dependency injection
- [x] Determine precedence order

### History
- [x] Repository
- [x] Service
- [x] Go routine for db transactions

