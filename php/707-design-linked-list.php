<?php
class Node{
    public $val = null;
    public $next = null;

    public function __construct($val = null)
    {
        $this->val = $val;
    }
}
class MyLinkedList {
    private $head = null;
    /**
     * Initialize your data structure here.
     */
    function __construct() {
        $this->head = new Node();
    }

    /**
     * Get the value of the index-th node in the linked list. If the index is invalid, return -1.
     * @param Integer $index
     * @return Integer
     */
    function get($index) {
        /**@var Node $tmp**/
        $tmp = $this->head->next;
        for ($i = 0; $i < $index; $i ++) {
            $tmp = $tmp->next;
            if ($tmp === null) {
                return -1;
            }

        }
        return $tmp->val;
    }

    /**
     * Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list.
     * @param Integer $val
     * @return NULL
     */
    function addAtHead($val) {
        $addNode = new Node($val);
        $addNode->next = $this->head->next;
        $this->head->next = $addNode;
    }

    /**
     * Append a node of value val to the last element of the linked list.
     * @param Integer $val
     * @return NULL
     */
    function addAtTail($val) {
        $tmp = $this->head->next;
        while ($tmp !== null) {
            $tmp = &$tmp->next;
        }
        $tmp = new Node($val);
    }

    /**
     * Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted.
     * @param Integer $index
     * @param Integer $val
     * @return NULL
     */
    function addAtIndex($index, $val) {
        $addNode = new Node($val);
        $tmp = $this->head;
        for ($i = 0; $i < $index; $i++) {
            $tmp = $tmp->next;
        }
        $addNode->next = $tmp->next;
        $tmp->next = $addNode;

    }

    /**
     * Delete the index-th node in the linked list, if the index is valid.
     * @param Integer $index
     * @return NULL
     */
    function deleteAtIndex($index) {
        $tmp = $this->head;
        for ($i = 0; $i < $index; $i++) {
            if ($tmp == null) {
                return ;
            }
            $tmp = $tmp->next;
        }
        $tmp->next =$tmp->next->next ?? null;
    }
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * $obj = MyLinkedList();
 * $ret_1 = $obj->get($index);
 * $obj->addAtHead($val);
 * $obj->addAtTail($val);
 * $obj->addAtIndex($index, $val);
 * $obj->deleteAtIndex($index);
 */

$obj = new MyLinkedList();
//$ret_1 = $obj->get($index);
$obj->addAtHead(4);
echo $obj->get(1);
$obj->addAtTail(3);
$obj->addAtIndex(1, 2);
echo $obj->get(1);
$obj->deleteAtIndex(0);
echo $obj->get(0);