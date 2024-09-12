// link to the problem: https://leetcode.com/problems/count-the-number-of-consistent-strings/description/?envType=daily-question&envId=2024-09-12
#include <iostream>
#include <vector>
#include <bitset>
#include <string>
using namespace std;


class Solution {
public:
    int countConsistentStrings(string allowed, vector<string>& words) {
        bitset<26> ASet = 0;
        for (char c: allowed) 
            ASet[c-'a']=1;
        int cnt = 0;
        for (string& w: words) {
            bool consistent = 1;
            for(char c: w) {
                if(ASet[c-'a']==0) {
                    consistent = 0;
                    break;
                }
            }
            cnt += consistent;
        }
        return cnt;
    }
};

int main() {
    Solution s;
    vector<string> words = {"ad","bd","aaab","baa","badab"};
    string allowed = "ab";
    cout << s.countConsistentStrings(allowed, words) << endl;
    return 0;
}