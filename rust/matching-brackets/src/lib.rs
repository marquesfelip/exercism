pub fn brackets_are_balanced(string: &str) -> bool {
    
    validate_expression(string)
    
}

fn validate_expression(s: &str) -> bool {
    let mut stack = Vec::new();

    for c in s.chars() {
        match c {
            '[' | '{' | '(' => stack.push(c),
            ']' => {
                if stack.pop() != Some('[') {
                    return false;
                }
            },
            '}' => {
                if stack.pop() != Some('{') {
                    return false;
                }
            }
            ')' => {
                if stack.pop() != Some('(') {
                    return false;
                }
            }
            _ => ()
        }
    }

    stack.is_empty()
}