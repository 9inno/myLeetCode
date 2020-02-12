<?php
class Solution {

    /**
     * @param String $str
     * @return String
     */
    function toLowerCase($str) {
        $len = strlen($str);
        for ($i = 0; $i < $len ; $i ++) {
            $ascii = ord($str[$i]);
            if ($ascii > 64 && $ascii < 91) {
                $str[$i] = chr($ascii + 32);
            }
        }
        return $str;
    }
}