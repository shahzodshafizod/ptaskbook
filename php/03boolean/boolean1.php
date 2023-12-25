<?php
function boolval($nat){
	return ($nat? "true": "false");
}
$A = mt_rand(-100, 100);
$plus = ($A > 0);
echo "Данное число: ".$A;
echo "<br>Положительность числа: ".boolval($plus);
?>

/*
<?php
	$A = intval(readline("A = "));
	$isPositive = $A > 0;
	echo var_export($isPositive);
?>
*/