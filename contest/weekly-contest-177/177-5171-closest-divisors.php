<?php
class Solution {

    /**
     * @param Integer $num
     * @return Integer[]
     */
    function closestDivisors($num) {


        $_closestDivisors = function ($tmp) {
            $start = ceil(sqrt($tmp));
            while (true) {
                if ($tmp % $start == 0) {
                    return [abs($start - $tmp/$start),[$start, $tmp/$start]];

                }
                $start -= 1;
            }
        };

        list($a, $resultA) = $_closestDivisors($num +1);
        list($b, $resultB) = $_closestDivisors($num +2);

        if ($a < $b) {
            return $resultA;
        }else{
            return $resultB;
        }
    }
}

$a = new Solution();
print_r($a->closestDivisors(170967091));