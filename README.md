# Inodebeat

Beat to monitor inodes on a machine.


## How to run

The easiest way to launch inodebeat is to run it in a Docker container:

```
docker run codingame/inodebeat
```


## Configuring inodebeat

To override the default configuration, just link yours to `/etc/inodebeat/inodebeat.yml`:

```
docker run -d \
  -v /directory/where/your/config/file/is/:/etc/inodebeat \
  -v /:/hostfs:ro \
  --name inodebeat \
  codingame/inodebeat
```

Otherwise, you could create your own image with your custom configuration with a Dockerfile like:

```Dockerfile
FROM codingame/inodebeat

COPY inodebeat.yml /etc/inodebeat/inodebeat.yml
```


## Exported fields

Example output:

```json
{
  "inodes": {
    "directory": "/hostfs",
    "total": 13582336,
    "free": {
      "pct": 81.895677,
      "count": 11123346
    },
    "used": {
      "pct": 18.104323,
      "count": 2458990
    }
  }
}
```

To get a detailed list of all generated fields, please read the [fields documentation page](docs/fields.asciidoc).


## Contributing to the project

See [contributing instructions](CONTRIBUTING.md) to set up the project and build it on your machine.
