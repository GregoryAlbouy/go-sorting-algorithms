# go-sorting-algorithms

I had never realized how fascinating sorting algorithms could be, until I saw [this amazing video](https://www.youtube.com/watch?v=kPRA0W1kECg). It motivated to implement some of them in my favorite language, Go.

I also tried to slightly improve the performance of some of them (see [mergesort.go](./algorithms/mergesort.go) and [quicksort.go](./algorithms/quicksort.go)) using concurrency, since it's so easy and pleasant to work with in Go. It can lead to up to 50% and 25% performance improvment respectively.

Finally, I wrote a CLI that benchmarks them and generates a .csv or .json output

## Algorithms

### Implemented

* Selection Sort
* Insertion Sort
* Bubble Sort
* Merge Sort / Merge Sort Conc (:heart:)
* Quick Sort / Quick Sort Conc

### Coming

* Cocktail Shaker Sort
* Radix Sort

## CLI

The CLI tool can generate some comparative data on demand. It takes three parameters:

| flag | usage | default value |
| ----- | ----- | ----- |
| `-a` | Names of algorithms to be tested(*) | `""` (all) |
| `-s` | Sizes of the arrays to be sorted | `"100 1000 10000"` |
| `-o` | Filename output (.csv or .json) | `"results.csv"` |

Each parameter accepts several values, in that case they must be **quoted** and separated by a space.

(*) Full names list (*not case-sensitive*):
* BubbleSort
* SelectionSort
* InsertionSort
* MergeSort
* MergeSortConc
* QuickSort
* QuickSortConc

### Examples

#### Minimal example
```console
go-sorting-algorithms$ go run .
```

#### Full example
```console
go-sorting-algorithms$ go run . -a "MergeSort MergeSortConc QuickSort QuickSortConc" -s "1000 1000000" -o "excel.csv results.json"
```

## Todo

* Implement more funky algorithms
* Add output `none`/`'cli` (no file generated) as default output value
* Add output `localhost:PORT` to get the json served on local http server