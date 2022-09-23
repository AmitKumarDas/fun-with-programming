fn main(){
    // &str; a string literal;
    // a string slice that refers to a preallocated text
    // that is stored in read-only memory as part of the executable
    // let my_name = "Amit";

    let my_name = "Amit".to_string(); // std::string::String
    greet(my_name.to_string());
    greet_and_drop();

    let name = "wow"; // a &str i.e. string literal; stored in preallocated read-only memory
    greet_it(name);
}

fn greet(name: String) { // std::string::String
    println!("Hello {}!", name);
}

fn greet_and_drop() {
    let s = "When am I dropped?".to_string();
    println!("{}", s); // s is dropped here
}

fn greet_it(who: &str) {
    println!("Hello {}!", who);
}
