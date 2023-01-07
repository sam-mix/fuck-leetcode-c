#include <iostream>
#include <vector>

class Solution {
 public:
  int minOperations(vector<int> &nums, int x) {
    int target = accumulate(nums.begin(), nums.end(), 0) - x;
    if (target < 0) return -1;  // 全部移除也无法满足要求
    int ans = -1, left = 0, sum = 0, n = nums.size();
    for (int right = 0; right < n; ++right) {
      sum += nums[right];
      while (sum > target) sum -= nums[left++];  // 缩小子数组长度
      if (sum == target) ans = max(ans, right - left + 1);
    }
    return ans < 0 ? -1 : n - ans;
  }
};

int main(int argc, char const *argv[]) {
  using namespace std;
  Solution s = new Solution();
  int result =
      s.minOperations({1, 2, 3, 43, 4, 5, 6, 6, 7, 234, 1, 23, 432}, 50);
  cout << result << endl;
  return 0;
}
