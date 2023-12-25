<?php
$A = mt_rand(-100, 100) / 4;
$B = mt_rand(-100, 100) / 4;
$C = mt_rand(-100, 100) / 4;
$A = number_format($A, 2);
$B = number_format($B, 2);
$C = number_format($C, 2);
echo "Первое число: ".$A;
echo "<br>Второе число: ".$B;
echo "<br>Третье число: ".$C;
$D = $A;
$A = $B;
$B = $C;
$C = $D;
echo "<br><br>Первое число: ".$A;
echo "<br>Второе число: ".$B;
echo "<br>Третье число: ".$C;
?>

/*
<?php
$A = mt_rand(-100, 100) / 4;
$B = mt_rand(-100, 100) / 4;
$C = mt_rand(-100, 100) / 4;
$A = number_format($A, 2);
$B = number_format($B, 2);
$C = number_format($C, 2);
echo "Адади якӯм: ".$A;
echo "<br>Адади дуюм: ".$B;
echo "<br>Адади сеюм: ".$C;
$D = $A;
$A = $B;
$B = $C;
$C = $D;
echo "<br><br>Адади якӯм: ".$A;
echo "<br>Адади дуюм: ".$B;
echo "<br>Адади сеюм: ".$C;
?>
*/