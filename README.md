## Getting Started

After retrieving the repo, compile and run the webapp

At the project root:

```
go build
./be-brucebeitman-interview
```

### Prerequisites

The only dependency is go-git

```
go get -u gopkg.in/src-d/go-git.v4/...
```

### Usage

The webapp runs on port `4242` and consists of the following endpoints

- `/repos` - returns a description of the repos for the given user
  - requires: ?user ([example](http://localhost:4242/repos?user=BruceABeitman))
  - example response:
  ```json
  [
    {
      "name": "ba-power-tutor",
      "description": "Power tutor app with added energy profile for nexus 5 phone, multi-core functionality, multi-threaded tracking, and additional fixes.",
      "url": "https://api.github.com/repos/BruceABeitman/ba-power-tutor",
      "size": 6724
    }
  ]
  ```
- `/branches` - returns a description of the branches for the given user & repo
  - requires: ?user&repo ([example](http://localhost:4242/branches?user=BruceABeitman&repo=svg_test))
  - example response:
  ```json
  [
    {
      "repo": {
        "name": "svg_test"
      },
      "name": "develop",
      "commit": "39c914c1031457e19da295d26b31a0e47c7457a6"
    }
  ]
  ```
- `/svgs` - returns a description of the svg files for the given user, repo, & branch
  - requires: ?user&repo&branch ([example](http://localhost:4242/svgs?user=BruceABeitman&repo=svg_test&branch=develop))
  - example response
  ```json
  {
    "repo": {
      "name": "svg_test"
    },
    "branch": {
      "name": "develop",
      "commit": "39c914c1031457e19da295d26b31a0e47c7457a6"
    },
    "total_files": 71,
    "filenames": [
      "icons/about-dot-me.svg",
      "icons/acm.svg",
      "icons/addthis.svg",
      "icons/adobe.svg"
    ]
  }
  ```
- `/svg/details` - returns a description of an svg file for the given user, repo, branch, & file
  - requires: ?user&repo&branch&file ([example](http://localhost:4242/svg/details?user=BruceABeitman&repo=svg_test&branch=develop&file=icons/americanexpress.svg))
  - example response:
  ```json
  {
    "repo": {
      "name": "svg_test"
    },
    "branch": {
      "name": "develop",
      "commit": "39c914c1031457e19da295d26b31a0e47c7457a6"
    },
    "filename": "icons/americanexpress.svg",
    "file_size": 583,
    "contents": "<svg width=\"400\" height=\"110\"><rect width=\"300\" height=\"100\" style=\"fill:rgb(200%, 200%, 1%);stroke-width:3;stroke:rgb(0,0,0)\" /></svg>\n<svg width=\"400\" height=\"110\"><rect width=\"300\" height=\"100\" style=\"fill:peru;stroke-width:3;stroke:rgb(0,0,0)\" /></svg>\n<svg width=\"400\" height=\"110\"><rect width=\"300\" height=\"100\" style=\"fill:rgb(0,0,255);stroke-width:3;stroke:rgb(0,0,0)\" /></svg>\n<circle fill=\"#CD853F icc-color(acmecmyk, 0.11, 0.48, 0.83, 0.00)\"/>\n<circle fill=\"rgb(205,133,63)\"/>\n<circle fill=\"peru\"/>\n<circle fill=\"rgb(80.392%, 52.157%, 24.706%)\"/>\n<circle fill=\"#CD853F\"/>\n",
    "colors": [
      {
        "raw": "rgb(200%, 200%, 1%)",
        "type": "rgb",
        "parsed": ["200%", "200%", "1%"]
      },
      {
        "raw": "peru",
        "type": "unknown"
      },
      {
        "raw": "rgb(0,0,255)",
        "type": "rgb",
        "parsed": ["0", "0", "255"]
      }
    ]
  }
  ```

## Project Layout

The project is layed out in 3 layers, each of which in it's own package.

- The `main` file/layer bootstraps the webapp, initializing the necessary pieces (e.g. `api` and `core`).
- The `api` layer interacts with the `core` interface. The `api` is responsible for all HTTP handling. It verifies requests are proper, responding with appropriate HTTP status codes, errors, and responses.
- The `core` layer handles business logic and handling external communication.

Additionally there are the `model` (JSON models) and the `utility` packages. I think of the `utility` package almost as a mini-library, that might end up getting broken out into a real library one day. In this case it holds some helper functions for interacting with `go-git`, and all logic for parsing the SVG color information.

## Tests

To run tests

At the project root:

```
go test ./...
```

Tests are not fully fleshed out, but the core logic and examples of each main area are covered (e.g. api, core, and utility). The most complex business logic being in the `utility/find_colors`. Examples of how the `api` and `core` can be tested are shown in `get_branches_test` in each package. Extending them to the other endpoints should be mostly straightforward.
