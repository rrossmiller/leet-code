/*
Given two strings s and t, return true if t is an anagram of s, and false otherwise.
 */
function isAnagram(s: string, t: string): boolean {
    if (s.length !== t.length) {
        return false;
    }
    const sSort = s.split('').sort();
    const tSort = t.split('').sort();

    for (let i = 0; i < sSort.length; i++) {
        if (sSort[i] !== tSort[i]) {
            return false;
        }
    }

    return true;
}

isAnagram('hello', 'world');
