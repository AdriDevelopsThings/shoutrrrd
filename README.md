# shoutrrr daemon
Orginal repositories: [shoutrrr repository](https://github.com/containrrr/shoutrrr) [shoutrrr documentation](https://containrrr.dev/shoutrrr)

## Installation & Configuration
Create a services.yaml with the following content (replace for your use):
```yaml
service_name_1: discord://token@id
service_name_2: smtp://username:password@host:port/?fromAddress=fromAddress&toAddresses=recipient1
service_name_3: gotify://gotify-host/token
```

```
docker-compose up -d
```

If you want to use it in your docker-compose file (for other services):
 services.yaml volume to `/services.yaml` and the port of the http server is 80.

```
docker pull ghcr.io/adridevelopsthings/shoutrrrd:main
```

## How to use?

Make a **POST** request to `/send` with the following *form url encoded* parameters:
- service (the name of the service: the key in the yaml file, the service url will be the value)
- message (the message that will be send to the service url)
