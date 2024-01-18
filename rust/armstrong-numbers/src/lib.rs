pub fn is_armstrong_number(num: u32) -> bool {
    let mut base = num;
    let power = num.to_string().len() as u32;
    let mut soma: u128 = 0;

    while base > 0 {
        let digito: u32 = base % 10;
        
        soma += digito.pow(power) as u128;

        base /= 10;
    }

    if soma == (num as u128) {
        true
    } else {
        false
    }
}
