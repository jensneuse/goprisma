# Linux Docker Example

## Prerequisites

Run the test database

## Build the application

```shell
docker build
```

## Run

```shell
docker run xxx
```

xxx = id of the build cmd

## Notes

In case you're not running on macOS, change "host.docker.internal" to point to the correct IP of the test database.
You might have to extend the "docker run" cmd to link to the database container.