<?php
class Solution {

    /**
     * @param Integer[] $nums
     * @return Boolean
     */
    function canJump($nums){
        $k =0;
        $count = count($nums);
        for ( $i = 0; $i < $count; $i ++) {
            if ($i > $k) {
                return false;
            }
            $k = max($k, $i+$nums[$i]);
        }
        return true;
    }


}

$arr = [3,2,1,0,4];
$a = new Solution();
var_dump($a->canJump($arr));

/**
 * 解题思路 ：
 *   用$k 记录当前元素能跳到的最远距离
 *   对每一个可以起跳的元素进行记录最远距离 如果一直可以跳到最后 则为true  若当前下标超过了能跳的最远距离则false
 *   时间复杂度O(n)
 */