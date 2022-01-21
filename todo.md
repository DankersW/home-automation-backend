# Todo

Fetch temperature from SMHI website: http://opendata.smhi.se/apidocs/metfcst/examples.html

## Needed

- logging
- package everything up
- neater endpoints

## Features

- CI: go reference, go reference, build, code coverage
- Unit tests
- Pass info from the HTTP request into the endpoints handlers
- Documentation
- [Linting action](https://github.com/wearerequired/lint-action)
- Orginize main src folder with [subfolders](https://stackoverflow.com/questions/23154898/break-up-go-project-into-subfolders)
- Get the average temp per device
- Re-orginize the /api call
- combine all the boiler plate code to create a new api handler
- reduce docker container size

## Bugs

- Fix connect to mongo bug when no mongo is present
