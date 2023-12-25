<?php

    require_once "Tree.php";

    $tree = new Tree();
    $tree->automake(rand(7, 14));
    $tree->display();
?>