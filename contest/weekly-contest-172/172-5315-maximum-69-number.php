<?php
class Solution {

    /**
     * @param Integer $num
     * @return Integer
     */
    function maximum69Number ($num) {
       $string = (string) $num;
       $len = strlen($string);
       for ($i = 0; $i < $len; $i ++) {
           if ($string[$i] === '6') {
               $string[$i] = '9';
               break;
           }
       }
       return (int) $string;
    }
}