<?php
$TF = mt_rand(-100, 100);
$TC = ($TF - 32) * 5 / 9;
$TC = number_format($TC, 2);
echo "Температура Фаренгейта: ".$TF;
echo "<br>Температура Селция: ".$TC;
?>

/*
<?php
$TF = mt_rand(-100, 100);
$TC = ($TF - 32) * 5 / 9;
$TC = number_format($TC, 2);
echo "Ҳарорати Фаренгейт: ".$TF;
echo "<br>Ҳарорати Селсий: ".$TC;
?>
*/