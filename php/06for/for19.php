<?php
$fact = 1;
$N = mt_rand(1, 50);
echo "Ҳудуди давр: ".$N."<br>";
for($i = 1; $i <= $N; $i++)
	$fact *= $i;
echo $N."! = ".$fact;
?>