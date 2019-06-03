# aws-kms-encrypter

[![GitHub release](http://img.shields.io/github/release/darkraiden/aws-kms-encrypter.svg?style=flat-square)][release]
[![CircleCI](https://img.shields.io/travis/tcnksm/ghr.svg?style=flat-square)](https://circleci.com/gh/darkraiden/workflows/aws-kms-encrypter)
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)][license]

[release]: https://github.com/darkraiden/aws-kms-encrypter/releases
[license]: https://github.com/darkraiden/aws-kms-encrypter/blob/master/LICENSE

AWS KMS Encrypter is a simple command line tool that can be used to Encrypt your passwords just providing a KMS ID, which of course implies the KMS resource is created beforehand.

## Installation

To install `aws-kms-encrypter`, please use `go get`. We tag versions so feel free to checkout that tag and compile.

```bash
$ go get github.com/darkraiden/aws-kms-encrypter
```

## Usage

This tool will first create a random string, and then use KMS algorithms to encrypt them, returning to the user a Cipher Text Blob and the plain password.

```bash
$ aws-kms-encrypter -kms-id="ThisIsTheIDOfYourKMSKey" -context="KMSEncryptionContext=valueOfTheContext" -length=12 -special-chars=false -digits=false
```

### Accepted Flags

|   Name    | Required |                          Description                          |
| :-------: | :------: | :-----------------------------------------------------------: |
| `-kms-id` |   Yes    | The ID of the KMS resource to be used to encrypt the password | `-context` | No | The KMS encryption context expressed in `key=value` format, as a string |

| `-length` | No | The length of the new password (default value: `12`) |
| `-special-chars` | No | Whether the new password should contain special characters or not (default value: `true`) |
| `-digits` | No | Whether the new password should contain numbers or not (default value: `true`) |

## Contributing

Please read [CONTRIBUTING.md](CONTRIBUTING.md) for details on our code of conduct, and the process for submitting pull requests to us.

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/darkraiden/sigmund/tags).

## Authors

-   [Davide Di Mauro](https://github.com/darkraiden)

See also the list of [contributors](contributors.md) who participated in this project.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details
