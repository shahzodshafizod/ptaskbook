<?php
$a = mt_rand(-100, 100);
$b = mt_rand(-100, 100);
echo "A = ".$a;
echo "<br>B = ".$b;
$c = $a;
$a = $b;
$b = $c;
echo "<br><br>A = ".$a;
echo "<br>B = ".$b;
?>