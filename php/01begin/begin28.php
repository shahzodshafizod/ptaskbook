<?php
$A = mt_rand(-100, 100);
echo "Данное число: ".$A;
$tagh = $A * $A;
$nat = $tagh;
echo "<br>A^2 = ".$tagh;
$nat *= $A;
echo "<br>A^3 = ".$nat;
$nat *= $tagh;
$tagh = $nat;
echo "<br>A^5 = ".$tagh;
$tagh *= $tagh;
echo "<br>A^10 = ".$tagh;        
$nat *= $tagh;
echo "<br>A^15 = ".$nat;
?>

/*
<?php
$A = mt_rand(-100, 100);
echo "Адади додашуда: ".$A;
$tagh = $A * $A;
$nat = $tagh;
echo "<br>A^2 = ".$tagh;
$nat *= $A;
echo "<br>A^3 = ".$nat;
$nat *= $tagh;
$tagh = $nat;
echo "<br>A^5 = ".$tagh;
$tagh *= $tagh;
echo "<br>A^10 = ".$tagh;        
$nat *= $tagh;
echo "<br>A^15 = ".$nat;
?>
*/