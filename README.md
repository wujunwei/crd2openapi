# crd2openapi
A tool for converting CRD (Custom Resource Definition) to open api json

## support flags

| name        | short | type   | default            | usage                                                      |
|-------------|-------|--------|--------------------|------------------------------------------------------------|
| file        | f     | string | -  (stdout)        | filename or path to the CRD to be converted.               |
| output      | o     | string | - (stdout)         | out openapi json file.                                     |
| pretty      | p     | bool   | true               | print the json pretty.                                     |
| title       | t     | string | empty string       | the tile of the swagger json.                              |
| indent      | I     | int    | 4                  | the indent of json line , only enable when pretty is true. |
| version     | v     | string | 1.0.0              | the version of the swagger json.                           |
| description | d     | string | kubernetes crd doc | the description of the swagger json.                       |

# todo

- support openapi style json.
- support Built-in swagger server.
- support post-render.