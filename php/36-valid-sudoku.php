<?php
class Solution {

    /**
     * @param String[][] $board
     * @return Boolean
     */
    function isValidSudoku($board) {
        $row = [];
        $column = [];
        $bulk = [];

        for ($i = 0; $i < 9; $i++ ) {
            for ($j = 0; $j < 9; $j++) {

                if (isset($row[$i][$board[$i][$j]])) {
                    return false;
                }
                if (isset($column[$j][$board[$i][$j]])) {
                    return false;
                }
                if (isset($bulk[intval($i / 3) * 3 + intval($j /3)][$board[$i][$j]])) {
                    return false;
                }
                if ($board[$i][$j] != '.') {
                    $row[$i][$board[$i][$j]] = $column[$j][$board[$i][$j]] = $bulk[intval($i / 3) * 3 + intval($j /3)][$board[$i][$j]] = true;
                }
            }
        }

        return true;
    }
}

$a = new Solution();
var_dump($a->isValidSudoku([
    ["5","3",".",".","7",".",".",".","."],
    ["6",".",".","1","9","5",".",".","."],
    [".","9","8",".",".",".",".","6","."],
    ["8",".",".",".","6",".",".",".","3"],
    ["4",".",".","8",".","3",".",".","1"],
    ["7",".",".",".","2",".",".",".","6"],
    [".","6",".",".",".",".","2","8","."],
    [".",".",".","4","1","9",".",".","5"],
    [".",".",".",".","8",".",".","7","9"]
]
));


/**
 * 解题思路 ：
 *   记录 行和列和3*3 块的值  对应做成hashmap  遍历整个数组
 *   时间复杂度O(1)  因固定大小9*9
 */