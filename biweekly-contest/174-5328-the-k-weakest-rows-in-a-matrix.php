<?php
class Solution {

    /**
     * @param Integer[][] $mat
     * @param Integer $k
     * @return Integer[]
     */
    function kWeakestRows($mat, $k) {

        $count = count($mat);
        $map = [];
        for($i = 0; $i < $count; $i++) {
            $tmpCount= 0;
            foreach ($mat[$i] as $item) {
                if ($item == 1) {
                    $tmpCount++;
                } else {
                    break;
                }
            }

            $map[$i]['key'] = $i;
            $map[$i]['count'] = $tmpCount;
        }

        array_multisort(array_column($map,'count'),SORT_ASC,array_column($map,'key'),SORT_ASC,$map);
        $result = [];
        for ($j = 0; $j < $k; $j++) {
            $result[] = $map[$j]['key'];
        }
        return $result;
    }
}


$a = new Solution();
print_r($a->kWeakestRows([[1,1,0,0,0],
    [1,1,1,1,0],
    [1,0,0,0,0],
    [1,1,0,0,0],
    [1,1,1,1,1]],3));