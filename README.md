# dotrr

Populate secrets in a dotenv file from secret providers. Currently supported providers are:

* AWS SSM Parameter Store
* AWS Secrets Manager

## Installing dotrr

Precompiled binaries are available from the [releases page](https://github.com/btisdall/dotrr/releases/latest). Alternatively, see [Building the Project](#building-the-project) below.

## Usage

1. Given a dotenv file name `local.env.tmpl` with these contents:

```
VAR1=aws-ssm-parameter:db-password
VAR2=aws-secretsmanager-secret:api-key
VAR3=not_a_secret
```

2. And you have valid AWS credentials in your AWS credential chain.

3. And following command is run:

```
$ dotrr resolve local.env/tmpl
```

4. The following will be output to stdout:

```
VAR1="The secret from the SSM Parameter named db-password"
VAR2="The secret from Secrets Manager secret named api-key"
VAR2="not_a_secret"
```

- `VAR1` is resolved from SSM Parameter store since it is prefixed with `aws-ssm-parameter:` in the template.
- `VAR2` is resolved from Secrets Manager since it is prefixed with `aws-secretsmanager-secret:` in the template.
- `VAR3` is left untouched as it has no known resolver prefix.

Note that items that collide with a resolver prefix can be double backslash
escaped to produce their intended form in the output, eg:

```
VAR4=\\aws-secretsmanager-secret:some-text
```

Results in:

```
VAR4="aws-secretsmanager-secret:some-text"
```

## Building the project.

This project uses go modules, so `go run` or `go build` will take care of resolving any dependencies.

Please inspect the [Makefile](./Makefile) to see the various build options. As an example, to build for macOS via
Docker:

```shell
make build_darwin
```

## License

Distributed under the MIT license. See `LICENSE` for more information.

## Contributing

1. Fork it (<https://github.com/yourname/yourproject/fork>)
2. Create your feature branch (`git checkout -b feature/fooBar`)
3. Commit your changes (`git commit -am 'Add some fooBar'`). Please make sure your changes are covered by tests where
   possible.
4. Push to the branch (`git push origin feature/fooBar`)
5. Create a new Pull Request

## Contact

Ben Tisdall – [@btisdall](https://twitter.com/btisdall) – ben@tisdall.org.uk

Project Link: [https://github.com/btisdall/dotrr](https://github.com/btisdall/dotrr)
