<?php
function PowerA3($A, &$B){
        $B = $A * $A * $A;
}
$hisob = 1;
while($hisob <= 5){
	$a = mt_rand(-40, 40);
	PowerA3($a, $b);
	echo "Адади додашуда: ".$a;
	echo "<br>Куби адади додашуда: ".$b."<br><br>";
	$hisob++;
}
?>

/*
<?php
	for ($i = 1; $i <= 5; ++$i) {
		$a = floatval(readline("A$i = "));
		$b = 0;
		PowerA3($a, $b);
		echo sprintf("B = %.2f\n", $b);
	}
	
	function PowerA3($a, &$b) {
		$b = $a * $a * $a;
	}
?>
*/