<?php
class Solution {

    /**
     * @param Integer[] $nums
     * @return Integer
     */
    function rob($nums) {
        return $this->recursion($nums, count($nums) - 1);
    }

    function recursion($nums,$i) {
        if ($i == 0) {
            return $nums[0];
        } elseif ($i == 1) {
            return max($nums[0], $nums[1]);
        }

        return max($this->recursion($nums,$i - 2) + $nums[$i], $this->recursion($nums,$i - 1));
    }


    function db_rob($nums) {
        $count = count($nums);
        if ($count == 0){
            return 0;
        }elseif ($count == 1) {
            return $nums[0];
        }elseif ($count == 2) {
            return max($nums[0], $nums[1]);
        }
        $opt = [];
        $opt[0] = $nums[0];
        $opt[1] = max($nums[0], $nums[1]);

        for ($i = 2; $i < $count; $i++) {
            $opt[$i] = max($opt[$i-2] + $nums[$i], $opt[$i-1]);
        }
        return $opt[$count-1];
    }
}

$arr =[1,2,3,1] ;
$a = new Solution();
echo $a->rob($arr);
echo "\n";
echo $a->db_rob($arr);

/**
 * 解题思路：
 *   动态规划 OPT($i) = 选i节点 OPT($i-2) + $num[$i]
 *                      不选i节点 OPT($i -1)
 *   取max
 *
 *   递归会超时超内存
 */