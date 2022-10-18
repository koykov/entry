# Entry

Collection of compact entries based on bit shift operations.

Currently, supports 3 variants:
* Entry12 (compacts uint16/uint8 lo/hi values)
* Entry16 (compacts uint8 lo/hi values)
* Entry24 (compacts uint32/uint8 lo/hi values)
* Entry32 (compacts uint16 lo/hi values)
* Entry48 (compacts uint64/uint16 lo/hi values)
* Entry64 (compacts uint32 lo/hi values)

### Usage

```go
buf := "Lorem ipsum dolor sit amet, consectetur adipiscing elit."
offset, length := uint16(6), uint16(11)
var e entry.Entry32
e.Encode(offset, length)
...
lo, hi := e.Decode()
println(buf[lo:hi]) // ipsum
```
