<?php
$N = mt_rand(1, 999);
echo "Адади ".$N." - адади ";
if(($N > 9) && ($N < 100))
	echo "дурақамаи ";
else if(($N > 0) && ($N < 10))
	echo "якрақамаи ";
else
	echo "серақамаи ";
if(($N % 2) == 0)
	echo "ҷуфт ";
else
	echo "тоқ ";
echo "аст.";
?>