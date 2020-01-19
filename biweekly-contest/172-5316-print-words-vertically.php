<?php
class Solution {

    /**
     * @param String $s
     * @return String[]
     */
    function printVertically($s) {
        $arr = explode(' ', $s);
        $result = [];
        $i = 0;
        while (true) {
            $string = '';
            foreach ($arr as $item) {
                $tmp =$item[$i] ?? ' ';
                $string .= $tmp;
            }
            $string = rtrim($string);
            if ($string === '') {
                break;
            }
            $result[] = $string;
            $i++;
        }
        return $result;
    }
}

$a = new Solution();
print_r($a->printVertically("HOW ARE YOU"));