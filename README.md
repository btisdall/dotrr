# dotrr

Populate secrets in a dotenv file from secret providers such as AWS SSM Parameter Store.

```
dottr resolve file.tmpl > dev.env
```

## Precompiled binaries

Precompiled binaries are available from the releases page.

### macOS

_macOS binaries are not notarized, downloading and extracting
with command line tools should avoid any problems._

If the binary has been downloaded via browser you can first run the binary by
the Finder context menu, or run the following from the command line:

```shell
xattr -d com.apple.quarantine <FILE>
```

Alternatively, see [Building the Project](#building-the-project) below.

## Usage example

1. Given a dotenv file name `local.env/tmpl` with these contents:

```
VAR1=aws-ssm-parameter:db-password
VAR2=aws-ssm-parameter:/another/password
VAR3=just_some_plain_text
VAR4=\\aws-ssm-parameter:some more text
```

2. And you have valid AWS credentials in your AWS credential chain.

3. And following command is run:

```
$ dotrr resolve local.env/tmpl
```

4. The following will be output to stdout:

```
VAR1="The_secret_stored_in_SSM_parameter_db-password"
VAR2="The_secret_stored_in_SSM_parameter_/another/password"
VAR2="just_some_plain_text"
VAR4="aws-ssm-parameter:some_more_text"

```

- `VAR1` and `VAR2` are resolved from SSM Parameter store parameters since they
  are prefixed with `aws-ssm-parameter:` in the template.
- `VAR3` is left untouched as it has no prefix.
- `VAR4` is resolved to the literal text supplied minus the double-slash escape.

## Supported Secret Providers

Currently only SSM Parameter Store is supported, but adding other providers
should be straightforward.

## Installation

This project uses go modules, so `go run` or `go build` will take care of resolving any dependencies.

## Building the project.

Please inspect the [Makefile](./Makefile) to see the various build options. As an example, to build for macOS via
Docker:

```shell
make build_darwin
```

## license

Distributed under the MIT license. See `LICENSE` for more information.

## Contributing

1. Fork it (<https://github.com/yourname/yourproject/fork>)
2. Create your feature branch (`git checkout -b feature/fooBar`)
3. Commit your changes (`git commit -am 'Add some fooBar'`). Please make sure
   your changes are covered by tests where possible.
4. Push to the branch (`git push origin feature/fooBar`)
5. Create a new Pull Request

## Contact

Ben Tisdall – [@btisdall](https://twitter.com/btisdall) – ben@tisdall.org.uk

Project Link: [https://github.com/btisdall/dotrr](https://github.com/btisdall/dotrr)
