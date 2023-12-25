<?php
$K = mt_rand(1, 100);
$N = mt_rand(1, 20);
echo "K = ".$K;
echo "<br>N = ".$N."<br><br>";
for($i = 1; $i <= $N; $i++)
	echo $K."<br>";
?>

/*
<?php
	$K = intval(readline("K = "));
	$N = intval(readline("N = "));
	for ($i = 0; $i < $N; ++$i) {
		echo $K."\t";
	}
?>
*/