pub fn nth(n: u32) -> u32 {

    let mut count = 0;
    let mut index = 0;
    let mut n_prim = 0;
    let mut is_prim = false;

    loop {
        (n_prim, is_prim) = is_prime(count);
        if is_prim {
            if index == n {
                break;
            } else {
                index += 1;
            }
        } else {
        }
        count += 1;
    }

    n_prim
}

pub fn is_prime(x: u32) -> (u32, bool) {

    if x <= 1 {
        return (0, false);
    }

    for i in 2..=(x / 2) {
        if x % i == 0 {
            return (0, false);
        }
    }

    (x, true)
}