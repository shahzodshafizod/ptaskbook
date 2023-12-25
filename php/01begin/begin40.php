<?php
$A1 = mt_rand(-100, 100);
$A2 = mt_rand(-100, 100);
$B1 = mt_rand(-100, 100);
$B2 = mt_rand(-100, 100);
$C1 = mt_rand(-100, 100);
$C2 = mt_rand(-100, 100);
$D = $A1 * $B2 - $A2 * $B1;
$x = ($C1 * $B2 - $C2 * $B1) / $D;
$y = ($A1 * $C2 - $A2 * $C1) / $D;
echo "A1 = ".$A1;
$x = number_format($x, 2);
$y = number_format($y, 2);
echo "<br>B1 = ".$B1;
echo "<br>C1 = ".$C1;
echo "<br>A2 = ".$A2;
echo "<br>B2 = ".$B2;
echo "<br>C2 = ".$C2;
echo "<br>D = ".$D;
echo "<br>x = ".$x;
echo "<br>y = ".$y;
?>

/*
<?php
$A1 = mt_rand(-100, 100);
$A2 = mt_rand(-100, 100);
$B1 = mt_rand(-100, 100);
$B2 = mt_rand(-100, 100);
$C1 = mt_rand(-100, 100);
$C2 = mt_rand(-100, 100);
$D = $A1 * $B2 - $A2 * $B1;
$x = ($C1 * $B2 - $C2 * $B1) / $D;
$y = ($A1 * $C2 - $A2 * $C1) / $D;
echo "A1 = ".$A1;
$x = number_format($x, 2);
$y = number_format($y, 2);
echo "<br>B1 = ".$B1;
echo "<br>C1 = ".$C1;
echo "<br>A2 = ".$A2;
echo "<br>B2 = ".$B2;
echo "<br>C2 = ".$C2;
echo "<br>D = ".$D;
echo "<br>x = ".$x;
echo "<br>y = ".$y;
?>
*/