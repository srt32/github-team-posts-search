# go CLI to search GitHub Team Posts (Discussions)

## Motivation

Want to search for content in GitHub team posts? Use this CLI.

## Usage:

Get a PAT with access to read team discussions.

Then run:

```
GITHUB_TOKEN=<token> go run main.go -org=github -team=engineering -query=mysql
```

## Limitations

* The CLI looks at only the most recent 100 team posts.
* The query is checked directly against the content doing an exact string match.

## Notes

Written mainly with copilot.
