pub fn add(left: usize, right: usize) -> usize {
    left + right
}

pub fn sqrt(number: f64) -> Result<f64, String> {
    if number > 0.0 {
        Ok(number.powf(0.5)) // Why no ;
    } else {
        Err("negative floats don't have square roots".to_owned()) // Where is ;
    }
}

#[cfg(test)]
mod tests {
    use super::*; // importing names from outer scope

    #[test]
    fn it_works() {
        assert_eq!(add(2, 2), 4);
    }

    #[test]
    fn test_sqrt() -> Result<(), String> { // enables use of ? within assert_eq
        let x = 4.0;
         // assert_eq panics if expression evaluates to false
         // So use ? to register custom error messages
        assert_eq!(sqrt(x)?.powf(2.0), x);
        Ok(()) // Where is ;
    }
}
