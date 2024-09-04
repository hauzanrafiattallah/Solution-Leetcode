#include <iostream>
#include <string>
#include <cstdlib>

class Solution {
public:
    int getLucky(std::string s, int k) {
        std::string number = "";
        for (char x : s) {
            number += std::to_string(x - 'a' + 1);
        }

        while (k > 0) {
            int temp = 0;
            for (char x : number) {
                temp += x - '0';
            }
            number = std::to_string(temp);
            k--;
        }
        return std::stoi(number);
    }
};

int main() {
    std::cout << "Program dimulai!" << std::endl;

    Solution sol;
    std::string s = "leetcode";  
    int k = 2;  

    std::cout << "Memanggil fungsi getLucky..." << std::endl;

    int result = sol.getLucky(s, k);

    std::cout << "Hasil: " << result << std::endl;

    return 0;
}

