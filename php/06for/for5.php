<?php
$narx = mt_rand(1, 100) / 4;
$narx = number_format($narx, 2);
echo "Нарх: ".$narx."<br><br>";
for($i = 0.1; $i < 1.1; $i += 0.1)
	echo $i." кг = ".number_format(($i * $narx),2)." сомонӣ<br>";
?>