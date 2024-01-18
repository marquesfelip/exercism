pub fn series(digits: &str, len: usize) -> Vec<String> {
    if len > digits.len() || digits.len() == 0 {
        return Vec::new()
    }

    digits
        .chars()
        .collect::<Vec<char>>()
        .windows(len)
        .map(|s| s.iter().collect::<String>())
        .collect()
}
