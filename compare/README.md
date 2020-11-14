# Directory Comparator

comparing source directory to find a new, modified, and deleted files on target directory.

## Build

you can build this source using Makefile by typing `make build` command or build using Go command: `go build -o compare .`

## Usages

| Args position | Descriptions                            |
|---------------|-----------------------------------------|
| 1             | source directory                        |
| 2             | target directory to compare with the source |

### Example

you can use `source` and `target` in `example/` folder for testing.

```sh
./compare example/source/ example/target/
```

result:

```sh
/dir_200/file_220.txt [DELETED]
/file_100.txt [MODIFIED]
/dir_200/dir_210/file_213.txt [NEW]
```
