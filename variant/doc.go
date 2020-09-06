/*
Package variant implements Variant data type.

Variant (also known as [tagged union](https://en.wikipedia.org/wiki/Tagged_union))
allows to store values of one of the following types:

- int,
- float64,
- string,
- []byte slice,
- ordered list of Variant,
- ordered key/value list of Variant, where key is a string.
- empty or no value.

Variant implementation is optimized for performance: for minimal CPU and
memory usage. The implementation currently targets amd64 or 386 GOARCH
only (it can be extended to other architectures).
*/
package variant