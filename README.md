# now

    go get -u github.com/chbrown/now

**Arguments**

| `now [...]`  | Output | Note |
|-------------:|:-------|:-----|
|              | `2016-11-20T18:46:29Z` |
|         `-d` | `2016-11-20` |
|         `-m` | `2016-11-20T18:46Z` |
|         `-s` | `2016-11-20T18:46:29Z` | same as no arguments
|        `-ms` | `2016-11-20T18:46:29.882Z` |
|        `-ns` | `2016-11-20T18:46:29.882112334Z` |
| `-local    ` | `2016-11-20T12:46:29-06:00` |
| `-local  -d` | `2016-11-20` |
| `-local  -m` | `2016-11-20T12:46-06:00` |
| `-local  -s` | `2016-11-20T12:46:29-06:00` | same as `-local` only
| `-local -ms` | `2016-11-20T12:46:29.882-06:00` |
| `-local -ns` | `2016-11-20T12:46:29.882112334-06:00` |
| `-epoch    ` | `1479667590` | same as `date +%s`
| `-epoch  -s` | `1479667589882` | same as `-epoch` only
| `-epoch -ms` | `1479667589882` |
| `-epoch -ns` | `1479667589882112256` |

## License

Copyright (c) 2017 Christopher Brown. [MIT Licensed](https://chbrown.github.io/licenses/MIT/#2017).
