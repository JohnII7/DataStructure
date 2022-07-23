#include <stdio.h>
int search(int *nums,int numsSize,int target){
    for (int i = 0; i < numsSize; ++i) {
        if(nums[i]==target) return i;
    }
    return -1;
}

int main() {
    int arr[] = {1,3,4,6,7,8,10,11,13,15},target = 3;
    printf("%d",search(arr,10,target));
}
