function groupAnagrams(strs: string[]): string[][] {
    let ans: string[][] = [[strs[0]]];

    // for every word
    for (let i = 1; i < strs.length; i++) {
        const s = strs[i];

        // for every array in ans
        let j = 0;
        let foundAnagram = false;
        for (; j < ans.length; j++) {
            // for every word in the current array
            for (let k = 0; k < ans[j].length; k++) {
                const e = ans[j][k];

                if (isAnagram(s, e)) {
                    foundAnagram = true;
                    ans[j].push(s);
                    break;
                }
            }
            if (foundAnagram) {
                break;
            }
        }

        if (!foundAnagram) {
            ans.push([s]);
        }
    }
    return ans;
}

// function isAnagram(s: string, t: string): boolean {
//     if (s.length !== t.length) {
//         return false;
//     }
//     const sSort = s.split('').sort();
//     const tSort = t.split('').sort();

//     for (let i = 0; i < sSort.length; i++) {
//         if (sSort[i] !== tSort[i]) {
//             return false;
//         }
//     }

//     return true;
// }

// const strs = ['eat', 'tea', 'tan', 'ate', 'nat', 'bat'];
// console.log(groupAnagrams(strs));
