<?php
class Solution {

    /**
     * @param Integer[][] $intervals
     * @return Integer[][]
     */
    function merge($intervals) {
        $count = count($intervals);
        $result = [];
        $column = array_column($intervals,0);
        array_multisort($intervals,SORT_ASC,$column);
        $j = 0;
        $result[$j] = $intervals[$j];
        for ($i = 1; $i < $count; $i++) {
            if ($result[$j][1] >= $intervals[$i][0]) {
                $result[$j] = [min($intervals[$i][0], $result[$j][0]), max($intervals[$i][1],$result[$j][1])];
            }else {
                $result[$j + 1] = $intervals[$i];
                $j++;
            }
        }
        return $result;
    }
}

$a = new Solution();
print_r($a->merge([[1,4],[0,4]]));

/**
 * 解题思路：
 *   按左区间排序  后对比 当前元素的左区间和下一元素的右区间 是否有重叠
 *
 */