<?php
class Solution {

    /**
     * @param Integer[] $arr
     * @return Integer
     */
    function minSetSize($arr) {
        $map = array_count_values($arr);
        $target = intval(count($arr)  / 2 );
        rsort($map);
        $min = 0;
        $count = 0;
        foreach ($map as $item) {
            if ($count < $target) {
                $min++;
                $count += $item;
            } else {
                break;
            }
        }
        return $min;
    }


}

$a = new Solution();
echo $a->minSetSize([1,2,3,4,5,6,7,8,9,10]);