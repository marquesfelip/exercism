pub fn private_key(p: u64) -> u64 {
    p - 1
}

pub fn public_key(p: u64, g: u64, a: u64) -> u64 {

    modular_pow(g, a, p)
}

pub fn secret(p: u64, b_pub: u64, a: u64) -> u64 {
    // todo!("Calculate secret key using prime number {p}, public key {b_pub}, and private key {a}")

    modular_pow(b_pub, a, p)
}

fn modular_pow(base: u64, exp: u64, modulus: u64) -> u64 {
    if modulus == 1 {
        return 0
    }

    let modulus_calc: u128 = modulus as u128;
    let mut result: u128 = 1;
    let mut base: u128 = (base as u128) % modulus_calc;
    let mut exponent = exp;

    while exponent > 0 {
        if exponent % 2 == 1 {
            result = (result * base) % modulus_calc
        }
        exponent >>= 1;
        base = (base * base) % modulus_calc
    }

    result as u64
}