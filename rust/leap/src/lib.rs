pub fn is_leap_year(year: u64) -> bool {

    let mut is_leap = false;

    if year % 4 == 0 {
        is_leap = true;

        if year % 100 == 0 {
            is_leap = false;

            if year % 400 == 0 {
                is_leap = true;
            }
        }
    }

    is_leap
}
