pub fn series(digits: &str, len: usize) -> Vec<String> {
    let mut vec_result: Vec<String> = Vec::new();

    if len > digits.len() || digits.len() == 0 {
        return vec_result
    }

    for (index, _digit) in digits.char_indices() {
        vec_result.push(digits[index..len+index].to_string());

        if len + index == digits.len() {
            break;
        }
    }

    vec_result
}
