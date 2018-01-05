# cuckoofilter
cuckoo filter implement in golang

**Cuckoo Hashing** is a technique for resolving collisions in hash tables that produces a dic- tionary with constant-time worst-case lookup and deletion operations as well as amortized constant-time insertion operations. First introduced by Pagh in 2001 as an extension of a previous static dictionary data structure, Cuckoo Hashing was the first such hash table with practically small constant factors. Here, we first outline existing hash table collision policies and go on to analyze the Cuckoo Hashing scheme in detail. Next, we give an overview of (c, k)-universal hash families, and finally, we summarize some selected experimental results comparing the varied collision policies.

![Golang version](https://img.shields.io/badge/golang-1.9.2-green.svg) ![passed](https://img.shields.io/badge/test-4%20passed%2C%200%20skipped-green.svg)[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
![PR](https://img.shields.io/badge/PRs-welcome-brightgreen.svg)


## Reference
- [Fastforwardlabs](https://github.com/fastforwardlabs/cuckoofilter)
- [https://coolshell.cn/articles/17225.html](https://coolshell.cn/articles/17225.html)
- [Cuckoo Filter: Practically Better Than Bloom](http://cs.stanford.edu/%7Erishig/courses/ref/l13a.pdf)

- [CS 166 Stanford lecture Cuckoo Hashing](http://web.stanford.edu/class/cs166/lectures/13/Small13.pdf)

According the CMU papers, they could archieve 96% above efficient storage usages. Here, we use simple examples to reach 98%. I guess that murmur3 brings us with great distributed.

Indeed, as the coolshell articles discussed, Bloom filter is faster and higher usage of storage than Cuckoo filter.
Bloom filter is bitset operation with O(1); on the contrary, Cuckoo Filter takes at least O(bucket_size) for bucket entries.

Cuckoo filter still gives us pretty useful hash-cache functions and **DETLETE** funcion instead.

the following shows the experiment with 98% usage (392/400)
```
go test -timeout 30s -run ^TestRehash$

fill 392
--- FAIL: TestRehash (0.01s)
Cuckoo Filter has filled up
```


## Usage


**Create CuckooFilter**

- NewCuckooFilter(capacity int, bucket_size int)

**CuckooFilter Operations**

- CuckooFilter.Insert(item string) (int,error)

- CuckooFilter.Remove(item string) bool

- CuckooFilter.Contains(item string) bool



```go
package main

import (
  "fmt"
  "github.com/yc0/cuckoofilter"
)

func main() {

  filter := NewCuckooFilter(1000, 4)
  item := "string"
  idx, err := filter.Insert(item)
  fmt.Println(idx,err)
  filter.Remove(item)
  filter.Contains(item)
}
```