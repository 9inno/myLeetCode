<?php
class Solution {

    /**
     * @param Integer[] $nums
     * @return Integer
     */
    function singleNumber($nums) {

        $tmp = [];
        $count = count($nums);
        while ($count > 0) {
            $num = array_shift($nums);
            if (!in_array($num,$nums) && !in_array($num, $tmp)) {
                return $num;
            }else {
                array_push($tmp,$num);
            }
            $count --;
        }
        return 0;
    }

    function singleNumber2($nums) {

        $tmp = [];
        foreach ($nums as $num) {
           $tmp[$num]++;
        }
        $result = array_flip($tmp);
        return $result[1] ?? 0;
    }
}


$a = new Solution();
echo $a->singleNumber2($arr = [2,2,2,1]);


/**
 * 解题思路：
 *    1 利用额外空数组 存储遍历过的元素 若元素在原数组和新数组都不存在的情况下即为出现一次
 *    2 遍历数组计算 出现次数
 *
 */