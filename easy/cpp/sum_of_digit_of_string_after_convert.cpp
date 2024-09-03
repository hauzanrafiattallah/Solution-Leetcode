// link to the problem: https://leetcode.com/problems/sum-of-digits-of-string-after-convert/

#include <string>  
#include <sstream> 
#include <cstdlib>

class Solution {
public:
    int getLucky(string s, int k) {
        // convert each character in the string to its corresponding numeric value
        string number = "";
        for (char x : s) {
            number += to_string(x - 'a' + 1);
        }

        // perform the transformation 'k' times
        while (k > 0) {
            int temp = 0;
            for (char x : number ){
                temp += x - '0'; // sum the digits of the current number
            }
            number = to_string(temp); // convert the sum back to a string
            k--;
        }
        return stoi(number); // return the final result an integer
    }
};