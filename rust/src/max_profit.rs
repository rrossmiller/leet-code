
pub fn max_profit(prices: Vec<i32>) -> i32 {
    let mut rtn = 0;
    let mut min = prices[0];
    for i in prices.into_iter() {
        if i < min {
            min = i;
        } else if i - min > rtn {
            rtn = i - min
        }
    }
    return rtn;
}
