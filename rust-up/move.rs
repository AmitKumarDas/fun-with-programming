fn main() {
    try_move();
    try_move_2();
}

fn try_move() {
    let name = "Amit keeps moving".to_string(); // name owns the value
    let _a = name; // a is the new owner; i.e. value is moved
    // let _b = name; // [ERROR] value used after move
}

fn try_move_2() {
    let name = "Amit keeps moving".to_string(); // name owns the value
    let a = &name; // a borrows from name; a is pointer to same data
    let _b = &name; // _b borrows from name; _b is pointer to same data
    borrow(a); // ok since this is still a borrow
    let _c = a; // borrow again; name never loses ownership
}

fn borrow(whom: &String) {
    println!("Hello {}!", whom);
}
