<?php
class Solution {

    /**
     * @param Integer[][] $grid
     * @return Integer
     */
    function countNegatives($grid) {
        $result = 0;
        foreach ($grid as $item) {
            $count = count($item);
            $tmp = 0;
            for ($i = 0; $i < $count; $i++) {
                if ($item[$i] < 0) {
                    $tmp = $count - $i;
                    break;
                }
            }
            $result += $tmp;
        }

        return $result;
    }
}