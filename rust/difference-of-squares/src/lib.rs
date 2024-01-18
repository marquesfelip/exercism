pub fn square_of_sum(n: u32) -> u32 {
    
    let mut soma: u32 = 0;
    
    for i in 1..=n {
        soma += i;
    }

    soma * soma
}

pub fn sum_of_squares(n: u32) -> u32 {
    
    let mut soma: u32 = 0;

    for i in 1..=n {
        soma = soma + (i * i)
    }

    soma
}

pub fn difference(n: u32) -> u32 {
    let square_of_sum = square_of_sum(n);
    let sum_of_squares = sum_of_squares(n);

    square_of_sum - sum_of_squares
}
