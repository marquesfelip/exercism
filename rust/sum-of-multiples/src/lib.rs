pub fn sum_of_multiples(limit: u32, factors: &[u32]) -> u32 {

    let mut vec_mutiples: Vec<u32> = Vec::new();

    for i in factors {
        let mut multiple: u32 = *i;
        
        if multiple != 0 {
            while multiple < limit {
                vec_mutiples.push(multiple);
    
                multiple += *i;
            }
        } else {
            vec_mutiples.push(0_u32);
        }
    }

    vec_mutiples.sort();
    vec_mutiples.dedup();

    vec_mutiples.iter().sum()

}
