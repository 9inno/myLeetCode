<?php
class Solution {

    private $map = [
        1000 => "M",
        900  => "CM",
        500  => "D",
        400  => "CD",
        100  => "C",
        90   => "XC",
        50   => "L",
        40   => "XL",
        10   => "X",
        9    => "IX",
        5    => "V",
        4    => "IV",
        1    => "I",
    ];

    /**
     * @param Integer $num
     * @return String
     */
    function intToRoman($num) {

        $str = '';

        foreach ($this->map as $key => $value) {
            if ($num < $key){
                continue;
            }
            $multi =intval($num / $key);
            while ($multi > 0){

                $str .= $this->map[$key];
                $num = $num - $key;
                $multi -- ;
            }
        }
        return $str;
    }
}

$a = new Solution();
echo $a->intToRoman(123);

/**
 *  解题思路：
 *   遍历枚举map 从高位向下算 商值 不为1的情况说明有连续的罗马数字
 *   需要循环补充
 *   时间复杂度O(1)  只用遍历既定的map  
 */


