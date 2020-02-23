<?php
class Solution {

    /**
     * @param String $date1
     * @param String $date2
     * @return Integer
     */
    function daysBetweenDates($date1, $date2) {
        $time1 = strtotime($date1);
        $time2 = strtotime($date2);
        return abs($time1 - $time2) / (24 * 60 * 60);
    }
}