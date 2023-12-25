<?php
do{
	$A = mt_rand(1, 100);
	$B = mt_rand(1, 100);
}
while($A <= $B);
echo "A = ".$A;
echo "<br>B = ".$B;
while($A >= $B)
	$A -= $B;
echo "<br>қисми банднабуда = ".$A;
?>

/*
<?php
	$A = floatval(readline("A = "));
	$B = floatval(readline("B = "));
	while ($A >= $B) {
		$A -= $B;
	}
	echo $A;
?>
*/