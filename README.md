# Berglas AWS
This is a sample project **for learning** mutating webhook implementation, and implemented referencing [Berglas](https://github.com/GoogleCloudPlatform/berglas)

Berglas AWS is a tool inspired by [GCP's Berglas](https://github.com/GoogleCloudPlatform/berglas) and command line tool and library for storing and retrieving secrets from AWS Secrets Manager. 

This repository is a CLI part of Berglas AWS. The mutation webhook implementation is [here](https://github.com/katainaka0503/berglas-aws-webhook)

## Settings
```$sh
$ export SOME_ENV_VAR=berglas-aws://<AWS Secrets Manager ARN>
```

## CLI usage
```sh
$ berglas-aws exec <some command>
```
