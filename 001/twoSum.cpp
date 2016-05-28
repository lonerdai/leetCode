class Solution {
public:
    vector<int> twoSum(vector<int>& nums, int target) {
        vector<int>res;
        map<int,int> m;
        for(vector<int>::size_type i=0;i<nums.size();i++){
            if(m.find(target - nums[i])!=m.end()){
                res.push_back(m[target - nums[i]]);
                res.push_back(i);
                break;
            }else{
                m[nums[i]] = i;
            }
        }
        return res;
    }
};