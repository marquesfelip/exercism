pub fn collatz(n: u64) -> Option<u64> {
    // todo!("return Some(x) where x is the number of steps required to reach 1 starting with {n}")

    let mut result = n;
    let mut steps = 0;

    while result > 1 {
        // par
        if result % 2 == 0 {
            result /= 2;
            steps += 1;
        }
        // impar
        else if result % 2 == 1 {
            match result.checked_mul(3) {
                Some(checked_result_mul) => {
                    match checked_result_mul.checked_add(1) {
                        Some(checked_result_add) => {
                            result = checked_result_add;
                        } None => {
                            return None
                        }
                    }
                    steps += 1;
                }
                None => {
                    return None
                }
            }
        }
    }

    if result < 1 {
        return None
    }
    Some(steps)
}