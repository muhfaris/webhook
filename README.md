# Webhook Documentation

## Table of Contents

- [Introduction](#introduction)
- [Features](#features)
- [Use Cases](#use-cases)
- [Installation](#installation)
- [Configuration](#configuration)
- [Usage](#usage)
- [Hash Endpoint](#hash-endpoint)
- [Docker](#docker)
- [Conclusion](#conclusion)

## Introduction

**Webhook** is a lightweight and user-friendly webhook server designed for executing commands in response to webhooks and generating SHA-1 hashes of strings.

## Features

Webhook offers the following key features:

- Execute commands in response to webhooks
- Generate SHA-1 hashes of strings
- Lightweight and easy to use
- Easy to configure

## Use Cases

Webhook is versatile and can be employed in various scenarios, such as:

- [x] Automating deployments and CI/CD workflows
- [x] Triggering builds and tests
- [ ] Sending notifications and alerts
- [x] Running custom scripts and tasks

## Installation

To install Webhook, execute the following command:

```bash
go get github.com/muhfaris/webhook
```

## Configuration

Once Webhook is installed, configure it by creating a `config.yaml` or `config.json` file in the same directory as the `webhook` executable.

### YAML Configuration Example

```yaml
port: 8080
webhooks:
  - id: dev-user-svc
    token: token
    execute_command: "docker service update"
    workdir: ""
    command_arguments:
      - name: "--image"
        source: "data.docker_image"
      - name: "captain-nginx"
        source: ""
```

### JSON Configuration Example

```json
{
  "port": 8080,
  "webhooks": [
    {
      "id": "dev-user-svc",
      "token": "token",
      "workdir": "",
      "execute_command": "docker service update",
      "command_arguments": [
        {
          "source": "data.docker_image",
          "name": "--image"
        },
        {
          "source": "",
          "name": "captain-nginx"
        }
      ]
    }
  ]
}
```

#### Configuration Parameters

- `port`: Specifies the port for Webhook to listen on.
- `webhooks`: Defines the accepted webhooks, each with `id`, `token`, `execute_command`, and `workdir` parameters.
- `execute_command`: Specifies the command executed upon receiving a webhook.
- `workdir`: Defines the working directory for executing the command.
- `command_arguments`: An optional list of command arguments, each with a `name` and `source` parameter for customizing command execution.

## Usage

Start the Webhook server with the following command:

```bash
go run main.go api
```

To specify a configuration file (e.g., `config.yaml`), use the `--config` parameter:

```bash
go run main.go api --config config.yaml
```

Send a webhook to the server using `curl` as shown below:

```bash
curl -X POST \
  -H "Content-Type: application/json" \
  -d '{
    "id": "dev-user-svc",
    "token": "hash-token",
    "data": {
      "docker_image": "image"
    }
  }' \
  http://localhost:8080/webhook
```

If the webhook is valid, the server will execute the configured command.

## Hash Endpoint

The Webhook server offers a hash endpoint to generate SHA-1 hashes of strings. Use the following `curl` command:

```bash
curl -X POST \
  -H "Content-Type: application/json" \
  -d '{
    "message": "my message"
  }' \
  http://localhost:8080/hash
```

The server will respond with the SHA-1 hash of the provided message.

## Docker

You can run Webhook using Docker with the following command:

```bash
docker run --rm -v /var/run/docker.sock:/var/run/docker.sock -v ./config.yaml:/app/config.yaml --publish 8080:8080 webhook:v1 api --config /app/config.yaml
```

## Conclusion

Webhook is a lightweight and accessible webhook server suitable for automating a wide range of tasks and processes. It is a valuable tool for DevOps teams and individuals looking to streamline and automate their workflows.
