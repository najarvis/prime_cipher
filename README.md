# Prime Cipher
A pretty basic implementation of the Sieve or Eratosthenes method for generating prime numbers.

## License
Copyright © 2018 Nick Jarvis <najarvis2016@gmail.com>

This work is free. You can redistribute it and/or modify it under the terms of the Do What The Fuck You Want To Public License, Version 2, as published by Sam Hocevar. See the COPYING file for more details.

## Usage
`./prime_cipher [max number]`

Will generate all prime numbers up until that number and print them to the standard output. If you just want to generate them, the function GetPrimes will create and return a set of all prime numbers up until the input max number in the form of an int->struct{} map.

### Why?
I found the wikipedia page and thought it was pretty neat. Also WTFPL because I didn't invent the Sieve of Eratosthenes, Eratosthenes did.
