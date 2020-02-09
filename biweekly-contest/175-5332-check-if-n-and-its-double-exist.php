<?php
class Solution {

    /**
     * @param Integer[] $arr
     * @return Boolean
     */
    function checkIfExist($arr) {
        foreach ($arr as $key => $value) {
            $tmp = $arr;
            unset($tmp[$key]);
            if (in_array($value*2,$tmp)) {
                return true;
            }
        }
        return false;
    }
}

$a = new Solution();
var_dump($a->checkIfExist([0,0]));