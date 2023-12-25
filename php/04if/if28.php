<?php
$sol = mt_rand(1, 3000);
if(($sol % 400) == 0)
	$ruz = 366;
else if(($sol % 100) == 0)
	$ruz = 365;
else if(($sol % 4) == 0)
	$ruz = 366;
else
	$ruz = 365;
echo "Соли ".$sol." дорои ".$ruz." рӯз аст";
?>