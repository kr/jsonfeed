/*

Package jsonfeed provides types that encode and decode
JSON Feed files, as specified in "JSON Feed Version 1".
See https://jsonfeed.org/version/1.

This package's interface might change in the future
in a way that requires updates to client code.
If you want to avoid having to update, you should
vendor this package in your application.

The spec mostly maps pretty cleanly
into Go's standard JSON encoding and decoding behavior,
with just a few special rules here and there.
The MarshalJSON and UnmarshalJSON methods handle those rules for you,
plus validation.
Clients of this package
should use the generic Marshal and Unmarshal functions
(or types Encoder and Decoder)
in package json.

*/
package jsonfeed
