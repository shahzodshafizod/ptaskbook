<?php
$N = mt_rand(6, 14);
$M = mt_rand(1, 4);
echo "N = ".$N;
echo "<br>M = ".$M."<br>";
switch($N){
	case 6:
		echo "шестерка";
		break;
	case 7:
		echo "семерка";
		break;
	case 8:
		echo "восьмерка";
		break;
	case 9:
		echo "девятка";
		break;
	case 10:
		echo "десятка";
		break;
	case 11:
		echo "валет";
		break;
	case 12:
		echo "дама";
		break;
	case 13:
		echo "кароль";
		break;
	case 14:
		echo "туз";
}
switch($M){
	case 1:
		echo " пиков";
		break;
	case 2:
		echo " треф";
		break;
	case 3:
		echo " бубен";
		break;
	case 4:
		echo " червей";
}
?>

/*
<?php
$N = mt_rand(6, 14);
$M = mt_rand(1, 4);
echo "N = ".$N;
echo "<br>M = ".$M."<br>";
switch($N){
	case 6:
		echo "шаш";
		break;
	case 7:
		echo "ҳафт";
		break;
	case 8:
		echo "ҳашт";
		break;
	case 9:
		echo "нӯҳ";
		break;
	case 10:
		echo "даҳ";
		break;
	case 11:
		echo "валет";
		break;
	case 12:
		echo "дама";
		break;
	case 13:
		echo "шоҳ";
		break;
	case 14:
		echo "туз";
}
echo "и ";
switch($M){
	case 1:
		echo "дил";
		break;
	case 2:
		echo "паша";
		break;
	case 3:
		echo "хишт";
		break;
	case 4:
		echo "таппон";
}
?>
*/