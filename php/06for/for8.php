<?php
$zarb = 1;
do{
	$A = mt_rand(1, 100);
	$B = mt_rand(1, 100);
}
while($A >= $B);
echo "A = ".$A;
echo "<br>B = ".$B."<br>";
for($i = $A; $i <= $B; $i++)
	$zarb *= $i;
echo "ҲосилиЗарб: ".$zarb;
?>