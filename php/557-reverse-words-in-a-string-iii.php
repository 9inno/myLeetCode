<?php
class Solution {

    /**
     * @param String $s
     * @return String
     */
    function reverseWords($s) {
        $arr = explode(' ', $s);
        foreach ($arr as &$value) {
            $j = strlen($value) - 1;
            $i = 0;
            while (true) {

                if ($j <= $i) {
                    break;
                }
                list($value[$i], $value[$j]) = [$value[$j], $value[$i]];
                $i++;
                $j--;
            }
        }

        return implode(' ', $arr);
    }
}

$a = new Solution();
echo $a->reverseWords('Let\'s take LeetCode contest');