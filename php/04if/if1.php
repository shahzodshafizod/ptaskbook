<?php
$N = mt_rand(-100, 100);
echo "Адади додашуда: ".$N;
if($N > 0)
	$N++;
echo "<br>Адад пас аз иҷрои шарт: ".$N;
?>

/*
<?php
	$number = intval(readline("number:\t"));
	if ($number > 0) {
		$number++;
	}
	echo $number;
?>
*/

/*
<?php
	$number = intval(readline("number:\t"));
	$number += $number > 0 ? 1 : 0;
	echo $number;
?>
*/