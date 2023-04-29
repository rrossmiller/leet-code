package roadmap;

import java.util.HashMap;
import java.util.Map;
import java.util.Stack;

class isValidParens {
    public static void main(String[] args) {
        System.out.println(isValid(""));
    }

    public static boolean isValid(String s) {
        Stack<Character> stack = new Stack<>();
        Map<Character, Character> pairs = new HashMap<>() {
            {
                put('(', ')');
                put('[', ']');
                put('{', '}');
            }
        };

        for (Character c : s.toCharArray()) {
            if (pairs.containsKey(c)) {
                stack.push(c);
            } else if (stack.peek() != c || stack.size() == 0) {
                return false;
            } else {
                stack.pop();
            }

        }
        return false;
    }
}
