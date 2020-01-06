<?php
class Solution {

    /**
     * @param Integer[][] $grid
     * @return Integer
     */
    function minPathSum($grid) {
        $row = count($grid);
        if ($row === 0) {
            return 0;
        }
        $column = count($grid[0]);
        $result = [];
        $result[0][0] = $grid[0][0];
        for ($i = 0; $i < $row ; $i++) {
            for ($j = 0; $j < $column; $j ++) {

                if ($i === 0 && $j ===0 ) {
                    continue;
                }
                $result[$i][$j] = min($result[$i][$j - 1] ?? $result[$i - 1][$j], $result[$i - 1][$j] ?? $result[$i][$j - 1] ) + $grid[$i][$j];

            }
        }

        return $result[$row - 1][$column - 1];
    }

}


$a = new Solution();
print_r($a->minPathSum([
    [1,3,1],
    [1,5,1],
    [4,2,1]
]));

/**
 * 解题思路：
 *   opt(i,j) = min(opt(i-1,j)  , opt(i,j-1)) + num(i,j)
 *   构造同样的二维数组 每个元素为当前位置最小路径 去右下角点即为结果
 *   时间复杂度O(m*n)
 */