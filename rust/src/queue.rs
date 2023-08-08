/**
 * Your MyQueue object will be instantiated and called as such:
 * let obj = MyQueue::new();
 * obj.push(x);
 * let ret_2: i32 = obj.pop();
 * let ret_3: i32 = obj.peek();
 * let ret_4: bool = obj.empty();
 */
struct MyQueue {
    vals: Vec<i32>,
}

/**
 * `&self` means the method takes an immutable reference.
 * If you need a mutable reference, change it to `&mut self` instead.
 */
impl MyQueue {
    fn new() -> Self {
        MyQueue { vals: vec![] }
    }

    fn push(&mut self, x: i32) {
        self.vals.push(x)
    }

    fn pop(&mut self) -> i32 {
        let popped = self.vals[0];
        self.vals = self.vals[1..].to_vec();

        popped
    }

    fn peek(&self) -> i32 {
        self.vals[0]
    }

    fn empty(&self) -> bool {
        if self.vals.len() == 0 {
            return true;
        }

        return false;
    }
}
