<?php
class Solution {

    /**
     * @param String $S
     * @return Integer[]
     */
    function diStringMatch($S) {

        $count = strlen($S);
        if ($count === 0) {
            return [0];
        }
        $arr = range(0, $count);
        $result = [];
        $left = 0;
        $right = $count;
        for ($i = 0; $i < $count ; $i ++) {
            if ($S[$i] == 'I') {
                if (empty($result)) {
                    $result[] = 0;
                }else {
                    $result[] = $arr[$left];
                }
                $left++;
            }
            if ($S[$i] == 'D') {
                if (empty($result)) {
                    $result[] = $arr[$right];
                }else {
                    $result[] = $arr[$right];
                }
                $right --;
            }
            if ($left == $right) {
                $result[] = $arr[$left];
            }
        }
        return $result;
    }
}

$a = new Solution();
print_r($a->diStringMatch("IDID"));