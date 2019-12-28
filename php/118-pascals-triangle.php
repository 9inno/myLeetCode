<?php
class Solution {

    /**
     * @param Integer $numRows
     * @return Integer[][]
     */
    function generate($numRows) {
        if ($numRows === 0) {
            return [];
        } elseif ($numRows === 1) {
            return [[1]];
        } elseif ($numRows === 2) {
            return [[1],[1,1]];
        }
        $result = [[1],[1,1]];
        for ($i = 2; $i < $numRows; $i ++) {
            $tmp =[];
            array_push($tmp, 1);
            for ($j = 1; $j < $i  ; $j++) {
                array_push($tmp, $result[$i-1][$j-1] + $result[$i-1][$j] );
            }
            array_push($tmp, 1);
            array_push($result, $tmp);
        }

        return $result;
    }
}

$a = new Solution();
print_r($a->generate(5));

/**
 * 解题思路:
 *   在杨辉三角中，每个数是它左上方和右上方的数的和。
 *   两边为 1 行中间元素为  $arr[$i-1][$j-1] + $arr[$i-1][$j]
 *   时间复杂度O(n^2)
 */