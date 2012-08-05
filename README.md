go-password-generator
=====================

This program can be used to generate passwords. By default it generates ten passwords with a length of 20 characters each, drawing from lower-case and upper-case characters, as well as numbers.

You can choose what character sets to use. You can also provide a custom set of characters. The length and number of passwords can be set, too.

Building from source
--------------------

```
$ git clone https://github.com/ChristianSiegert/go-password-generator.git
$ cd ./go-password-generator
$ go build
```

You should now see an executable named `go-password-generator`.

Command-line arguments
----------------------

You can modify the output with these flags:

- **--count**	Number of passwords to generate. Default is `10`.
- **--custom**	Custom set of characters to include.
- **--length**	Length of passwords. Default is `20`.
- **--lower**	Include lower-case characters `a-z`. Default is `true`.
- **--numbers**	Include numbers `0-9`. Default is `true`.
- **--special**	Include special characters `!§$%&=?,.-;:_`. Default is `false`.
- **--upper**	Include upper-case characters `A-Z`. Default is `true`.
- **--verbose**	Display additional information.

Usage examples
--------------

Generate passwords with default settings:

	$ ./go-password-generator

Generate passwords that only contain lower-case characters:

	$ ./go-password-generator --upper=false --numbers=false

Generate 5 passwords with a length of 16 that can contain lower-case characters, upper-case characters, numbers and characters from the custom set `ñäöüéèê£€@`:

	$ ./go-password-generator --count 5 --length 16 --custom "ñäöüéèê£€@"
