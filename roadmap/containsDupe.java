package roadmap;

import java.util.HashMap;
import java.util.Map;

class constainsDupe {
    public static void main(String[] args) {

    }
}

class Solution {
    /*
     * Given an integer array nums, return true
     * if any value appears at least twice in the
     * array, and return false if every element is distinct.
     * 
     */
    public boolean containsDuplicate(int[] nums) {
        Map<Integer, Integer> counter = new HashMap<>();
        for (int i : nums) {
            int v = counter.getOrDefault(i, 0);
            if (v == 1) {
                return true;
            }
            counter.put(i, v + 1);
        }

        return false;
    }
}