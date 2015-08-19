# JunoLab test


## Implementation requrements

- <any> *nix
- docker
- golang
- <any> DB
- API docs


## Features

common:
- registration (with confirmation token)
- registration confirmation by token

anonymous user features:
- show list of registered users
- show any particular user's profile

authenticated user features:
- update my profile info {full name / addr / phone / etc.}
- show updates history


## Security requirements

- external access: static container, APP container
- DB container should be isolated from the outer world
- TLS
- XSS protection (1-2 cases)


## Infrastructure

```
     +-----+      +----+
---> | APP | <--> | DB |
     +-----+      +----+

     +--------+
---> | Static |
     +--------+
```

Static container:
- serves css/js/static elements

App container:
- runs appserver's binary
- goes to DB container for data

DB container:
- runs DB server (choose any)
