pub fn square(s: u32) -> u64 {

    match s {
        1..=64 => {},
        _ => panic!("Square must be between 1 and 64")
    }

    let mut soma = 1;
    let mut count = 1;

    while count < s {

        soma = soma * 2;

        count += 1;
    }

    soma
}

pub fn total() -> u64 {

    let mut soma = 0;

    for i in 1..=64 {
        soma += square(i)
    }

    soma
}

