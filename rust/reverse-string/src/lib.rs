use unicode_segmentation::UnicodeSegmentation;

pub fn reverse(input: &str) -> String {
    
    if !is_grapheme_cluster(input) {
        let string_rev: String = input.chars().rev().collect();
        string_rev
    } else {
        let string_rev: String = UnicodeSegmentation::graphemes(input, true)
            .rev()
            .collect();
        string_rev
    }

}

fn is_grapheme_cluster(s: &str) -> bool {
    // Use o método graphemes para obter um iterador de grafemas
    let graphemes = UnicodeSegmentation::graphemes(s, true);

    // Verifica se há pelo menos um grafema no iterador
    graphemes.count() > 0
}