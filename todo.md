# Todo

## Needed

- Create API call to fetch the device states, returing a list of all device names and there status
- Create API call to return a summary of all the devices states, 4 online, 7 offline
- API call to return the current temperature and humidity inside (averging all sensors) and outside
- API call for the daily/weekly temperature and cpu usage of the server.

## Features

- CI: go reference, go reference, build, code coverage
- Unit tests
- More generic mongo driver
- Mongo filter generator that is reusable
- Pass info from the HTTP request into the endpoints handlers
- Documentation
- [Linting action](https://github.com/wearerequired/lint-action)

## Bugs

- Fix connect to mongo bug when no mongo is present
