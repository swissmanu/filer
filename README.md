# 🗄 filer

> A web application to make filing scanned PDF documents simpler.

## Configuration

### Environment Variables

| Environment Variable | Default       | Description                                                                                                                  |
| -------------------- | ------------- | ---------------------------------------------------------------------------------------------------------------------------- |
| `FILER_ADDR`         | `:8000`       | A network interface and port where filer will provide its API and UI via HTTP.                                               |
| `FILER_INBOX_PATH`   | `./inbox`     | Path to the inbox directory.                                                                                                 |
| `FILER_DATA_PATH`    | `./data`      | Path to the data directory. Rule target paths are always evaluated relative to the data path.                                |
| `FILER_RULES_PATH`   | `./rules.yml` | Path to a YAML file containing rule definitions.                                                                             |
| `FILER_UI_PATH`      | `./ui`        | Path to filers web user interface. This variable is useful during development; you can ignore it in productive environments. |
| `UMASK_SET`          | `-022`        | Set the `umask` value for files created by the filer application.                                                            |

### Rules

```yaml
rules:
	# Moves a document from the inbox to "FILER_DATA_PATH/Receipts":
  - name: "Receipts"
  	description: "Receipts and bills"
    actions:
      - type: "move"
        target: "Receipts"
	# Moves a document from the inbox to "FILER_DATA_PATH/Insurances/Health":
  - name: "Health Insurance"
  	description: "All documents related to our health insurance"
    actions:
      - type: "move"
        target: "Insurances/Health"
```

## Development

### Start development

```shell
make start-ui			# Start rollup in watch mode
make start-server # Start http server
```

### Publish a new Version

```shell
git tag vX.Y.Z
VERSION=X.Y.Z make publish-docker-image
```
