# Berglas AWS
This is a sample project **for learning** mutating webhook implementation.

Berglas AWS is a tool inspired by [GCP's Berglas](https://github.com/GoogleCloudPlatform/berglas) and command line tool and library for storing and retrieving secrets from AWS Secrets Manager. 


## Settings
```$sh
$ export SOME_ENV_VAR=berglas-aws://<AWS Secrets Manager ARN>
```

## CLI usage
```sh
$ berglas-aws exec <some command>
```
