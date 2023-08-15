import java.util.HashMap;
import java.util.Map;
import java.util.Stack;

class isValidParens {
    public static void main(String[] args) {
        System.out.println(isValid("([])"));
    }

    public static boolean isValid(String s) {
        if (s.length() == 0 || s.length() % 2 != 0) {
            return false;
        }

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
            } else if (stack.size() == 0 || pairs.get(stack.peek()) != c) {
                return false;
            } else {
                stack.pop();
            }

        }
        return stack.size() == 0;
    }
}
