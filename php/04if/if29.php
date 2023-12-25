<?php
$N = mt_rand(-1000, 1000);
echo "Адади ".$N." - адади ";
if ($N != 0){
	if(($N % 2) == 0)
		echo "ҷуфти ";
	else
		echo "тоқи ";
	if($N > 0)
		echo "мусбӣ ";
	else
		echo "манфӣ ";
}
else
	echo "нулӣ ";
echo "аст.";
?>