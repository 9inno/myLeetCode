<?php
class Solution {

    /**
     * @param Integer[][] $events
     * @return Integer
     */
    function maxEvents($events) {
        sort($events);
        $map = [];
        foreach ($events as $item) {
            foreach (range($item[0], $item[1]) as $value) {
                if (!isset($map[$value])) {
                    $map[$value] = true;
                }
            }
        }
        return  count($map);
    }
}

$a = new Solution();
echo $a->maxEvents([[27,27],[8,10],[9,11],[20,21],[25,29],[17,20],[12,12],[12,12],[10,14],[7,7],[6,10],[7,7],[4,8],[30,31],[23,25],[4,6],[17,17],[13,14],[6,9],[13,14]]);