<?php
class Solution {

    /**
     * @param String $s
     * @param Integer $numRows
     * @return String
     */
    function convert($s, $numRows) {

        $flag = true;
        $result = [];
        $row = 1;
        for ($i = 0; $i< strlen($s); $i++) {

            $result[$row] .= $s[$i];
            if ($row === $numRows) {
                $flag = false;
            } elseif ($row ===1) {
                $flag = true;
            }

            if ($flag) {
                $row ++;
            } else {
                $row --;
            }

        }
        $str ='';
        foreach ($result as $value) {
            $str .= $value;
        }
        return $str;
    }
}


$a = new Solution();
echo $a->convert('LEETCODEISHIRING',3);


/**
 * 解题思路：
 *   遍历字符串  按z字形拼接数组元素 用flag判断 数组元素的顺序
 *   后合并数组元素
 *
 *   时间复杂度O(n)
 */