# xbar plugin for the [Bitrise](https://www.bitrise.io) build status

## Functions
- Display the number of running builds
- Display the list of builds
  - Running
  - Finished in a day with the status
- Display the last updated time

## Settings

| key | description |
| -- | -- |
| TOKEN | Set your Bitrise personal access token |
| APP | Set list of your Bitrise application name and slug <br> eg: `{"iOS app": "abcdef", "Android app": "012345"}`  |

## How to use
- `make build`
  - Create an executable binary in the `bin` directory
- `make clean`
  - Remove the `bin` directory