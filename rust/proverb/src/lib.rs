pub fn build_proverb(list: &[&str]) -> String { 
    
    let mut vec_pieces: Vec<String> = Vec::new();

    for (index, _) in list.iter().enumerate() {
        if index == list.len() - 1 {
            vec_pieces.push(format!("And all for the want of a {}.", list[0]));
        } else {
            vec_pieces.push(format!("For want of a {} the {} was lost.", list[index], list[index + 1]));
        }
    }

    vec_pieces.join("\n")
}
