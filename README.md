# 🗄 filer

> A web application to make filing scanned PDF documents simpler.

## Use Case

We use a scanning app on our smartphones to scan documents to an inbox on a local file server at home. The `filer` web application shows an inbox document and provides a set of pre-configured rules to categorize it. Selecting a rule moves the document to its final destination in the document archive directory structure and the next inbox document is displayed.

![process](./docs/process.png)

### Features

- Preview PDF document from an inbox directory
- Move inbox document, based on a set of configurable rules, to another directory
- Rename inbox document before moving it to the document archive
- Delete inbox document without moving

### The Story aka "Why!?"

We went paperless some years ago and made scanning receipts and other documents a habit. To make the scanning as effortless as possible, we decided to upload the resulting PDF files automatically to a shared inbox directory on our file server. Unfortunately, filing and categorizing the scans afterwards turned out to be too cumbersome to do regularly: Sitting down and moving files from one folder to another was just nothing we enjoy that much 😇

`filer` makes this task less tedious for us: Open the web application, have a glance on the PDF preview, select one of the preset actions, repeat until the inbox is empty.

## Usage

1. Clone this repository.
2. Install dependencies: `make install`
3. Build the server: `make build-server`
4. Build the UI: `make build-ui`
5. Build a Docker image using the provided `Dockerfile.`
6. Create a `rules.yml` file with your categorization rules of choice (see Rules section below).
7. Start Docker container, see Environment Variables section below for available configuration options.

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
make install      # Instal dependencies
make start-ui     # Start rollup in watch mode
make start-server # Start http server
```

### Publish a new Version

The `make` target `publish-docker-image` builds the `Dockerfile` using `buildx` for multiple CPU architectures and publishes the resulting Docker image. Update the `DOCKER_REPO` variable in `Makefile` to your own Docker repository destination.

```shell
git tag vX.Y.Z  # Optional
VERSION=X.Y.Z make publish-docker-image
```
