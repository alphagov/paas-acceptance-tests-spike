# PaaS Acceptance Tests

This repo contains the results of a spike into writing additional CloudFoundry
acceptance tests for our customisations of the deployment.

## Running the tests

You must set an environment variable `$CONFIG` which points to a JSON file that
contains several pieces of data that will be used to configure the acceptance
tests, e.g. telling the tests how to target your running Cloud Foundry
deployment.

Example:
```json
{
  "api": "api.foo.cf.example.com",
  "admin_user": "admin",
  "admin_password": "secret",
  "apps_domain": "foo.cf.example.com",
  "skip_ssl_validation": true,
  "use_http": false
}
```

Run the tests using the `run_tests.sh` script.
