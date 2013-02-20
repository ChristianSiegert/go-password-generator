#go-password-generator

This program can be used to generate passwords. By default it generates ten passwords with a length of 20 characters each, drawing from lower-case and upper-case characters, as well as numbers.

You can choose what character sets to use. You can also provide a custom set of characters. The length and number of passwords can be changed, too.

## Building from source

Check out the project with:

	$ go get github.com/ChristianSiegert/go-password-generator

You should now see the project files at `$GOPATH/src/github.com/ChristianSiegert/go-password-generator`. It was also built and then installed at `$GOPATH/bin`. You can execute it now with:

	$ go-password-generator

## Command-line flags

You can modify the output with these flags:

- **--count**	Number of passwords to generate. Default is `10`.
- **--custom**	Custom set of characters to include.
- **--length**	Length of passwords. Default is `20`.
- **--lower**	Include lower-case characters `a-z`. Default is `false`.
- **--numbers**	Include numbers `0-9`. Default is `false`.
- **--special**	Include special characters `!§$%&=?,.-;:_`. Default is `false`.
- **--upper**	Include upper-case characters `A-Z`. Default is `false`.
- **--verbose**	Display additional information.

For convenience, if no character set is specified, the password generator automatically includes lower-case and upper-case characters, as well as numbers.

## Usage examples

Generate passwords with default settings:

	$ go-password-generator

Generate passwords that only contain lower-case characters:

	$ go-password-generator --lower

Generate 5 passwords with a length of 16 characters each and include the number set and custom set `ñäöüéèê£€@`:

	$ go-password-generator --count 5 --length 16 --numbers --custom "ñäöüéèê£€@"
