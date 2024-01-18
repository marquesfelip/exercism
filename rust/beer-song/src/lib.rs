use regex::Regex;
//    1          2                             3          4
// {No more} {bottles} of beer on the wall, {no more} {bottles} of beer. 
//   {Go to the store and buy some more}, {99} {bottles} of beer on the wall.
//                  5                       6      7

// {1} {bottle} of beer on the wall, {1} {bottle} of beer. 
//   {Take it down and pass it around}, {no more} {bottles} of beer on the wall.

// {2} {bottles} of beer on the wall, {2} {bottles} of beer. 
//   {Take one down and pass it around}, {1} {bottle} of beer on the wall.

pub fn verse(n: u32) -> String {
    let mut word_one: String = String::new();
    let mut word_two: String = String::new();
    let mut word_three: String = String::new();
    let mut word_four: String = String::new();
    let mut word_five: String = String::new();
    let mut word_six: String = String::new();
    let mut word_seven: String = String::new();

    if n == 0 {
        word_one = String::from("no more");
        word_two = String::from("bottles");
        word_three = String::from("no more");
        word_four = String::from("bottles");
        word_five = String::from("go to the store and buy some more");
        word_six = String::from("99");
        word_seven = String::from("bottles");
    } else if n == 1 {
        word_one = String::from(n.to_string());
        word_two = String::from("bottle");
        word_three = String::from(n.to_string());
        word_four = String::from("bottle");
        word_five = String::from("take it down and pass it around");
        word_six = String::from("no more");
        word_seven = String::from("bottles");
    } else {
        word_one = String::from(n.to_string());
        word_two = String::from("bottles");
        word_three = String::from(n.to_string());
        word_four = String::from("bottles");
        word_five = String::from("take one down and pass it around");
        word_six = String::from((n - 1).to_string());
        if (n - 1) == 1 {
            word_seven = String::from("bottle");
        } else {
            word_seven = String::from("bottles");
        }
    };

    let text = format!("{word_one} {word_two} of beer on the wall, {word_three} {word_four} of beer.\n{word_five}, {word_six} {word_seven} of beer on the wall.\n");

    capitalize(&text)

}

pub fn sing(start: u32, end: u32) -> String {
    let mut string_sing: String = String::new();
    let mut str_verse = String::new();
    let mut count = start;

    while count >= end {
        str_verse = verse(count);

        string_sing = format!("{string_sing}{str_verse}");
        
        if count != end {
            string_sing = format!("{string_sing}\n");
        }
        
        if count == 0 {
            break;
        }
        count -= 1;
    }

    string_sing
}

pub fn capitalize(text: &str) -> String {
    // let texto = "ola. este é um exemplo. outra frase.";
    
    let re = Regex::new(r"(\.\s+|^)([a-z])").unwrap();
    let result = re.replace_all(text, |caps: &regex::Captures| {
        let letter = caps[2].to_uppercase();
        format!("{}{}", &caps[1], letter)
    });
    
    result.to_string()
}