<?php
$A = mt_rand(-100, 100) / 4;
$B = mt_rand(-100, 100) / 4;
$A = number_format($A, 2);
$B = number_format($B, 2);
echo "Первое число: ".$A;
echo "<br>Второе число: ".$B;
$C = $A;
$A = $B;
$B = $C;
echo "<br><br>Первое число: ".$A;
echo "<br>Второе число: ".$B;
?>

/*
<?php
$A = mt_rand(-100, 100) / 4;
$B = mt_rand(-100, 100) / 4;
$A = number_format($A, 2);
$B = number_format($B, 2);
echo "Адади якӯм: ".$A;
echo "<br>Адади дуюм: ".$B;
$C = $A;
$A = $B;
$B = $C;
echo "<br><br>Адади якӯм: ".$A;
echo "<br>Адади дуюм: ".$B;
?>
*/