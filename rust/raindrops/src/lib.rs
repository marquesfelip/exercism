pub fn raindrops(n: u32) -> String {
    let mut result = String::new();

    let mut if_is_factor_push = |factor: u32, sound: &str| {
        if n % factor == 0 {
            result.push_str(sound)
        }
    };

    if_is_factor_push(3, "Pling");
    if_is_factor_push(5, "Plang");
    if_is_factor_push(7, "Plong");

    if result.is_empty() { result = n.to_string(); }

    result
}