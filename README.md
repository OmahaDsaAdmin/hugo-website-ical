# Hugo Website iCal

Orgs can use this project to generate .ics / iCal files from a list of Action Network events given the following:

1. Export Path (where to store the .ics files)
2. One (or more) Action Network API Keys

Running the script will create a single .ics file containing every Action Network event from all given API keys. That file can be added / hosted from a hugo website (preferably in a static location) and refreshed by other calendar apps automatically.

### Overwriting Existing

The ics file is overwritten every time this program runs.

## How to run locally

1. `go build`
2. `sudo ./hugo-website-ical ./exports first_api_key optional_next_api_key`

## How to run via Github Actions

Github Actions can run this program and commit any new events into your hugo repository directly. It is advised to do this once per day to save on monthly GA limits. To do this, you will need:

1. A Github Secret called `BOT_PAT`
2. An export path within the workspace (not local)
3. One (or more) Action Network API Keys (preferably stored as repo secrets)

See the `example-pull-events.txt` file for a `.yml` example of how to implement this action. It is necessary to store this entire repo (minus the git stuff) in your hugo repo under .github/scripts/hugo-website-events.
