<?php
class Solution {
    private $map = [
        '2' => ['a', 'b', 'c'],
        '3' => ['d', 'e', 'f'],
        '4' => ['g', 'h', 'i'],
        '5' => ['j', 'k', 'l'],
        '6' => ['m', 'n', 'o'],
        '7' => ['p', 'q', 'r', 's'],
        '8' => ['t', 'u', 'v'],
        '9' => ['w', 'x', 'y', 'z']
    ];
    /**
     * @param String $digits
     * @return String[]
     */
    function letterCombinations($digits) {
        $len = strlen($digits);
        $result = [];
        $this->recursion($result, '', 0, $len, $digits);
        return $result;
    }

    public function recursion(&$result, $string, $n, $len, $digits) {

        if ($n === $len && $len != 0) {
            array_push($result, $string);
            return;
        }
        if ($n < $len) {
            foreach ($this->map[$digits[$n]] as $value) {
                $this->recursion($result, $string.$value, $n + 1, $len, $digits);
            }
        }
    }
}


$a = new Solution();
print_r($a->letterCombinations('23'));


/**
 * 解题思路 ：
 *               2
 *             /|\
 *           / | \
 *         a  b  c
 *       /   |   \
 *     3    3    3
 *    /|\  /|\  /|\
 *   def   def  def
 *  指针从字符串头开始 遍历对应的字母 并 向下递归 传入下一指针与当前字母字符串
 *  递归出口为 指针指向字符串最后 存入数组
 *  注意边界值 当字符串为空字符串时 不计结果
 *  时间复杂度 O(3^N * 4^M)
 * 
 */