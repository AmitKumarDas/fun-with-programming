fn main() {
    let mut name = "Hello".to_string();
    name.push_str(" World"); // can be mutated since its mutable
    let s = &name[7..]; // a string slice referencing text owned by name; does not store capacity info on stack
    println!("Hello {}!", s)
}