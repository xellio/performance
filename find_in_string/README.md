## Find string in string

Benchmark for checking if a given needle is in a given haystack

```
Haystack: "1f8b 0800 4466 295a 0003 ed1a 6b6f db46ac5f a55f 71d3 8742 4a1c 59f2 13c8 1a6c4ee2 74dd 563b 8b9d 6140 1a04 b27c 9635cb3a 4f3a 2fcd 86fc f791 f790 6527 6d5a2049 b7d5 57a0 b278 3c92 471e 1f47 25cc82eb 8466 d550 3edd 88bd 78f4 e179 5eabd120 f0f4 db4d affc 94a3 5ef7 895f aff9b55a abe6 b7ea c4f3 1b8d b6f7 8278 8f2fcadd b1cc 7990 8128 59bc c83e 86f7 d0bcdc0a 299e ff91 b108 c259 1051 a2ec 6f9af17c c132 4e6c d3b0 a298 4f97 2337 64f3ea7b 9a24 31d3 a7a4 fa3b 1b59 8090 b0c8321d d3e4 370b 4a8e e41c c979 b60c 39f9db34 7e64 a39c 5c5c ee00 b60b bfcd 5bd39c2c d390 d821 d951 d80e 1980 fab9 ed109a65 2cc3 5540 d43d cde2 9427 a96d 8959295c 9c46 aeeb 5a8e 6954 ab24 0af8 14788d03 1e90 09ac 0316 a691 51be cc52 92c609b0 fad2 8afd 8f0c 6dd1 703e 2eac fbd83cc0 1fda cde6 87fd 1ffd 45f9 7fdd 6bb7d0ff db7e fd05 693e b620 f78d afdc ffcbf67f 2a1e 9f6e ff66 adde f2c0 fef5 5aa3b5b5 ff73 8ca7 f2f9 f2f8 0cfb 379a 3ee6ff5a d36b 6eed ff1c a394 d1ff 15fe efd77ce1 ff4d 6f6b ffe7 1845 fc67 e924 8ed4e391 6f01 0fd5 ff10 ecb5 fddb 8d3a fa7fb3d9 6c6c ebff e718 45fd 2f0c 6f9a 7f0619d6 fec7 dd93 cef9 cfc3 abe3 c3ab 1ffa8321 3920 965f 6bbb 1efc f3ad b5e9 d3fe9998 aed7 bd96 6542 d1df 3fbc 1a0c 3bc3f3c1 55ff b4db 236a 1c10 6f6d ee4d efeaf4ac fffa ac3b 18c0 9cbf 3677 dcef 7557"
Needle: "ee4d"
```
```
go test -bench=. -benchtime=5s

goos: linux
goarch: amd64
pkg: github.com/xellio/performance/find_in_string
BenchmarkStringContains-4               10000000           881 ns/op
BenchmarkBytesContains-4                10000000           909 ns/op
BenchmarkStringContainsWithCast-4        5000000          1536 ns/op
BenchmarkBytesContainsWithCast-4         5000000          1587 ns/op
BenchmarkRegexpMatchString-4              500000         11894 ns/op
BenchmarkRegexpMatch-4                    500000         11780 ns/op
BenchmarkRegexpMatchCompiled-4            500000         11770 ns/op
PASS
ok      github.com/xellio/performance/find_in_string    56.613s
```