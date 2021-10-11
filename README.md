# Minicel - Minimal Excel Engine

**A simple [excel](https://en.wikipedia.org/wiki/Microsoft_Excel) engine without ui to parse [.csv](https://en.wikipedia.org/wiki/Comma-separated_values) files.**\
**Written in [Go](https://golang.org/)**

### Input

```csv
A      , B      , C
1      , 2      , 3
4      , 5      , 6
=A1+C2 , =A2+B2 , =C1-B1
```

### Output

```csv
====================
A    | B    | C    |
1    | 2    | 3    |
4    | 5    | 6    |
7    | 9    | 1    |
====================
```

### Usage

-   **Build**

```console
$ make build
```

-   **Run**

```console
$ ./minicel input.csv
```

> Replace `input.csv` with your input file name
