<?php
class Solution {

    /**
     * @param Integer $n
     * @return String[]
     */
    function generateParenthesis($n) {
        $result = [];
        if ($n === 0) {
            return [];
        }
        $this->recursion($result, $n, $n, '');
        return $result;

    }

    public function recursion(&$result, $left, $right, $string) {
        if ($left > $right) {
            return;
        }
        if( $left === 0 && $right === 0) {
            array_push($result, $string);
        }
        if ($left > 0) {
            $this->recursion($result, $left-1, $right,$string."(");
        }
        if ($right > 0) {
            $this->recursion($result, $left, $right - 1, $string.")");
        }
    }

}

$a = new Solution();
print_r($a->generateParenthesis(3));