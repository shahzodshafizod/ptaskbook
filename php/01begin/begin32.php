<?php
$TC = mt_rand(-100, 100);
$TF = 9 * $TC / 5 + 32;
$TF = number_format($TF, 2);
echo "Температура Селция: ".$TC;
echo "<br>Температура Фаренгейта: ".$TF;
?>

/*
<?php
$TC = mt_rand(-100, 100);
$TF = 9 * $TC / 5 + 32;
$TF = number_format($TF, 2);
echo "Ҳарорати Селсий: ".$TC;
echo "<br>Ҳарорати Фаренгейт: ".$TF;
?>
*/