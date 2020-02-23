<?php
class TweetCounts {
    public $tweet = [];
    /**
     */
    function __construct() {

    }

    /**
     * @param String $tweetName
     * @param Integer $time
     * @return NULL
     */
    function recordTweet($tweetName, $time) {
        $this->tweet[$tweetName][$time]= $this->tweet[$tweetName][$time] ?? 0;
        $this->tweet[$tweetName][$time] ++;
    }

    /**
     * @param String $freq
     * @param String $tweetName
     * @param Integer $startTime
     * @param Integer $endTime
     * @return Integer[]
     */
    function getTweetCountsPerFrequency($freq, $tweetName, $startTime, $endTime) {
        $result = [];
        switch ($freq) {
            case "minute":
                $tmp = [];
                for($i = $startTime; $i< $endTime ; $i ++) {
                    $mod = intval($i / 60);
                    if(!isset($tmp[$mod])) {
                        $tmp[$mod] = 0;
                    }
                    if (isset($this->tweet[$tweetName][$i])) {
                        $tmp[$mod] ++;
                    }
                }
                $result = array_merge($result, $tmp);
                break;
            case "hour":
                $tmp = [];
                for($i = $startTime; $i< $endTime ; $i ++) {
                    $mod = intval($i / 3600);
                    if(!isset($tmp[$mod])) {
                        $tmp[$mod] = 0;
                    }
                    if (isset($this->tweet[$tweetName][$i])) {
                        $tmp[$mod] ++;
                    }
                }
                $result = array_merge($result, $tmp);
                break;
            case "day":
                $tmp = [];
                for($i = $startTime; $i< $endTime ; $i ++) {
                    $mod = intval($i / (3600 * 24));
                    if(!isset($tmp[$mod])) {
                        $tmp[$mod] = 0;
                    }
                    if (isset($this->tweet[$tweetName][$i])) {
                        $tmp[$mod] ++;
                    }
                }
                $result = array_merge($result, $tmp);
                break;
        }
        return $result;
    }
}

/**
 * Your TweetCounts object will be instantiated and called as such:
 * $obj = TweetCounts();
 * $obj->recordTweet($tweetName, $time);
 * $ret_2 = $obj->getTweetCountsPerFrequency($freq, $tweetName, $startTime, $endTime);
 */

$a = new TweetCounts();
$a->recordTweet("tweet3", 0);
$a->recordTweet("tweet3", 60);
$a->recordTweet("tweet3", 10);
print_r($a->getTweetCountsPerFrequency("minute", "tweet3", 0, 59));
print_r($a->getTweetCountsPerFrequency("minute", "tweet3", 0, 60));
$a->recordTweet("tweet3", 120);
print_r($a->getTweetCountsPerFrequency("hour", "tweet3", 0, 210));