pub fn factors(mut n: u64) -> Vec<u64> {
    let mut vec_prime_factors: Vec<u64> = Vec::new();
    let mut divisor = 2;

    while n > 1 {
        if n % divisor == 0 {
            vec_prime_factors.push(divisor);
            n /= divisor

        } else {
            divisor += 1;
        }
    }

    vec_prime_factors
}
