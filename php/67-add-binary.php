<?php
class Solution {

    /**
     * @param String $a
     * @param String $b
     * @return String
     */
    function addBinary($a, $b) {
        $result = '';
        $tmp = 0;
        for ($i = strlen($a) - 1, $j = strlen($b) - 1; $i >=0 || $j >=0 ; $i--, $j--) {
            $sum = $tmp;
            $sum += $i >= 0 ? $a[$i] - '0' : 0;
            $sum += $j >= 0 ? $b[$j] - '0' : 0;
            $result .= (int)($sum % 2);
            $tmp = (int) ($sum / 2);
        }
        $result .= $tmp == 1 ? $tmp : "";
        return implode("", array_reverse(str_split($result)));
    }
}