// References
// - https://www.thecodedmessage.com/posts/strong-typing/

// A Result can fail
// pub fn from_slice(v: &[u8]) -> Result<Value>

// 1. Crash on error
//
// let input = from_slice(&input).expect("Invalid JSON");

// Hey, can you do above in rust?
//
// Can input be reused? Yes
// Reusing the name input like this with a different type is allowed in Rust
// This declares a new variable that shadows the old one
// This is idiomatic when the value is being transformed and we don’t need the old form anymore

// 2. Bubble the error up to the caller
//
// let input = from_slice(&input)?;

// 3. Compiler Error
//
// let input = match from_slice(&input) {
//     Ok(parsed_value) => parsed_value, // This is the parsed value, type `Value`
//     Err(_) => input, // This is the raw `Vec<u8>` data... TYPE MISMATCH!
// }

// 4. Avoid Compiler Error
//
// enum IncomingMessage {
//     Parsed(Value),
//     Unparsed(Vec<u8>),
// }
//
// let input = match input {
//     Parsed(value) => value,
//     Unparsed(_) => {
//         // return an error JSON blob
//     }
// }
