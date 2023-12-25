<?php
$rebmuN = 0;
$Number = mt_rand(1, 1000);
echo "Number = ".$Number;
while($Number > 0){
	$baq = $Number % 10;
	$rebmuN = $rebmuN * 10 + $baq;
	$Number = intval($Number / 10);
}
echo "<br>rebmuN = ".$rebmuN;
?>